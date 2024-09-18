package messaging

import (
	"fmt"

	"github.com/guidomantilla/go-feather-lib/pkg/common/errors"
)

func ErrMessagingFailed(errs ...error) error {
	return fmt.Errorf("messaging failed: %s", errors.ErrJoin(errs...).Error())
}

func ErrMessageDeliveryFailed(errs ...error) error {
	return fmt.Errorf("message delivery failed: %s", errors.ErrJoin(errs...).Error())
}

func ErrMessageHandlingFailed(errs ...error) error {
	return fmt.Errorf("message handling failed: %s", errors.ErrJoin(errs...).Error())
}
