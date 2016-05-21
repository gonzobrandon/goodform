package users

import (
	"net/http"

	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/users"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type EmailVerificationController struct {
	base.UnsupportedMethodController
}

func (evc *EmailVerificationController) Path() string {
	return settings.SiteMap["email-verification"]
}

func (evc *EmailVerificationController) Update(id string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	result, err := datastore.Get(new(users.EmailVerification), id)
	if err != nil {
		messages.AddErrorMessage("Cannot load email verification: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	} else if result == nil {
		message := "Invalid email verification id"
		messages.AddErrorMessage(message)
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, message, settings.FullLinks)
	}
	ev := result.(*users.EmailVerification)
	if err := ev.Verify(); err != nil {
		messages.AddErrorMessage("Cannot verify email: " + err.Error())
	}
	result, err = datastore.Get(new(users.Account), ev.User.Id)
	if err != nil {
		messages.AddErrorMessage("Cannot load user: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	} else if result == nil {
		messages.AddErrorMessage("User not found")
		return web_responders.Respond(ctx, http.StatusNotFound, messages, "User not found", settings.FullLinks)
	}
	datastore.Delete(ev)
	return web_responders.Respond(ctx, http.StatusOK, messages, result.(*users.Account), settings.FullLinks)
}
