// +build !heroku

package main

import (
	"log"

	"github.com/Radiobox/web-frontend/settings"
	"github.com/stretchr/goweb"
)

func init() {
	log.Print("Starting development server...\n")
}

func MapStaticFiles() {
	goweb.MapStatic("/js", settings.ProjectPath+"/public/src/js")
	goweb.MapStatic("/css", settings.ProjectPath+"/public/src/css")
	goweb.MapStatic("/img", settings.ProjectPath+"/public/src/img")
	goweb.MapStatic("/fonts", settings.ProjectPath+"/public/src/fonts")
	goweb.MapStatic("/partials", settings.ProjectPath+"/public/src/partials")
	goweb.MapStatic("/media", settings.ProjectPath+"/public/src/media")
	goweb.MapStaticFile("crossdomain.xml", settings.ProjectPath+"/public/src/crossdomain.xml")
	goweb.MapStaticFile("***", settings.ProjectPath+"/public/src/index.html")

}
