// The datastore package provides to functions to enforce a
// singleton-style database connection.  It may also extend some base
// functionality from the database libraries that we're using.
//
// All functions in this package are just built around gorp's methods,
// being used on a singleton gorp.DbMap value.  See
// github.com/nelsam/gorp_queries for documentation of those functions.
package datastore

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"os"

	"github.com/Radiobox/web-frontend/settings"
	"github.com/coopernurse/gorp"
	"github.com/nelsam/gorp_queries"
)

var (
	// defaultConnection should store the connection *only* if the
	// default connection is being used.  If a connection has been
	// provided, this should be nil.  This is to ensure that we don't
	// close any connections that weren't created in this package.
	defaultConnection *sql.DB
	dbMap             *gorp_queries.DbMap
)

// DefaultDbMap returns the current gorp singleton.
func DefaultDbMap() *gorp_queries.DbMap {
	if dbMap == nil {
		log.Print("Default DB Map is nil, initializing...")
		err := initializeDefaultDbMap()
		if err != nil {
			panic(err)
		}
		log.Print("DB Map initialized.")
	}
	return dbMap
}

// SetDefaultDbMap changes the current gorp singleton to
// another connection.  This is especially useful if you don't want to
// use postgres (our current default) as your database engine.
func SetDefaultDbMap(newDbMap *gorp_queries.DbMap) {
	if defaultConnection != nil {
		// If the default connection was being used, close it;
		// otherwise, leave the dbMap.Db connection open, since it was
		// probably created by a user.
		defaultConnection.Close()
		defaultConnection = nil
	}
	dbMap = newDbMap
}

// initializeDefaultDbmap initializes a connection using the
// default settings.
func initializeDefaultDbMap() error {
	defaultConnection, err := sql.Open("postgres", settings.DbConnectionString)
	if err != nil {
		return err
	}
	dbMap = &gorp_queries.DbMap{gorp.DbMap{
		Db:            defaultConnection,
		Dialect:       gorp.PostgresDialect{},
		TypeConverter: SelfConverterTypeConverter{},
	}}
	dbMap.TraceOn("DB: ", log.New(os.Stdout, "", log.LstdFlags))
	return nil
}
