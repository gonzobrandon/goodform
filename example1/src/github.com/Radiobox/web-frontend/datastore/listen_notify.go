package datastore

import (
	"log"
	"sync"
	"time"

	"github.com/Radiobox/web-frontend/settings"
	"github.com/lib/pq"
)

var defaultListenMap *ListenMapper

// reportListenerProblem reports any problems that a listener runs in
// to, since it won't be running in a main thread and its errors need
// to be reported somewhere.
func reportListenerProblem(ev pq.ListenerEventType, err error) {
	if err != nil {
		go log.Print("Listener error: ", err)
	}
}

// A Listener is a process which listens to a single channel name.
// When the Go code receives a postgres NOTIFY event, it walks through
// a list of registered Listeners and attempts to send the event on
// one of the Listener's channels.
//
// When a reconnect event is received by the Go code, *all* Listeners
// are sent a true value on Listener.Reconnect.
//
// If either Listener.Reconnect or Listener.Payloads blocks for more
// than five seconds, the Listener will be closed.
type Listener struct {
	ChannelName string
	Reconnect   chan bool
	Payloads    chan string
	Closed      bool
	cancel      chan bool
	wait        *sync.WaitGroup
	mapper      *ListenMapper
}

func (listener *Listener) Close() {
	listener.Closed = true
	listener.cancel <- true
	listener.wait.Wait()
	close(listener.Reconnect)
	close(listener.Payloads)
	close(listener.cancel)
	listener.mapper.Prune()
}

func (listener *Listener) Send(payload string) {
	if !listener.Closed {
		listener.wait.Add(1)
		select {
		case listener.Payloads <- payload:
		case <-time.After(5 * time.Second):
		case <-listener.cancel:
			listener.cancel <- true
		}
		listener.wait.Done()
	}
}

func (listener *Listener) SendReconnect() {
	if !listener.Closed {
		listener.wait.Add(1)
		select {
		case listener.Reconnect <- true:
		case <-time.After(5 * time.Second):
		case <-listener.cancel:
			listener.cancel <- true
		}
		listener.wait.Done()
	}
}

// Our ListenerMapper is a type of listener that notifies different
// channels based on the channel that a notification was received on.
type ListenMapper struct {
	*pq.Listener
	listeners    []*Listener
	newListeners chan *Listener
	shutdown     chan bool
	listening    map[string]bool
}

func DefaultListenMapper() *ListenMapper {
	if defaultListenMap == nil {
		pqListener := pq.NewListener(settings.DbConnectionString,
			10*time.Second, time.Minute, reportListenerProblem)
		defaultListenMap = &ListenMapper{
			Listener:  pqListener,
			listening: make(map[string]bool),
		}
	}
	return defaultListenMap
}

// Only one instance of run() can be running per ListenMapper
func (listenMapper *ListenMapper) run() {
	var notification *pq.Notification
	for {
		select {
		case notification = <-listenMapper.Listener.Notify:
			if notification == nil {
				go listenMapper.NotifyReconnect()
			} else {
				go listenMapper.Notify(notification.Channel, notification.Extra)
			}
		case newListener := <-listenMapper.newListeners:
			listenMapper.listeners = append(listenMapper.listeners, newListener)
		case <-listenMapper.shutdown:
			for _, listener := range listenMapper.listeners {
				listener.Close()
			}
			close(listenMapper.newListeners)
			close(listenMapper.shutdown)
			return
		}
		// Get rid of dead connections while we're here.
		listenMapper.Prune()
	}
}

func (listenMapper *ListenMapper) Shutdown() {
	listenMapper.shutdown <- true
}

// Prune goes through the list of Listeners and removes any that are
// closed.
func (listenMapper *ListenMapper) Prune() {
	// Mark all listening values as false
	for channel := range listenMapper.listening {
		listenMapper.listening[channel] = false
	}
	for index := 0; index < len(listenMapper.listeners); {
		listener := listenMapper.listeners[index]
		if listener.Closed {
			listenMapper.listeners = append(listenMapper.listeners[:index], listenMapper.listeners[index+1:]...)
		} else {
			// If the listener is still listening, mark the listening
			// value as true for that channel.
			listenMapper.listening[listener.ChannelName] = true
			index++
		}
	}
	// Unlisten all channels that have no listeners that are currently
	// being used.
	for channel, listen := range listenMapper.listening {
		if !listen {
			listenMapper.Listener.Unlisten(channel)
			delete(listenMapper.listening, channel)
		}
	}
}

// NotifyReconnect runs SendReconnect() on every listener that has
// been mapped.
func (listenMapper *ListenMapper) NotifyReconnect() {
	for _, listener := range listenMapper.listeners {
		go listener.SendReconnect()
	}
}

// Notify runs Send(payload) for each listener that has been mapped
// and has a ChannelName that matches channel.
func (listenMapper *ListenMapper) Notify(channel string, payload string) {
	for _, listener := range listenMapper.listeners {
		if listener.ChannelName == channel {
			go listener.Send(payload)
		}
	}
}

// Listen creates, maps, and returns a new Listener for the requested
// channel name.
func (listenMapper *ListenMapper) Listen(channel string) (*Listener, error) {
	if _, ok := listenMapper.listening[channel]; !ok {
		if err := listenMapper.Listener.Listen(channel); err != nil {
			return nil, err
		}
		listenMapper.listening[channel] = true
	}
	if listenMapper.newListeners == nil && listenMapper.listeners == nil && listenMapper.shutdown == nil {
		listenMapper.newListeners = make(chan *Listener, 5)
		listenMapper.listeners = make([]*Listener, 0, 5)
		listenMapper.shutdown = make(chan bool)
		go listenMapper.run()
	}
	newListener := &Listener{
		ChannelName: channel,
		Reconnect:   make(chan bool, 2),
		Payloads:    make(chan string, 5),
		Closed:      false,
		cancel:      make(chan bool, 1),
		wait:        new(sync.WaitGroup),
		mapper:      listenMapper,
	}
	listenMapper.newListeners <- newListener
	return newListener, nil
}

// Listen is a convenience function that returns
// DefaultListenMapper().Listen(channel).
func Listen(channel string) (*Listener, error) {
	return DefaultListenMapper().Listen(channel)
}

func Notify(channel, payload string) {
	query := "SELECT pg_notify($1, $2)"
	_, err := Exec(query, channel, payload)
	if err != nil {
		log.Print("Could not notify: ", err)
	}
}
