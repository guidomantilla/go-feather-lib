package assert

import (
	"log"
	"reflect"

	"github.com/guidomantilla/go-feather-lib/pkg/common/utils"
)

func NotEmpty(object any, message string) {
	if utils.IsEmpty(object) {
		log.Fatal(message)
	}
}

func NotNil(object any, message string) {
	if isNil(object) {
		log.Fatal(message)
	}
}

func Equal(val1 any, val2 any, message string) {
	if !isEqual(val1, val2) {
		log.Fatal(message)
	}
}

func NotEqual(val1 any, val2 any, message string) {
	if isEqual(val1, val2) {
		log.Fatal(message)
	}
}

func True(condition bool, message string) {
	if !condition {
		log.Fatal(message)
	}
}

func False(condition bool, message string) {
	if condition {
		log.Fatal(message)
	}
}

//

func nil(object any) bool {
	return isNil(object)
}

//

func isNil(object any) bool {
	value := reflect.ValueOf(object)
	if !value.IsValid() {
		return true
	}
	switch value.Kind() {
	case
		reflect.Chan, reflect.Func,
		reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:

		return value.IsNil()
	}

	return false
}

func isEqual(val1, val2 any) bool {
	v1 := reflect.ValueOf(val1)
	v2 := reflect.ValueOf(val2)

	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}

	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}

	if !v1.IsValid() && !v2.IsValid() {
		return true
	}

	switch v1.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v1.IsNil() {
			v1 = reflect.ValueOf(nil)
		}
	}

	switch v2.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v2.IsNil() {
			v2 = reflect.ValueOf(nil)
		}
	}

	v1Underlying := reflect.Zero(reflect.TypeOf(v1)).Interface()
	v2Underlying := reflect.Zero(reflect.TypeOf(v2)).Interface()

	if v1 == v1Underlying {
		if v2 == v2Underlying {
			goto CASE4
		} else {
			goto CASE3
		}
	} else {
		if v2 == v2Underlying {
			goto CASE2
		} else {
			goto CASE1
		}
	}

CASE1:
	return reflect.DeepEqual(v1.Interface(), v2.Interface())
CASE2:
	return reflect.DeepEqual(v1.Interface(), v2)
CASE3:
	return reflect.DeepEqual(v1, v2.Interface())
CASE4:
	return reflect.DeepEqual(v1, v2)
}
