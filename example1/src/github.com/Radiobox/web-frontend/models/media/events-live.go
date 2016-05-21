package media

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/base"
	"github.com/Radiobox/web-frontend/models/images"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/coopernurse/gorp"
	"github.com/nelsam/gorp_queries/extensions"
	query_interfaces "github.com/nelsam/gorp_queries/interfaces"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
)

type priority int32

func (p priority) DefaultValue() interface{} {
	return 1000
}

type EventRoute struct {
	base.CreatedUpdatedModel
	Id           *string    `db:"event_live_route_id" request:"-"`
	Event        *LiveEvent `db:"event_live_id" response:"event_live"`
	Broadcast    *Broadcast `db:"event_live_provisioned_broadcast_id" response:"event_live_provisioned_broadcast"`
	Priority     priority   `request:",optional"`
	Primary      bool       `db:"is_primary" request:",optional"`
	MaxListeners int64      `db:"listeners_max" request:",optional"`
	Listeners    int64      `db:"listeners_now" request:",optional"`
}

// A Live event is an event that is currently live, was streamed live,
// or will be streamed live.
type LiveEvent struct {
	base.CreatedUpdatedModel

	Id string `db:"event_live_id" request:"-"`

	InProgress *bool `db:"is_in_progress" request:",optional"`
	Concluded  *bool `db:"is_concluded" request:",optional"`

	OffScheduleEnd   *int64 `db:"off_schedule_end_seconds" request:",optional"`
	OffScheduleStart *int64 `db:"off_schedule_start_seconds" request:",optional"`

	// Title is the title of the event.
	Title base.NonEmptyString

	// TitleFrom is the sub-title, usually "Live From Somewhere"
	TitleFrom *string `db:"title_from" response:"title_from" request:",optional"`

	// Caption is another small sub-title, for a blurb in the live
	// player.
	Caption *string `request:",optional"`

	// Images
	PicSquare *images.Image `db:"pic_square_id" response:"pic_square" request:",optional"`

	// Artist is the primary artist performing during the event.
	Artist *Artist `db:"artist_id" response:"artist" request:",optional"`

	// Collaborators is a json list of other artists also performing
	// during the event.
	Collaborators *base.JsonArray `db:"artist_collaborators_id" response:"artist_collaborators" request:",optional"`

	// Messages is a list of messages for the event.
	Messages *base.JsonArray `request:",optional"`

	// Address is the address of the event (which may be different
	// than the Venue's location).
	Address base.Address `db:"location_address" request:",optional"`

	// Venue is the Venue that is hosting this event.
	Venue *Venue `db:"venue_id" response:"venue" request:",optional"`

	// Active determines whether or not the event is currently live.
	Active base.FakeBoolean `db:"is_active" request:"-"`

	// Archiving determines whether or not the event is currently
	// ended and being archived as a track.
	Archiving *bool `request:",optional"`

	// ArchiveTrack is the track containing the live recording of this
	// event.
	ArchiveTrack *Track `db:"archive_track_id" response:"archive_track" request:",optional"`

	// ScheduledStart and ScheduledEnd represent when this event is
	// intended to start and end.
	ScheduledStart *base.DbTime `db:"scheduled_start" response:"scheduled_start"`
	ScheduledEnd   *base.DbTime `db:"scheduled_end" response:"scheduled_end"`

	// ActualStart and ActualEnd represent when this event actually
	// ran.
	ActualStart *base.DbTime `db:"actual_start" response:"actual_start" request:",optional"`
	ActualEnd   *base.DbTime `db:"actual_end" response:"actual_end" request:",optional"`

	// Listeners stores the current count of users listening to this event.
	Listeners int `db:"-" response:"-"`

	// MaxListeners stores the maximum number of users that were
	// listening to this event at one time.
	MaxListeners *int64 `db:"listeners_max" response:"max_listeners" request:",optional"`

	// Broadcasts contains the connected broadcast objects, if any
	// exist.
	Broadcasts []*Broadcast `db:"-" response:"allocated_broadcasts" request:"-"`
}

func NewLiveEvent() *LiveEvent {
	return new(LiveEvent)
}

