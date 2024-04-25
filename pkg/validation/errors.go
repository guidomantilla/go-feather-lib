package validation

import (
	"fmt"
)

type ErrFieldMustBeUndefined struct {
	structName string
	field      string
}

func NewErrFieldMustBeUndefined(structName string, field string) *ErrFieldMustBeUndefined {
	return &ErrFieldMustBeUndefined{structName: structName, field: field}
}

func (e *ErrFieldMustBeUndefined) Error() string {
	return fmt.Sprintf("%s.%s must be undefined", e.structName, e.field)
}

type ErrFieldIsRequired struct {
	structName string
	field      string
}

func NewErrFieldIsRequired(structName string, field string) *ErrFieldIsRequired {
	return &ErrFieldIsRequired{structName: structName, field: field}
}

func (e *ErrFieldIsRequired) Error() string {
	return fmt.Sprintf("%s.%s is required", e.structName, e.field)
}
