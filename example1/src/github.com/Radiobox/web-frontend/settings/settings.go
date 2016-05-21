// The settings package contains settings for this application.
// Things like the bcrypt cost, database connection strings, etc.  All
// production-specific code should go in the production.go file, all
// development-specific code should go in the development.go file, and
// anything that is universal across development and production should
// go in settings.go.
package settings

import (
	"fmt"
	"os"
	"path"
	"reflect"
)

const (
	// PasswordHashCost is the cost to pass to our hashing libraries
	// when hashing a password.
	PasswordHashCost = 10

	// GoPathEnv is the environment variable that stores the go path.
	GoPathEnv = "GOPATH"

	// ProjectImportPath is the import path to the base package for
	// this project.
	ProjectImportPath = "github.com/Radiobox/web-frontend"

	// FullLinks just stores whether we're returning full
	// protocol://host/path links, or just /path.
	FullLinks = true
)

var (
	// SiteMap is a map of controller names to URLs, for easy remapping.
	SiteMap = map[string]string{
		"forgot-password":        "/api/password-reset",
		"email-verification":     "/api/email-verification",
		"user-accounts":          "/api/user",
		"user-profiles":          "/api/users",
		"users":                  "/api/users",
		"beta-signup":            "/api/beta-signup",
		"artists":                "/api/artists",
		"tracks":                 "/api/tracks",
		"media":                  "/api/media",
		"albums":                 "/api/albums",
		"venues":                 "/api/venues",
		"events":                 "/api/events",
		"stream_name_validation": "/api/stream_name",
		"slugs":                  "/api/slugs",
		"images":                 "/api/images",
		"logs":                   "/api/logs",
	}

	// ProjectPath is the path to this project.
	ProjectPath string
)

func init() {
	goPath := os.Getenv(GoPathEnv)
	if goPath == "" {
		ProjectPath = "."
	} else {
		ProjectPath = path.Join(goPath, "src", ProjectImportPath)
	}
}

// UrlFor returns the url for a given SiteMap name and target value
// (usually an ID).
func UrlFor(mapName string, target interface{}) string {
	targetVal := reflect.ValueOf(target)
	for targetVal.Kind() == reflect.Ptr {
		targetVal = targetVal.Elem()
	}
	target = targetVal.Interface()
	return fmt.Sprintf("%s/%v", SiteMap[mapName], target)
}
