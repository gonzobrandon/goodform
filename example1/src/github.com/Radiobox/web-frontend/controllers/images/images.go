package images

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/Radiobox/web-frontend/controllers/base"
	errorControllers "github.com/Radiobox/web-frontend/controllers/errors"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/images"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
)

type ImageController struct {
	base.UnsupportedMethodController
}

func (controller *ImageController) Path() string {
	return settings.SiteMap["images"]
}

func (controller *ImageController) WantsJpeg(ctx context.Context) bool {
	switch ctx.FileExtension() {
	case ".jpg":
		fallthrough
	case ".jpeg":
		return true
	}
	return strings.Contains(ctx.HttpRequest().Header.Get("Accept"), "image/jpeg")
}

func (controller *ImageController) Read(id string, ctx context.Context) error {
	result, err := datastore.Get(new(images.Image), id)
	if err != nil {
		log.Print("Error while retrieving image: " + err.Error())
		return errorControllers.ServerError(ctx)
	}
	if result == nil {
		log.Print("Cannot find image")
		return errorControllers.NotFound(ctx)
	}
	image := result.(*images.Image)
	var imageOptions objx.Map
	if err := ctx.HttpRequest().ParseForm(); err == nil {
		if imageOptions, err = controller.ParseImageOptions(ctx.QueryParams()); err != nil {
			log.Print("Could not parse image options: " + err.Error())
		}
	}
	if controller.WantsJpeg(ctx) {
		_, contentType, err := image.Image(imageOptions)
		if err != nil {
			log.Print("Cannot get image object: " + err.Error())
			return errorControllers.NotFound(ctx)
		}
		ctx.HttpResponseWriter().Header().Set("Content-Type", contentType)
		if err = image.Encode(ctx.HttpResponseWriter()); err != nil {
			log.Print("Failed to write image to response body")
			return errorControllers.NotFound(ctx)
		}
		return nil
	}
	return web_responders.Respond(ctx, http.StatusOK, web_responders.NewMessageMap(), image, settings.FullLinks)
}

func (controller *ImageController) ParseImageOptions(params objx.Map) (objx.Map, error) {

	var sizeOptions, cropOptions map[string]interface{}
	for key, values := range params {
		var (
			strValue   string
			floatValue float64
		)
		switch tmp := values.(type) {
		case string:
			strValue = tmp
		case []string:
			strValue = tmp[0]
		case float64:
			floatValue = tmp
		default:
			continue
		}

		if strings.HasPrefix(key, "crop-") {
			cropKey := strings.TrimPrefix(key, "crop-")
			if strValue != "" {
				var err error
				floatValue, err = parsePercent(strValue)
				if err != nil {
					return nil, errors.New("Could not parse " + cropKey + "crop percent: " + err.Error())
				}
			} else {
				floatValue /= 100
			}
			if cropOptions == nil {
				cropOptions = make(map[string]interface{})
			}
			cropOptions[cropKey] = floatValue
		} else if key == "width" || key == "height" {
			if strValue != "" {
				var err error
				floatValue, err = strconv.ParseFloat(strValue, 64)
				if err != nil {
					return nil, errors.New("Could not parse " + key + " size value: " + err.Error())
				}
			}
			if sizeOptions == nil {
				sizeOptions = make(map[string]interface{})
			}
			sizeOptions[key] = floatValue
		}
	}
	var options objx.Map
	if sizeOptions != nil || cropOptions != nil {
		// Only allocate memory if there are options passed
		options = make(objx.Map)
	}
	if sizeOptions != nil {
		options.Set("size", sizeOptions)
	}
	if cropOptions != nil {
		options.Set("crop-percent", cropOptions)
	}

	return options, nil
}

func (controller *ImageController) Create(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		messages.AddErrorMessage("Failed to parse input params: " + err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
	}

	image := images.New()
	var fileMap map[string][]*multipart.FileHeader
	if !params.Get("files").IsNil() {
		fileMap = params.Get("files").Inter().(map[string][]*multipart.FileHeader)
	}
	var imageFound bool
	for _, files := range fileMap {
		for _, fileHeader := range files {
			imageFound = true
			imageFile, err := fileHeader.Open()
			if err != nil {
				messages.AddErrorMessage("Could not open file: " + err.Error())
				return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
			}
			defer imageFile.Close()
			if err := image.SetImage(imageFile); err != nil {
				messages.AddErrorMessage("Could not set image: " + err.Error())
				return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
			}
			break
		}
		if imageFound {
			break
		}
	}
	if !imageFound {
		// The file must have already been stored on S3
		url := params.Get("url").Str()
		if url == "" {
			messages.AddErrorMessage("Image upload requires either a raw image or the url of an existing image")
			return web_responders.Respond(ctx, http.StatusBadRequest, messages, "No image found.", settings.FullLinks)
		}
		image.SecretUrl = url
	}

	imageOptions, err := controller.ParseImageOptions(params)
	if err != nil {
		messages.AddErrorMessage(err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
	}
	crop := objx.Map(imageOptions.Get("crop-percent").MSI())
	if len(crop) > 0 {
		image.SetCropOptions(crop)
	}
	resize := objx.Map(imageOptions.Get("size").MSI())
	if len(resize) > 0 {
		image.SetSizeOptions(resize)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Print("Caught panic: ", r)
			panic(r)
		}
	}()
	if err := image.Save(imageFound); err != nil {
		messages.AddErrorMessage("Could not save image: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err.Error(), settings.FullLinks)
	}

	response := map[string]interface{}{
		"image_path": image.ResponseValue(nil),
		"image_id":   image.Id,
	}
	return web_responders.Respond(ctx, http.StatusOK, messages, response, settings.FullLinks)
}

func (controller *ImageController) Update(id string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	result, err := datastore.Get(images.New(), id)
	if err != nil {
		messages.AddErrorMessage("Cannot load image: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err.Error(), settings.FullLinks)
	} else if result == nil {
		return errorControllers.NotFound(ctx)
	}
	image := result.(*images.Image)
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		messages.AddErrorMessage("Cannot parse params: " + err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err.Error(), settings.FullLinks)
	}
	newOptions, err := controller.ParseImageOptions(params)
	if err != nil {
		messages.AddErrorMessage("Cannot parse image options: " + err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err.Error(), settings.FullLinks)
	}
	crop := objx.Map(newOptions.Get("crop-percent").MSI())
	if len(crop) > 0 {
		image.SetCropOptions(crop)
	}
	resize := objx.Map(newOptions.Get("size").MSI())
	if len(resize) > 0 {
		image.SetSizeOptions(resize)
	}

	if err = image.Save(false); err != nil {
		messages.AddErrorMessage("Could not save image: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err.Error(), settings.FullLinks)
	}

	messages.AddInfoMessage("Image successfully updated")
	return web_responders.Respond(ctx, http.StatusOK, messages, map[string]interface{}{
		"image_path": image.ResponseValue(nil),
		"image_id":   image.Id,
	})
}

func parsePercent(percentStr string) (float64, error) {
	percentStr = strings.TrimSpace(percentStr)
	if percentStr[len(percentStr)-1] == '%' {
		percentStr = percentStr[:len(percentStr)-1]
	}
	percent, err := strconv.ParseFloat(percentStr, 64)
	if percent > 0 {
		percent /= 100
	}
	return percent, err
}
