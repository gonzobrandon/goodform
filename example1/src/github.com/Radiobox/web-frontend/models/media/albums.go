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
	"github.com/Radiobox/web-frontend/settings"
	"github.com/coopernurse/gorp"
	query_interfaces "github.com/nelsam/gorp_queries/interfaces"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
)

type NumberedTrack struct {
	Track
	TrackNumber *int64 `db:"track_number"`
}

func (track *NumberedTrack) Valid() bool {
	if track == nil {
		return false
	}
	return track.Track.Valid()
}

func (track *NumberedTrack) ResponseObject() interface{} {
	if !track.Valid() {
		return nil
	}
	return track
}

type Album struct {
	base.CreatedUpdatedModel

	Id string `db:"album_id" request:"-"`

	// Downloads is the count of full album downloads for this Album.
	Downloads *int64 `request:"-"`

	// AlbumArtist is the Artist that owns this Album.
	AlbumArtist *Artist `db:"artist_id" response:"album_artist"`

	// Title is the title of this Album.
	Title base.NonEmptyString

	// Length is the length of the album, in seconds
	Length *int64 `db:"time_length" request:",optional"`

	// Active determines whether or not this Album is currently
	// available.
	Active bool `db:"is_active" request:"-"`

	PicSquare *images.AlbumImage `db:"pic_square_id" response:"pic_square" request:",optional"`

	// Tracks contains a list of tracks in the album.  It is lazy
	// loaded, so you shouldn't usually read this value directly.
	Tracks []*NumberedTrack `db:"-" request:"-"`

	// dbSynced stores whether or not this album was loaded from
	// a database query.
	dbSynced bool `db:"-" response:"-"`
}

func NewAlbum() *Album {
	return new(Album)
}

func (album *Album) ExampleOutput() map[string]interface{} {
	return map[string]interface{}{
		"album_artist": nil,
		"time_length":  nil,
		"pic_square":   nil,
		"id":           "",
		"is_active":    true,
	}
}

func (album *Album) Valid() bool {
	return album != nil && album.Id != ""
}

func (album *Album) RelatedLinks() map[string]string {
	links := make(map[string]string)
	if album.AlbumArtist.Valid() {
		links["artist"] = album.AlbumArtist.Location()
	}
	// TODO: link to tracks
	return links
}

func (album *Album) Location() string {
	return settings.UrlFor("albums", album.Id)
}

func (album *Album) PostGet(exec gorp.SqlExecutor) error {
	album.Active = true
	album.dbSynced = true
	return nil
}

func (album *Album) PreInsert(exec gorp.SqlExecutor) error {
	if err := album.CreatedUpdatedModel.PreInsert(exec); err != nil {
		return err
	}
	album.Active = false
	return nil
}

func (album *Album) PreUpdate(exec gorp.SqlExecutor) error {
	if err := album.CreatedUpdatedModel.PreInsert(exec); err != nil {
		return err
	}
	album.Active = false
	return nil
}

func (album *Album) Authorize(access *auth.AccessData, requestType auth.RequestType) error {
	switch requestType {
	case auth.REQUEST_READ_ONE, auth.REQUEST_READ_MANY:
	case auth.REQUEST_CREATE, auth.REQUEST_UPDATE:
		if access == nil {
			return errors.New("You must be logged in to create an album")
		}
		if album.AlbumArtist == nil {
			return errors.New("Cannot create an album without an artist")
		}
		if !album.AlbumArtist.IsAdmin(access.UserId) {
			return errors.New("You are not an admin of this artist")
		}
	default:
		return errors.New("Invalid request type")
	}
	return nil
}

func (album *Album) Receive(value interface{}) error {
	album.dbSynced = false
	src, ok := value.(string)
	if !ok {
		return errors.New("Album ID must be a string type")
	}
	album.Id = src
	return nil
}

// Tracks returns a slice of Track values that are part of this album.
func (album *Album) GetTracks() []*NumberedTrack {
	if album.Valid() && album.Tracks == nil {
		dbTarget := new([]*NumberedTrack)
		trackRef := new(Track)
		mapRef := new(AlbumTrack)
		err := datastore.Query(trackRef).
			Join(mapRef).On().
			Equal(&trackRef.Id, &mapRef.Track).
			Where().
			Equal(&mapRef.Album, album.Id).
			OrderBy(&mapRef.TrackNumber, "ASC").
			SelectToTarget(dbTarget)
		if err != nil {
			log.Print("Album.GetTracks Error: " + err.Error())
		}
		for index, track := range *dbTarget {
			track.Album = album
			if track.TrackNumber == nil {
				newNumber := int64(index + 1)
				track.TrackNumber = &newNumber
			}
		}
		album.Tracks = *dbTarget
	}
	return album.Tracks
}

