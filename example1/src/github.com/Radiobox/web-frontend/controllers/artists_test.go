package controllers

import (
	"testing"

	goweb_controllers "github.com/stretchr/goweb/controllers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArtistControllerTestSuite struct {
	suite.Suite
}

func TestArtistController(t *testing.T) {
	suite.Run(t, new(ArtistControllerTestSuite))
}

func (suite *ArtistControllerTestSuite) TestImplements() {
	controller := new(ArtistController)
	assert.Implements(suite.T(), (*goweb_controllers.RestfulController)(nil), controller)
	assert.Implements(suite.T(), (*goweb_controllers.RestfulCreator)(nil), controller)
	assert.Implements(suite.T(), (*goweb_controllers.RestfulReader)(nil), controller)
	assert.Implements(suite.T(), (*goweb_controllers.RestfulManyReader)(nil), controller)
	assert.Implements(suite.T(), (*goweb_controllers.RestfulDeletor)(nil), controller)
	assert.Implements(suite.T(), (*goweb_controllers.RestfulManyDeleter)(nil), controller)
	assert.Implements(suite.T(), (*goweb_controllers.RestfulUpdater)(nil), controller)
	assert.Implements(suite.T(), (*goweb_controllers.RestfulReplacer)(nil), controller)
	assert.Implements(suite.T(), (*goweb_controllers.RestfulManyUpdater)(nil), controller)
}
