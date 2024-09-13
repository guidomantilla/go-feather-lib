package errors

import (
	"errors"
	"strings"
)

func ErrJoin(errs ...error) error {
	return errors.New(strings.Replace(errors.Join(errs...).Error(), "\n", ": ", -1))
}
