package venues

import (
	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/models/media"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type VenuesController struct {
	base.UnsupportedMethodController
}

func (controller *VenuesController) Path() string {
	return settings.SiteMap["venues"]
}

func (controller *VenuesController) Create(ctx context.Context) error {
	return base.Create(ctx, media.NewVenue(), web_responders.NewMessageMap())
}

func (controller *VenuesController) Read(id string, ctx context.Context) error {
	return base.Read(ctx, media.NewVenue(), id, web_responders.NewMessageMap())
}

func (controller *VenuesController) ReadMany(ctx context.Context) error {
	return base.ReadMany(ctx, media.NewVenue(), "venues", web_responders.NewMessageMap())
}

func (controller *VenuesController) Update(id string, ctx context.Context) error {
	return base.Update(ctx, media.NewVenue(), id, web_responders.NewMessageMap())
}
