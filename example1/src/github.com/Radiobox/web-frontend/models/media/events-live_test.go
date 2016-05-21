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

type LiveEventTestSuite struct {
	suite.Suite
}

func TestLiveEvents(t *testing.T) {
	suite.Run(t, new(LiveEventTestSuite))
}

func (suite *LiveEventTestSuite) TestImplements() {
	liveEvent := new(LiveEvent)
	assert.Implements(suite.T(), (*web_responders.LazyLoader)(nil), liveEvent)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), liveEvent)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), liveEvent)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), liveEvent)
	assert.Implements(suite.T(), (*auth.Authorizer)(nil), liveEvent)
	assert.Implements(suite.T(), (*web_responders.RelatedLinker)(nil), liveEvent)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), liveEvent)

	liveEventCollection := new(LiveEventCollection)
	assert.Implements(suite.T(), (*datastore.Queryer)(nil), liveEventCollection)
}
