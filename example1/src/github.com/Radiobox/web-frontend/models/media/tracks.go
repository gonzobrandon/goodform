package media

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Radiobox/web-frontend/buckets"
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

const (
	ENCODING_ENDPOINT = "https://manage.encoding.com"
	ENCODING_USER     = "30889"
	ENCODING_KEY      = "0d67bc2ace63ac903c4efb5dad42a800"
	CALLBACK_PATH     = "/api/track_encode"
	CALLBACK_DOMAIN   = "https://www.theradiobox.com"
	CALLBACK          = CALLBACK_DOMAIN + CALLBACK_PATH
)

var (
	// Valid MP3 bitrates, in kb/s
	mp3Bitrates = []int{
		128,
		144,
		160,
		192,
		224,
		256,
		320,
	}

	// We should normally only need one encoder, so the DefaultEncoder
	// is here for making encode requests.
	DefaultEncoder = &Encoder{User: ENCODING_USER, Key: ENCODING_KEY}
)

func EncodingChannelName(mediaId ...string) string {
	name := "encoding_"
	if len(mediaId) > 0 {
		name += mediaId[0] + "_"
	}
	name += "status_update"
	return name
}

// Encoder is a type that sends and receives data to and from
// encoding.com.
type Encoder struct {
	User string
	Key  string
}

func (encoder *Encoder) send(request map[string]map[string]interface{}) (map[string]interface{}, error) {
	bodyBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body := append([]byte("json="), bodyBytes...)
	log.Print("Sending request with body:\n", string(body))
	bodyBuffer := bytes.NewBuffer(body)
	resp, err := http.Post("https://manage.encoding.com", "application/x-www-form-urlencoded", bodyBuffer)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	log.Print("Response from encoding.com:\n", string(respBytes))
	respValues := make(map[string]interface{})
	if err := json.Unmarshal(respBytes, &respValues); err != nil {
		return nil, err
	}
	actualResponse := respValues["response"].(map[string]interface{})
	if errs, ok := actualResponse["errors"]; ok {
		errMap := errs.(map[string]interface{})
		return nil, errors.New(errMap["error"].(string))
	}
	return actualResponse, nil
}

func (encoder *Encoder) makeQuery(query map[string]interface{}) map[string]map[string]interface{} {
	if _, ok := query["userid"]; !ok {
		query["userid"] = encoder.User
		query["userkey"] = encoder.Key
	}
	return map[string]map[string]interface{}{"query": query}
}

func (encoder *Encoder) AddMedia(source string, destFormats ...map[string]string) (map[string]interface{}, error) {
	query := encoder.makeQuery(map[string]interface{}{
		"action":        "addMedia",
		"source":        source,
		"region":        "us-east-1",
		"notify":        CALLBACK,
		"notify_format": "json",
		"format":        destFormats,
	})
	return encoder.send(query)
}

func (encoder *Encoder) Status(mediaId string) (map[string]interface{}, error) {
	query := encoder.makeQuery(map[string]interface{}{
		"action":  "GetStatus",
		"mediaid": mediaId,
	})
	return encoder.send(query)
}

type AlbumTrack struct {
	base.CreatedUpdatedModel

	Id          string `db:"album_track_id"`
	Album       *Album `db:"album_id"`
	Track       *Track `db:"track_id"`
	TrackNumber int64  `db:"track_number"`
}

