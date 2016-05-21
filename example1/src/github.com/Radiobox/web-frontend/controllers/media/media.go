package media

import (
	"errors"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"

	"github.com/Radiobox/web-frontend/buckets"
	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/models/media"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
)

type MediaController struct {
	base.UnsupportedMethodController
}

func (controller *MediaController) Path() string {
	return settings.SiteMap["media"]
}

func (controller *MediaController) Create(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		messages.AddErrorMessage("Failed to parse input params: " + err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
	}

	newMedia := media.NewMedia(buckets.Uploads())
	var fileMap map[string][]*multipart.FileHeader
	if !params.Get("files").IsNil() {
		fileMap = params.Get("files").Inter().(map[string][]*multipart.FileHeader)
		delete(params, "files")
	}
	var mediaFound bool
	for _, files := range fileMap {
		for _, fileHeader := range files {
			mediaFound = true
			newFile, err := fileHeader.Open()
			if err != nil {
				messages.AddErrorMessage("Could not open file: " + err.Error())
				return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
			}
			allBytes, err := ioutil.ReadAll(newFile)
			if err != nil {
				messages.AddErrorMessage("Could not read file: " + err.Error())
			}
			newMedia.ContentLength = int64(len(allBytes))
			newFile.Seek(0, 0)
			newMedia.SetFile(newFile)

			contentType := fileHeader.Header.Get("Content-Type")
			if contentType == "" {
				err = errors.New("The uploaded file has no Content-Type header")
				messages.AddErrorMessage(err.Error())
				return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
			}
			newMedia.ContentType, _, err = mime.ParseMediaType(contentType)
			if err != nil {
				return err
			}

			go newMedia.WriteToS3()
			newMedia.Status = "uploading"
			messages.AddInfoMessage("Upload started")
		}
		if mediaFound {
			break
		}
	}
	return base.Create(ctx, newMedia, messages)
}

func (controller *MediaController) Read(idString string, ctx context.Context) error {
	return base.Read(ctx, new(media.Media), idString, web_responders.NewMessageMap())
}

// TODO: Security
func (controller *MediaController) Update(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Update(ctx, new(media.Media), idString, messages)
}
