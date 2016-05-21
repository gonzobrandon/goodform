package datastore

import (
	"errors"
	"reflect"

	"github.com/coopernurse/gorp"
)

type ConverterToDb interface {
	// ToDb should convert the SelfConverter to its equivalent
	// database value.
	ToDb() interface{}
}

type ConverterFromDb interface {
	// DefaultValue should return a pointer to a value of the same
	// type as the final database value.
	DefaultDbValue() interface{}

	// FromDb should take a database value and convert it to the final
	// Go value, returning an error if conversion fails at any point.
	FromDb(interface{}) error
}

type SelfConverter interface {
	ConverterToDb
	ConverterFromDb
}

// SelfConverterTypeConverter is a gorp.TypeConverter that converts
// any values matching the SelfConverter interface.
type SelfConverterTypeConverter struct {
}

// ToDb will attempt to typecast the passed in value to SelfConverter
// and return the result of calling ToDb() on that value.  If the
// passed in value does not successfully typecast to SelfConverter, it
// will instead just return the passed in value.
func (converter SelfConverterTypeConverter) ToDb(value interface{}) (interface{}, error) {
	refValue := reflect.ValueOf(value)
	if refValue.Kind() == reflect.Ptr && refValue.IsNil() {
		return nil, nil
	}
	if converter, ok := value.(ConverterToDb); ok {
		return converter.ToDb(), nil
	}
	return value, nil
}

// FromDb will attempt to typecast the passed in target to
// SelfConverter and return a CustomScanner that uses the
// SelfConverter's DefaultDbValue and FromDb methods to convert from
// the database value.  If the passed in target does not successfully
// typecast to SelfConverter, it returns an empty scanner and false.
func (converter SelfConverterTypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	scanner := gorp.CustomScanner{}
	convert := true
	targetIsPtr := false

	// target will always be a pointer to the actual target, which
	// itself could be a pointer.
	targetVal := reflect.ValueOf(target)
	if targetVal.Elem().Kind() == reflect.Ptr {
		targetIsPtr = true
		targetVal = targetVal.Elem()
		if targetVal.IsNil() {
			targetVal.Set(reflect.New(targetVal.Type().Elem()))
		}
		target = targetVal.Interface()
	}

	switch converter := target.(type) {
	case ConverterFromDb:
		scanner.Target = converter
		holder := converter.DefaultDbValue()
		holderVal := reflect.ValueOf(holder)
		newHolder := reflect.New(holderVal.Type())
		scanner.Holder = newHolder.Interface()
		scanner.Binder = func(interface{}, interface{}) error {
			holderVal = newHolder.Elem()
			if targetIsPtr {
				// If the DB value is nil, set the value to nil and
				// return; otherwise, initialize the pointer.
				if holderVal.IsNil() {
					if !targetVal.IsNil() {
						targetVal.Set(reflect.Zero(targetVal.Type()))
					}
					return nil
				}
			} else if holderVal.IsNil() {
				return errors.New("Non-pointer types cannot be nil")
			}
			return converter.FromDb(holderVal.Interface())
		}
	default:
		convert = false
	}
	return scanner, convert
}
