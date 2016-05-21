package media

import (
	"testing"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type VenueTestSuite struct {
	suite.Suite
}

func TestVenues(t *testing.T) {
	suite.Run(t, new(VenueTestSuite))
}

func (suite *VenueTestSuite) TestImplements() {
	venue := new(Venue)
	assert.Implements(suite.T(), (*web_responders.LazyLoader)(nil), venue)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), venue)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), venue)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), venue)
	assert.Implements(suite.T(), (*auth.Authorizer)(nil), venue)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), venue)
}
