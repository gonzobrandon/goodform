package media

import (
	"testing"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MediaTestSuite struct {
	suite.Suite
}

func TestMedia(t *testing.T) {
	suite.Run(t, new(MediaTestSuite))
}

func (suite *MediaTestSuite) TestImplements() {
	media := new(Media)
	assert.Implements(suite.T(), (*web_responders.LazyLoader)(nil), media)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), media)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), media)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), media)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), media)
}
