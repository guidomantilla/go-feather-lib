package security

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/log"
	"github.com/guidomantilla/go-feather-lib/pkg/validation"
)

type DefaultAuthenticationService struct {
	passwordEncoder  PasswordEncoder
	principalManager PrincipalManager
	tokenManager     TokenManager
}

func NewDefaultAuthenticationService(passwordEncoder PasswordEncoder, principalManager PrincipalManager, tokenManager TokenManager) *DefaultAuthenticationService {

	if passwordEncoder == nil {
		log.Fatal("starting up - error setting up authenticationService: passwordEncoder is nil")
	}

	if principalManager == nil {
		log.Fatal("starting up - error setting up authenticationService: principalManager is nil")
	}

	if tokenManager == nil {
		log.Fatal("starting up - error setting up authenticationService: tokenManager is nil")
	}

	return &DefaultAuthenticationService{
		passwordEncoder:  passwordEncoder,
		principalManager: principalManager,
		tokenManager:     tokenManager,
	}
}

func (service *DefaultAuthenticationService) Authenticate(ctx context.Context, principal *Principal) error {

	var err error
	var user *Principal
	if user, err = service.principalManager.Find(ctx, *principal.Username); err != nil {
		return ErrAuthenticationFailed(err)
	}

	var needsUpgrade *bool
	if needsUpgrade, err = service.passwordEncoder.UpgradeEncoding(*(user.Password)); err != nil || *(needsUpgrade) {
		return ErrAuthenticationFailed(ErrAccountExpiredPassword)
	}

	var matches *bool
	if matches, err = service.passwordEncoder.Matches(*(user.Password), *principal.Password); err != nil || !*(matches) {
		return ErrAuthenticationFailed(ErrAccountInvalidPassword)
	}

	principal.Password, principal.Passphrase = nil, nil
	principal.Role, principal.Resources = user.Role, user.Resources
	if principal.Token, err = service.tokenManager.Generate(principal); err != nil {
		return ErrAuthenticationFailed(err)
	}

	return nil
}

func (service *DefaultAuthenticationService) Validate(principal *Principal) []error {

	var errors []error

	if err := validation.ValidateFieldIsRequired("this", "username", principal.Username); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "role", principal.Role); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldIsRequired("this", "password", principal.Password); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "passphrase", principal.Passphrase); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "enabled", principal.Enabled); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "non_locked", principal.NonLocked); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "non_expired", principal.NonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "password_non_expired", principal.PasswordNonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "signup_done", principal.SignUpDone); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateStructMustBeUndefined("this", "resources", principal.Resources); err != nil {
		errors = append(errors, err)
		return errors
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "token", principal.Token); err != nil {
		errors = append(errors, err)
	}

	return errors
}
