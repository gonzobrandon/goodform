// The models package contains structs intended to be used as data
// models.  They should be mapped to database tables and care should
// be taken to keep their data up to date in the database, even if
// it's expected that this application will never read that data.
//
// When you are working with foreign-key relationships, remember that
// queries will usually only populate the Id field of foreign key
// entries.  If you need the full value from the foreign key
// relationship, models *should* (usually) provide a Get<field>
// method, which will run the correct query to retrieve the rest of
// the information for the field before returning it.  If you only
// need the Id field, though, you can access it easily directly
// through the field value.
package models

import (
	"log"
	"reflect"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/images"
	"github.com/Radiobox/web-frontend/models/logs"
	"github.com/Radiobox/web-frontend/models/media"
	"github.com/Radiobox/web-frontend/models/slugs"
	"github.com/Radiobox/web-frontend/models/users"
)

var (
	TableModelMap = [][]interface{}{
		[]interface{}{"users", &users.Account{}},
		[]interface{}{"users", &users.Profile{}},
		[]interface{}{"beta_signups", &users.BetaSignup{}},
		[]interface{}{"artists", &media.Artist{}},
		[]interface{}{"events_live", &media.LiveEvent{}},
		[]interface{}{"venues", &media.Venue{}},
		[]interface{}{"tracks", &media.Track{}},
		[]interface{}{"media", &media.Media{}},
		[]interface{}{"albums", &media.Album{}},
		[]interface{}{"slugs", &slugs.Slug{}},
		[]interface{}{"photos", &images.Image{}},
		[]interface{}{"album_track", &media.AlbumTrack{}},
		[]interface{}{"artist_user", &media.ArtistUser{}},
		[]interface{}{"logs", &logs.LogMessage{}},
		[]interface{}{"event_live_provisioned_broadcasts", &media.Broadcast{}},
		[]interface{}{"event_live_routes", &media.EventRoute{}},
		[]interface{}{"tokens", &users.EmailVerification{}},
		[]interface{}{"tokens", &users.PasswordReset{}},
	}
)

// Map our models to tables.
func MapModels() {

	// These tables all have drastically different primary keys, so
	// it's hard to set them up in a loop.
	datastore.AddTableWithName(auth.Client{}, "oauth_clients").SetKeys(false, "Id")
	datastore.AddTableWithName(auth.AuthorizeData{}, "oauth_auth_data").SetKeys(false, "Code")
	datastore.AddTableWithName(auth.AccessData{}, "oauth_access_data").SetKeys(false, "AccessToken")

	// For the most part, we want to let the code generate the table
	// structure for auth tables.  This won't be the case for most
	// other tables.
	if err := datastore.CreateTablesIfNotExists(); err != nil {
		log.Fatal(err)
	}

	// Now add all the tables that should be created externally.
	for _, mapping := range TableModelMap {
		table := mapping[0].(string)
		modelPtr := mapping[1]
		model := reflect.ValueOf(modelPtr).Elem().Interface()
		datastore.AddTableWithName(model, table).SetKeys(true, "Id")

		if slugType, ok := modelPtr.(slugs.SlugValue); ok {
			slugs.RegisterSlugType(slugType)
		}
	}
}
