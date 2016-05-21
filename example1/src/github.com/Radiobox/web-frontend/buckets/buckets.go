package buckets

import (
	"errors"
	"log"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
)

const (
	UPLOADS_BUCKET_NAME = "uploads.theradiobox.com"
	ASSETS_BUCKET_NAME  = "static1.theradiobox.com"
	WALLED_BUCKET_NAME  = "walled.theradiobox.com"
)

var (
	server                  *s3.S3
	uploads, assets, walled *DbBucket
)

type DbBucket struct {
	*s3.Bucket
}

func (bucket DbBucket) ToDb() interface{} {
	return bucket.Name
}

func (bucket *DbBucket) DefaultDbValue() interface{} {
	return new(string)
}

func (bucket *DbBucket) FromDb(value interface{}) error {
	if value == nil {
		return nil
	}
	if bucket == nil {
		return errors.New("Cannot load into nil bucket")
	}
	if Server() == nil {
		return errors.New("Cannot load bucket: no S3 connection present")
	}
	namePtr := value.(*string)
	bucket.Bucket = Server().Bucket(*namePtr)
	return nil
}

func Server() *s3.S3 {
	if server == nil {
		auth, err := aws.EnvAuth()
		if err != nil {
			log.Printf("Cannot connect to S3, image uploading disabled: %s", err)
			return nil
		}
		server = s3.New(auth, aws.USEast)
	}
	return server
}

func Uploads() *DbBucket {
	if uploads == nil {
		if Server() == nil {
			return nil
		}
		uploads = &DbBucket{Server().Bucket(UPLOADS_BUCKET_NAME)}
	}
	return uploads
}

func Assets() *DbBucket {
	if assets == nil {
		if Server() == nil {
			return nil
		}
		assets = &DbBucket{Server().Bucket(ASSETS_BUCKET_NAME)}
	}
	return assets
}

func Walled() *DbBucket {
	if walled == nil {
		if Server() == nil {
			return nil
		}
		walled = &DbBucket{Server().Bucket(WALLED_BUCKET_NAME)}
	}
	return walled
}
