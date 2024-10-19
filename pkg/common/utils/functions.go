package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func ToString(value any) string {
	return strings.TrimSpace(fmt.Sprintf("%v", value))
}

//

func IsEmpty(object any) bool {
	objValue := reflect.ValueOf(object)
	if !objValue.IsValid() {
		return true
	}

	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return objValue.Len() == 0
	// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return IsEmpty(deref)
	case reflect.String:
		value := objValue.String()
		return strings.TrimSpace(value) == ""
	// for all other types, compare against the zero value
	// array types are empty when they match their zero-initialized state
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}
