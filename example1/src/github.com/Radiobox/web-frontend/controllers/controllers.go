package controllers

import (
	"html/template"

	"github.com/Radiobox/web-frontend/controllers/errors"
	"github.com/Radiobox/web-frontend/controllers/images"
	"github.com/Radiobox/web-frontend/controllers/logs"
	"github.com/Radiobox/web-frontend/controllers/media"
	"github.com/Radiobox/web-frontend/controllers/slugs"
	"github.com/Radiobox/web-frontend/controllers/users"
	"github.com/Radiobox/web-frontend/controllers/util"
	"github.com/Radiobox/web-frontend/controllers/venues"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
)

var controllers = []interface{}{
	new(users.AccountController),
	new(users.ProfileController),
	new(util.StreamNameValidationController),
	new(users.EmailVerificationController),
	new(media.MediaController),
	new(media.TrackController),
	new(media.AlbumController),
	new(media.LiveEventsController),
	new(venues.VenuesController),
	new(ArtistController),
	new(slugs.SlugController),
	new(images.ImageController),
	new(logs.LogController),
	new(users.ForgotPasswordController),
	new(users.BetaSignupController),
}

func DocumentationRoot(ctx context.Context) error {
	docsTemplate, err := template.ParseFiles("./docs/documentation_template.html")
	if err != nil {
		return err
	}
	context := map[string]interface{}{
		"SiteMap": settings.SiteMap,
	}
	err = docsTemplate.Execute(ctx.HttpResponseWriter(), context)
	if err != nil {
		return err
	}
	return nil
}

func MapApi() {
	goweb.Map("/api", DocumentationRoot)
	// Controllers
	for _, controller := range controllers {
		goweb.MapController(controller)
	}

	// Oauth2 paths
	goweb.Map("GET", "/api/authorize", Authorize)
	goweb.Map("POST", "/api/authorize", Authorize)
	goweb.Map("OPTIONS", "/api/authorize", DisplayLoginOptions)

	goweb.Map("POST", "/api/token", Token)
	goweb.Map("OPTIONS", "/api/token", DisplayLoginOptions)

	goweb.MapController(new(Oauth2TestController))

	// All remaining paths get mapped to a 404 controller
	goweb.Map("/api/***", errors.NotFound)
}