func (event *LiveEvent) ExampleOutput() map[string]interface{} {
	return map[string]interface{}{
		"active":        false,
		"actual_end":    nil,
		"actual_start":  nil,
		"archive_track": nil,
		"archiving":     false,
		"artist": map[string]interface{}{
			"id":       1,
			"link":     "/api/artists/1",
			"username": "",
		},
		"caption":              "",
		"artist_collaborators": []interface{}{},
		"created_at":           "2014-03-05 20:54:17.9677599 +0000 UTC",
		"id":                   "ad0a3f2a-6ffb-4b10-a086-41118b2f1f66",
		"location": map[string]interface{}{
			"address1":       "123 Street",
			"address2":       "Unit D",
			"city":           "Denver",
			"country":        "USA",
			"state_province": "Colorado",
		},
		"max_listeners":   70,
		"messages":        []interface{}{},
		"scheduled_end":   "2014-03-04 16:00:00 -0700 -0700",
		"scheduled_start": "2014-03-04 15:00:00 -0700 -0700",
		"title":           "Test Event (45mw)",
		"title_from":      "Live from Your House",
		"updated_at":      "2014-03-05 20:54:17.9677599 +0000 UTC",
		"venue":           nil,
	}
}

func (event *LiveEvent) Valid() bool {
	return event != nil && event.Id != ""
}

func (event *LiveEvent) Connect(broadcast *Broadcast, exec gorp.SqlExecutor) error {
	if broadcast.InProgress || broadcast.Concluded || !bool(*broadcast.Available) {
		return errors.New("Cannot connect a broadcast that is in progress, concluded, or not available")
	}
	broadcast.InProgress = true
	*broadcast.Available = base.DefaultTrueBoolean(false)
	if _, err := exec.Update(broadcast); err != nil {
		return err
	}
	route := &EventRoute{
		Event:        event,
		Broadcast:    broadcast,
		Priority:     0,
		Primary:      len(event.GetBroadcasts()) == 0,
		MaxListeners: 100,
		Listeners:    0,
	}
	return exec.Insert(route)
}

func (event *LiveEvent) Start(exec gorp.SqlExecutor) error {
	broadRef := new(Broadcast)
	err := datastore.Query(broadRef).
		Where().
		NotEqual(&broadRef.InProgress, true).
		NotEqual(&broadRef.Concluded, true).
		Equal(&broadRef.Available, true).
		Limit(1).
		SelectToTarget(&event.Broadcasts)
	if err != nil {
		return err
	}
	if len(event.Broadcasts) == 0 {
		return errors.New("No available allocated broadcasts")
	}
	return event.Connect(event.Broadcasts[0], exec)
}

func (event *LiveEvent) End(exec gorp.SqlExecutor) error {
	broadRef := new(Broadcast)
	routeRef := new(EventRoute)
	_, err := datastore.Query(broadRef).
		Extend().(extensions.Postgres).
		Assign(&broadRef.InProgress, false).
		Assign(&broadRef.Concluded, true).
		Join(routeRef).On().
		Equal(&broadRef.Id, &routeRef.Broadcast).
		Where().
		Equal(&routeRef.Event, event.Id).
		Update()
	if err != nil {
		return err
	}
	_, err = datastore.Query(routeRef).
		Where().
		Equal(&routeRef.Event, event.Id).
		Delete()
	if err != nil {
		return err
	}
	event.Broadcasts = nil
	return nil
}

func (event *LiveEvent) GetBroadcasts() []*Broadcast {
	if event.Broadcasts == nil {
		broadRef := new(Broadcast)
		routeRef := new(EventRoute)
		err := datastore.Query(broadRef).
			Join(routeRef).On().
			Equal(&routeRef.Broadcast, &broadRef.Id).
			Where().
			Equal(&routeRef.Event, event.Id).
			SelectToTarget(&event.Broadcasts)
		if err != nil {
			return nil
		}
	}
	return event.Broadcasts
}

