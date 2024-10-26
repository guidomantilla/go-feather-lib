package mongo

import (
	"errors"
	"fmt"
)

func ErrDBConnectionFailed(errs ...error) error {
	return fmt.Errorf("db connection failed: %s", errors.Join(errs...).Error())
}
