// The models/base package has some basic model functionality for
// other models to include when it makes sense.
package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/Radiobox/web_responders"
	"github.com/coopernurse/gorp"
	"github.com/lib/pq"
	"github.com/stretchr/objx"
)

type Duration struct {
	time.Duration
}

// Store the unit types of a typical duration string sequence,
// usually HH:MM:SS.MS.
var unitSequence = []string{
	"h",
	"m",
	"s",
}

func ParseDuration(durationStr string) (Duration, error) {
	currentUnit := 0
	currentValue := ""
	for _, character := range durationStr {
		if character == ':' {
			currentValue += unitSequence[currentUnit]
			currentUnit++
		} else {
			currentValue += string(character)
		}
	}
	currentValue += unitSequence[currentUnit]
	timeDuration, err := time.ParseDuration(currentValue)
	if err != nil {
		return Duration{}, err
	}
	return Duration{timeDuration}, nil
}

func (duration Duration) ToDb() interface{} {
	return duration.String()
}

func (duration *Duration) DefaultDbValue() interface{} {
	return new(string)
}

func (duration *Duration) FromDb(value interface{}) error {
	if value == nil {
		return nil
	}
	durationStr, ok := value.(*string)
	if !ok {
		return errors.New("Cannot parse non-string duration")
	}
	newDuration, err := ParseDuration(*durationStr)
	if err != nil {
		return err
	}
	*duration = newDuration
	return nil
}

type CountryCode int32

func (code CountryCode) ToDb() interface{} {
	if code == 0 {
		return 1
	}
	return code
}

type DefaultTrueBoolean bool

func (boolean DefaultTrueBoolean) DefaultValue() interface{} {
	return true
}

func (boolean DefaultTrueBoolean) ToDb() interface{} {
	return bool(boolean)
}

func (boolean *DefaultTrueBoolean) DefaultDbValue() interface{} {
	return new(bool)
}

func (boolean *DefaultTrueBoolean) FromDb(value interface{}) error {
	b := value.(*bool)
	*boolean = DefaultTrueBoolean(*b)
	return nil
}

// We're using a fake boolean value for most active and verified
// columns, right now.  It always returns true in responses, but is
// always stored to the db as false.
type FakeBoolean struct{}

func (boolean FakeBoolean) ToDb() interface{} {
	return false
}

func (boolean FakeBoolean) DefaultDbValue() interface{} {
	return new(bool)
}

func (boolean FakeBoolean) FromDb(value interface{}) error {
	// Whee.  Database value ignored.
	return nil
}

func (boolean FakeBoolean) ResponseObject() interface{} {
	return true
}

func (boolean FakeBoolean) ResponseValue(options objx.Map) interface{} {
	return true
}

// DbColumns takes a type and a tablename and returns the column names
// to be used in a query for the type, in the style of
// tableName.columnName.
//
// TODO: This is duplicate functionality (and worse performing, at
// that) of parts of our database libraries (gorp).  We should create
// an issue with gorp to see if we can get some of that functionality
// exported, so that we don't have to use this.
func DbColumns(targetType reflect.Type, tablename string) []string {
	// A common scenario is to have a *[]*Model type, so keep digging
	// until we have the actual base model type.
	for targetType.Kind() == reflect.Ptr || targetType.Kind() == reflect.Slice {
		targetType = targetType.Elem()
	}
	columns := make([]string, 0, targetType.NumField())
	for i := 0; i < targetType.NumField(); i++ {
		targetField := targetType.Field(i)
		if targetField.Anonymous {
			embeddedColumns := DbColumns(targetField.Type, tablename)

			// Don't query the same column twice
			for _, embedded := range embeddedColumns {
				shouldAppend := true
				for _, existing := range columns {
					if existing == embedded {
						shouldAppend = false
					}
				}
				if shouldAppend {
					columns = append(columns, embedded)
				}
			}
			continue
		}
		name := targetField.Tag.Get("db")
		switch name {
		case "-":
			continue
		case "":
			name = strings.ToLower(targetField.Name)
			fallthrough
		default:
			name = fmt.Sprintf("%s.%s", tablename, name)

			// Don't query the same column twice
			shouldAppend := true
			for _, column := range columns {
				if column == name {
					shouldAppend = false
				}
			}
			if shouldAppend {
				columns = append(columns, name)
			}
		}
	}
	return columns
}

// A DbTime is a time object that handles conversion to and from the
// database a little better than the standard NullTime.  It's also
// able to receive RFC3339 values from requests.
type DbTime struct {
	pq.NullTime
}

func Now() *DbTime {
	return &DbTime{
		pq.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
}

func (dbTime *DbTime) Receive(data interface{}) error {
	dateString := strings.TrimSpace(data.(string))
	if strings.ContainsRune(dateString, ' ') {
		dateString = strings.Replace(dateString, " ", "T", 1)
	}

	if len(dateString) == len(time.RFC3339)-len("+00:00") {
		dateString += "+00:00"
	}

	if len(dateString) == len(time.RFC3339)-len(":00") {
		dateString += ":00"
	}

	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		log.Print("Time parse error: " + err.Error())
		return err
	}
	dbTime.NullTime.Time = date
	dbTime.NullTime.Valid = true
	return nil
}