func (event *LiveEvent) GetArtist() *Artist {
	if event.Artist.Valid() && event.Artist.Username == "" {
		result, err := datastore.Get(event.Artist, event.Artist.Id)
		if err != nil {
			log.Print("Error getting event artist: ", err)
			return nil
		}
		event.Artist = result.(*Artist)
	}
	return event.Artist
}

func (event *LiveEvent) GetVenue() *Venue {
	if event.Venue.Valid() && event.Venue.Name == "" {
		result, err := datastore.Get(event.Venue, event.Venue.Id)
		if err != nil {
			log.Print("Error getting event venue: ", err)
			return nil
		}
		event.Venue = result.(*Venue)
	}
	return event.Venue
}

func (event *LiveEvent) Authorize(access *auth.AccessData, requestType auth.RequestType) error {
	switch requestType {
	case auth.REQUEST_READ_ONE, auth.REQUEST_READ_MANY:
	case auth.REQUEST_CREATE, auth.REQUEST_UPDATE:
		if access == nil || access.UserId <= 0 {
			return errors.New("Only logged in users can create events.")
		}
		if !event.GetArtist().IsAdmin(access.UserId) {
			return errors.New("You can only create events for artists you are an admin of.")
		}
	default:
		return errors.New("Invalid request type")
	}
	return nil
}

func (event *LiveEvent) LazyLoad(options objx.Map) {
	if event.Title == "" {
		result, err := datastore.Get(event, event.Id)
		if err != nil {
			panic(err)
		}
		fullEvent := result.(*LiveEvent)
		*event = *fullEvent
	}
	event.GetBroadcasts()
	event.GetArtist()
	event.GetVenue()
}

func (event *LiveEvent) Running() bool {
	return event.InProgress != nil && *event.InProgress &&
		!(event.Concluded != nil && *event.Concluded)
}

func (event *LiveEvent) PreUpdate(exec gorp.SqlExecutor) error {
	if event.ActualStart == nil && event.Running() && len(event.GetBroadcasts()) == 0 {
		event.ActualStart = base.Now()
	} else if event.ActualEnd == nil && !event.Running() && len(event.GetBroadcasts()) > 0 {
		event.ActualEnd = base.Now()
	}
	return nil
}

func (event *LiveEvent) PostUpdate(exec gorp.SqlExecutor) error {
	if event.Running() && len(event.GetBroadcasts()) == 0 {
		if err := event.Start(exec); err != nil {
			return err
		}
	} else if !event.Running() && len(event.GetBroadcasts()) > 0 {
		if err := event.End(exec); err != nil {
			return err
		}
	}
	return nil
}

func (event *LiveEvent) ToDb() interface{} {
	return event.Id
}

func (event *LiveEvent) DefaultDbValue() interface{} {
	return new(string)
}

func (event *LiveEvent) FromDb(value interface{}) error {
	idPtr := value.(*string)
	event.Id = *idPtr
	return nil
}

func (event *LiveEvent) ResponseValue(options objx.Map) interface{} {
	return map[string]interface{}{
		"id":    event.Id,
		"link":  event.Location(),
		"title": event.Title,
		"schedule": map[string]interface{}{
			"start": event.ScheduledStart.ResponseObject(),
			"end":   event.ScheduledEnd.ResponseObject(),
		},
	}
}

func (event *LiveEvent) RelatedLinks() map[string]string {
	links := make(map[string]string)
	if event.Artist.Valid() {
		links["artist"] = event.Artist.Location()
	}
	// TODO: Link to collaborators
	if event.Venue.Valid() {
		links["venue"] = event.Venue.Location()
	}
	if event.ArchiveTrack.Valid() {
		links["archive_track"] = event.ArchiveTrack.Location()
	}
	return links
}

func (event *LiveEvent) Location() string {
	return settings.UrlFor("events", event.Id)
}

func (event *LiveEvent) Receive(value interface{}) error {
	event.Id = value.(string)
	return nil
}

type LiveEventCollection struct {
	events  []*LiveEvent
	options objx.Map
}

func (events *LiveEventCollection) DefaultPageSize() int {
	return 50
}

