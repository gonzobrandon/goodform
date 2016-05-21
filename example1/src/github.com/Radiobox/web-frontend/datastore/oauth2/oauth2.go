package oauth2

import (
	"errors"
	"log"
	"net/http"

	"github.com/Radiobox/osin"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
)

var (
	defaultOauth2Storage *Oauth2Storage
	AuthNotFound         = errors.New("No authorization found in storage.")
)

// DefaultOauth2Storage returns the current default storage singleton
// for oauth2 data.
func DefaultOauth2Storage() *Oauth2Storage {
	if defaultOauth2Storage == nil {
		defaultOauth2Storage = new(Oauth2Storage)
	}
	return defaultOauth2Storage
}

// Oauth2Storage is a mapper between our database and oauth2
// libraries.
//
// TODO: Add refresh token support.
type Oauth2Storage struct {
}

func (storage *Oauth2Storage) GetClient(id string) (osin.Client, error) {
	result, err := datastore.Get(new(auth.Client), id)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*auth.Client), nil
}

func (storage *Oauth2Storage) SaveAuthorize(data osin.AuthorizeData) error {
	count, err := datastore.Update(data)
	if err == nil && count == 0 {
		err = datastore.Insert(data)
	}
	return err
}

func (storage *Oauth2Storage) LoadAuthorize(code string) (osin.AuthorizeData, error) {
	result, err := datastore.Get(new(auth.AuthorizeData), code)
	if err != nil || result == nil {
		return nil, err
	}
	data := result.(*auth.AuthorizeData)
	fullClient, err := storage.GetClient(data.GetClient().GetId())
	if err != nil || fullClient == nil {
		return nil, err
	}
	data.SetClient(fullClient)
	return data, nil
}

func (storage *Oauth2Storage) RemoveAuthorize(code string) error {
	data, err := storage.LoadAuthorize(code)
	if err != nil {
		return err
	}
	_, err = datastore.Delete(data)
	return err
}

func (storage *Oauth2Storage) SaveAccess(access osin.AccessData) error {
	count, err := datastore.Update(access)
	if err == nil && count == 0 {
		err = datastore.Insert(access)
	}
	return err
}

func (storage *Oauth2Storage) LoadAccess(token string) (osin.AccessData, error) {
	result, err := datastore.Get(new(auth.AccessData), token)
	if err != nil || result == nil {
		log.Print("Could not get access data")
		return nil, err
	}
	access := result.(*auth.AccessData)
	fullClient, err := storage.GetClient(access.GetClient().GetId())
	if err != nil || fullClient == nil {
		log.Print("Could not get full client")
		return nil, err
	}
	access.SetClient(fullClient)
	return access, nil
}

func (storage *Oauth2Storage) RemoveAccess(token string) error {
	access, err := storage.LoadAccess(token)
	if err != nil {
		return err
	}
	_, err = datastore.Delete(access)
	return err
}

func (storage *Oauth2Storage) LoadRefresh(refreshToken string) (osin.AccessData, error) {
	refData := new(auth.AccessData)
	results, err := datastore.Query(refData).
		Where().
		Equal(&refData.RefreshToken, refreshToken).
		Select()
	if err != nil {
		return nil, err
	}
	if len(results) < 1 {
		return nil, &osin.HttpError{
			Status:  http.StatusNotFound,
			Message: "No access found",
		}
	}
	if len(results) > 1 {
		return nil, errors.New("More than one access data found for that refresh")
	}
	return results[0].(*auth.AccessData), nil
}

func (storage *Oauth2Storage) RemoveRefresh(refreshToken string) error {
	// This is really a flaw in the osin library.  It removes the
	// refresh token only when it removes the access data, so there's
	// no point in actually removing it.
	return nil
}