func (album *Album) LazyLoad(options objx.Map) {
	if album.Valid() {
		if album.Title == "" {
			result, err := datastore.Get(album, album.Id)
			if err != nil {
				panic(err)
			}
			newAlbum := result.(*Album)
			newAlbum.Tracks = album.Tracks
			*album = *newAlbum
		}
		album.GetTracks()
	}
}

// Response is used by our responder when using a *Album as a value in
// a response.
func (album *Album) ResponseValue(options objx.Map) interface{} {
	if album.Valid() {
		return map[string]string{
			"id":    album.Id,
			"title": string(album.Title),
			"link":  settings.UrlFor("albums", album.Id),
		}
	} else {
		return nil
	}
}

// ToDb returns this Album's Id for storing in the database.
func (album *Album) ToDb() interface{} {
	// Once this album is saved, it's no longer from a request.
	album.dbSynced = true
	return album.Id
}

// DefaultDbValue returns the default value for Album.Id.
func (album *Album) DefaultDbValue() interface{} {
	return new(string)
}

// FromDb stores the Album.Id value read from the database.
func (album *Album) FromDb(id interface{}) error {
	idStr := id.(*string)
	album.dbSynced = true
	album.Id = *idStr
	return nil
}

type AlbumCollection struct {
	albums  []*Album
	options objx.Map
}

func (albums *AlbumCollection) LazyLoad(options objx.Map) {
	if albums != nil {
		for _, album := range albums.albums {
			if album.AlbumArtist != nil {
				album.AlbumArtist.Albums = albums
			}
		}
	}
}

func (albums *AlbumCollection) RelatedLinks() map[string]string {
	page := 1
	if albums.options.Has("page") {
		pageStr := albums.options.Get("page").Str()
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil {
			page = parsedPage
		}
	}
	basePath := settings.SiteMap["albums"]
	nextOptions := albums.options.Copy()
	nextOptions.Set("page", []string{strconv.Itoa(page + 1)})
	nextQueryString, err := nextOptions.URLQuery()
	if err != nil {
		// This should never happen
		panic(err)
	}
	links := map[string]string{
		"next": fmt.Sprintf("%s?%s", basePath, nextQueryString),
	}
	if page > 1 {
		prevOptions := albums.options.Copy()
		prevOptions.Set("page", []string{strconv.Itoa(page - 1)})
		prevQueryString, err := prevOptions.URLQuery()
		if err != nil {
			// Again, this should never happen
			panic(err)
		}
		links["prev"] = fmt.Sprintf("%s?%s", basePath, prevQueryString)
	}
	return links
}

func (albums *AlbumCollection) ResponseObject() interface{} {
	if albums == nil {
		return []*Album{}
	}
	if albums.albums == nil {
		albums.albums = make([]*Album, 0, 10)
	}
	return &albums.albums
}

// Where generates a where clause (including an ORDER BY if
// appropriate) for an album.
func (albums *AlbumCollection) Query(ctx context.Context) (query_interfaces.SelectQuery, error) {
	var err error
	albums.options, err = objx.FromURLQuery(ctx.HttpRequest().URL.RawQuery)
	if err != nil {
		return nil, err
	}
	queryType := ""
	if albums.options.Has("type") {
		queryType = strings.ToLower(albums.options.Get("type").Str())
	}
	var query query_interfaces.WhereQuery
	ref := new(Album)
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
			Equal(&ref.AlbumArtist, &mapRef.Artist).
			Where().
			Equal(&mapRef.User, userId).
			Equal(&mapRef.Admin, true)
	case "":
		query = baseQuery.Where()
	default:
		message := "Error: Live event type " + albums.options.Get("type").Str() + " not understood."
		return nil, errors.New(message)
	}

	if albums.options.Has("artist") {
		artist := albums.options.Get("artist").Str()
		query = query.Equal(&ref.AlbumArtist, artist)
	}

	return query.(query_interfaces.SelectQuery), nil
}