// A Track is an audio recording available for listening at any time.
type Track struct {
	base.CreatedUpdatedModel

	Id string `db:"track_id" request:"-"`

	// Title is the title of this track.
	Title base.NonEmptyString

	// Artist is the artist who created the track.
	Artist *Artist `db:"artist_id" response:"artist" request:",optional"`

	// Media is this track's actual media file.
	Media *Media `db:"media_id" response:"media"`

	// Preview is a smaller sample of this track's media.
	Preview     *Media `db:"preview_media_id" response:"preview_media"`
	PreviewOnly *bool  `db:"preview_only" request:",optional"`

	// Album is the album that this track is a part of.
	Album *Album `db:"-" request:",optional"`

	// Checksum is the md5 of the file.
	Checksum *string `db:"checksum_md5" request:",optional"`

	// Downloads is a count of how many times the track has been
	// downloaded.
	Downloads int64 `request:"-"`

	TrackNumber *int64 `db:"-" response:"-" request:"track_number,optional"`

	DownloadLinks map[string]string `db:"-" response:"media-links" request:"-"`

	// Active represents whether or not this track is available to be
	// streamed or downloaded.
	Active base.FakeBoolean `db:"is_active" request:"-"`

	// Length is the length of the track, in seconds
	Length int64 `db:"time_length" request:",optional"`

	// Should be moved to a separate type
	PicSquare *images.Image `db:"pic_square_id" response:"pic_square" request:",optional"`

	PriceUSD    string                  `db:"price_usd" request:",optional"`
	PriceMinUSD string                  `db:"price_usd_min" request:",optional"`
	PriceNYO    base.DefaultTrueBoolean `db:"price_nyo" request:",optional"`

	Purchases int64 `request:",optional"`

	EncodingStatus  string `db:"encoding_status" request:"-"`
	EncodingMediaId string `db:"encoding_media_id" response:"-"`
}

func NewTrack() *Track {
	return new(Track)
}

func (track *Track) ExampleOutput() map[string]interface{} {
	return map[string]interface{}{
		"id":            "",
		"artist":        nil,
		"media":         nil,
		"preview_media": nil,
		"preview_only":  nil,
		"checksum_md5":  "",
		"is_active":     false,
		"time_length":   nil,
		"pic_big":       nil,
		"pic_small":     nil,
		"pic_square":    nil,
		"price_usd":     nil,
		"price_usd_min": nil,
		"price_nyo":     nil,
	}
}

func (track *Track) Valid() bool {
	return track != nil && track.Id != ""
}

func (track *Track) Authorize(access *auth.AccessData, requestType auth.RequestType) error {
	switch requestType {
	case auth.REQUEST_READ_ONE, auth.REQUEST_READ_MANY:
	case auth.REQUEST_CREATE, auth.REQUEST_UPDATE:
		if access == nil || access.UserId <= 0 {
			return errors.New("You must be logged in to create a track")
		}
		if track.Artist == nil {
			return errors.New("You cannot create a track without an artist")
		}
		albumOrArtistAdmin := track.Artist.IsAdmin(access.UserId) || (track.Album != nil && track.Album.AlbumArtist.IsAdmin(access.UserId))
		if !albumOrArtistAdmin {
			return errors.New("You are not an admin of the artist you are trying to create a track for.")
		}
	default:
		return errors.New("Invalid request type")
	}
	return nil
}

// Encode attempts to encode this track's media as any missing values
// through our encoding service.  If this track.Media.Status !=
// "done", then it will wait until the status has been updated to
// "done" before starting the upload.
func (track *Track) Encode() {
	// Listen first
	notifier, err := datastore.Listen(track.Media.ChannelName())
	if err != nil {
		log.Print("Cannot listen for media status updates: ", err)
		return
	}
	defer notifier.Close()

	// Then check media status
	if track.mediaStatus() == "done" {
		track.finalize()
		return
	}

	// Wait for notification of done status
	log.Print("Media not uploaded, waiting for signal...")
	for {
		select {
		case <-notifier.Reconnect:
			log.Print("Reconnecting...")
			if track.mediaStatus() == "done" {
				track.finalize()
				return
			}
		case status := <-notifier.Payloads:
			log.Print("New media status: ", status)
			if status == "done" {
				track.finalize()
				return
			}
			if status == "failed" {
				track.updateEncodingStatus("Failed: media could not be uploaded")
				return
			}
		case <-time.After(5 * time.Minute):
			log.Print("Track encode error: timed out waiting for media to finish uploading.")
			track.updateEncodingStatus("Failed: track media did not finish uploading")
			return
		}
	}
}

func (track *Track) GetMedia() *Media {
	if track.Media.S3Path == "" && track.Media.Id != "" {
		fullMedia, err := datastore.Get(new(Media), track.Media.Id)
		if err != nil || fullMedia == nil {
			log.Print("Cannot load media from database: ", err)
			return nil
		}
		track.Media = fullMedia.(*Media)
	}
	return track.Media
}

