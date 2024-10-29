package security

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/rest"
)

var (
	_ AuthenticationEndpoint = (*DefaultAuthenticationEndpoint)(nil)
	_ AuthenticationService  = (*DefaultAuthenticationService)(nil)
	_ AuthenticationEndpoint = (*MockAuthenticationEndpoint)(nil)
	_ AuthenticationService  = (*MockAuthenticationService)(nil)
)

type AuthenticationEndpoint interface {
	Authenticate(ctx rest.Context)
}

type AuthenticationService interface {
	Authenticate(ctx context.Context, principal *Principal) error
	Validate(principal *Principal) []error
}
