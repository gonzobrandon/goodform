package media

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"

	"code.google.com/p/go-uuid/uuid"
	"github.com/Radiobox/web-frontend/buckets"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/base"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
	"github.com/stretchr/objx"
)

// Write a log message every time we've uploaded the following number
// of bytes.
const BYTE_LOG_WINDOW = 1 << 20

type Media struct {
	Id      string  `db:"media_id" request:"-"`
	FileId  *string `db:"file_store_id" request:",optional"`
	FileKey *string `db:"file_store_key" request:",optional"`
	Preview *bool   `db:"is_preview" response:"-"`

	OriginalFileName string            `db:"original_file_name" request:"-"`
	S3Path           string            `db:"secret_url" response:"-"`
	Bucket           *buckets.DbBucket `db:"bucket_name" response:"-"`

	ContentType   string        `db:"content_type" response:"-"`
	ContentLength int64         `db:"content_length" response:"-"`
	file          io.ReadCloser `db:"-" response:"-"`

	// Values from ffmpeg.  Empty strings here mean that ffmpeg could
	// not respond with these values.
	Bitrate  string        `db:"bitrate" response:"-"`
	Duration base.Duration `response:"-"`

	writtenBytes int64 `db:"-" response:"-"`
	lastLogBytes int64 `db:"-" response:"-"`

	Status string `request:"-"`
}

func NewMedia(bucket *buckets.DbBucket) *Media {
	return &Media{Bucket: bucket}
}

func (media *Media) Authorize(access *auth.AccessData, requestType auth.RequestType) error {
	switch requestType {
	case auth.REQUEST_READ_ONE, auth.REQUEST_READ_MANY:
	case auth.REQUEST_CREATE, auth.REQUEST_UPDATE:
		if access == nil || access.UserId <= 0 {
			return errors.New("Only logged in users can create media.")
		}
	default:
		return errors.New("Invalid request type")
	}
	return nil
}

func (media *Media) ExampleOutput() map[string]interface{} {
	return map[string]interface{}{
		"file_store_id":  "test",
		"file_store_key": "test",
		"id":             "4daf7038-2848-4330-a812-7c7b5b5186e9",
	}
}

// ToDb returns this Media's Id for storing in the database.
func (media *Media) ToDb() interface{} {
	return media.Id
}

// DefaultDbValue returns the default value for Media.Id.
func (media *Media) DefaultDbValue() interface{} {
	return new(string)
}

// FromDb stores the Media.Id value read from the database.
func (media *Media) FromDb(id interface{}) error {
	idStr, ok := id.(*string)
	if !ok {
		return errors.New("Could not read Media ID")
	}
	media.Id = *idStr
	return nil
}

func (media *Media) ValidateInput(value interface{}) error {
	if value == nil {
		return nil
	}
	if err := media.Receive(value); err != nil {
		return err
	}
	if existingValue, err := datastore.Get(media, media.Id); err != nil {
		return errors.New("Internal error: cannot query for user value: " + err.Error())
	} else if existingValue == nil {
		return errors.New("No media found by that ID")
	}
	return nil
}

func (media *Media) LazyLoad(options objx.Map) {
	if media.Status == "" {
		response, err := datastore.Get(media, media.Id)
		if err != nil {
			panic(err)
		}
		fullMedia := response.(*Media)
		*media = *fullMedia
	}
}

func (media *Media) Location() string {
	return settings.UrlFor("media", media.Id)
}

func (media *Media) Receive(value interface{}) error {
	if value != nil {
		media.Id = value.(string)
	}
	return nil
}

func (media *Media) ResponseValue(options objx.Map) interface{} {
	return settings.UrlFor("media", media.Id)
}

func (media *Media) GenerateName() error {
	if media.Bucket == nil {
		return errors.New("Cannot generate name: no S3 connection present")
	}

	// Just trying to lower the length of a line that was a good 200
	// characters.
	exists, err := media.Bucket.Exists(media.S3Path)
	for ; media.S3Path == "" || exists; exists, err = media.Bucket.Exists(media.S3Path) {
		if err != nil {
			return err
		}
		media.S3Path = aws.Encode(uuid.New())
	}
	return nil
}

func (media *Media) SetFile(f io.ReadCloser) {
	media.file = f
}

func (media *Media) ChannelName() string {
	return "media_" + media.Id + "_status_update"
}

func (media *Media) Open() error {
	var err error
	media.file, err = media.Bucket.GetReader(media.S3Path)
	return err
}