func (track *Track) mediaStatus() string {
	// Replace the current Media object with a new Media object that
	// just has its Id set, to force a reload from the DB.
	track.Media = &Media{Id: track.Media.Id}
	return track.GetMedia().Status
}

func retry(task func() error, retries time.Duration, retryErrors ...string) error {
	err := task()
	for count := time.Duration(1); err != nil && count <= retries; count++ {
		if len(retryErrors) > 0 {
			shouldBreak := true
			for _, message := range retryErrors {
				if message == err.Error() {
					shouldBreak = false
					break
				}
			}
			if shouldBreak {
				break
			}
		}
		log.Printf("Task failed [%s], retrying in %d minutes", err, count)
		<-time.After(count * time.Minute)
		err = task()
	}
	return err
}

func (track *Track) updateEncodingStatus(status string) {
	statusUpdateQuery := "UPDATE tracks SET encoding_status = $1 WHERE track_id = $2"
	track.EncodingStatus = status
	if _, err := datastore.Exec(statusUpdateQuery, track.EncodingStatus, track.Id); err != nil {
		log.Print("Could not update track status: ", err)
	}
}

func (track *Track) finalize() {
	track.GetMedia()
	if track.Media.Status != "done" {
		// We need to reload this media.
		if track.mediaStatus() != "done" {
			log.Print(errors.New("Media either cannot be loaded or is not uploaded"))
			return
		}
	}
	fullpath := track.Media.S3Path
	if !strings.HasPrefix(fullpath, "tracks/") {
		fullpath = "tracks/" + fullpath
	}
	switch track.Media.ContentType {
	case "audio/flac":
		fullpath += ".flac"
	case "audio/mpeg":
		fallthrough
	case "audio/mp3":
		fullpath += ".mp3"
	case "audio/aac":
		fullpath += ".aac"
	default:
		log.Print(errors.New("MIME Type " + track.Media.ContentType + " is not supported as a track."))
		track.updateEncodingStatus("Failed: invalid media type [" + track.Media.ContentType + "]")
		return
	}
	track.updateEncodingStatus("Moving to secure track location")
	if err := retry(func() error { return track.Media.Move(buckets.Walled(), fullpath) }, 5); err != nil {
		log.Printf("Move to walled bucket failed: %s", err)
		track.updateEncodingStatus("Failed: could not move media")
		return
	}
	track.updateEncodingStatus("Queuing encode")
	if err := retry(track.sendToEncoder, 5, "Request rate over limit"); err != nil {
		log.Printf("Encoding task could not be queued: %s", err)
		track.updateEncodingStatus("Failed: encode task was rejected")
		return
	}
	return
}

func (track *Track) OpenMediaType(fileExtension string) (*Media, error) {
	media := track.GetMedia()
	if media == nil {
		return nil, errors.New("Track not found: This track has no media")
	}
	if track.EncodingStatus != "Finished" {
		return nil, errors.New("Track not found: This track has not finished encoding yet")
	}
	if !strings.HasSuffix(media.S3Path, fileExtension) {
		if fileExtension == ".flac" {
			return nil, errors.New("Track not found: This track was uploaded in a lossy format and cannot be converted to flac")
		}
		i := len(media.S3Path) - 1
		for ; i >= 0 && media.S3Path[i] != '.'; i-- {
		}
		media.S3Path = media.S3Path[:i] + fileExtension
	}
	exists, err := media.Exists()
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Track not found: This track does not exist with file extension: " + fileExtension)
	}
	if err = media.Open(); err != nil {
		return nil, err
	}
	return media, nil
}

