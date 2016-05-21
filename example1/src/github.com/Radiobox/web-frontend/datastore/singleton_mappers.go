package datastore

import (
	"database/sql"
	"log"
	"reflect"

	"github.com/coopernurse/gorp"
	"github.com/nelsam/gorp_queries"
	query_interfaces "github.com/nelsam/gorp_queries/interfaces"
)

func deRefArgs(args []interface{}) []interface{} {
	for index, arg := range args {
		argVal := reflect.ValueOf(arg)
		for argVal.Kind() == reflect.Ptr {
			argVal = argVal.Elem()
		}
		args[index] = argVal.Interface()
	}
	return args
}

func AddTable(i interface{}) *gorp.TableMap {
	return DefaultDbMap().AddTable(i)
}

func AddTableWithName(i interface{}, name string) *gorp.TableMap {
	return DefaultDbMap().AddTableWithName(i, name)
}

func Begin() (*gorp_queries.Transaction, error) {
	return DefaultDbMap().Begin()
}

func CreateTables() error {
	return DefaultDbMap().CreateTables()
}

func CreateTablesIfNotExists() error {
	return DefaultDbMap().CreateTablesIfNotExists()
}

func Delete(list ...interface{}) (int64, error) {
	return DefaultDbMap().Delete(list...)
}

func DropTables() error {
	return DefaultDbMap().DropTables()
}

func DropTablesIfExists() error {
	return DefaultDbMap().DropTablesIfExists()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return DefaultDbMap().Exec(query, deRefArgs(args)...)
}

func Get(i interface{}, keys ...interface{}) (interface{}, error) {
	return DefaultDbMap().Get(i, deRefArgs(keys)...)
}

func Insert(list ...interface{}) error {
	return DefaultDbMap().Insert(list...)
}

func Query(target interface{}) query_interfaces.Query {
	return DefaultDbMap().Query(target)
}

func Select(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
	return DefaultDbMap().Select(i, query, deRefArgs(args)...)
}

func SelectFloat(query string, args ...interface{}) (float64, error) {
	return DefaultDbMap().SelectFloat(query, deRefArgs(args)...)
}

func SelectInt(query string, args ...interface{}) (int64, error) {
	return DefaultDbMap().SelectInt(query, deRefArgs(args)...)
}

func SelectNullFloat(query string, args ...interface{}) (sql.NullFloat64, error) {
	return DefaultDbMap().SelectNullFloat(query, deRefArgs(args)...)
}

func SelectNullInt(query string, args ...interface{}) (sql.NullInt64, error) {
	return DefaultDbMap().SelectNullInt(query, deRefArgs(args)...)
}

func SelectNullStr(query string, args ...interface{}) (sql.NullString, error) {
	return DefaultDbMap().SelectNullStr(query, deRefArgs(args)...)
}

func SelectStr(query string, args ...interface{}) (string, error) {
	return DefaultDbMap().SelectStr(query, deRefArgs(args)...)
}

func TraceOff() {
	DefaultDbMap().TraceOff()
}

func TraceOn(prefix string, logger *log.Logger) {
	DefaultDbMap().TraceOn(prefix, logger)
}

func TruncateTables() error {
	return DefaultDbMap().TruncateTables()
}

func Update(list ...interface{}) (int64, error) {
	return DefaultDbMap().Update(list...)
}
