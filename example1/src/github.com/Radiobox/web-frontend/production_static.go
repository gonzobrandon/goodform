// +build heroku

package main

import (
	"log"

	"github.com/stretchr/goweb"
)

func init() {
	log.Print("Starting production server...\n")
}

func MapStaticFiles() {
	goweb.MapStatic("/js", "./public/dist/js")
	goweb.MapStatic("/css", "./public/dist/css")
	goweb.MapStatic("/img", "./public/dist/img")
	goweb.MapStatic("/fonts", "./public/dist/fonts")
	goweb.MapStatic("/partials", "./public/dist/partials")
	goweb.MapStatic("/media", "./public/dist/media")
	goweb.MapStaticFile("crossdomain.xml", "./public/dist/crossdomain.xml")
	goweb.MapStaticFile("***", "./public/dist/index.html")
}
