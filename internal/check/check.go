package check

import (
	"reflect"
)

// IsPointer can be used to check arguments passed as type any are in fact of type pointer
func IsPointer(v any) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Pointer
}

// IsNil can be used to check arguments passed as type any are in fact nil
func IsNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer {
		return false
	}
	return rv.IsNil()
}
