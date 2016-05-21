package logs

import (
	"log"

	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/base"
	"github.com/stretchr/objx"
)

type LogMessage struct {
	base.CreatedUpdatedModel
	Id   int64         `db:"log_id" request:"-"`
	Type string        `db:"log_type"`
	Log  *base.JsonMap `db:"log_details"`
}

func New(logType string, details objx.Map) *LogMessage {
	return &LogMessage{
		Type: logType,
		Log:  &base.JsonMap{Map: details},
	}
}

func WriteLog(args objx.Map) {
	logType := args.Get("type").Str()
	delete(args, "type")
	logMessage := New(logType, args)
	if err := datastore.Insert(logMessage); err != nil {
		log.Print("Error while writing log message: " + err.Error())
	}
}
