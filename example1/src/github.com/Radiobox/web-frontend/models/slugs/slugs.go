package slugs

import (
	"errors"
	"strings"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/base"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/objx"
)

// slugTypes stores all registered slug types, using their
// IdentifierColumn value as their key to make it easier to load them
// from the database.
var slugTypes = map[string]SlugValue{}

// A SlugValue contains methods to identify itself to the slug
// wrapper.
type SlugValue interface {

	// Identifier should return the value of this SlugValue's primary
	// key.
	Identifier() interface{}

	// IdentifierName should return the name that should be used when
	// referencing this SlugValue in the database.  Often, this will
	// be the name of the primary key column of the SlugValue.
	IdentifierName() string

	// TypeName should return a name used to present this SlugValue in
	// a response.  The client will use this value to know which
	// interface to load.
	TypeName() string

	// SetSlug should store the Slug in the SlugValue's fields, if it
	// has a slug field.
	SetSlug(*Slug)
}

func RegisterSlugType(slugType SlugValue) {
	if _, ok := slugTypes[slugType.IdentifierName()]; ok {
		panic("Slug type for column name " + slugType.IdentifierName() + "was already registered.")
	}
	slugTypes[slugType.IdentifierName()] = slugType
}

type Target struct {
	base.JsonMap
	value SlugValue
}

func (target *Target) ToDb() interface{} {
	if target.Map == nil {
		target.Map = objx.Map{}
	}
	if target.value != nil {
		target.Map[target.value.IdentifierName()] = target.value.Identifier()
	}
	return target.JsonMap.ToDb()
}

func (target *Target) Value() SlugValue {
	if target.value == nil {
		for key, value := range target.JsonMap.Map {
			if slugType, ok := slugTypes[key]; ok {
				result, err := datastore.Get(slugType, int(value.(float64)))
				if err != nil {
					return nil
				}
				target.value = result.(SlugValue)
			}
		}
	}
	return target.value
}

func (target *Target) SetValue(value SlugValue) {
	target.value = value
}

func (target *Target) ResponseValue(options objx.Map) interface{} {
	return web_responders.CreateResponse(target.Value(), options)
}

type Slug struct {
	Id     int64 `response:"-"`
	Target *Target
	Slug   string `response:"-"`
	Type   string `db:"-"`
}

func New(target SlugValue, slug string) *Slug {
	return &Slug{
		Target: &Target{
			value: target,
		},
		Slug: slug,
	}
}

func GetSlug(slug string) (*Slug, error) {
	query := "SELECT * FROM slugs WHERE lower(slug) = $1"
	results, err := datastore.Select(new(Slug), query, strings.ToLower(slug))
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	if len(results) > 1 {
		return nil, errors.New("Too many slugs found")
	}
	return results[0].(*Slug), nil
}

func (slug *Slug) ValidateInput(input interface{}) error {
	if slugExists, err := GetSlug(slug.Slug); err != nil {
		return errors.New("Internal error: could not load slug from database")
	} else if slugExists != nil {
		return errors.New("Duplicate slug found")
	}
	return nil
}

func (slug *Slug) Authorize(access *auth.AccessData, requestType auth.RequestType) error {
	target := slug.GetTarget()
	if authorizer, ok := target.(auth.Authorizer); ok {
		return authorizer.Authorize(access, requestType)
	}
	return nil
}

func (slug *Slug) GetTarget() SlugValue {
	return slug.Target.Value()
}

func (slug *Slug) LazyLoad(options objx.Map) {
	target := slug.GetTarget()
	target.SetSlug(slug)
	if target != nil {
		slug.Type = target.TypeName()
	}
}

func (slug *Slug) SetTarget(target SlugValue) {
	if slug.Target == nil {
		slug.Target = new(Target)
	}
	slug.Target.SetValue(target)
}

func (slug *Slug) Receive(value interface{}) error {
	if value == nil {
		return errors.New("Slug values cannot be null")
	}
	valueStr, ok := value.(string)
	if !ok {
		return errors.New("Slug values must be a string type")
	}
	if valueStr == "" {
		return errors.New("Slug values cannot be empty")
	}
	slug.Slug = valueStr
	return nil
}

func (slug *Slug) ResponseValue(options objx.Map) interface{} {
	if slug != nil {
		return slug.Slug
	}
	return nil
}

func (slug *Slug) RelatedLinks() map[string]string {
	if slug.Type == "" {
		slug.LazyLoad(nil)
	}
	return map[string]string{
		slug.Type: settings.UrlFor(slug.Type+"s", slug.GetTarget().Identifier()),
	}
}

func (slug *Slug) Location() string {
	return settings.UrlFor("slugs", slug.Slug)
}
