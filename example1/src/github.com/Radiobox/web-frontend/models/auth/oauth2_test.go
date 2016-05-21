package auth

import (
	"testing"

	"github.com/Radiobox/osin"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Oauth2TestSuite struct {
	suite.Suite
}

func TestModelSuite(t *testing.T) {
	suite.Run(t, new(Oauth2TestSuite))
}

func (suite *Oauth2TestSuite) TestOauth2_ClientInterfaces() {
	var (
		testPtr interface{} = &Client{}
		ok      bool
	)
	_, ok = testPtr.(osin.Client)
	assert.True(suite.T(), ok, "Client pointers should match the osin.Client interface")
	_, ok = testPtr.(datastore.ConverterToDb)
	assert.True(suite.T(), ok, "Client pointers should match the ConverterToDb interface")
	_, ok = testPtr.(datastore.ConverterFromDb)
	assert.True(suite.T(), ok, "Client pointers should match the ConverterFromDb interface")
}

func (suite *Oauth2TestSuite) TestOauth2_AuthDataInterfaces() {
	var (
		testPtr interface{} = &AuthorizeData{}
		ok      bool
	)
	_, ok = testPtr.(osin.AuthorizeData)
	assert.True(suite.T(), ok, "AuthorizeData pointers should match the osin.AuthorizeData interface")
	_, ok = testPtr.(datastore.ConverterToDb)
	assert.True(suite.T(), ok, "AuthorizeData pointers should match the ConverterToDb interface")
	_, ok = testPtr.(datastore.ConverterFromDb)
	assert.True(suite.T(), ok, "AuthorizeData pointers should match the ConverterFromDb interface")
}

func (suite *Oauth2TestSuite) TestOauth2_AccessDataInterfaces() {
	var (
		testPtr interface{} = &AccessData{}
		ok      bool
	)
	_, ok = testPtr.(osin.AccessData)
	assert.True(suite.T(), ok, "AccessData pointers should match the osin.AccessData interface")
	_, ok = testPtr.(datastore.ConverterToDb)
	assert.True(suite.T(), ok, "AccessData pointers should match the ConverterToDb interface")
	_, ok = testPtr.(datastore.ConverterFromDb)
	assert.True(suite.T(), ok, "AccessData pointers should match the ConverterFromDb interface")
}
