package logs

import (
	"net/http"

	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/models/logs"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type LogController struct {
	base.UnsupportedMethodController
}

func (controller *LogController) Path() string {
	return settings.SiteMap["logs"]
}

func (controller *LogController) Create(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		messages.AddErrorMessage("Failed to parse parameters: " + err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err.Error(), settings.FullLinks)
	}

	go logs.WriteLog(params)
	return web_responders.Respond(ctx, http.StatusOK, messages, "Log write in progress.", settings.FullLinks)
}
