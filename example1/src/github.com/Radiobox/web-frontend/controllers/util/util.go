// The controllers/util package contains utility functions and types
// for controllers to use when reading and responding to requests.
package util

import (
	"log"
	"net/http"
	"strings"

	"github.com/Radiobox/web-frontend/datastore/oauth2"
	"github.com/Radiobox/web-frontend/errors"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

func Authorize(ctx context.Context) (*auth.AccessData, error) {
	authHeader := ctx.HttpRequest().Header.Get("Authorization")
	authParts := strings.SplitN(authHeader, " ", 2)
	if authParts[0] != "Bearer" {
		return nil, &errors.HttpError{
			Status:  http.StatusBadRequest,
			Message: "Invalid authorization type; must use Bearer authorization",
		}
	}

	if len(authParts) < 2 {
		return nil, &errors.HttpError{
			Status:  http.StatusBadRequest,
			Message: "Bearer authorization requested but no token present",
		}
	}
	token := authParts[1]
	access, err := oauth2.DefaultOauth2Storage().LoadAccess(token)
	if err != nil {
		log.Print(err)
		return nil, &errors.HttpError{
			Status:  http.StatusInternalServerError,
			Message: "Could not load access",
		}
	} else if access == nil || access.GetClient() == nil {
		return nil, &errors.HttpError{
			Status:  http.StatusUnauthorized,
			Message: "Access not found",
		}
	}

	if access.IsExpired() {
		return nil, &errors.HttpError{
			Status:  http.StatusUnauthorized,
			Message: "Access expired",
		}
	}

	return access.(*auth.AccessData), nil
}

type StreamNameValidationController struct {
}

func (controller *StreamNameValidationController) Path() string {
	return settings.SiteMap["stream_name_validation"]
}

func (controller *StreamNameValidationController) Create(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		messages.AddErrorMessage("Could not parse parameters: " + err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
	}
	name := params.Get("name").Str()
	result := false
	if strings.HasPrefix(name, "test") {
		result = true
	}
	return web_responders.Respond(ctx, http.StatusOK, messages, result, settings.FullLinks)
}
