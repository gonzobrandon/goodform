package media

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Radiobox/web-frontend/controllers/util"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/base"
	"github.com/Radiobox/web-frontend/models/images"
	"github.com/Radiobox/web-frontend/models/slugs"
	"github.com/Radiobox/web-frontend/models/users"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/coopernurse/gorp"
	query_interfaces "github.com/nelsam/gorp_queries/interfaces"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
)

type ArtistUser struct {
	Id     string         `db:"artist_user_id"`
	Artist *Artist        `db:"artist_id"`
	User   *users.Profile `db:"user_id"`
	Admin  bool           `db:"is_admin"`
}

type Artist struct {
	base.CreatedUpdatedModel
	Id int64 `db:"artist_id" request:"-"`

	HometownAddress base.Address    `db:"hometown_address"`
	Keywords        *base.JsonArray `request:",optional"`
	ContactEmail    *string         `db:"contact_email"`
	Description     *string         `request:",optional"`

	// Should be moved to a separate type
	PicCover  *images.ArtistImage `db:"pic_cover_id" response:"pic_cover" request:"pic_cover,optional"`
	PicSquare *images.ArtistImage `db:"pic_square_id" response:"pic_square" request:"pic_square,optional"`

	Username base.Username
	Slug     *slugs.Slug `db:"-" request:",optional"`

	Albums *AlbumCollection     `db:"-" request:"-"`
	Tracks *TrackCollection     `db:"-" request:"-"`
	Events *LiveEventCollection `db:"-" request:"-"`

	// Should be moved to a separate type
	Phone       *int64  `db:"contact_phone" request:",optional"`
	CountryCode *int64  `db:"contact_phone_country_code" request:",optional"`
	Extension   *string `db:"contact_phone_extension" request:",optional"`

	Owner *users.Profile `db:"-" response:"-"`

	Subscribers *int64  `db:"subscriber_count" request:"-"`
	Timezone    *string `db:"timezone" request:",optional"`
	Website     *string `request:",optional"`

	Active   base.FakeBoolean `db:"is_active" request:"-"`
	Verified base.FakeBoolean `db:"is_verified" request:"-"`

	// Facebook
	FacebookPage *string `db:"facebook_page_id" request:",optional"`

	// Lazy loading values - these will only have Ids loaded, at
	// first, so you will need to call the corresponding Get* method
	// to retrieve anything from the DB beyond that.

	Members       *base.JsonMap `db:"band_user_ids" response:"band_users"`
	membersLoaded bool          `db:"-"`

	BookingAgent       *users.Profile `db:"booking_user_id" response:"booking_user"`
	bookingAgentLoaded bool           `db:"-"`

	Manager       *users.Profile `db:"manager_user_id" response:"manager_user"`
	managerLoaded bool           `db:"-"`
}

func NewArtist() *Artist {
	return new(Artist)
}

func (artist *Artist) ExampleOutput() map[string]interface{} {
	return map[string]interface{}{
		"albums":        []interface{}{},
		"booking_user":  nil,
		"contact_email": "test_band@test.com",
		"created_at":    "2014-03-03 23:19:30.739165 +0000 UTC",
		"description":   "This is totally a band.",
		"hometown_address": map[string]interface{}{
			"address1":       "test",
			"address2":       "place",
			"city":           "thing",
			"country":        "USA",
			"state_province": "CO",
			"zip_postal":     "80231",
		},
		"id": 1,
		"keywords": []interface{}{
			"test_band",
			"test",
			"band",
		},
		"manager_user": nil,
		"band_users":   map[string]interface{}{},
		"pic":          nil,
		"pic_cover":    nil,
		"pic_square":   nil,
		"slug":         "test_band",
		"tracks":       []interface{}{},
		"updated_at":   "2014-03-04 00:35:54.804424 +0000 UTC",
		"username":     "test_band",
	}
}

func (artist *Artist) Valid() bool {
	return artist != nil && artist.Id > 0
}