func (media *Media) Read(buffer []byte) (int, error) {
	if media.file == nil {
		var err error
		media.file, err = media.Bucket.GetReader(media.S3Path)
		if err != nil {
			return 0, err
		}
	}
	bytes, err := media.file.Read(buffer)
	media.writtenBytes += int64(bytes)
	if media.lastLogBytes+BYTE_LOG_WINDOW <= media.writtenBytes {
		media.lastLogBytes = media.writtenBytes
		percentDone := (float64(media.writtenBytes) / float64(media.ContentLength)) * 100.0
		go datastore.Notify(media.ChannelName(), fmt.Sprintf("%f%%", percentDone))
	}
	return bytes, err
}

func (media *Media) Close() error {
	if media.file == nil {
		return nil
	}
	if err := media.file.Close(); err != nil {
		return err
	}
	media.file = nil
	return nil
}

func (media *Media) Exists() (bool, error) {
	return media.Bucket.Exists(media.S3Path)
}

func (media *Media) Move(bucket *buckets.DbBucket, newPath string) error {
	if err := media.Open(); err != nil {
		return err
	}
	originalPath := media.S3Path
	originalBucket := media.Bucket
	media.Bucket = bucket
	media.S3Path = newPath
	media.Status = "moving"
	query := "UPDATE media SET status = $1 WHERE media_id = $2"
	if _, err := datastore.Exec(query, media.Status, media.Id); err != nil {
		return err
	}
	if err := media.WriteToS3(); err != nil {
		media.Status = "s3_error: " + err.Error()
		datastore.Exec(query, media.Status, media.Id)
		return err
	}
	media.Status = "done"
	// No turning back now.
	if _, err := datastore.Update(media); err != nil {
		return err
	}
	return originalBucket.Del(originalPath)
}

// findDuration returns the duration parsed from the ffmpeg output
// (usually from STDERR).  The format will be e.g. 01:02:04.88,
// meaning one hour, two minutes, four point eight eight seconds.
func findDuration(output string) (duration string) {
	remaining := output
	startPattern := "Duration: "
	startIdx := strings.Index(remaining, startPattern)
	if startIdx == -1 {
		return
	}
	startIdx += len(startPattern)
	remaining = remaining[startIdx:]
	endIdx := strings.IndexRune(remaining, ',')
	if endIdx == -1 {
		return remaining
	}
	return remaining[:endIdx]
}

// findBitrate returns the bitrate parsed from ffmpeg output (usually
// from STDERR).  The format will be e.g. 320k - the b/s won't be
// included, because our encoding service doesn't expect b/s.
func findBitrate(output string) (bitrate string) {
	remaining := output
	endIdx := strings.Index(remaining, "b/s")
	if endIdx == -1 {
		return
	}
	// Walk backward, looking for a comma
	idx := endIdx
	for ; idx >= 0; idx-- {
		if remaining[idx] == ',' {
			// Don't include the comma
			idx++
			break
		}
	}
	bitrate = remaining[idx:endIdx]
	// Get rid of spaces
	for spaceIdx := strings.IndexRune(bitrate, ' '); spaceIdx >= 0; spaceIdx = strings.IndexRune(bitrate, ' ') {
		bitrate = bitrate[:spaceIdx] + bitrate[spaceIdx+1:]
	}
	return
}

func (media *Media) WriteToS3() error {
	if readSeeker, ok := media.file.(io.ReadSeeker); ok {
		// Attempt to read bitrate and length of audio/video
		cmd := exec.Command("ffmpeg", "-i", "-")
		cmd.Stdin = readSeeker

		// ffmpeg without an output file will always error, so ignore
		// the error response from the command.
		output, _ := cmd.CombinedOutput()
		media.Bitrate = findBitrate(string(output))
		var err error
		media.Duration, err = base.ParseDuration(findDuration(string(output)))
		if err != nil {
			log.Print("Could not read ffmpeg duration output: ", err)
		}
		readSeeker.Seek(0, 0)
	}
	if media.Bucket == nil {
		log.Print("Error: cannot write to nil bucket")
		return errors.New("Cannot upload media: no S3 connection present")
	}
	if media.file == nil {
		log.Print("Error: cannot write nil file")
		return errors.New("No file to write")
	}
	defer media.Close()

	if media.S3Path == "" {
		if err := media.GenerateName(); err != nil {
			log.Print("Error generating name: ", err)
			return err
		}
		if _, err := datastore.Update(media); err != nil {
			return err
		}
	}
	err := media.Bucket.PutReader(media.S3Path, media,
		media.ContentLength, media.ContentType, s3.AuthenticatedRead, s3.Options{})
	if err != nil {
		log.Print("Error uploading to S3: ", err)
		return err
	}

	query := "UPDATE media SET status = 'done' WHERE media_id = $1"
	_, err = datastore.Exec(query, media.Id)
	datastore.Notify(media.ChannelName(), "done")
	return err
}
