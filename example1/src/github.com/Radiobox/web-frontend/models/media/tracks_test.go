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

type TrackTestSuite struct {
	suite.Suite
}

func TestTracks(t *testing.T) {
	suite.Run(t, new(TrackTestSuite))
}

func (suite *TrackTestSuite) TestImplements() {
	track := new(Track)
	assert.Implements(suite.T(), (*web_responders.LazyLoader)(nil), track)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), track)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), track)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), track)
	assert.Implements(suite.T(), (*auth.Authorizer)(nil), track)
	assert.Implements(suite.T(), (*web_responders.RelatedLinker)(nil), track)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), track)

	trackCollection := new(TrackCollection)
	assert.Implements(suite.T(), (*datastore.Queryer)(nil), trackCollection)
}
