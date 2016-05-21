// The test package contains helpers for tests.
package test

import (
	"database/sql"

	"github.com/Radiobox/web-frontend/datastore"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nelsam/gorp"
	"github.com/nelsam/gorp_queries"
)

func SetupTestingDb() {
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	dbMap := &gorp.DbMap{
		Db:            conn,
		Dialect:       gorp_queries.SqliteDialect{},
		TypeConverter: datastore.SelfConverterTypeConverter{},
	}
	datastore.SetDefaultDbMap(dbMap)
}
