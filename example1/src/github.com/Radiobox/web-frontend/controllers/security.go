package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/Radiobox/osin"
	"github.com/Radiobox/web-frontend/controllers/util"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/datastore/oauth2"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/users"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
)

var (
	oauth2Server *osin.Server
)

func DefaultOauth2Config() *osin.ServerConfig {
	config := osin.NewServerConfig()
	config.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	config.AllowedAccessTypes = osin.AllowedAccessType{
		osin.AUTHORIZATION_CODE,
		osin.REFRESH_TOKEN,
		osin.PASSWORD,
		osin.CLIENT_CREDENTIALS,
		osin.FB_TOKEN,
	}
	return config
}

func DefaultOauth2Server() *osin.Server {
	if oauth2Server == nil {
		log.Print("Creating default oauth2 server...")
		oauth2Server = osin.NewServer(DefaultOauth2Config(), oauth2.DefaultOauth2Storage())
		log.Print("Default oauth2 server created.")
	}
	return oauth2Server
}

// Authorize is for token and code authorization requests, as defined
// by oauth2.  We currently don't use it, so honestly, it probably
// doesn't work too well.
func Authorize(ctx context.Context) error {
	params, err := web_request_readers.ParseParams(ctx)
	if err != nil {
		return err
	}
	server := DefaultOauth2Server()
	var userId int64
	authRequest, err := server.HandleAuthorizeRequest(params)
	if err != nil {
		return err
	}
	userId = HandleLoginPage(authRequest, ctx)
	if userId == -1 {
		// The user hasn't logged in yet.
		return nil
	}
	authRequest.Authorized = true

	var target interface{}
	responseType := ctx.FormValues("response_type")[0]
	switch osin.AuthorizeRequestType(responseType) {
	case osin.CODE:
		target = &auth.AuthorizeData{UserId: userId}
	case osin.TOKEN:
		target = &auth.AccessData{UserId: userId}
	default:
		errMessage := fmt.Sprintf("Value for response_type [%s] not understood", responseType)
		log.Print(errMessage)
		return errors.New(errMessage)
	}
	redirect, err := server.FinishAuthorizeRequest(params, authRequest, target)
	if err != nil {
		return err
	}
	ctx.HttpResponseWriter().Header().Add("Location", redirect)
	ctx.HttpResponseWriter().WriteHeader(http.StatusFound)
	return nil
}

func passwordAuth(accessRequest *osin.AccessRequest, access *auth.AccessData, params objx.Map) error {
	username, password := params.Get("username").Str(), params.Get("password").Str()
	user, err := users.Login(username, password)
	if user == nil || err != nil {
		message := "Login failed: "
		if _, ok := err.(users.TooManyAccounts); ok {
			message += err.Error()
		} else {
			if err != nil {
				log.Print(err)
			}
			// We don't want to tell them which part of the
			// request was bad, in case someone's trying to
			// brute force usernames and passwords.
			message += "bad username or password"
		}
		return users.LoginError{message}
	}
	access.UserId = user.Id
	accessRequest.Authorized = true
	return nil
}

func facebookAuth(accessRequest *osin.AccessRequest, access *auth.AccessData, params objx.Map) error {
	id, token := params.Get("facebook_id").Str(), params.Get("facebook_token").Str()
	if id == "" || token == "" {
		err := errors.New("Cannot log in with empty facebook id or token")
		return err
	}
	user, err := users.FacebookLogin(id, token)
	if err != nil {
		return err
	}
	access.UserId = user.Id
	accessRequest.Authorized = true
	return nil
}

func defaultAuth(accessRequest *osin.AccessRequest, access *auth.AccessData, params objx.Map) error {
	accessRequest.Authorized = true
	return nil
}

func Token(ctx context.Context) error {
	messages := web_responders.NewMessageMap()
	params, parseErr := web_request_readers.ParseParams(ctx)
	if parseErr != nil {
		messages.AddErrorMessage("Could not parse parameters: " + parseErr.Error())
		return web_responders.Respond(ctx, http.StatusBadRequest, messages, parseErr, settings.FullLinks)
	}
	server := DefaultOauth2Server()
	accessRequest, err := server.HandleAccessRequest(ctx.HttpRequest(), params)
	if err != nil {
		messages.AddErrorMessage(err.Message)
		return web_responders.Respond(ctx, err.Status, messages, err, settings.FullLinks)
	}
	access := new(auth.AccessData)
	grantType := params.Get("grant_type").Str()
	var authFunc func(*osin.AccessRequest, *auth.AccessData, objx.Map) error
	switch grantType {
	case "password":
		authFunc = passwordAuth
	case "facebook":
		authFunc = facebookAuth
	default:
		authFunc = defaultAuth
	}
	if err := authFunc(accessRequest, access, params); err != nil {
		status := http.StatusInternalServerError
		if _, ok := err.(users.LoginError); ok {
			status = http.StatusUnauthorized
		}
		messages.AddErrorMessage(err.Error())
		return web_responders.Respond(ctx, status, messages, err, settings.FullLinks)
	}
	response, err := server.FinishAccessRequest(params, accessRequest, access)
	if err != nil {
		messages.AddErrorMessage(err.Message)
		return web_responders.Respond(ctx, err.Status, messages, err, settings.FullLinks)
	}
	if access.GetAccessData() != nil {
		access.UserId = access.GetAccessData().(*auth.AccessData).UserId
		datastore.Update(access)
	}
	if access.UserId > 0 {
		response.Set("user_account", fmt.Sprintf("%s/%d", settings.SiteMap["user-accounts"], access.UserId))
		response.Set("user_profile", fmt.Sprintf("%s/%d", settings.SiteMap["user-profiles"], access.UserId))
		response.Set("user_id", access.UserId)
	}
	return web_responders.Respond(ctx, http.StatusOK, messages, response, settings.FullLinks)
}