func (artist *Artist) Authorize(access *auth.AccessData, requestType auth.RequestType) error {
	switch requestType {
	case auth.REQUEST_READ_ONE, auth.REQUEST_READ_MANY:
	case auth.REQUEST_CREATE:
		if access == nil || access.UserId <= 0 {
			return errors.New("Only logged in users can create artists")
		}
		artist.Owner = &users.Profile{Id: access.UserId}
	case auth.REQUEST_UPDATE:
		if access == nil || access.UserId <= 0 {
			return errors.New("You must be logged in to update an artist profile")
		}
		if !artist.IsAdmin(access.UserId) {
			return errors.New("You are not an admin of this artist")
		}
	default:
		return errors.New("Invalid request type")
	}
	return nil
}

func (artist *Artist) Identifier() interface{} {
	if !artist.Valid() {
		return nil
	}
	return artist.Id
}

func (artist *Artist) IdentifierName() string {
	return "artist_id"
}

func (artist *Artist) TypeName() string {
	return "artist"
}

func (artist *Artist) IsAdmin(userId int64) bool {
	query := "SELECT count(*) FROM artist_user WHERE artist_id = $1 AND user_id = $2 AND is_admin"
	count, err := datastore.SelectInt(query, artist.Id, userId)
	if err != nil {
		log.Print("Error while checking for admin: " + err.Error())
		return false
	}
	return count > 0
}

func (artist *Artist) PostUpdate(exec gorp.SqlExecutor) error {
	if artist.Slug != nil {
		artist.Slug.SetTarget(artist)
		_, err := exec.Update(artist.Slug)
		if err != nil {
			return err
		}
	}
	return nil
}

func (artist *Artist) PostInsert(exec gorp.SqlExecutor) error {
	if artist.Slug == nil {
		artist.Slug = slugs.New(artist, string(artist.Username))
	} else {
		artist.Slug.SetTarget(artist)
	}
	if err := exec.Insert(artist.Slug); err != nil {
		return err
	}

	mapper := &ArtistUser{
		Artist: artist,
		User:   artist.Owner,
		Admin:  true,
	}
	return exec.Insert(mapper)
}

func (artist *Artist) RelatedLinks() map[string]string {
	links := make(map[string]string)
	if artist.Slug != nil {
		links["slug"] = artist.Slug.Location()
	}
	// TODO: Link to albums
	// TODO: Link to tracks
	links["events"] = fmt.Sprintf("%s?artist=%d", settings.SiteMap["events"], artist.Id)
	if artist.Owner.Valid() {
		links["owner"] = artist.Owner.Location()
	}
	if artist.BookingAgent.Valid() {
		links["booking_user"] = artist.BookingAgent.Location()
	}
	if artist.Manager.Valid() {
		links["manager_user"] = artist.Manager.Location()
	}
	return links
}

func (artist *Artist) Location() string {
	return settings.UrlFor("artists", artist.Id)
}

func (artist *Artist) GetEvents(options objx.Map) *LiveEventCollection {
	// Default is to return any event(s) that is/are currently live,
	// followed by upcoming events
	artist.Events = new(LiveEventCollection)

	query, err := artist.Events.QueryFromParams(objx.Map{"type": "upcoming_not_ended", "artist": artist.Id})
	if err != nil {
		log.Print("Received error generating where clause for upcoming events: ", err)
		return nil
	}

	err = query.SelectToTarget(artist.Events.ResponseObject())
	if err != nil {
		log.Print("Received error trying to query live-now artist events: ", err)
	}

	return artist.Events
}

func (artist *Artist) GetSlug() *slugs.Slug {
	if artist.Slug == nil {
		query := "SELECT * FROM slugs WHERE (target->>'artist_id')::int = $1"
		results, err := datastore.Select(new(slugs.Slug), query, artist.Id)
		if err != nil || len(results) != 1 {
			if err != nil {
				log.Print("Got error trying to select artist slug: " + err.Error())
			} else {
				log.Print("No slug found for artist " + artist.Username)
			}
			return nil
		}
		artist.Slug = results[0].(*slugs.Slug)
		artist.Slug.SetTarget(artist)
	}
	return artist.Slug
}

func (artist *Artist) SetSlug(slug *slugs.Slug) {
	artist.Slug = slug
}

