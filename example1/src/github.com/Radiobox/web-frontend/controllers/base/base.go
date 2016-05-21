package base

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/Radiobox/web-frontend/controllers/util"
	"github.com/Radiobox/web-frontend/datastore"
	httpErrors "github.com/Radiobox/web-frontend/errors"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	query_interfaces "github.com/nelsam/gorp_queries/interfaces"
	"github.com/stretchr/goweb/context"
)

const methodNotFoundMessage = "[%s]: Unsupported method at this endpoint."

type UnsupportedMethodController struct {
	web_responders.BaseRestController
}

func (controller *UnsupportedMethodController) Read(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	message := fmt.Sprintf(methodNotFoundMessage, "GET :id")
	messages.AddErrorMessage(message)
	return web_responders.Respond(ctx, http.StatusMethodNotAllowed, messages, message)
}

func (controller *UnsupportedMethodController) ReadMany(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	message := fmt.Sprintf(methodNotFoundMessage, "GET")
	messages.AddErrorMessage(message)
	return web_responders.Respond(ctx, http.StatusMethodNotAllowed, messages, message)
}

func (controller *UnsupportedMethodController) Update(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	message := fmt.Sprintf(methodNotFoundMessage, "PATCH :id")
	messages.AddErrorMessage(message)
	return web_responders.Respond(ctx, http.StatusMethodNotAllowed, messages, message)
}

func (controller *UnsupportedMethodController) UpdateMany(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	message := fmt.Sprintf(methodNotFoundMessage, "PATCH")
	messages.AddErrorMessage(message)
	return web_responders.Respond(ctx, http.StatusMethodNotAllowed, messages, message)
}

func (controller *UnsupportedMethodController) Replace(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	message := fmt.Sprintf(methodNotFoundMessage, "PUT :id")
	messages.AddErrorMessage(message)
	return web_responders.Respond(ctx, http.StatusMethodNotAllowed, messages, message)
}

func (controller *UnsupportedMethodController) Create(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	message := fmt.Sprintf(methodNotFoundMessage, "POST")
	messages.AddErrorMessage(message)
	return web_responders.Respond(ctx, http.StatusMethodNotAllowed, messages, message)
}

func (controller *UnsupportedMethodController) Delete(idString string, ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	message := fmt.Sprintf(methodNotFoundMessage, "DELETE :id")
	messages.AddErrorMessage(message)
	return web_responders.Respond(ctx, http.StatusMethodNotAllowed, messages, message)
}

func (controller *UnsupportedMethodController) DeleteMany(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	message := fmt.Sprintf(methodNotFoundMessage, "DELETE")
	messages.AddErrorMessage(message)
	return web_responders.Respond(ctx, http.StatusMethodNotAllowed, messages, message)
}

func Create(ctx context.Context, target interface{},
	messages web_responders.MessageMap) error {

	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		log.Print("Create Error: " + err.Error())
		messages.AddErrorMessage("Could not read parameters: " + err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err, settings.FullLinks)
	}

	if err := web_request_readers.UnmarshalParams(params, target); err != nil {
		messages.AddErrorMessage("Could not create model: " + err.Error())
		return web_responders.RespondWithInputErrors(ctx, messages, target, true)
	}

	if authorizer, ok := target.(auth.Authorizer); ok {
		authData, err := util.Authorize(ctx)
		var httpErr *httpErrors.HttpError
		if err != nil {
			httpErr = err.(*httpErrors.HttpError)
			authData = nil
		}
		if err := authorizer.Authorize(authData, auth.REQUEST_CREATE); err != nil {
			status := http.StatusUnauthorized
			if httpErr != nil {
				status = httpErr.Status
				messages.AddErrorMessage(httpErr.Error())
			}
			messages.AddErrorMessage("Permission denied: " + err.Error())
			return web_responders.Respond(ctx, status, messages, err, settings.FullLinks)
		}
	}

	transaction, err := datastore.Begin()
	if err != nil {
		log.Print("Create Error: " + err.Error())
		messages.AddErrorMessage("Could not open transaction: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err.Error(), settings.FullLinks)
	}
	if err := transaction.Insert(target); err != nil {
		transaction.Rollback()
		log.Print("Create Error: " + err.Error())
		messages.AddErrorMessage("Could not write model: " + err.Error())
		return web_responders.RespondWithInputErrors(ctx, messages, target, true)
	}
	transaction.Commit()
	messages.AddInfoMessage("Model successfully created")
	return web_responders.Respond(ctx, http.StatusOK, messages, target, settings.FullLinks)
}