// Test login page.  Returns the user's id if login is successful, -1
// otherwise.
func HandleLoginPage(ar *osin.AuthorizeRequest, ctx context.Context) int64 {
	r := ctx.HttpRequest()
	w := ctx.HttpResponseWriter()
	var message string
	if r.Method == "POST" {
		params, err := web_request_readers.ParseParams(ctx)
		if err != nil {
			return -1
		}
		user, err := users.Login(params.Get("username").Str(), params.Get("password").Str())
		if err != nil {
			message = "Login Failed: " + err.Error()
		} else if user != nil {
			return user.Id
		}
	}

	loginTemplate, err := template.ParseFiles("./public/src/partials/go_login.html")
	if err != nil {
		panic(err)
	}
	context := map[string]interface{}{
		"form_action": fmt.Sprintf("/api/authorize?response_type=%s&client_id=%s&state=%s&redirect_uri=%s",
			ar.Type, ar.Client.GetId(), ar.State, url.QueryEscape(ar.RedirectUri)),
		"client_id": ar.Client.GetId(),
		"message":   message,
	}
	err = loginTemplate.Execute(w, context)
	if err != nil {
		panic(err)
	}
	return -1
}

// Test oauth2 authentication on a REST request.
type Oauth2TestController struct {
}

func (controller *Oauth2TestController) Path() string {
	return "/api/secure_test"
}

// Use a bearer auth type before returning data.
func (controller *Oauth2TestController) ReadMany(ctx context.Context) error {
	response := map[string]interface{}{
		"success": false,
	}

	if _, err := util.Authorize(ctx); err != nil {
		response["error"] = err.Error()
		return goweb.API.WriteResponseObject(ctx, http.StatusBadRequest, response)
	}

	response["success"] = true
	response["data"] = "I am some data that you are authorized to see."
	return goweb.API.WriteResponseObject(ctx, http.StatusOK, response)
}

func DisplayLoginOptions(ctx context.Context) error {
	docs := "The /api/authorize and /api/token " +
		"endpoints are used for oauth2 logins.  The authorize endpoint is used for " +
		"authorizing a user's login, while the access endpoint is used for " +
		"retrieving an access token (with one exception)." +
		"\n\n" +
		"The authorize endpoint expects you to redirect users to it as a GET " +
		"request and expects the following parameters: client_id (the ID of the " +
		"client application) and response_type.  Once the user has successfully " +
		"logged in, they will be redirected back to the web page that was specified " +
		"at client signup.  If you redirect to this endpoint using " +
		"response_type=" + string(osin.CODE) + ", then the redirect back to the client site " +
		"will include a code that must be sent to the /api/token endpoint in a " +
		"POST request to retrieve an access token.  If you are using response_type=" +
		string(osin.TOKEN) + ", then the redirect back to the client site will include an " +
		"access token which can be used immediately for making secure requests." +
		"\n\n" +
		"The token endpoint, on the other hand, expects a POST request with a " +
		"grant_type parameter.  If your grant_type is " + string(osin.AUTHORIZATION_CODE) +
		", " + string(osin.REFRESH_TOKEN) + ", or " + string(osin.CLIENT_CREDENTIALS) + ", then you should " +
		"include an Authorization header of 'Basic: encoded_credentials', where " +
		"encoded_credentials is a base64-encoded client_id:client_secret string." +
		"\n\n" +
		"There is a special case for requests to /api/token with a grant_type of " +
		string(osin.PASSWORD) + ", where you skip the /api/authorize endpoint and directly " +
		"send the username and password of the user as POST parameters.  This is " +
		"useful for situations where a redirect is difficult or impossible.  If " +
		"you use this method, the request expects parameters of: client_id, " +
		"username, password, and grant_type=password.  The response will include " +
		"an access token upon successful login."
	return web_responders.Respond(ctx, http.StatusOK, nil, docs, settings.FullLinks)
}