func (dbTime *DbTime) ToDb() interface{} {
	if dbTime != nil {
		return dbTime.NullTime
	}
	return nil
}

func (dbTime *DbTime) DefaultDbValue() interface{} {
	return new(pq.NullTime)
}

func (dbTime *DbTime) FromDb(value interface{}) error {
	timeFromDb := value.(*pq.NullTime)
	dbTime.NullTime = *timeFromDb
	return nil
}

func (dbTime *DbTime) ResponseObject() interface{} {
	if dbTime != nil && dbTime.Valid {
		return dbTime.Time.Format(time.RFC3339)
	}
	return nil
}

type Location struct {
	time.Location
}

func (l *Location) Receive(value interface{}) error {
	if value == nil {
		return errors.New("Time zone is null, but must be a string")
	}
	locationStr, ok := value.(string)
	if !ok {
		return errors.New("Time zone must be a string")
	}
	if locationStr == "" {
		return errors.New("Time zone cannot be empty")
	}
	newLoc, err := time.LoadLocation(locationStr)
	if err != nil {
		return err
	}
	l.Location = *newLoc
	return nil
}

func (l *Location) ResponseObject() interface{} {
	return l.ToDb()
}

func (l *Location) ResponseValue(options objx.Map) interface{} {
	return l.ToDb()
}

func (l *Location) ToDb() interface{} {
	if l.String() != "" {
		return l.String()
	}
	return nil
}

func (l *Location) DefaultDbValue() interface{} {
	return new(string)
}

func (l *Location) FromDb(value interface{}) error {
	strPtr := value.(*string)
	return l.Receive(*strPtr)
}

// CreatedUpdatedModel includes fields and logic for storing a
// record's creation and most recent update timestamps.
type CreatedUpdatedModel struct {
	Created     *DbTime `db:"created_at" response:"created_at" request:"-"`
	LastUpdated *DbTime `db:"updated_at" response:"updated_at" request:"-"`
}

// PreInsert is used by gorp, and will be run prior to an insert.
func (model *CreatedUpdatedModel) PreInsert(exec gorp.SqlExecutor) error {
	now := DbTime{pq.NullTime{Time: time.Now(), Valid: true}}
	model.Created = &now
	model.LastUpdated = &now
	return nil
}

// PreUpdate is used by gorp, and will be run prior to an update.
func (model *CreatedUpdatedModel) PreUpdate(exec gorp.SqlExecutor) error {
	now := DbTime{pq.NullTime{Time: time.Now(), Valid: true}}
	model.LastUpdated = &now
	return nil
}

// A JsonMap is a map of key, value pairs that is stored in the
// database as a json datatype.
type JsonMap struct {
	objx.Map
}

func (jsonMap JsonMap) ToDb() interface{} {
	if jsonMap.Map == nil {
		return nil
	}
	output, err := json.Marshal(jsonMap.Map)
	if err != nil {
		panic(err)
	}
	return output
}

func (jsonMap *JsonMap) DefaultDbValue() interface{} {
	return new([]byte)
}

func (jsonMap *JsonMap) FromDb(value interface{}) error {
	if value != nil {
		jsonPtr := value.(*[]byte)
		if *jsonPtr == nil {
			return nil
		}
		return json.Unmarshal(*jsonPtr, &jsonMap.Map)
	}
	return nil
}

func (jsonMap *JsonMap) Receive(value interface{}) error {
	if value != nil {
		jsonMap.Map = objx.Map(value.(objx.Map))
	}
	return nil
}

func (jsonMap JsonMap) ResponseValue(options objx.Map) interface{} {
	return web_responders.CreateResponse(jsonMap.Map, options)
}

// A JsonArray is a list of values that is stored in the database as a
// json datatype.
type JsonArray []string

func (jsonArray *JsonArray) Receive(value interface{}) error {
	switch values := value.(type) {
	case []string:
		*jsonArray = values
	case []interface{}:
		for _, value := range values {
			*jsonArray = append(*jsonArray, value.(string))
		}
	case string:
		*jsonArray = append(*jsonArray, values)
	default:
		return errors.New("Could not parse value to JsonArray")
	}
	return nil
}

func (jsonArray JsonArray) ToDb() interface{} {
	if jsonArray == nil {
		return nil
	}
	output, err := json.Marshal(jsonArray)
	if err != nil {
		panic(err)
	}
	return output
}

func (jsonArray *JsonArray) DefaultDbValue() interface{} {
	return new([]byte)
}

func (jsonArray *JsonArray) FromDb(value interface{}) error {
	if value != nil {
		jsonArrayPtr := value.(*[]byte)
		if *jsonArrayPtr == nil {
			return nil
		}
		return json.Unmarshal(*jsonArrayPtr, jsonArray)
	}
	return nil
}
