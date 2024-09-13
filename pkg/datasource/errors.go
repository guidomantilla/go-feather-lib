package datasource

import (
	"errors"
)

func ErrDBConnectionFailed(errs ...error) error {
	return errors.New("db connection failed: " + errors.Join(errs...).Error())
}
