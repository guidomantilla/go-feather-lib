package security

import (
	"context"
	"net/http"
)

var (
	_ AuthenticationEndpoint = (*DefaultAuthenticationEndpoint)(nil)
	_ AuthenticationService  = (*DefaultAuthenticationService)(nil)
	_ AuthenticationEndpoint = (*MockAuthenticationEndpoint)(nil)
	_ AuthenticationService  = (*MockAuthenticationService)(nil)
)

type AuthenticationEndpoint interface {
	Authenticate(response http.ResponseWriter, request *http.Request)
}

type AuthenticationService interface {
	Authenticate(ctx context.Context, principal *Principal) error
	Validate(principal *Principal) []error
}
