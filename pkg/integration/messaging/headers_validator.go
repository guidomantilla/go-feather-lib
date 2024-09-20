package messaging

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

//

type FunctionAdapterHeadersValidator struct {
	validator HeadersValidatorHandler
}

func NewFunctionAdapterHeadersValidator(validator HeadersValidatorHandler) *FunctionAdapterHeadersValidator {
	assert.NotNil(validator, "integration messaging error - validator is required")
	return &FunctionAdapterHeadersValidator{
		validator: validator,
	}
}

func (validator *FunctionAdapterHeadersValidator) Validate(ctx context.Context, headers Headers) error {
	return validator.validator(ctx, headers)
}
