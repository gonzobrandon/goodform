package slugs

import (
	"testing"

	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SlugTestSuite struct {
	suite.Suite
}

func TestSlugs(t *testing.T) {
	suite.Run(t, new(SlugTestSuite))
}

// TODO: Make Slug a SelfConverter and add slug_id to database values
// that need slugs.
func (suite *SlugTestSuite) TestImplements() {
	slug := new(Slug)
	assert.Implements(suite.T(), (*web_responders.LazyLoader)(nil), slug)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), slug)
	assert.Implements(suite.T(), (*auth.Authorizer)(nil), slug)
	assert.Implements(suite.T(), (*web_responders.RelatedLinker)(nil), slug)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), slug)
	//	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), slug)
}