func (track *Track) sendToEncoder() error {
	bucketUrl := fmt.Sprintf("http://%s.s3.amazonaws.com", buckets.Walled().Name)
	log.Print("Bucket URL: ", bucketUrl)
	sourcepath := fmt.Sprintf("%s/%s", bucketUrl, track.Media.S3Path)
	var mp3path, aacpath, bitrate string
	switch track.Media.ContentType {
	case "audio/flac":
		mp3path = strings.TrimSuffix(sourcepath, ".flac") + ".mp3"
		aacpath = strings.TrimSuffix(sourcepath, ".flac") + ".aac"
		bitrate = "320k"
	case "audio/mpeg":
		fallthrough
	case "audio/mp3":
		aacpath = strings.TrimSuffix(sourcepath, ".mp3") + ".aac"
	case "audio/aac":
		mp3path = strings.TrimSuffix(sourcepath, ".aac") + ".mp3"
	}
	if bitrate == "" && track.Media.Bitrate != "" && track.Media.Bitrate != "N/A" {
		bitrate = track.Media.Bitrate
	}

	track.Length = int64(track.Media.Duration.Seconds())

	formats := make([]map[string]string, 0, 2)
	if mp3path != "" {
		mp3format := map[string]string{
			"output":                "mp3",
			"audio_channels_number": "2",
			"audio_sample_rate":     "44100",
			"audio_codec":           "libmp3lame",
			"two_pass":              "no",
			"turbo":                 "no",
			"twin_turbo":            "no",
			"metadata_copy":         "yes",
			"destination":           mp3path,
		}
		if bitrate != "" {
			bitrateInt, err := strconv.Atoi(strings.TrimSuffix(bitrate, "k"))
			if err != nil {
				return err
			}
			// Find a valid mp3 bitrate - if we can't find one that
			// exactly matches the requested bitrate, go one step
			// higher.
			idx := 0
			for ; mp3Bitrates[idx] < bitrateInt && idx < len(mp3Bitrates); idx++ {
			}
			newBitrate := fmt.Sprintf("%dk", mp3Bitrates[idx])
			mp3format["audio_bitrate"] = newBitrate
		}
		formats = append(formats, mp3format)
	}
	if aacpath != "" {
		aacformat := map[string]string{
			"output":                "m4a",
			"audio_channels_number": "2",
			"audio_sample_rate":     "44100",
			"audio_codec":           "dolby_heaac",
			"two_pass":              "no",
			"turbo":                 "no",
			"twin_turbo":            "no",
			"metadata_copy":         "yes",
			"destination":           aacpath,
		}
		if bitrate != "" {
			aacformat["audio_bitrate"] = bitrate
		}
		formats = append(formats, aacformat)
	}
	response, err := DefaultEncoder.AddMedia(sourcepath, formats...)
	if err != nil {
		return err
	}
	if message, ok := response["message"]; ok {
		track.EncodingStatus = message.(string)
	}
	if mediaId, ok := response["MediaID"]; ok {
		track.EncodingMediaId = mediaId.(string)
	}
	if count, err := datastore.Update(track); err != nil {
		return err
	} else if count != 1 {
		return errors.New("Expected exactly one track to be updated")
	}
	go track.MonitorEncode()
	return nil
}

func (track *Track) EncodingChannelName() string {
	return EncodingChannelName(track.EncodingMediaId)
}

func (track *Track) Status() (string, error) {
	response, err := DefaultEncoder.Status(track.EncodingMediaId)
	if err != nil {
		return "", err
	}
	status, ok := response["status"].(string)
	if !ok {
		return "", errors.New("No status found in response")
	}
	if status == "Processing" {
		progressPercent, ok := response["progress"].(string)
		if !ok {
			return "", errors.New("No progress found in response")
		}
		status = fmt.Sprintf("%s (%s%% finished)", status, progressPercent)
	}
	return status, nil
}

// Intended to be run in a new thread.  This process periodically
// checks encoding.com to see how
func (track *Track) MonitorEncode() {
	listener, err := datastore.Listen(track.EncodingChannelName())
	if err != nil {
		log.Print("Cannot monitor encoding updates: ", err)
		return
	}
	defer listener.Close()
	var newStatus string
	for {
		err = nil
		newStatus = track.EncodingStatus
		select {
		case newStatus = <-listener.Payloads:
		case <-listener.Reconnect:
			newStatus, err = track.Status()
		case <-time.After(10 * time.Second):
			newStatus, err = track.Status()
		}
		if err != nil {
			log.Print("Error checking status: ", err)
			continue
		}
		if track.EncodingStatus != newStatus {
			track.updateEncodingStatus(newStatus)
		}
		if track.EncodingStatus == "Finished" || track.EncodingStatus == "Error" {
			return
		}
	}
}

