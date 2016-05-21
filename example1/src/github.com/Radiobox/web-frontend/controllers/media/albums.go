package media

import (
	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/models/media"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type AlbumController struct {
	base.UnsupportedMethodController
}

func (controller *AlbumController) Path() string {
	return settings.SiteMap["albums"]
}

func (controller *AlbumController) Create(ctx context.Context) error {
	return base.Create(ctx, media.NewAlbum(), web_responders.NewMessageMap())
}

func (controller *AlbumController) Read(idString string, ctx context.Context) error {
	return base.Read(ctx, media.NewAlbum(), idString, web_responders.NewMessageMap())
}

// TODO: Security
func (controller *AlbumController) Update(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Update(ctx, media.NewAlbum(), idString, messages)
}

func (controller *AlbumController) ReadMany(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.ReadMany(ctx, new(media.AlbumCollection), "albums", messages)
}