func (events *LiveEventCollection) RelatedLinks() map[string]string {
	page := 1
	if events.options.Has("page") {
		pageStr := events.options.Get("page").Str()
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil {
			page = parsedPage
		}
	}
	basePath := settings.SiteMap["events"]
	links := make(map[string]string)
	pageSize := events.DefaultPageSize()
	if sizeStr := events.options.Get("page_size").Str(); sizeStr != "" {
		size, err := strconv.Atoi(sizeStr)
		if err == nil {
			pageSize = size
		}
	}
	if len(events.events) == pageSize {
		// This is probably not the last page.
		nextOptions := events.options.Copy()
		nextOptions.Set("page", page+1)
		nextQueryString, err := nextOptions.URLQuery()
		if err != nil {
			// This should never happen
			panic(err)
		}
		links["next"] = fmt.Sprintf("%s?%s", basePath, nextQueryString)
	}
	if page > 1 {
		prevOptions := events.options.Copy()
		prevOptions.Set("page", page-1)
		prevQueryString, err := prevOptions.URLQuery()
		if err != nil {
			// Again, this should never happen
			panic(err)
		}
		links["prev"] = fmt.Sprintf("%s?%s", basePath, prevQueryString)
	}
	return links
}

func (events *LiveEventCollection) Location() string {
	query, err := events.options.URLQuery()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s?%s", settings.SiteMap["events"], query)
}

func (events *LiveEventCollection) ResponseObject() interface{} {
	if events.events == nil {
		events.events = make([]*LiveEvent, 0, 10)
	}
	return &events.events
}

// Where generates a where clause (including an ORDER BY if
// appropriate) for a live event.
func (events *LiveEventCollection) Query(ctx context.Context) (query_interfaces.SelectQuery, error) {
	params, err := objx.FromURLQuery(ctx.HttpRequest().URL.RawQuery)
	if err != nil {
		panic(err)
	}
	return events.QueryFromParams(params)
}

func (events *LiveEventCollection) QueryFromParams(params objx.Map) (query_interfaces.SelectQuery, error) {
	events.options = params.Copy()
	queryType := strings.ToLower(events.options.Get("type").Str())
	ref := new(LiveEvent)
	orderBy := &ref.ScheduledStart
	orderDirection := "DESC"
	query := datastore.Query(ref).Where()
	switch queryType {
	case "live-now":
		query.Equal(&ref.InProgress, true)
		orderBy = &ref.ActualStart
	case "range":
		var (
			start = new(base.DbTime)
			end   = new(base.DbTime)
		)
		startErr := start.Receive(events.options.Get("start").Str())
		endErr := end.Receive(events.options.Get("end").Str())
		if startErr != nil || endErr != nil {
			message := "Error: Parsing range failed, 'end' and 'start' values " +
				"could not be parsed as RFC3339 timestamps."
			return nil, errors.New(message)
		}
		query.Greater(&ref.ScheduledStart, start).
			Less(&ref.ScheduledStart, end)
		orderDirection = events.options.Get("order").Str()
		switch strings.ToLower(orderDirection) {
		case "desc", "asc":
		default:
			orderDirection = "DESC"
		}
		orderBy = &ref.ScheduledStart
	case "upcoming":
		query.Greater(&ref.ScheduledStart, time.Now())
		orderBy = &ref.ScheduledStart
		orderDirection = "ASC"
	case "upcoming_not_ended":
		query.Greater(&ref.ScheduledEnd, time.Now())
		orderBy = &ref.ScheduledEnd
		orderDirection = "ASC"
	case "recent":
		query.NotNull(&ref.ActualEnd).
			Less(&ref.ActualEnd, time.Now())
		orderBy = &ref.ActualStart
	case "":
	default:
		message := "Error: Live event type " + events.options.Get("type").Str() + " not understood."
		return nil, errors.New(message)
	}

	if events.options.Has("artist") {
		artistId := events.options.Get("artist").Int64(-1)
		if artistId == -1 {
			var err error
			artistId, err = strconv.ParseInt(events.options.Get("artist").Str(), 10, 64)
			if err != nil {
				return nil, errors.New("Could not parse artist ID")
			}
		}
		query.Equal(&ref.Artist, artistId)
	}

	return query.OrderBy(orderBy, orderDirection), nil
}
