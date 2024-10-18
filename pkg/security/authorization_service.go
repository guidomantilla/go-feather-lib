package security

import (
	"context"
	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

type DefaultAuthorizationService struct {
	tokenManager     TokenManager
	principalManager PrincipalManager
}

func NewDefaultAuthorizationService(tokenManager TokenManager, principalManager PrincipalManager) *DefaultAuthorizationService {
	assert.NotNil(tokenManager, "starting up - error setting up authorization service: tokenManager is nil")
	assert.NotNil(principalManager, "starting up - error setting up authorization service: principalManager is nil")

	return &DefaultAuthorizationService{
		tokenManager:     tokenManager,
		principalManager: principalManager,
	}
}

func (service *DefaultAuthorizationService) Authorize(ctx context.Context, tokenString string) (*Principal, error) {
	assert.NotNil(ctx, "authorization service - error authorizing: context is nil")
	assert.NotEmpty(tokenString, "authorization service - error authorizing: tokenString is empty")

	var err error
	var principal *Principal
	if principal, err = service.tokenManager.Validate(tokenString); err != nil {
		return nil, err
	}

	var user *Principal
	if user, err = service.principalManager.Find(ctx, *principal.Username); err != nil {
		return nil, ErrAuthorizationFailed(err)
	}

	if *(user.Role) != *(principal.Role) {
		return nil, ErrAuthorizationFailed(ErrAccountInvalidRole)
	}

	var value any
	if value = ctx.Value(ResourceCtxKey{}); value == nil {
		return nil, ErrAuthorizationFailed(ErrAccountEmptyAuthorities)
	}

	var ok bool
	var resource string
	if resource, ok = value.(string); !ok {
		return nil, ErrAuthorizationFailed(ErrAccountEmptyResource)
	}

	if err = service.principalManager.VerifyResource(ctx, *user.Username, resource); err != nil {
		return nil, ErrAuthorizationFailed(err)
	}

	principal.Password, principal.Passphrase, principal.Token = nil, nil, &tokenString
	return principal, nil
}
