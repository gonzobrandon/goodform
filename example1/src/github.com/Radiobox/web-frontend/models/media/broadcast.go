package media

import (
	"github.com/Radiobox/web-frontend/models/base"
)

type Broadcast struct {
	base.CreatedUpdatedModel

	Id *string `db:"event_live_provisioned_broadcast_id" request:"-"`

	Type *string `request:",optional"`

	IsVideo bool `db:"is_video" request:",optional"`

	// Should be merged into new type
	FirstUrl     *string `db:"broadcast_url_1" request:",optional"`
	FirstUrlPort *int64  `db:"broadcast_url_1_port" request:",optional"`

	SecondUrl     *string `db:"broadcast_url_2" request:",optional"`
	SecondUrlPort *int64  `db:"broadcast_url_2_port" request:",optional"`

	StreamName     *string `db:"broadcast_stream_name" request:",optional"`
	ProviderStream *string `db:"provider_stream_id" request:",optional"`

	Username *string `db:"broadcast_username" request:",optional"`
	Password *string `db:"broadcast_password" request:",optional"`

	SuggestedParams *base.JsonMap `db:"encode_suggested_params" request:",optional"`
	Params          *base.JsonMap `db:"encode_accepted_params" request:",optional"`

	HDS          *string `db:"client_hds" request:",optional"`
	HLS          *string `db:"client_hls" request:",optional"`
	HDFlash      *string `db:"client_hdflash1" request:",optional"`
	ShoutcastUrl *string `db:"client_shoutcast_url" request:",optional"`

	ReservedExpiration *base.DbTime             `db:"reserved_until" request:",optional"`
	InProgress         bool                     `db:"is_in_progress" request:",optional"`
	Concluded          bool                     `db:"is_concluded" request:",optional"`
	Available          *base.DefaultTrueBoolean `db:"is_available" request:",optional"`

	MaxListeners int64 `db:"listeners_max" request:",optional"`
	Listeners    int64 `db:"listeners_now" request:",optional"`
}

func NewBroadcast() *Broadcast {
	return new(Broadcast)
}

func (broadcast *Broadcast) ToDb() interface{} {
	return broadcast.Id
}