// Start the encoding process before insert.
func (track *Track) PreInsert(exec gorp.SqlExecutor) error {
	media := track.GetMedia()
	if media == nil {
		return errors.New("Media could not be found.")
	}
	switch media.ContentType {
	case "audio/flac":
		// flac is always going to have a high enough bitrate.
	case "audio/mpeg":
		fallthrough
	case "audio/mp3":
		fallthrough
	case "audio/aac":
		// ffmpeg reports mp3 and aac bitrates in kb/s; however, it is
		// possible that we'll support a codec that will be returned
		// in Mb/s in the future, so just assume if ffmpeg is
		// reporting in anything other than kb/s, it's probably good
		// enough.
		if strings.HasSuffix(media.Bitrate, " kb/s") {
			bitrateStr := strings.TrimSuffix(media.Bitrate, " kb/s")
			bitrate, err := strconv.Atoi(bitrateStr)
			if err != nil {
				return errors.New("Cannot read bitrate: " + err.Error())
			}
			if bitrate < 128 {
				return errors.New("Tracks must be a minimum of 128 kb/s")
			}
		}
	default:
		return errors.New("Tracks can only be created with flac, mpeg, or aac media files.")
	}
	track.EncodingStatus = "Uploading track"

	go track.Encode()
	return track.CreatedUpdatedModel.PreInsert(exec)
}

// Insert into album_track if album is not null on insert.
func (track *Track) PostInsert(exec gorp.SqlExecutor) error {
	if track.Album.Valid() {
		if track.TrackNumber == nil {
			track.TrackNumber = new(int64)
			sql := "SELECT " +
				"CASE WHEN max(track_number) IS NULL " +
				"THEN 1 " +
				"ELSE max(track_number) + 1 " +
				"END " +
				"FROM album_track WHERE album_id = $1"
			nextNumber, err := exec.SelectInt(sql, track.Album.Id)
			if err != nil {
				return err
			}
			*track.TrackNumber = nextNumber
		}
		mapper := &AlbumTrack{
			Album:       track.Album,
			Track:       track,
			TrackNumber: *track.TrackNumber,
		}
		return exec.Insert(mapper)
	}
	return nil
}

