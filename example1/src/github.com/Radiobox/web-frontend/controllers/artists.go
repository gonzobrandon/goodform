package controllers

import (
	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/models/media"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type ArtistController struct {
	base.UnsupportedMethodController
}

func (controller *ArtistController) Path() string {
	return settings.SiteMap["artists"]
}

func (controller *ArtistController) Create(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	artist := media.NewArtist()
	return base.Create(ctx, artist, messages)
}

func (controller *ArtistController) Update(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Update(ctx, media.NewArtist(), idString, messages)
}

func (controller *ArtistController) Read(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Read(ctx, media.NewArtist(), idString, messages)
}

func (controller *ArtistController) ReadMany(ctx context.Context) error {
	return base.ReadMany(ctx, new(media.ArtistCollection), "artists", web_responders.NewMessageMap())
}
