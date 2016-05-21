package images

import (
	"bytes"
	"errors"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"strings"

	"code.google.com/p/go-uuid/uuid"
	"github.com/Radiobox/web-frontend/buckets"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/base"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
	"github.com/disintegration/imaging"
	"github.com/stretchr/objx"
)

const (
	ANON_ARTIST = "http://static1.theradiobox.com/inhouse/web/anon-artist.png"
	ANON_USER   = "http://static1.theradiobox.com/inhouse/web/anon-user.png"
	ANON_ALBUM  = "http://static1.theradiobox.com/inhouse/web/anon-album.png"
	MAX_WIDTH   = 1200
)

// An ArtistImage is an Image that will return the ANON_ARTIST image
// if it is null.
type ArtistImage struct {
	Image
}

func (i *ArtistImage) ResponseValue(options objx.Map) interface{} {
	if i == nil {
		return ANON_ARTIST
	}
	return i.Image.ResponseValue(options)
}

// A UserImage is an Image that will return the ANON_USER image if it
// is null.
type UserImage struct {
	Image
}

func (i *UserImage) ResponseValue(options objx.Map) interface{} {
	if i == nil {
		return ANON_USER
	}
	return i.Image.ResponseValue(options)
}

// An AlbumImage is an Image that will return the ANON_ALBUM image if
// it is null.
type AlbumImage struct {
	Image
}

func (i *AlbumImage) ResponseValue(options objx.Map) interface{} {
	if i == nil {
		return ANON_ALBUM
	}
	return i.Image.ResponseValue(options)
}

type Image struct {
	base.CreatedUpdatedModel

	Id        string `db:"photo_id"`
	SecretUrl string `db:"secret_url" response:"-"`

	image      image.Image `db:"-" response:"-"`
	rawImage   image.Image `db:"-" response:"-"`
	imageBytes []byte      `db:"-" response:"-"`

	Options *base.JsonMap
}

func New() *Image {
	return &Image{
		Options: new(base.JsonMap),
	}
}

func (i Image) ToDb() interface{} {
	return i.Id
}

func (i *Image) DefaultDbValue() interface{} {
	return new(string)
}

func (i *Image) FromDb(value interface{}) error {
	idPtr := value.(*string)
	i.Id = *idPtr
	return nil
}

func (i *Image) Receive(value interface{}) error {
	id, ok := value.(string)
	if !ok {
		return errors.New("Image IDs must be of string types")
	}
	i.Id = id
	return nil
}

func (i *Image) GenerateName() error {
	if buckets.Assets() == nil {
		return errors.New("Cannot generate name: no S3 connection present")
	}
	exists, err := buckets.Assets().Exists(i.SecretUrl)
	for ; i.SecretUrl == "" || exists; exists, err = buckets.Assets().Exists(i.SecretUrl) {
		if err != nil {
			return err
		}
		i.SecretUrl = aws.Encode(uuid.New())
	}
	return nil
}

func (i *Image) ResponseValue(options objx.Map) interface{} {
	if i != nil {
		return settings.UrlFor("images", i.Id+".jpg")
	}
	return nil
}

func (i *Image) SetImage(imageReader io.Reader) error {
	var err error
	i.rawImage, _, err = image.Decode(imageReader)
	return err
}

func (i *Image) RelatedLinks() map[string]string {
	return map[string]string{
		"options": i.Location(),
		"image":   i.Location() + ".jpg",
	}
}

func (i *Image) Location() string {
	return settings.UrlFor("images", i.Id)
}

func (i *Image) OptionsMap() objx.Map {
	if i.Options == nil {
		i.Options = new(base.JsonMap)
	}
	if i.Options.Map == nil {
		i.Options.Map = make(objx.Map)
	}
	return i.Options.Map
}

func (i *Image) CropOptions(externalOptions ...objx.Map) objx.Map {
	options := i.OptionsMap()
	if len(externalOptions) == 1 {
		options = externalOptions[0]
	}
	return objx.Map(options.Get("crop-percent").MSI())
}

func (i *Image) SetCropOptions(newOptions objx.Map) {
	oldOptions := i.CropOptions()
	if oldOptions == nil {
		oldOptions = make(objx.Map)
	}
	oldOptions.MergeHere(newOptions)
	i.Options.Set("crop-percent", map[string]interface{}(oldOptions))
}

func (i *Image) SizeOptions(externalOptions ...objx.Map) objx.Map {
	options := i.OptionsMap()
	if len(externalOptions) == 1 {
		options = externalOptions[0]
	}
	return objx.Map(options.Get("size").MSI())
}

func (i *Image) SetSizeOptions(newOptions objx.Map) {
	oldOptions := i.SizeOptions()
	if oldOptions == nil {
		oldOptions = make(objx.Map)
	}
	oldOptions.MergeHere(newOptions)
	i.Options.Set("size", map[string]interface{}(oldOptions))
}

