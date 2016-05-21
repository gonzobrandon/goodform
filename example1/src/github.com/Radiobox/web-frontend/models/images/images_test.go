package images

import (
	"testing"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ImageTestSuite struct {
	suite.Suite
}

func TestImages(t *testing.T) {
	suite.Run(t, new(ImageTestSuite))
}

func (suite *ImageTestSuite) TestImplements() {
	image := new(Image)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), image)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), image)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), image)
	assert.Implements(suite.T(), (*web_responders.RelatedLinker)(nil), image)
	assert.Implements(suite.T(), (*web_responders.Locationer)(nil), image)
}
