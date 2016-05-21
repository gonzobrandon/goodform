package base

import (
	"testing"
	"time"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/test"
	"github.com/Radiobox/web_request_readers"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DbTimeTestSuite struct {
	suite.Suite
}

func TestDbTime(t *testing.T) {
	suite.Run(t, new(DbTimeTestSuite))
}

func (suite *DbTimeTestSuite) TestImplements() {
	dbTime := new(DbTime)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), dbTime)
	assert.Implements(suite.T(), (*web_responders.ResponseObjectCreator)(nil), dbTime)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), dbTime)
}

type TestCreatedUpdatedModel struct {
	CreatedUpdatedModel
	Id int64
}

type CreatedUpdatedTestSuite struct {
	suite.Suite
}

func TestCreatedUpdated(t *testing.T) {
	suite.Run(t, new(CreatedUpdatedTestSuite))
}

// SetupTest changes the default connection to a memory-only sqlite
// connection for testing.
func (suite *CreatedUpdatedTestSuite) SetupTest() {
	test.SetupTestingDb()
	datastore.AddTable(TestCreatedUpdatedModel{}).SetKeys(true, "Id")
	datastore.CreateTablesIfNotExists()
}

// TearDownTest resets the default connection.
func (suite *CreatedUpdatedTestSuite) TearDownTest() {
	datastore.SetDefaultDbMap(nil)
}

func (suite *CreatedUpdatedTestSuite) TestCreatedUpdated_Insert() {
	test := new(TestCreatedUpdatedModel)
	datastore.Insert(test)
	assert.NotNil(suite.T(), test.Created, "Created should be set on insert")
	assert.Equal(suite.T(), test.Created, test.LastUpdated,
		"Created and LastUpdated should be equal after insert")
}

func (suite *CreatedUpdatedTestSuite) TestCreatedUpdated_Update() {
	test := new(TestCreatedUpdatedModel)
	datastore.Insert(test)
	time.Sleep(5 * time.Second)
	assert.NotNil(suite.T(), test.Created)
	datastore.Update(test)
	assert.NotNil(suite.T(), test.LastUpdated)
	assert.NotEqual(suite.T(), test.Created, test.LastUpdated,
		"Created and LastUpdated should be different after update")
}

type JsonTestSuite struct {
	suite.Suite
}

func TestJsonTypes(t *testing.T) {
	suite.Run(t, new(JsonTestSuite))
}

func (suite *JsonTestSuite) TestMapImplements() {
	m := new(JsonMap)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), m)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), m)
	assert.Implements(suite.T(), (*web_responders.ResponseValueCreator)(nil), m)
}

func (suite *JsonTestSuite) TestArrayImplements() {
	a := new(JsonArray)
	assert.Implements(suite.T(), (*datastore.SelfConverter)(nil), a)
	assert.Implements(suite.T(), (*web_request_readers.RequestValueReceiver)(nil), a)
}