func Update(ctx context.Context, target interface{}, id interface{},
	messages web_responders.MessageMap) error {

	result, err := datastore.Get(target, id)
	if err != nil {
		log.Print("Update Error: " + err.Error())
		messages.AddErrorMessage(err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err.Error(), settings.FullLinks)
	} else if result == nil {
		messages.AddErrorMessage("Not found")
		return web_responders.Respond(ctx, http.StatusNotFound, messages, "404: Not found", settings.FullLinks)
	}

	target = result
	if authorizer, ok := target.(auth.Authorizer); ok {
		authData, err := util.Authorize(ctx)
		var httpErr *httpErrors.HttpError
		if err != nil {
			httpErr = err.(*httpErrors.HttpError)
			authData = nil
		}
		if err := authorizer.Authorize(authData, auth.REQUEST_UPDATE); err != nil {
			status := http.StatusUnauthorized
			if httpErr != nil {
				status = httpErr.Status
				messages.AddErrorMessage(httpErr.Error())
			}
			messages.AddErrorMessage("Permission denied: " + err.Error())
			return web_responders.Respond(ctx, status, messages, err, settings.FullLinks)
		}
	}

	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		log.Print("Update Error: " + err.Error())
		messages.AddErrorMessage(err.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, err.Error(), settings.FullLinks)
	}

	if err = web_request_readers.UnmarshalParams(params, target); err != nil {
		if _, ok := err.(web_request_readers.MissingFields); !ok {
			messages.AddErrorMessage("Could not update model: " + err.Error())
			return web_responders.RespondWithInputErrors(ctx, messages, target, false)
		}
	}

	transaction, err := datastore.Begin()
	if err != nil {
		log.Print("Update Error: " + err.Error())
		messages.AddErrorMessage("Could not open transaction: " + err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err.Error(), settings.FullLinks)
	}
	if count, err := transaction.Update(target); err != nil || count != 1 {
		transaction.Rollback()
		if err != nil {
			log.Print("Received error from update: " + err.Error())
		} else {
			log.Printf("Received unexpected update count: %d", count)
		}
		message := "Could not update model: " + err.Error()
		messages.AddErrorMessage(message)
		return web_responders.RespondWithInputErrors(ctx, messages, target, false)
	}
	transaction.Commit()
	messages.AddInfoMessage("Update successful")
	return web_responders.Respond(ctx, http.StatusOK, messages, target, settings.FullLinks)
}

func Read(ctx context.Context, target interface{}, id interface{},
	messages web_responders.MessageMap, authDataList ...*auth.AccessData) error {

	result, err := datastore.Get(target, id)
	if err != nil {
		log.Print("Read Error: " + err.Error())
		messages.AddErrorMessage(err.Error())
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, err.Error(), settings.FullLinks)
	}
	if result == nil {
		messages.AddErrorMessage("Not found")
		return web_responders.Respond(ctx, http.StatusNotFound, messages, "Not found", settings.FullLinks)
	}
	if authorizer, ok := result.(auth.Authorizer); ok {
		authData, err := util.Authorize(ctx)
		var httpErr *httpErrors.HttpError
		if err != nil {
			httpErr = err.(*httpErrors.HttpError)
			authData = nil
		}
		if err := authorizer.Authorize(authData, auth.REQUEST_READ_ONE); err != nil {
			status := http.StatusUnauthorized
			if httpErr != nil {
				status = httpErr.Status
				messages.AddErrorMessage(httpErr.Error())
			}
			messages.AddErrorMessage("Permission denied: " + err.Error())
			return web_responders.Respond(ctx, status, messages, err, settings.FullLinks)
		}
	}

	return web_responders.Respond(ctx, http.StatusOK, messages, result, settings.FullLinks)
}

// TODO: get table name and columns from gorp.  This is currently not
// possible because too many fields are unexported.
func ReadMany(ctx context.Context, target interface{}, table string,
	messages web_responders.MessageMap, authDataList ...*auth.AccessData) error {

	if authorizer, ok := target.(auth.Authorizer); ok {
		authData, err := util.Authorize(ctx)
		var httpErr *httpErrors.HttpError
		if err != nil {
			httpErr = err.(*httpErrors.HttpError)
			authData = nil
		}
		if err := authorizer.Authorize(authData, auth.REQUEST_READ_MANY); err != nil {
			status := http.StatusUnauthorized
			if httpErr != nil {
				status = httpErr.Status
				messages.AddErrorMessage(httpErr.Error())
			}
			messages.AddErrorMessage("Permission denied: " + err.Error())
			return web_responders.Respond(ctx, status, messages, err, settings.FullLinks)
		}
	}

	dbTarget := target
	if responder, ok := target.(web_responders.ResponseObjectCreator); ok {
		dbTarget = responder.ResponseObject()
	}

	targetVal := reflect.ValueOf(dbTarget)
	targetIsSlicePtr := targetVal.Kind() == reflect.Ptr && targetVal.Elem().Kind() == reflect.Slice

	var (
		query query_interfaces.SelectQuery
		err   error
	)
	if queryer, ok := target.(datastore.Queryer); ok {
		query, err = queryer.Query(ctx)
		if err != nil {
			status := http.StatusBadRequest
			message := err.Error()
			if httpErr, ok := err.(*httpErrors.HttpError); ok {
				status = httpErr.Status
				message = httpErr.Message
			}
			messages.AddErrorMessage(message)
			return web_responders.Respond(ctx, status, messages, message, settings.FullLinks)
		}
	} else {
		query = datastore.Query(dbTarget)
	}
	defaultPageSize := 50
	if defaultPageSizer, ok := target.(DefaultPageSizer); ok {
		defaultPageSize = defaultPageSizer.DefaultPageSize()
	}
	offset, limit, err := web_request_readers.ParsePage(ctx.QueryParams(), defaultPageSize)
	if err != nil {
		messages.AddWarningMessage("Could not parse page and page_size values: " + err.Error())
		limit = defaultPageSize
	}
	query = query.Offset(int64(offset))
	query = query.Limit(int64(limit))

	var results interface{}
	if targetIsSlicePtr {
		results = dbTarget
		err = query.SelectToTarget(dbTarget)
	} else {
		results, err = query.Select()
	}
	if err != nil {
		log.Print("Read Many Error: " + err.Error())
		messages.AddErrorMessage("Could not load results from database")
		return web_responders.Respond(ctx, http.StatusInternalServerError, messages, "Could not load results from database", settings.FullLinks)
	}
	return web_responders.Respond(ctx, http.StatusOK, messages, results, settings.FullLinks)
}
