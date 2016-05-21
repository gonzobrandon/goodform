package users

import (
	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/models/users"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type BetaSignupController struct {
	base.UnsupportedMethodController
}

func (controller *BetaSignupController) Path() string {
	return settings.SiteMap["beta-signup"]
}

func (controller *BetaSignupController) Create(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Create(ctx, new(users.BetaSignup), messages)
}
