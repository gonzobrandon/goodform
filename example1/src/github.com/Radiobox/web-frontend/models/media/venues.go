package media

import (
	"errors"
	"strconv"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/base"
	"github.com/Radiobox/web-frontend/models/images"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/stretchr/objx"
)

// A Venue is a location that hosts Artists during live events.
type Venue struct {
	base.CreatedUpdatedModel

	Id int64 `db:"venue_id" request:"-"`

	// Address is the physical address of this Venue.
	Address base.Address `request:",optional"`

	// TimeZone is the time zone that this Venue is in.
	TimeZone *base.Location `request:",optional"`

	// Email is the email address to send to when contacting this
	// Venue.
	Email *string

	// Name is the name of this Venue.
	Name base.NonEmptyString

	// Phone is the phone number to call when contacting this Venue.
	// Should be moved to a separate type.
	Phone       *int64  `db:"contact_phone" request:",optional"`
	CountryCode *int64  `db:"contact_phone_country_code" request:",optional"`
	Extension   *string `db:"contact_phone_extension" request:",optional"`

	// Active determines whether or not this Venue is still around.
	Active base.FakeBoolean `db:"is_active" request:"-"`

	// Verified determines whether or not we have verified that this
	// Venue exists and that the account for this Venue is managed by
	// someone who works there.
	Verified base.FakeBoolean `db:"is_verified" request:"-"`

	// Url is the URL for this Venue's website, if any.
	Url *string `request:",optional"`

	// Blurb is a short marketing description for this Venue.
	Blurb *string `db:"venue_blurb" request:",optional"`

	// Facebook
	FbUser *string `db:"facebook_user" request:",optional"`

	// Should be moved to a separate type
	PicCover  *images.Image `db:"pic_cover_id" response:"pic_cover" request:"pic_cover,optional"`
	PicSquare *images.Image `db:"pic_square_id" response:"pic_square" request:"pic_square,optional"`
}

func NewVenue() *Venue {
	return new(Venue)
}

func (venue *Venue) ExampleOutput() map[string]interface{} {
	return map[string]interface{}{
		"address": map[string]interface{}{
			"address1":       "foo",
			"address2":       "bar",
			"city":           "baz",
			"country":        "bazinga",
			"state_province": "fizz",
			"zip_postal":     "buzz",
		},
		"created_at":    "2014-03-27 09:29:39.165325937 -0600 MDT",
		"email":         "test@place.com",
		"facebook_user": nil,
		"id":            1,
		"is_active":     true,
		"is_verified":   true,
		"name":          "The Totally Awesome Thingy",
		"phone":         "1234567890",
		"pic":           nil,
		"pic_big":       nil,
		"pic_cover":     nil,
		"pic_small":     nil,
		"pic_square":    nil,
		"postal_code":   nil,
		"state":         nil,
		"updated_at":    "2014-03-27 09:29:39.165325937 -0600 MDT",
		"url":           "http://some.horrible.website/from/1990",
		"venue_blurb":   "WE FIGURED OUT HOW TO USE CAPSLOCK AND IT MAKES IT EASIER TO BE OBNOXIOUS NOW",
	}
}

func (venue *Venue) Valid() bool {
	return venue != nil && venue.Id > 0
}

func (venue *Venue) Authorize(access *auth.AccessData, requestType auth.RequestType) error {
	switch requestType {
	case auth.REQUEST_READ_ONE, auth.REQUEST_READ_MANY:
	case auth.REQUEST_CREATE, auth.REQUEST_UPDATE:
		// Check that access.UserId is an admin of the venue
	default:
		return errors.New("Invalid request type")
	}
	return nil
}

// ToDb returns this Venue's Id for storing in the database.
func (venue *Venue) ToDb() interface{} {
	return venue.Id
}

// DefaultDbValue returns the default value for Venue.Id.
func (venue *Venue) DefaultDbValue() interface{} {
	return new(int64)
}

// FromDb stores the Venue.Id value read from the database.
func (venue *Venue) FromDb(id interface{}) error {
	idInt := id.(*int64)
	venue.Id = *idInt
	return nil
}

func (venue *Venue) LazyLoad(options objx.Map) {
	if venue.Valid() && venue.Name == "" {
		result, err := datastore.Get(venue, venue.Id)
		if err != nil {
			panic(err)
		}
		fullVenue := result.(*Venue)
		*venue = *fullVenue
	}
}

func (venue *Venue) Receive(value interface{}) error {
	switch source := value.(type) {
	case string:
		id64, err := strconv.ParseInt(source, 10, 64)
		if err != nil {
			return err
		}
		venue.Id = id64
	case int:
		venue.Id = int64(source)
	case int8:
		venue.Id = int64(source)
	case int16:
		venue.Id = int64(source)
	case int32:
		venue.Id = int64(source)
	case int64:
		venue.Id = source
	case float32:
		venue.Id = int64(source)
	case float64:
		venue.Id = int64(source)
	default:
		return errors.New("Type of venue value not understood")
	}
	return nil
}

func (venue *Venue) ResponseValue(options objx.Map) interface{} {
	if venue.Valid() {
		return map[string]interface{}{
			"id":       venue.Id,
			"name":     venue.Name,
			"address":  venue.Address,
			"timezone": venue.TimeZone,
			"link":     settings.UrlFor("venues", venue.Id),
		}
	}
	return nil
}

func (venue *Venue) Location() string {
	return settings.UrlFor("venues", venue.Id)
}
