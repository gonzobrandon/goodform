package users

import (
	"testing"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/slugs"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProfileTestSuite struct {
	suite.Suite
}

func TestProfiles(t *testing.T) {
	suite.Run(t, new(ProfileTestSuite))
}

func (suite *ProfileTestSuite) TestImplements() {
	profile := new(Profile)
	assert.Implements(suite.T(), (*slugs.SlugValue)(nil), profile)
	assert.Implements(suite.T(), (*web_responders.LazyLoader)(nil), profile)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), profile)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), profile)
	assert.Implements(suite.T(), (*datastore.Queryer)(nil), profile)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), profile)
	assert.Implements(suite.T(), (*auth.Authorizer)(nil), profile)
	assert.Implements(suite.T(), (*web_responders.RelatedLinker)(nil), profile)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), profile)
}