func (i *Image) RawImage() (image.Image, string, error) {
	if i.rawImage == nil {
		var (
			imageReader io.ReadCloser
			err         error
		)
		if strings.HasPrefix(i.SecretUrl, "http") {
			resp, err := http.Get(i.SecretUrl)
			if err != nil {
				return nil, "", err
			}
			imageReader = resp.Body
		} else {
			if buckets.Assets() == nil {
				return nil, "", errors.New("Cannot get raw image: no S3 connection present")
			}
			imageReader, err = buckets.Assets().GetReader(i.SecretUrl)
			if err != nil {
				return nil, "", err
			}
		}
		defer imageReader.Close()

		i.rawImage, _, err = image.Decode(imageReader)
		if err != nil {
			return nil, "", err
		}
	}
	return i.rawImage, "image/jpeg", nil
}

// Image returns the actual image, formatted based on Image.Options,
// as well as the image format name, for use in a Content-Type header.
func (i *Image) Image(externalOptions ...objx.Map) (image.Image, string, error) {
	var options objx.Map
	if len(externalOptions) == 1 {
		options = externalOptions[0]
	}
	// If the image is not nil, but there were non-nil external
	// options passed, regenerate the image.
	if i.image == nil || options != nil {
		if options == nil {
			options = i.Options.Map
		} else {
			options = i.Options.Merge(options)
		}

		imageObj, _, err := i.RawImage()
		if err != nil {
			return nil, "", err
		}

		cropOptions := i.CropOptions(options)
		if cropOptions != nil {
			leftMult := cropOptions.Get("left").Float64()
			rightMult := cropOptions.Get("right").Float64()
			topMult := cropOptions.Get("top").Float64()
			bottomMult := cropOptions.Get("bottom").Float64()

			imageSize := imageObj.Bounds()
			imageWidth := imageSize.Max.X - imageSize.Min.X
			imageHeight := imageSize.Max.Y - imageSize.Min.Y
			imageSize.Min.X += int(leftMult * float64(imageWidth))
			imageSize.Max.X -= int(rightMult * float64(imageWidth))
			imageSize.Min.Y += int(topMult * float64(imageHeight))
			imageSize.Max.Y -= int(bottomMult * float64(imageHeight))

			imageObj = imaging.Crop(imageObj, imageSize)
		}

		resizeOptions := i.SizeOptions(options)
		if resizeOptions != nil {
			oldHeight := imageObj.Bounds().Max.Y - imageObj.Bounds().Min.Y
			oldWidth := imageObj.Bounds().Max.X - imageObj.Bounds().Min.X

			// If either width or height comes out to zero, assume
			// that we should preserve the aspect ratio.
			width := int(resizeOptions.Get("width").Float64())
			height := int(resizeOptions.Get("height").Float64())

			if width <= 0 {
				width = int(float32(oldWidth) / float32(oldHeight) * float32(height))
			}
			if height <= 0 {
				height = int(float32(oldHeight) / float32(oldWidth) * float32(width))
			}

			// If they're still zero, then someone's trying to resize
			// an image to 0x0, or very close to it.
			if height == 0 || width == 0 {
				return nil, "", errors.New("Cannot resize to zero width or height")
			}

			imageObj = imaging.Resize(imageObj, width, height, imaging.Box)
		}

		i.image = imageObj
	}
	return i.image, "image/jpeg", nil
}

func (i *Image) Encode(w io.Writer, externalOptions ...objx.Map) error {
	_, _, err := i.Image(externalOptions...)
	if err != nil {
		return err
	}

	return jpeg.Encode(w, i.image, &jpeg.Options{Quality: 80})
}

func (i *Image) Bytes(externalOptions ...objx.Map) ([]byte, error) {
	if i.imageBytes == nil {
		buffer := new(bytes.Buffer)
		if err := i.Encode(buffer, externalOptions...); err != nil {
			return nil, err
		}
		i.imageBytes = buffer.Bytes()
	}
	return i.imageBytes, nil
}

func (i *Image) Save(writeToS3 bool) error {
	if writeToS3 {
		if buckets.Assets() == nil {
			return errors.New("Cannot upload image: no S3 connection present")
		}
		rawImage, _, err := i.RawImage()
		if err != nil {
			return err
		}
		options := objx.Map{
			"crop-percent": nil,
			"size":         nil,
		}
		bounds := rawImage.Bounds()
		if bounds.Max.X-bounds.Min.X > MAX_WIDTH {
			options.Set("size", map[string]interface{}{
				"width":  float64(MAX_WIDTH),
				"height": float64(-1),
			})
		}
		imageBytes, err := i.Bytes(options)
		if err != nil {
			return err
		}
		if i.SecretUrl == "" {
			if err := i.GenerateName(); err != nil {
				return err
			}
		}
		if err := buckets.Assets().Put(i.SecretUrl, imageBytes, "image/jpeg", s3.Private, s3.Options{}); err != nil {
			return err
		}
	}
	if i.Id != "" {
		if _, err := datastore.Update(i); err != nil {
			return err
		}
	} else {
		if err := datastore.Insert(i); err != nil {
			if writeToS3 {
				buckets.Assets().Del(i.SecretUrl)
			}
			return err
		}
	}
	return nil
}
