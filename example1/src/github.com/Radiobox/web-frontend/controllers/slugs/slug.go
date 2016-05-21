package slugs

import (
	"net/http"

	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/models/slugs"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type SlugController struct {
	base.UnsupportedMethodController
}

func (controller *SlugController) Path() string {
	return settings.SiteMap["slugs"]
}

func (controller *SlugController) Read(slug string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	slugObj, err := slugs.GetSlug(slug)
	if err != nil {
		messages.AddErrorMessage(err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	} else if slugObj == nil {
		messages.AddErrorMessage("Not found")
		return web_responders.Respond(ctx, http.StatusNotFound, messages, "404: Not found", settings.FullLinks)
	}
	return web_responders.Respond(ctx, http.StatusOK, messages, slugObj, settings.FullLinks)
}