// Load all the sub-bits of this artist.
//
// PERF: We don't necessarily know for certain that all of these
// elements are needed, so when this is taking too long we can use the
// options value to see if they're necessary.
func (artist *Artist) LazyLoad(options objx.Map) {
	if artist.Valid() {
		if artist.Username == "" {
			if result, err := datastore.Get(artist, artist.Id); err == nil && result != nil {
				fullArtist := result.(*Artist)
				fullArtist.Albums = artist.Albums
				fullArtist.Tracks = artist.Tracks
				*artist = *fullArtist
			}
		}
		artist.GetSlug()
		artist.GetManager()
		artist.GetMembers()
		artist.GetBookingAgent()
		artist.GetAlbums()
		artist.GetTracks()
		artist.GetEvents(objx.Map(options.Get("events").MSI()))
	}
}

func (artist *Artist) ResponseValue(options objx.Map) interface{} {
	if artist.Valid() {
		return nil
	}
	return map[string]interface{}{
		"id":       artist.Id,
		"username": artist.Username,
		"link":     settings.UrlFor("artists", artist.Id),
	}
}

func (artist Artist) ToDb() interface{} {
	return artist.Id
}

func (artist *Artist) DefaultDbValue() interface{} {
	return new(int64)
}

func (artist *Artist) FromDb(id interface{}) error {
	intIdPtr := id.(*int64)
	artist.Id = *intIdPtr
	return nil
}

func (artist *Artist) Receive(value interface{}) error {
	switch idValue := value.(type) {
	case string:
		id64, err := strconv.ParseInt(idValue, 10, 64)
		if err != nil {
			return err
		}
		artist.Id = id64
	case float32:
		artist.Id = int64(idValue)
	case float64:
		artist.Id = int64(idValue)
	case int64:
		artist.Id = idValue
	case int32:
		artist.Id = int64(idValue)
	case int:
		artist.Id = int64(idValue)
	default:
		return errors.New(fmt.Sprintf("Don't know how to handle ID value %v", idValue))
	}
	return nil
}

func (artist *Artist) ValidateInput(value interface{}) error {
	if value == nil {
		return nil
	}
	if err := artist.Receive(value); err != nil {
		return errors.New("Cannot parse input as artist ID")
	}
	if existingValue, err := datastore.Get(artist, artist.Id); err != nil {
		return errors.New("Internal error: cannot query for artist value: " + err.Error())
	} else if existingValue == nil {
		return errors.New("No artist found by that ID")
	}
	return nil
}

// GetMembers retrieves the full map of members to user profile
// objects.  It may not be very useful to have in the backend, though.
func (artist *Artist) GetMembers() *base.JsonMap {
	if !artist.membersLoaded {
		var (
			err    error
			result interface{}
		)
		for index, entry := range artist.Members.Map {
			if floatId, ok := entry.(float64); ok {
				memberId := int(floatId)
				result, err = datastore.Get(new(users.Profile), memberId)
				if err == nil {
					artist.Members.Set(index, result)
				}
			} else {
				artist.Members.Set(index, map[string]interface{}{
					"name": entry.(string),
					"link": nil,
				})
			}
		}
		artist.membersLoaded = true
	}
	return artist.Members
}

// GetBookingAgent returns the full user profile for this Artist's
// BookingAgent.  Much like GetMembers, it may not be that useful,
// because angular might be doing the heavy lifting here.
func (artist *Artist) GetBookingAgent() *users.Profile {
	if !artist.bookingAgentLoaded && artist.BookingAgent.Valid() {
		result, err := datastore.Get(new(users.Profile), artist.BookingAgent.Id)
		if err == nil {
			artist.BookingAgent = result.(*users.Profile)
		}
		artist.bookingAgentLoaded = true
	}
	return artist.BookingAgent
}

// GetManager returns the full user profile for this Artist's Manager.
// Again, it might not be useful.
func (artist *Artist) GetManager() *users.Profile {
	if !artist.managerLoaded && artist.Manager.Valid() {
		result, err := datastore.Get(new(users.Profile), artist.Manager.Id)
		if err == nil {
			artist.Manager = result.(*users.Profile)
		}
		artist.managerLoaded = true
	}
	return artist.Manager
}

