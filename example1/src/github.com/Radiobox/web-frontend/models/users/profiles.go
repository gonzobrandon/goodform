package users

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/base"
	"github.com/Radiobox/web-frontend/models/images"
	"github.com/Radiobox/web-frontend/models/slugs"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/coopernurse/gorp"
	query_interfaces "github.com/nelsam/gorp_queries/interfaces"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
)

// Profile is for public user information - anything that is intended
// to be presented to other people as a means of social interaction.
type Profile struct {
	base.CreatedUpdatedModel

	Id                 int64         `db:"user_id" response:"-"`
	DisplayName        string        `db:"username" response:"display_name" request:"-"`
	Slug               *slugs.Slug   `db:"-"`
	FirstName          *string       `db:"first_name" response:"first_name"`
	LastName           *string       `db:"last_name" response:"last_name"`
	BirthDate          *base.DbTime  `db:"birthday_date" response:"birthday_date"`
	Description        *string       `db:"profile_blurb"`
	CurrentAddress     *base.Address `db:"current_address" request:",optional"`
	CurrentLocation    *base.JsonMap `db:"current_location" request:",optional"`
	LastLocationUpdate *base.DbTime  `db:"current_location_updated" request:",optional"`

	// Should probably be moved to a phone type
	Phone       *int64  `db:"contact_phone" request:",optional"`
	CountryCode *int64  `db:"contact_phone_country_code" request:",optional"`
	Extension   *string `db:"contact_phone_extension" request:",optional"`

	EmailProxy         *string         `db:"email_proxy" request:",optional"`
	EmailPendingUpdate *string         `db:"email_pending_update" request:",optional"`
	Interests          *string         `request:",optional"`
	Keywords           *base.JsonArray `request:",optional"`
	Locale             *string         `request:",optional"`

	// Images
	PicCover  *images.UserImage `db:"pic_cover_id" response:"pic_cover" request:",optional"`
	PicSquare *images.UserImage `db:"pic_square_id" response:"pic_square" request:",optional"`

	Politics     *string `db:"political" request:",optional"`
	Relationship *string `db:"relationship_status" request:",optional"`
	Sex          *string

	RbSubscribers *int64 `db:"rb_subscriber_count" request:"-"`

	// Facebook
	FbUser        *string      `db:"facebook_user" request:",optional"`
	FbSubscribers *int64       `db:"facebook_subscriber_count" request:",optional"`
	FbLastPull    *base.DbTime `db:"facebook_last_data_pull" request:",optional"`

	// LinkedIn
	LiUser     *string      `db:"linkedin_user" request:",optional"`
	LiLastPull *base.DbTime `db:"linkedin_last_data_pull" request:",optional"`

	// Twitter
	TwSubscribers *int64       `db:"twitter_subscriber_count" request:",optional"`
	TwLastPull    *base.DbTime `db:"twitter_last_data_pull" request:",optional"`

	Timezone  *string `request:",optional"`
	Active    *bool   `db:"is_active" request:"-"`
	Verified  *bool   `db:"is_verified" request:"-"`
	WallCount *string `db:"wall_count" request:",optional"`
	Work      *string `db:"work" request:",optional"`
}

func (profile *Profile) Valid() bool {
	return profile != nil && profile.Id > 0
}

func (profile *Profile) Authorize(access *auth.AccessData, requestType auth.RequestType) error {
	switch requestType {
	case auth.REQUEST_READ_ONE, auth.REQUEST_READ_MANY:
	case auth.REQUEST_CREATE:
		return errors.New("Cannot create user profile, try " + settings.SiteMap["user-accounts"])
	case auth.REQUEST_UPDATE:
		if access == nil || access.UserId != profile.Id {
			return errors.New("Only users can update their profile")
		}
	default:
		return errors.New("Invalid request type")
	}
	return nil
}

func (profile *Profile) Identifier() interface{} {
	return profile.Id
}

func (profile *Profile) IdentifierName() string {
	return "user_id"
}

func (profile *Profile) TypeName() string {
	return "user"
}

func (profile *Profile) RelatedLinks() map[string]string {
	links := make(map[string]string)
	if profile.Slug != nil {
		links["slug"] = profile.Slug.Location()
	}
	return links
}

func (profile *Profile) Location() string {
	return settings.UrlFor("user-profiles", profile.Id)
}

func (profile *Profile) GetSlug() *slugs.Slug {
	if profile.Slug == nil && profile.Valid() {
		query := "SELECT * FROM slugs WHERE (target->>'user_id')::int = $1"
		results, err := datastore.Select(new(slugs.Slug), query, profile.Identifier())
		if err != nil || len(results) != 1 {
			return nil
		}
		profile.Slug = results[0].(*slugs.Slug)
		profile.Slug.SetTarget(profile)
	}
	return profile.Slug
}

func (profile *Profile) SetSlug(slug *slugs.Slug) {
	profile.Slug = slug
}

func (profile *Profile) LazyLoad(options objx.Map) {
	profile.GetSlug()
}

// ToDb returns this Profile's Id for storing in the database.
func (profile *Profile) ToDb() interface{} {
	return profile.Id
}

// DefaultDbValue returns the default value for Profile.Id.
func (profile *Profile) DefaultDbValue() interface{} {
	return new(int64)
}

// FromDb stores the Profile.Id value read from the database.
func (profile *Profile) FromDb(id interface{}) error {
	idInt := id.(*int64)
	profile.Id = *idInt
	return nil
}

func (profile *Profile) PostUpdate(exec gorp.SqlExecutor) error {
	if profile.Slug != nil {
		profile.Slug.SetTarget(profile)
		_, err := exec.Update(profile.Slug)
		if err != nil {
			return err
		}
	}
	return nil
}

func (profile *Profile) ResponseValue(options objx.Map) interface{} {
	if profile.Valid() {
		return map[string]interface{}{
			"id":   profile.Id,
			"name": profile.DisplayName,
			"link": settings.UrlFor("user-profiles", profile.Id),
		}
	}
	return nil
}

func (profile *Profile) ResponseObject() interface{} {
	if profile.Valid() {
		return profile
	}
	return (*Profile)(nil)
}

func (profile *Profile) Receive(value interface{}) error {
	var err error
	switch idValue := value.(type) {
	case string:
		profile.Id, err = strconv.ParseInt(idValue, 10, 64)
		if err != nil {
			return err
		}
	case float32:
		profile.Id = int64(idValue)
	case float64:
		profile.Id = int64(idValue)
	case int64:
		profile.Id = idValue
	case int32:
		profile.Id = int64(idValue)
	case int:
		profile.Id = int64(idValue)
	default:
		return errors.New(fmt.Sprintf("Don't know how to handle ID value %v", idValue))
	}
	return nil
}

func (profile *Profile) ValidateInput(value interface{}) error {
	if value == nil {
		return nil
	}
	if err := profile.Receive(value); err != nil {
		return errors.New("Cannot parse input as user ID")
	}
	if existingValue, err := datastore.Get(profile, profile.Id); err != nil {
		return errors.New("Internal error: cannot query for user value: " + err.Error())
	} else if existingValue == nil {
		return errors.New("No user found by that ID")
	}
	return nil
}

func (profile *Profile) Query(ctx context.Context) (query_interfaces.SelectQuery, error) {
	return datastore.Query(profile).(query_interfaces.SelectQuery), nil
}
