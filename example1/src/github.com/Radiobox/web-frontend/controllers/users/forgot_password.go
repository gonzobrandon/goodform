package users

import (
	"net/http"

	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/datastore/oauth2"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/users"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type ForgotPasswordController struct {
	base.UnsupportedMethodController
}

func (controller *ForgotPasswordController) Path() string {
	return settings.SiteMap["forgot-password"]
}

func (controller *ForgotPasswordController) Create(ctx context.Context) error {
	var user *users.Account
	messages := web_responders.NewMessageMap()
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		messages.AddErrorMessage("Could not parse input: " + err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
	}
	emailOrUsername := params.Get("email_or_username").Str()
	user, err = users.GetByUsername(emailOrUsername)
	if user == nil {
		user, err = users.GetByEmail(emailOrUsername)
	}
	if err != nil {
		messages.AddErrorMessage("Could not load user: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	if user == nil {
		messages.AddErrorMessage("No user found")
		return web_responders.Respond(ctx, http.StatusNotFound, messages, "No user found", settings.FullLinks)
	}
	client, err := oauth2.DefaultOauth2Storage().GetClient(params.Get("client_id").Str())
	if err != nil {
		messages.AddErrorMessage("Could not load client: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}

	transaction, err := datastore.Begin()
	if err != nil {
		messages.AddErrorMessage("Could not open database transaction: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	reset := &users.PasswordReset{users.Token{Id: new(string), User: user}}
	if err = transaction.Insert(reset); err != nil {
		transaction.Rollback()
		messages.AddErrorMessage("Could not create reset token: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	if err = transaction.Insert(reset.CreateAccess(client.(*auth.Client))); err != nil {
		transaction.Rollback()
		messages.AddErrorMessage("Could not create access token: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	if err = reset.SendResetEmail(); err != nil {
		transaction.Rollback()
		messages.AddErrorMessage("Could not send password reset email: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	transaction.Commit()
	return web_responders.Respond(ctx, http.StatusOK, messages, "Password reset email sent", settings.FullLinks)
}
