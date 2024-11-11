package errors

import (
	"errors"
	"strings"
)

func ErrJoin(errs ...error) error {
	var err error
	if err = errors.Join(errs...); err == nil {
		return nil
	}
	return errors.New(strings.Replace(err.Error(), "\n", ": ", -1))
}
