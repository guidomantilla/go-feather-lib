package security

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/rest"
)

var (
	_ AuthorizationFilter  = (*DefaultAuthorizationFilter)(nil)
	_ AuthorizationService = (*DefaultAuthorizationService)(nil)
	_ AuthorizationFilter  = (*MockAuthorizationFilter)(nil)
	_ AuthorizationService = (*MockAuthorizationService)(nil)
)

type ResourceCtxKey struct{}

type AuthorizationFilter interface {
	Authorize(ctx rest.Context)
}

type AuthorizationService interface {
	Authorize(ctx context.Context, tokenString string) (*Principal, error)
}