// GetAlbums returns the list of albums that this artist has released.
func (artist *Artist) GetAlbums() *AlbumCollection {
	if artist.Albums == nil {
		artist.Albums = new(AlbumCollection)
		dbTarget := artist.Albums.ResponseObject()
		albumRef := new(Album)
		err := datastore.Query(albumRef).
			Where().
			Equal(&albumRef.AlbumArtist, artist.Id).
			SelectToTarget(dbTarget)
		if err != nil {
			artist.Albums = nil
			return nil
		}
		for _, album := range artist.Albums.albums {
			album.AlbumArtist = artist
		}
	}
	return artist.Albums
}

// GetTracks returns the list of tracks that this artist has released.
//
// TODO: This is hitting a minor error.  Figure out why.
func (artist *Artist) GetTracks() *TrackCollection {
	if artist.Tracks == nil {
		artist.Tracks = new(TrackCollection)
		dbTarget := artist.Tracks.ResponseObject()
		trackRef := new(Track)
		err := datastore.Query(trackRef).
			Where().
			Equal(&trackRef.Artist, artist.Id).
			SelectToTarget(dbTarget)
		if err != nil {
			log.Print("Artist.GetTracks Error: " + err.Error())
			return nil
		}
		for _, track := range artist.Tracks.tracks {
			track.Artist = artist
		}
	}
	return artist.Tracks
}

type ArtistCollection struct {
	artists []*Artist
	options objx.Map
}

func (artists *ArtistCollection) DefaultPageSize() int {
	return 50
}

func (artists *ArtistCollection) RelatedLinks() map[string]string {
	page := 1
	if artists.options.Has("page") {
		pageStr := artists.options.Get("page").Str()
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil {
			page = parsedPage
		}
	}
	basePath := settings.SiteMap["artists"]
	links := make(map[string]string)
	pageSize := artists.DefaultPageSize()
	if sizeStr := artists.options.Get("page_size").Str(); sizeStr != "" {
		size, err := strconv.Atoi(sizeStr)
		if err == nil {
			pageSize = size
		}
	}
	if len(artists.artists) == pageSize {
		// This is probably not the last page.
		nextOptions := artists.options.Copy()
		nextOptions.Set("page", page+1)
		nextQueryString, err := nextOptions.URLQuery()
		if err != nil {
			// This should never happen
			panic(err)
		}
		links["next"] = fmt.Sprintf("%s?%s", basePath, nextQueryString)
	}
	if page > 1 {
		prevOptions := artists.options.Copy()
		prevOptions.Set("page", page-1)
		prevQueryString, err := prevOptions.URLQuery()
		if err != nil {
			// Again, this should never happen
			panic(err)
		}
		links["prev"] = fmt.Sprintf("%s?%s", basePath, prevQueryString)
	}
	return links
}

func (artists *ArtistCollection) Location() string {
	query, err := artists.options.URLQuery()
	if err != nil {
		// Something is horribly wrong.
		panic(err)
	}
	return fmt.Sprintf("%s?%s", settings.SiteMap["artists"], query)
}

func (artists *ArtistCollection) ResponseObject() interface{} {
	if artists.artists == nil {
		artists.artists = make([]*Artist, 0, 10)
	}
	return &artists.artists
}

// Where generates a where clause (including an ORDER BY if
// appropriate) for an artist.
func (artists *ArtistCollection) Query(ctx context.Context) (query_interfaces.SelectQuery, error) {
	var err error
	artists.options, err = objx.FromURLQuery(ctx.HttpRequest().URL.RawQuery)
	if err != nil {
		return nil, err
	}

	queryType := ""
	if artists.options.Has("type") {
		queryType = strings.ToLower(artists.options.Get("type").Str())
	}
	var query query_interfaces.WhereQuery
	ref := new(Artist)
	baseQuery := datastore.Query(ref)
	switch queryType {
	case "admin":
		access, err := util.Authorize(ctx)
		if err != nil {
			return nil, err
		}
		userId := int64(-1)
		if access != nil {
			userId = access.UserId
		}
		mapRef := new(ArtistUser)
		query = baseQuery.Join(mapRef).
			On().
			Equal(&ref.Id, &mapRef.Artist).
			Where().
			Equal(&mapRef.User, userId).
			Equal(&mapRef.Admin, true)
	case "":
		query = baseQuery.Where()
	default:
		message := "Error: Live event type " + artists.options.Get("type").Str() + " not understood."
		return nil, errors.New(message)
	}
	return query.(query_interfaces.SelectQuery), nil
}
