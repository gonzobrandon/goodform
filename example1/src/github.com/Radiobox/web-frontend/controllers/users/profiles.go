package users

import (
	"net/http"

	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/models/users"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type ProfileController struct {
	base.UnsupportedMethodController
}

func (controller *ProfileController) Path() string {
	return settings.SiteMap["user-profiles"]
}

func (controller *ProfileController) Options(ctx context.Context) error {
	return web_responders.Respond(ctx, http.StatusOK, nil, map[string]string{
		"description": "The users endpoint is for retrieving or updating " +
			"public information about a user.  User creation is handled at " +
			"the user endpoint, since the users endpoint is only supposed to " +
			"deal with public details, and therefor can't set email address, " +
			"username, or password details.",
		"GET " + controller.Path():         "Retrieve a list of users.",
		"GET " + controller.Path() + "/id": "Retrieve the details about a user.",
		"PATCH " + controller.Path() + "/id": "Update a user's details.  This " +
			"request requires an Authorization header with a valid access " +
			"token.  Allowed parameters: first_name, last_name, sex, birth_date.",
	})
}

func (controller *ProfileController) Update(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Update(ctx, new(users.Profile), idString, messages)
}

func (controller *ProfileController) Read(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Read(ctx, new(users.Profile), idString, messages)
}

func (controller *ProfileController) ReadMany(ctx context.Context) error {
	return base.ReadMany(ctx, new(users.Profile), "users", web_responders.NewMessageMap())
}
