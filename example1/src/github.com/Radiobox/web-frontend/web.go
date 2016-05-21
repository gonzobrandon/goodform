package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/Radiobox/web-frontend/controllers"
	"github.com/Radiobox/web-frontend/models"
	rbcodecs "github.com/Radiobox/web_responders/codecs"
	"github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/responders"
)

// AddCodecs is a workaround to default to our codec, but still
// support all other codecs in the stretchr/codecs package.
func AddCodecs() {
	goweb.CodecService = new(services.WebCodecService)
	rbcodecs.AddCodecs()
	for _, defCodec := range services.DefaultCodecs {
		goweb.CodecService.AddCodec(defCodec)
	}
	goweb.API = responders.NewGowebAPIResponder(goweb.CodecService, goweb.Respond)
}

// main initializes and starts the web server.
func main() {
	// PARALLELISM, BITCHES
	runtime.GOMAXPROCS(runtime.NumCPU())

	models.MapModels()
	controllers.MapApi()
	AddCodecs()
	MapStaticFiles()

	log.Print("Mapped URLs, creating server...")
	address := ":" + os.Getenv("PORT")
	server := &http.Server{
		Addr:           address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   48 * time.Hour,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("Created server, listening for requests.")
	listener, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		log.Printf("Could not listen for tcp connections on address %s: %s", address, listenErr.Error())
		panic(listenErr)
	}
	server.Serve(listener)
}
