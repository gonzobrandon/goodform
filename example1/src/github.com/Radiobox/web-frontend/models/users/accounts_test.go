package users

import (
	"testing"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TooManyAccountsTestSuite struct {
	suite.Suite
}

func TestTooManyAccounts(t *testing.T) {
	suite.Run(t, new(TooManyAccountsTestSuite))
}

func (suite *TooManyAccountsTestSuite) TestImplements() {
	err := TooManyAccounts{}
	assert.Implements(suite.T(), (*error)(nil), err)
}

type PasswordTestSuite struct {
	suite.Suite
}

func TestPasswords(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}

func (suite *PasswordTestSuite) TestImplements() {
	password := new(Password)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), password)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), password)
}

type AccountTestSuite struct {
	suite.Suite
}

func TestAccounts(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}

func (suite *AccountTestSuite) TestImplements() {
	account := new(Account)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), account)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), account)
	assert.Implements(suite.T(), (*auth.Authorizer)(nil), account)
	assert.Implements(suite.T(), (*web_responders.RelatedLinker)(nil), account)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), account)
}
