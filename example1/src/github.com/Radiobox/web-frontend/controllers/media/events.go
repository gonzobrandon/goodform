package media

import (
	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/models/media"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type LiveEventsController struct {
	base.UnsupportedMethodController
}

func (controller *LiveEventsController) Path() string {
	return settings.SiteMap["events"]
}

func (controller *LiveEventsController) Create(ctx context.Context) error {
	return base.Create(ctx, media.NewLiveEvent(), web_responders.NewMessageMap())
}

func (controller *LiveEventsController) Read(idString string, ctx context.Context) error {
	return base.Read(ctx, media.NewLiveEvent(), idString, web_responders.NewMessageMap())
}

// TODO: Security
func (controller *LiveEventsController) Update(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Update(ctx, media.NewLiveEvent(), idString, messages)
}

// ReadMany responds to requests to retrieve multiple results.
func (controller *LiveEventsController) ReadMany(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.ReadMany(ctx, new(media.LiveEventCollection), "events_live", messages)
}
