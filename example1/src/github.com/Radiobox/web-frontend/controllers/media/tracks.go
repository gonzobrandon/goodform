package media

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Radiobox/web-frontend/controllers/base"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/media"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
)

type TrackController struct {
	base.UnsupportedMethodController
}

func (controller *TrackController) Path() string {
	return settings.SiteMap["tracks"]
}

func (controller *TrackController) Create(ctx context.Context) error {
	return base.Create(ctx, media.NewTrack(), web_responders.NewMessageMap())
}

func (controller *TrackController) Read(idString string, ctx context.Context) error {
	switch ctx.FileExtension() {
	case ".aac", ".mp3", ".flac":
		return controller.ReadMediaFile(idString, ctx)
	}
	return base.Read(ctx, media.NewTrack(), idString, web_responders.NewMessageMap())
}

func (controller *TrackController) ReadMediaFile(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	result, err := datastore.Get(media.NewTrack(), idString)
	if err != nil {
		messages.AddErrorMessage("Could not load track from database: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, nil, settings.FullLinks)
	} else if result == nil {
		messages.AddErrorMessage("No track found")
		return web_responders.Respond(ctx, http.StatusNotFound, messages, nil, settings.FullLinks)
	}
	track := result.(*media.Track)
	f, err := track.OpenMediaType(ctx.FileExtension())
	if err != nil {
		messages.AddErrorMessage("Track could not be opened: " + err.Error())
		if strings.HasPrefix(err.Error(), "Track not found") {
			return web_responders.Respond(ctx, http.StatusNotFound, messages, err, settings.FullLinks)
		}
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	if strings.HasSuffix(track.Media.S3Path, ".flac") {
		ctx.HttpResponseWriter().Header().Add("Content-Length", fmt.Sprintf("%d", track.Media.ContentLength))
	}
	buffer := make([]byte, 1<<19)
	var (
		bytesRead         int
		readErr, writeErr error
	)
	for {
		bytesRead, readErr = f.Read(buffer)
		_, writeErr = ctx.HttpResponseWriter().Write(buffer[:bytesRead])
		if readErr == io.EOF {
			break
		} else if readErr != nil {
			return readErr
		} else if writeErr != nil {
			return writeErr
		}
	}
	return nil
}

func (controller *TrackController) Update(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	return base.Update(ctx, media.NewTrack(), idString, messages)
}

func (controller *TrackController) UpdateMany(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	body, err := web_request_readers.ParseBody(ctx)
	if err != nil {
		messages.AddErrorMessage("Could not parse body:", err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
	}
	tracks, ok := body.([]interface{})
	if !ok {
		err = errors.New("Request body must be an array for batch update")
		messages.AddErrorMessage(err)
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
	}
	tx, err := datastore.Begin()
	if err != nil {
		messages.AddErrorMessage("Could not start transaction:", err)
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	response := make([]*media.Track, 0, len(tracks))
	sql := "CREATE TEMP TABLE track_update_tmp (id uuid, track_number smallint)"
	if _, err := tx.Exec(sql); err != nil {
		messages.AddErrorMessage("Could not create temp table:", err)
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	insertQuery := "INSERT INTO track_update_tmp (id, track_number) VALUES ($1, $2)"
	for _, trackInput := range tracks {
		trackMap, ok := trackInput.(objx.Map)
		if !ok {
			err = errors.New("Entries in request body must be maps for batch update")
			messages.AddErrorMessage(err)
			tx.Rollback()
			return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
		}
		id := trackMap.Get("id").Str()
		if id == "" {
			err = errors.New("Each entry must contain an ID for batch update")
			messages.AddErrorMessage(err)
			tx.Rollback()
			return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
		}
		delete(trackMap, "id")
		trackNum := trackMap.Get("track_number")
		if !trackNum.IsNil() {
			if _, err := tx.Exec(insertQuery, id, int64(trackNum.Float64())); err != nil {
				messages.AddErrorMessage("Could not insert into temp table:", err)
				return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
			}
			delete(trackMap, "track_number")
		}
		result, err := tx.Get(new(media.Track), id)
		if err != nil {
			messages.AddErrorMessage("Could not get track: " + err.Error())
			tx.Rollback()
			return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
		}
		track := result.(*media.Track)
		if track == nil {
			err = errors.New("Some of the passed in tracks do not exist")
			messages.AddErrorMessage(err)
			tx.Rollback()
			return web_responders.Respond(ctx, http.StatusNotFound, messages, err, settings.FullLinks)
		}
		if err = web_request_readers.UnmarshalParams(trackMap, track); err != nil {
			if _, ok := err.(web_request_readers.MissingFields); !ok {
				messages.AddErrorMessage("Could not update model: " + err.Error())
				ctx.Data().Set("params", trackMap)
				return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
			}
		}
		if _, err := tx.Update(track); err != nil {
			messages.AddErrorMessage("Could not update track:", err)
			return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
		}
		if !trackNum.IsNil() {
			track.TrackNumber = new(int64)
			*track.TrackNumber = int64(trackNum.Float64())
		}
		response = append(response, track)
	}
	sql = "UPDATE album_track SET track_number = null FROM track_update_tmp AS tmp WHERE track_id = tmp.id"
	if _, err := tx.Exec(sql); err != nil {
		messages.AddErrorMessage("Could not run update:", err)
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	sql = "UPDATE album_track SET track_number = tmp.track_number FROM track_update_tmp AS tmp WHERE track_id = tmp.id"
	if _, err := tx.Exec(sql); err != nil {
		messages.AddErrorMessage("Could not run update:", err)
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	tx.Exec("DROP TABLE track_update_tmp")
	err = tx.Commit()
	if err != nil {
		messages.AddErrorMessage("Could not commit transaction:", err)
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err, settings.FullLinks)
	}
	messages.AddInfoMessage("Update finished")
	return web_responders.Respond(ctx, http.StatusOK, messages, response, settings.FullLinks)
}

type TrackEncodeCallbackController struct {
	base.UnsupportedMethodController
}

func (controller *TrackEncodeCallbackController) Path() string {
	return media.CALLBACK_PATH
}

func (controller *TrackEncodeCallbackController) Create(ctx context.Context) error {
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		return err
	}
	reqParams := make(map[string]interface{})
	if jsonStr := params.Get("json").Str(); jsonStr != "" {
		if err := json.Unmarshal([]byte(jsonStr), reqParams); err != nil {
			return err
		}
	}
	if xmlStr := params.Get("xml").Str(); xmlStr != "" {
		log.Print("ERROR: Callback used XML")
		return errors.New("XML is not supported")
	}
	result := reqParams["result"].(map[string]interface{})
	datastore.Notify(media.EncodingChannelName(result["mediaid"].(string)), result["status"].(string))
	return web_responders.Respond(ctx, http.StatusOK, web_responders.NewMessageMap(), nil, settings.FullLinks)
}
