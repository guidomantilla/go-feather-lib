package security

import "context"

var (
	_ PrincipalManager = (*BasePrincipalManager)(nil)
	_ PrincipalManager = (*GormPrincipalManager)(nil)
	_ PrincipalManager = (*MockPrincipalManager)(nil)
)

type Principal struct {
	Username           *string  `json:"username,omitempty" binding:"required"`
	Role               *string  `json:"role,omitempty"`
	Password           *string  `json:"password,omitempty" binding:"required"`
	Passphrase         *string  `json:"passphrase,omitempty" `
	Enabled            *bool    `json:"enabled,omitempty"`
	NonLocked          *bool    `json:"non_locked,omitempty"`
	NonExpired         *bool    `json:"non_expired,omitempty"`
	PasswordNonExpired *bool    `json:"password_non_expired,omitempty"`
	SignUpDone         *bool    `json:"signup_done,omitempty"`
	Resources          []string `json:"resources,omitempty"`
	Token              *string  `json:"token,omitempty"`
}

type PrincipalManager interface {
	Create(ctx context.Context, principal *Principal) error
	Update(ctx context.Context, principal *Principal) error
	Delete(ctx context.Context, username string) error
	Find(ctx context.Context, username string) (*Principal, error)
	Exists(ctx context.Context, username string) error

	ChangePassword(ctx context.Context, username string, password string) error
	VerifyResource(ctx context.Context, username string, resource string) error
}
