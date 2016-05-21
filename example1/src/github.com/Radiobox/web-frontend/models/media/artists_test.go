package media

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

type ArtistTestSuite struct {
	suite.Suite
}

func TestArtists(t *testing.T) {
	suite.Run(t, new(ArtistTestSuite))
}

func (suite *ArtistTestSuite) TestImplements() {
	artist := new(Artist)
	assert.Implements(suite.T(), (*slugs.SlugValue)(nil), artist)
	assert.Implements(suite.T(), (*web_responders.LazyLoader)(nil), artist)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), artist)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), artist)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), artist)
	assert.Implements(suite.T(), (*auth.Authorizer)(nil), artist)
	assert.Implements(suite.T(), (*web_responders.RelatedLinker)(nil), artist)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), artist)

	artistCollection := new(ArtistCollection)
	assert.Implements(suite.T(), (*datastore.Queryer)(nil), artistCollection)
}
