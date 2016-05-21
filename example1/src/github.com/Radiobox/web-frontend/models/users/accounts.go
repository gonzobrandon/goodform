// The users package contains all logic related to user accounts,
// profiles, preferences, payments, etc, etc.
package users

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/base"
	"github.com/Radiobox/web-frontend/models/slugs"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/coopernurse/gorp"
	"github.com/stretchr/objx"
)

const (
	FbGraphHome = "https://graph.facebook.com"
)

type TooManyAccounts struct {
	username string
}

func (err TooManyAccounts) Error() string {
	return "More than one account found for username " + err.username
}

type Password string

func (pass *Password) Set(rawPassword string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), settings.PasswordHashCost)
	if err != nil {
		return err
	}
	*pass = Password(passHash)
	return nil
}

func (pass *Password) Receive(newPass interface{}) error {
	if newPass == nil {
		return nil
	}
	newPassString, ok := newPass.(string)
	if !ok {
		return errors.New("Password must be a string type")
	}
	if newPassString == "" {
		return errors.New("Password cannot be empty")
	}
	return pass.Set(newPassString)
}

func (pass *Password) ToDb() interface{} {
	if pass == nil {
		// Store a password meaning that this account should never be
		// logged in to directly.
		return "*"
	}
	return string(*pass)
}

func (pass *Password) DefaultDbValue() interface{} {
	return new(string)
}

func (pass *Password) FromDb(value interface{}) error {
	passHashString := value.(*string)
	if *passHashString == "*" {
		return nil
	}
	*pass = Password(*passHashString)
	return nil
}

func (pass Password) ValidateInput(input interface{}) error {
	inputStr, ok := input.(string)
	if !ok {
		return errors.New("Passwords must be string types")
	}
	if inputStr == "" {
		return errors.New("Passwords cannot be empty")
	}
	return nil
}

type userEmail string

func (email userEmail) ToDb() interface{} {
	return string(email)
}

func (email *userEmail) DefaultDbValue() interface{} {
	return new(string)
}

func (email *userEmail) FromDb(value interface{}) error {
	newEmailPtr := value.(*string)
	*email = userEmail(*newEmailPtr)
	return nil
}

func (email *userEmail) Receive(value interface{}) error {
	newEmail, err := base.ValidateEmail(value)
	if err != nil {
		switch err.Error() {
		case "mail: no address":
			return errors.New("Cannot create an account with an empty email address")
		case "mail: missing phrase":
			fallthrough
		case "mail: missing @ in addr-spec":
			return errors.New("Malformed email address: no @ symbol found")
		case "mail: no domain in addr-spec":
			return errors.New("Malformed email address: no mail domain found")
		default:
			return err
		}
	}
	*email = userEmail(newEmail)
	return nil
}

func (email userEmail) ValidateInput(input interface{}) error {
	emailPtr := &email
	if err := emailPtr.Receive(input); err != nil {
		return err
	}
	query := "SELECT count(*) FROM users WHERE email = $1"
	if count, err := datastore.SelectInt(query, input); err != nil {
		return errors.New("Internal error: could not check for duplicate: " + err.Error())
	} else if count > 0 {
		return errors.New("Duplicate email found")
	}
	return nil
}

type facebookUserId string

func (fbId facebookUserId) ToDb() interface{} {
	return string(fbId)
}

func (fbId *facebookUserId) DefaultDbValue() interface{} {
	return new(string)
}

func (fbId *facebookUserId) FromDb(value interface{}) error {
	idPtr := value.(*string)
	*fbId = facebookUserId(*idPtr)
	return nil
}

func (fbId *facebookUserId) Receive(value interface{}) error {
	if value == nil {
		return nil
	}
	id, ok := value.(string)
	if !ok {
		return errors.New("Facebook ID must be a string type")
	}
	if id == "" {
		return errors.New("Facebook ID cannot be empty")
	}
	*fbId = facebookUserId(id)
	return nil
}

type LoginError struct {
	Message string
}

func (err LoginError) Error() string {
	return err.Message
}

func GetFacebookIdFromToken(token string) (string, error) {
	url, err := url.Parse(FbGraphHome + "/me")
	if err != nil {
		return "", err
	}
	query := url.Query()
	query.Add("access_token", token)
	url.RawQuery = query.Encode()
	response, err := http.Get(url.String())
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	blob := make(objx.Map)
	if err := json.Unmarshal(body, &blob); err != nil {
		return "", err
	}
	if errorBlob := blob.Get("error"); !errorBlob.IsNil() {
		return "", LoginError{errorBlob.MSI()["message"].(string)}
	}
	return blob.Get("id").Str(), nil
}

