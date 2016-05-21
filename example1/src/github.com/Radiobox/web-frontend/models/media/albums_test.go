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

type AlbumTestSuite struct {
	suite.Suite
}

func TestAlbums(t *testing.T) {
	suite.Run(t, new(AlbumTestSuite))
}

func (suite *AlbumTestSuite) TestImplements() {
	album := new(Album)
	assert.Implements(suite.T(), (*web_responders.LazyLoader)(nil), album)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), album)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), album)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), album)
	assert.Implements(suite.T(), (*auth.Authorizer)(nil), album)
	assert.Implements(suite.T(), (*web_responders.RelatedLinker)(nil), album)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), album)

	albumCollection := new(AlbumCollection)
	assert.Implements(suite.T(), (*datastore.Queryer)(nil), albumCollection)
}
