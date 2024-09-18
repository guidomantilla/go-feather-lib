package security

import (
	"errors"
	"fmt"

	secerrors "github.com/guidomantilla/go-feather-lib/pkg/common/errors"
)

var (
	ErrAccountExistingUsername    = errors.New("principal username already exists")
	ErrAccountInvalidUsername     = errors.New("principal username is invalid")
	ErrAccountEmptyRole           = errors.New("principal role is empty")
	ErrAccountInvalidRole         = errors.New("principal role is invalid")
	ErrAccountEmptyPassword       = errors.New("principal password is empty")
	ErrAccountInvalidPassword     = errors.New("principal password is invalid")
	ErrAccountEmptyPassphrase     = errors.New("principal passphrase is empty")
	ErrAccountInvalidPassphrase   = errors.New("principal passphrase is invalid")
	ErrAccountDisabled            = errors.New("principal is disabled")
	ErrAccountLocked              = errors.New("principal is locked")
	ErrAccountExpired             = errors.New("principal has expired")
	ErrAccountExpiredPassword     = errors.New("principal password has expired")
	ErrAccountEmptyAuthorities    = errors.New("principal authorities are empty")
	ErrAccountInvalidAuthorities  = errors.New("principal authorities are invalid")
	ErrAccountEmptyResource       = errors.New("principal resource is empty")
	ErrTokenFailedParsing         = errors.New("token failed to parse")
	ErrTokenInvalid               = errors.New("token is invalid")
	ErrTokenEmptyClaims           = errors.New("token claims is empty")
	ErrTokenEmptyUsernameClaim    = errors.New("token username claim is empty")
	ErrTokenInvalidUsernameClaim  = errors.New("token username claim is invalid")
	ErrTokenEmptyRoleClaim        = errors.New("token role claim is empty")
	ErrTokenInvalidRoleClaim      = errors.New("token role claim is invalid")
	ErrTokenEmptyResourcesClaim   = errors.New("token resources claim is empty")
	ErrTokenInvalidResourcesClaim = errors.New("token resources claim is invalid")
	ErrPasswordEncoderNotFound    = errors.New("password encoder not found")
	ErrPasswordLength             = errors.New("password length is too short")
	ErrPasswordSpecialChars       = errors.New("password must contain at least 2 special characters")
	ErrPasswordNumbers            = errors.New("password must contain at least 2 numbers")
	ErrPasswordUppercaseChars     = errors.New("password must contain at least 2 uppercase characters")
	ErrRawPasswordIsEmpty         = errors.New("rawPassword cannot be empty")
	ErrSaltIsNil                  = errors.New("salt cannot be nil")
	ErrSaltIsEmpty                = errors.New("salt cannot be empty")
	ErrHashFuncIsNil              = errors.New("hashFunc cannot be nil")
	ErrEncodedPasswordIsEmpty     = errors.New("encodedPassword cannot be empty")
	ErrEncodedPasswordNotAllowed  = errors.New("encodedPassword format not allowed")
	ErrBcryptCostNotAllowed       = errors.New("bcryptCost not allowed")
)

func ErrAuthenticationFailed(errs ...error) error {
	return fmt.Errorf("authentication failed: %s", secerrors.ErrJoin(errs...).Error())
}

func ErrAuthorizationFailed(errs ...error) error {
	return fmt.Errorf("authorization failed: %s", secerrors.ErrJoin(errs...).Error())
}

func ErrTokenGenerationFailed(errs ...error) error {
	return fmt.Errorf("token generation failed: %s", secerrors.ErrJoin(errs...).Error())
}

func ErrTokenValidationFailed(errs ...error) error {
	return fmt.Errorf("token validation failed: %s", secerrors.ErrJoin(errs...).Error())
}

func ErrPasswordEncodingFailed(errs ...error) error {
	return fmt.Errorf("password encoding failed: %s", secerrors.ErrJoin(errs...).Error())
}

func ErrPasswordMatchingFailed(errs ...error) error {
	return fmt.Errorf("password matching failed: %s", secerrors.ErrJoin(errs...).Error())
}

func ErrPasswordUpgradeEncodingValidationFailed(errs ...error) error {
	return fmt.Errorf("password upgrade encoding validation failed: %s", secerrors.ErrJoin(errs...).Error())
}

func ErrPasswordValidationFailed(errs ...error) error {
	return fmt.Errorf("password validation failed: %s", secerrors.ErrJoin(errs...).Error())
}
