package validation

import (
	"reflect"
	"strings"
)

// Structs

func ValidateStructMustBeUndefined(structName string, field string, value any) error {
	if !reflect.ValueOf(value).IsNil() {
		return NewErrFieldMustBeUndefined(structName, field)
	}
	return nil
}

func ValidateStructIsRequired(structName string, field string, value any) error {
	if reflect.ValueOf(value).IsNil() {
		return NewErrFieldIsRequired(structName, field)
	}

	reflectedValue := reflect.ValueOf(value)
	if reflect.TypeOf(reflectedValue.Interface()).String() == "*string" {
		str := reflectedValue.Interface().(*string)
		if *(str) == "" || strings.TrimSpace(*(str)) == "" {
			return NewErrFieldIsRequired(structName, field)
		}
	}

	return nil
}

// Fields

type AllowedFieldTypes interface {
	*int64 | *float64 | *string | *bool
}

func ValidateFieldMustBeUndefined[T AllowedFieldTypes](structName string, field string, value T) error {
	if value != nil {
		return NewErrFieldMustBeUndefined(structName, field)
	}
	return nil
}

func ValidateFieldIsRequired[T AllowedFieldTypes](structName string, field string, value T) error {
	if value == nil {
		return NewErrFieldIsRequired(structName, field)
	}

	reflectedValue := reflect.ValueOf(value)
	if reflect.TypeOf(reflectedValue.Interface()).String() == "*string" {
		str := reflectedValue.Interface().(*string)
		if *(str) == "" || strings.TrimSpace(*(str)) == "" {
			return NewErrFieldIsRequired(structName, field)
		}
	}

	return nil
}
