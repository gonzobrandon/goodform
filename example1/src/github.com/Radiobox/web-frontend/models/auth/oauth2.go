package auth

import (
	"github.com/Radiobox/osin"
)

// RequestType is a type of request to authorize.
type RequestType uint

const (
	REQUEST_INVALID RequestType = iota
	REQUEST_CREATE
	REQUEST_UPDATE
	REQUEST_READ_ONE
	REQUEST_READ_MANY
	REQUEST_DELETE
)

type Authorizer interface {
	Authorize(*AccessData, RequestType) error
}

// OSIN (our oauth2 library) provides structs that are built for
// oauth2 use.  However, they're not in a state that will make it easy
// for us to directly use them with gorp (our database library).  So
// let's fiddle with them a little bit and make 'em work.

// Client is osin.Client with database type conversion.
type Client struct {
	osin.BasicClient
}

// ToDb returns this Client's Id for storing in the database.
func (client Client) ToDb() interface{} {
	return client.Id
}

// DefaultDbValue returns the default value for Client.Id.
func (client *Client) DefaultDbValue() interface{} {
	return new(string)
}

// FromDb stores the Client.Id value read from the database.
func (client *Client) FromDb(id interface{}) error {
	idPtr := id.(*string)
	client.SetId(*idPtr)
	return nil
}

func (client *Client) MakeEmpty() interface{} {
	return new(Client)
}

// AuthorizeData is osin.AuthorizeData with database type conversion.
// It overrides some of the values from the osin library to ensure
// that the database libraries will handle reading data correctly.
type AuthorizeData struct {
	osin.BasicAuthorizeData

	// Override the pulled in Client attribute with one that is
	// *specifically* our Client type, for database loading purposes.
	Client *Client

	// Add a User ID
	UserId int64
}

func (data *AuthorizeData) SetClient(client osin.Client) {
	data.Client = client.(*Client)
}

func (data *AuthorizeData) GetClient() osin.Client {
	if data.Client == nil {
		return nil
	}
	return data.Client
}

// ToDb returns this AuthorizeData's Code for storing in the database.
func (data *AuthorizeData) ToDb() interface{} {
	return data.Code
}

// DefaultDbValue returns the default value for AuthorizeData.Code.
func (data *AuthorizeData) DefaultDbValue() interface{} {
	return new(string)
}

// FromDb stores the AuthorizeData.Code value read from the database.
func (data *AuthorizeData) FromDb(code interface{}) error {
	codePtr := code.(*string)
	data.SetCode(*codePtr)
	return nil
}

func (data *AuthorizeData) MakeEmpty() interface{} {
	return new(AuthorizeData)
}

// AccessData is osin.AccessData with database type conversion.  It
// overrides some of the values from the osin library to ensure that
// the database libraries will handle reading data correctly.
type AccessData struct {
	osin.BasicAccessData

	Client        *Client
	AuthorizeData *AuthorizeData
	AccessData    *AccessData

	UserId int64
}

func (data *AccessData) SetClient(client osin.Client) {
	if client != nil {
		data.Client = client.(*Client)
	}
}

func (data *AccessData) GetClient() osin.Client {
	if data.Client == nil {
		return nil
	}
	return data.Client
}

func (data *AccessData) SetAuthorizeData(authData osin.AuthorizeData) {
	if authData != nil {
		data.AuthorizeData = authData.(*AuthorizeData)
	}
}

func (data *AccessData) GetAuthorizeData() osin.AuthorizeData {
	if data.AuthorizeData == nil {
		return nil
	}
	return data.AuthorizeData
}

func (data *AccessData) SetAccessData(access osin.AccessData) {
	if access != nil {
		data.AccessData = access.(*AccessData)
	}
}

func (data *AccessData) GetAccessData() osin.AccessData {
	if data.AccessData == nil {
		return nil
	}
	return data.AccessData
}

// ToDb returns this AccessData's AccessToken for storing in the
// database.
func (data AccessData) ToDb() interface{} {
	return data.AccessToken
}

// DefaultDbValue returns the default value for
// AccessData.AccessToken.
func (data *AccessData) DefaultDbValue() interface{} {
	return new(string)
}

// FromDb stores the AccessData.AccessToken from the database.
func (data *AccessData) FromDb(accessToken interface{}) error {
	tokenPtr := accessToken.(*string)
	data.SetAccessToken(*tokenPtr)
	return nil
}

func (data *AccessData) MakeEmpty() interface{} {
	return new(AccessData)
}
