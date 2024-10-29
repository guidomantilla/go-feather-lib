package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func ToString(value any) string {
	return strings.TrimSpace(fmt.Sprintf("%v", value))
}

func ToBool(value any) bool {
	return strings.ToLower(ToString(value)) == "true"
}

func ToType[T any](value any) T {
	objValue := reflect.ValueOf(value)
	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return value.(T)
		// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			var zero T
			return zero
		}
		return value.(T)
		// for all other types, compare against the zero value
		// array types are empty when they match their zero-initialized state
	default:
		return value.(T)
	}
}

//

func IsEmpty(value any) bool {
	objValue := reflect.ValueOf(value)
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
		return reflect.DeepEqual(value, zero.Interface())
	}
}
