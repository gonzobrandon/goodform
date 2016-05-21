package base

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/Radiobox/web-frontend/datastore"
)

type NonEmptyString string

func (s NonEmptyString) ToDb() interface{} {
	return string(s)
}

func (s *NonEmptyString) DefaultDbValue() interface{} {
	return new(string)
}

func (s *NonEmptyString) FromDb(value interface{}) error {
	if value == nil {
		*s = "N/A"
	}
	*s = NonEmptyString(*(value.(*string)))
	return nil
}

func (s *NonEmptyString) Receive(value interface{}) error {
	if value == nil {
		return errors.New("Value is null, but must be string")
	}
	valueStr, ok := value.(string)
	if !ok {
		return errors.New("Value is not a string")
	}
	if valueStr == "" {
		return errors.New("Value cannot be empty")
	}
	*s = NonEmptyString(valueStr)
	return nil
}

type Address struct {
	JsonMap
}

func (a *Address) Receive(value interface{}) error {
	if err := a.JsonMap.Receive(value); err != nil {
		return err
	}
	errValues := make([]string, 0, 3)
	if a.Get("city").Str() == "" {
		errValues = append(errValues, "city")
	}
	if a.Get("state_province").Str() == "" {
		errValues = append(errValues, "state or province")
	}
	if a.Get("country").Str() == "" {
		errValues = append(errValues, "country")
	}
	if len(errValues) > 0 {
		return errors.New("Missing value for field(s): " + strings.Join(errValues, ", "))
	}
	return nil
}

type Username string

func (name Username) ToDb() interface{} {
	return string(name)
}

func (name *Username) DefaultDbValue() interface{} {
	return new(string)
}

func (name *Username) FromDb(value interface{}) error {
	newNamePtr := value.(*string)
	*name = Username(*newNamePtr)
	return nil
}

func (name *Username) Receive(value interface{}) error {
	username, ok := value.(string)
	if !ok {
		return errors.New("Username values must be a string type")
	} else if username == "" {
		return errors.New("Username cannot be an empty string")
	}
	*name = Username(username)
	return nil
}

func (name Username) ValidateInput(value interface{}) error {
	namePtr := &name
	if err := namePtr.Receive(value); err != nil {
		return err
	}
	query := "SELECT count(*) FROM usernames WHERE username = $1"
	if count, err := datastore.SelectInt(query, value); err != nil {
		return errors.New("Internal error: could not check for duplicate: " + err.Error())
	} else if count > 0 {
		return errors.New("Duplicate username found")
	}
	return nil
}

// *Very* basic email validation.  This is not intended to be thorough
// - if you want to validate an email address, you should attempt to
// send an email to that address, and the response will tell you
// whether or not the email is valid (and exists).
func ValidateEmail(value interface{}) (string, error) {
	email, ok := value.(string)
	if !ok {
		return "", errors.New("Email values must be a string type")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return "", err
	}
	return email, nil
}