// Account is for private user details - login information,
// preferences, etc.
type Account struct {
	base.CreatedUpdatedModel

	Id           int64 `db:"user_id" request:"-"`
	Username     base.Username
	Email        userEmail
	FacebookUser *facebookUserId `db:"facebook_user" request:",optional"`
	PassHash     *Password       `db:"pass_hash" response:"-" request:"password,optional"`
}

// New creates a new, empty account.
func New() *Account {
	account := new(Account)
	return account
}

// Login takes a username and password and returns that user's
// account upon successful login.  Upon a failed login, a nil account
// pointer and an error will be returned instead.
func Login(username, password string) (*Account, error) {
	user, err := GetByUsername(username)
	if user == nil || err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PassHash), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}

func FacebookLogin(facebookId, token string) (*Account, error) {
	user, err := GetByFacebookId(facebookId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, LoginError{"No account found for that facebook ID"}
	}
	id, err := GetFacebookIdFromToken(token)
	if id != facebookId {
		return nil, LoginError{"Facebook ID does not match access token"}
	}
	// TODO: kick off thread to get back long term token
	return user, nil
}

func GetByFacebookId(id string) (*Account, error) {
	emptyAccount := new(Account)
	results, err := datastore.Query(emptyAccount).
		Where().
		Equal(&emptyAccount.FacebookUser, id).
		Select()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	return results[0].(*Account), nil
}

func GetByEmail(email string) (*Account, error) {
	emptyAccount := new(Account)
	results, err := datastore.Query(emptyAccount).
		Where().
		Equal(&emptyAccount.Email, email).
		Select()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	return results[0].(*Account), nil
}

func GetByUsername(username string) (*Account, error) {
	emptyAccount := new(Account)
	results, err := datastore.Query(emptyAccount).
		Where().
		Equal(&emptyAccount.Username, username).
		Select()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	return results[0].(*Account), nil
}

func (user *Account) RelatedLinks() map[string]string {
	links := make(map[string]string)
	links["profile"] = user.Profile().Location()
	return links
}

func (user *Account) Location() string {
	return settings.UrlFor("user-accounts", user.Id)
}

func (user *Account) Profile() *Profile {
	result, err := datastore.Get(new(Profile), user.Id)
	if err != nil {
		log.Print("Error from db: ", err)
	}
	if err != nil || result == nil {
		return nil
	}
	return result.(*Profile)
}

func (user *Account) Authorize(access *auth.AccessData, requestType auth.RequestType) error {
	switch requestType {
	case auth.REQUEST_CREATE:
	case auth.REQUEST_READ_MANY:
		return errors.New("Cannot list accounts, try " + settings.SiteMap["user-profiles"])
	case auth.REQUEST_UPDATE, auth.REQUEST_READ_ONE:
		if access == nil || access.UserId != user.Id {
			return errors.New("Only users can view/update their private data")
		}
	default:
		return errors.New("Invalid access method")
	}
	return nil
}

func (user *Account) PreInsert(exec gorp.SqlExecutor) error {
	if err := user.CreatedUpdatedModel.PreInsert(exec); err != nil {
		return err
	}
	if user.PassHash == nil && user.FacebookUser == nil {
		return errors.New("New accounts must have either a password or a facebook ID")
	}
	return nil
}

func (user *Account) PreUpdate(exec gorp.SqlExecutor) error {
	err := user.CreatedUpdatedModel.PreUpdate(exec)
	if err != nil {
		return err
	}
	query := "UPDATE slugs SET slug = $1 " +
		"FROM users " +
		"WHERE users.user_id = (slugs.target->>'user_id')::integer " +
		"AND users.username = slugs.slug " +
		"AND users.user_id = $2"
	_, err = exec.Exec(query, user.Username, user.Id)
	return err
}

func (user *Account) PostInsert(exec gorp.SqlExecutor) error {
	slug := slugs.New(&Profile{Id: user.Id}, string(user.Username))
	if err := exec.Insert(slug); err != nil {
		return nil
	}
	verify := &EmailVerification{Token{Id: new(string), User: user}}
	if err := exec.Insert(verify); err != nil {
		return err
	}
	return verify.SendVerification()
}

func (user *Account) ResponseValue(options objx.Map) interface{} {
	return map[string]string{
		"username":     string(user.Username),
		"email":        string(user.Email),
		"profile-link": settings.UrlFor("user-profiles", user.Id),
	}
}

func (user *Account) SetPassword(rawPassword string) error {
	return user.PassHash.Receive(rawPassword)
}

func (user *Account) ToDb() interface{} {
	return user.Id
}

func (user *Account) DefaultDbValue() interface{} {
	return new(int64)
}

func (user *Account) FromDb(value interface{}) error {
	idPtr, ok := value.(*int64)
	if !ok {
		return errors.New("Could not parse non-int ID for user account")
	}
	id := *idPtr
	user.Id = id
	return nil
}
