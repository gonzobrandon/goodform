package users

import (
	"errors"
	"net/http"

	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/datastore"
	weberrors "github.com/Radiobox/web-frontend/errors"
	"github.com/Radiobox/web-frontend/models/users"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
)

const PasswordResetKey = "reset_token"

var tokenInvalidErr = &weberrors.HttpError{
	Status:  http.StatusBadRequest,
	Message: "Password reset token is not valid",
}

func LoadPasswordReset(params objx.Map) (*users.PasswordReset, *weberrors.HttpError) {
	if reset_token := params.Get(PasswordResetKey).Str(); reset_token != "" {
		result, err := datastore.Get(new(users.PasswordReset), reset_token)
		if err != nil {
			return nil, &weberrors.HttpError{
				Status:  http.StatusInternalServerError,
				Message: "Could not load password reset token",
			}
		}
		if result == nil {
			return nil, tokenInvalidErr
		}
		reset := result.(*users.PasswordReset)
		access, err := reset.Access()
		if err != nil {
			return nil, &weberrors.HttpError{
				Status:  http.StatusInternalServerError,
				Message: "Could not load access token",
			}
		}
		if access == nil {
			return nil, tokenInvalidErr
		}
		return reset, nil
	}
	return nil, nil
}

type AccountController struct {
	base.UnsupportedMethodController
}

func (controller *AccountController) Path() string {
	return settings.SiteMap["user-accounts"]
}

func (controller *AccountController) Options(ctx context.Context) error {
	return web_responders.Respond(ctx, http.StatusOK, nil, map[string]string{
		"description": "The user endpoint is for interactions with the logged " +
			"in user's account information.  Changing passwords, confirming " +
			"email addresses, and creating new users can be done here.",
		"POST " + controller.Path(): "Create a new user.  Expected parameters: " +
			"username, email, password.",
		"GET " + controller.Path() + "/id": "Retrieve a user's " +
			"account details.  You must pass an Authorization header with a " +
			"valid access token.  The username and email of the user will be " +
			"returned.",
		"PATCH " + controller.Path() + "/id": "Update a user's account details.  " +
			"You must pass an Authorization header with a valid access token.  " +
			"Allowed parameters: username, email, password.",
	})
}

func (controller *AccountController) Create(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		messages.AddErrorMessage(err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err)
	}
	if fbUser := params.Get("facebook_user").Str(); fbUser != "" {
		token := params.Get("facebook_token").Str()
		if token == "" {
			err := errors.New("Cannot create facebook_user without facebook_token")
			messages.AddErrorMessage(err.Error())
			return web_responders.Respond(ctx, http.StatusBadRequest, messages, err)
		}
		delete(params, "facebook_token")
		tokenUserId, err := users.GetFacebookIdFromToken(token)
		if err != nil {
			status := http.StatusInternalServerError
			if _, ok := err.(users.LoginError); ok {
				status = http.StatusUnauthorized
			}
			messages.AddErrorMessage("Could not get facebook ID from token: " + err.Error())
			return web_responders.Respond(ctx, status, messages, err)
		}
		if tokenUserId != fbUser {
			err := errors.New("Token does not match facebook user")
			messages.AddErrorMessage(err.Error())
			return web_responders.Respond(ctx, http.StatusUnauthorized, messages, err)
		}
	}
	return base.Create(ctx, users.New(), web_responders.NewMessageMap())
}

func (controller *AccountController) Update(id string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Update(ctx, new(users.Account), id, messages)
}

// UpdateMany is here just to handle password resets.
func (controller *AccountController) UpdateMany(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		messages.AddErrorMessage("Could not read input parameters: " + err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
	}
	reset, authErr := LoadPasswordReset(params)
	password := params.Get("password").Str()
	if reset == nil {
		messages.SetInputMessage(PasswordResetKey, "No reset key found")
	}
	if len(password) < 2 {
		messages.SetInputMessage("password", "Password is too short")
	}
	if authErr != nil {
		messages.AddErrorMessage("Could not load password reset token: " + authErr.Message)
		messages.SetInputMessage(PasswordResetKey, authErr.Message)
		return web_responders.Respond(ctx, authErr.Status, messages, authErr, settings.FullLinks)
	}
	if reset == nil {
		messages.AddErrorMessage("No password reset found matching the input token")
		return web_responders.Respond(ctx, http.StatusNotFound, messages, "No password reset found matching the input token", settings.FullLinks)
	}
	if len(password) < 2 {
		messages.AddErrorMessage("Invalid password")
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, "Invalid password", settings.FullLinks)
	}
	user := reset.GetUser()
	user.SetPassword(password)
	transaction, err := datastore.Begin()
	if err != nil {
		messages.AddErrorMessage("Could not start database transaction: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	if _, err = transaction.Update(user); err != nil {
		transaction.Rollback()
		messages.AddErrorMessage("Could not update user password: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	access, _ := reset.Access()
	if _, err = transaction.Delete(access, reset); err != nil {
		transaction.Rollback()
		messages.AddErrorMessage("Could not revoke password reset: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	transaction.Commit()
	return web_responders.Respond(ctx, http.StatusOK, messages, user, settings.FullLinks)
}

func (controller *AccountController) Read(id string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Read(ctx, new(users.Account), id, messages)
}
