package errors

import (
	"net/http"

	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

const (
	NotFoundMessage  = "Requested endpoint not found."
	ServerErrMessage = "An unexpected error occurred."
)

func NotFound(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	messages.AddErrorMessage(NotFoundMessage)
	return web_responders.Respond(ctx, http.StatusNotFound, messages, NotFoundMessage, settings.FullLinks)
}

func ServerError(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	messages.AddErrorMessage(ServerErrMessage)
	return web_responders.Respond(ctx, http.StatusInternalServerError, messages, ServerErrMessage, settings.FullLinks)
}