func (track *Track) PostUpdate(exec gorp.SqlExecutor) error {
	if track.Album != nil && !track.Album.dbSynced {
		// This album was assigned in a request.
		query := "DELETE FROM album_track WHERE track_id = $1"
		if _, err := exec.Exec(query, track.Id); err != nil {
			log.Print("Could not delete old track map: ", err)
		}
		return track.PostInsert(exec)
	} else if track.TrackNumber != nil {
		// The track number was changed
		album := track.GetAlbum()
		if !album.Valid() {
			return errors.New("Cannot update track number without album_id")
		}
		query := "UPDATE album_track SET track_number = $1 WHERE track_id = $2 AND album_id = $3"
		_, err := datastore.Exec(query, track.TrackNumber, track.Id, track.Album.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

// ToDb returns this Track's Id for storing in the database.
func (track *Track) ToDb() interface{} {
	return track.Id
}

// DefaultDbValue returns the default value for Track.Id.
func (track *Track) DefaultDbValue() interface{} {
	return new(string)
}

// FromDb stores the Track.Id value read from the database.
func (track *Track) FromDb(id interface{}) error {
	idString := id.(*string)
	track.Id = *idString
	return nil
}

func (track *Track) GetAlbum() *Album {
	if track.Album.Valid() {
		refAlbum := new(Album)
		refMap := new(AlbumTrack)
		results, err := datastore.Query(refAlbum).
			Join(refMap).On().
			Equal(&refMap.Album, &refAlbum.Id).
			Where().
			Equal(&refMap.Track, track.Id).
			Select()
		if err == nil {
			if len(results) == 0 {
				track.Album = nil
			} else if len(results) == 1 {
				track.Album = results[0].(*Album)
			} else if len(results) > 1 {
				panic("More than one album per track is not yet supported")
			}
		}
	}
	return track.Album
}

func (track *Track) GetArtist() *Artist {
	if track.Artist.Valid() && track.Artist.Username == "" {
		response, err := datastore.Get(track.Artist, track.Artist.Id)
		if err != nil {
			return nil
		}
		track.Artist = response.(*Artist)
	}
	return track.Artist
}

func (track *Track) LazyLoad(options objx.Map) {
	if track.Valid() && track.Title == "" {
		response, err := datastore.Get(track, track.Id)
		if err != nil {
			panic(err)
		}
		fullTrack := response.(*Track)
		fullTrack.Album = track.Album
		*track = *fullTrack
	}
	if track.EncodingStatus == "Finished" {
		track.DownloadLinks = track.MediaLinks()
	}
	track.GetAlbum()
	track.GetArtist()
}

func (track *Track) Receive(value interface{}) error {
	idString, ok := value.(string)
	if !ok {
		return errors.New("Track ID values must be of a string type")
	}
	track.Id = idString
	return nil
}

func (track *Track) ResponseValue(options objx.Map) interface{} {
	if track.Valid() {
		return map[string]interface{}{
			"id":          track.Id,
			"title":       string(track.Title),
			"link":        settings.UrlFor("tracks", track.Id),
			"media-links": track.MediaLinks(),
		}
	}
	return nil
}

func (track *Track) ResponseObject() interface{} {
	if track.Valid() {
		return track
	}
	return nil
}

func (track *Track) MediaLinks() map[string]string {
	links := make(map[string]string)
	if track.Valid() && track.EncodingStatus == "Finished" {
		baseUrl := settings.UrlFor("tracks", track.Id)
		links["mp3"] = baseUrl + ".mp3"
		links["aac"] = baseUrl + ".aac"
		if track.GetMedia().ContentType == "audio/flac" {
			links["flac"] = baseUrl + ".flac"
		}
	}
	return links
}

func (track *Track) RelatedLinks() map[string]string {
	links := track.MediaLinks()
	if track.Artist.Valid() {
		links["artist"] = track.Artist.Location()
	}
	if track.Album.Valid() {
		links["album"] = track.Album.Location()
	}
	return links
}

func (track *Track) Location() string {
	return settings.UrlFor("tracks", track.Id)
}

type TrackCollection struct {
	tracks  []*Track
	options objx.Map
}

func (tracks *TrackCollection) RelatedLinks() map[string]string {
	page := 1
	if tracks.options.Has("page") {
		pageStr := tracks.options.Get("page").StrSlice()[0]
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil {
			page = parsedPage
		}
	}
	basePath := settings.SiteMap["tracks"]
	nextOptions := tracks.options.Copy()
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
		prevOptions := tracks.options.Copy()
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

func (tracks *TrackCollection) ResponseObject() interface{} {
	if tracks == nil {
		return nil
	}
	if tracks.tracks == nil {
		// It's very common for albums to contain roughly 10-12 tracks.
		// Since many TrackCollections will be collections of tracks
		// in an album, using a base capacity of 13 gives just enough
		// buffer to handle all of the most common albums, without
		// allocating far more memory than necessary.
		tracks.tracks = make([]*Track, 0, 13)
	}
	return &tracks.tracks
}

// Where generates a where clause (including an ORDER BY if
// appropriate) for an track.
func (tracks *TrackCollection) Query(ctx context.Context) (query_interfaces.SelectQuery, error) {
	tracks.options = ctx.QueryParams()
	queryType := ""
	if tracks.options.Has("type") {
		queryType = strings.ToLower(tracks.options.Get("type").StrSlice()[0])
	}
	var query query_interfaces.WhereQuery
	ref := new(Track)
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
			Equal(&ref.Artist, &mapRef.Artist).
			Where().
			Equal(&mapRef.User, userId)
	case "":
		query = baseQuery.(query_interfaces.WhereQuery)
	default:
		message := "Error: Live event type " + tracks.options.Get("type").Str() + " not understood."
		return nil, errors.New(message)
	}
	return query, nil
}
