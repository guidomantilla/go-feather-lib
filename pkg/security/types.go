package security

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type ResourceCtxKey struct{}

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

//

var (
	_ PrincipalManager       = (*InMemoryPrincipalManager)(nil)
	_ AuthenticationEndpoint = (*DefaultAuthenticationEndpoint)(nil)
	_ AuthenticationService  = (*DefaultAuthenticationService)(nil)
	_ AuthorizationFilter    = (*DefaultAuthorizationFilter)(nil)
	_ AuthorizationService   = (*DefaultAuthorizationService)(nil)
	_ TokenManager           = (*JwtTokenManager)(nil)
	_ PrincipalManager       = (*MockPrincipalManager)(nil)
	_ AuthenticationEndpoint = (*MockAuthenticationEndpoint)(nil)
	_ AuthenticationService  = (*MockAuthenticationService)(nil)
	_ AuthorizationFilter    = (*MockAuthorizationFilter)(nil)
	_ AuthorizationService   = (*MockAuthorizationService)(nil)
	_ TokenManager           = (*MockTokenManager)(nil)
)

type PrincipalManager interface {
	Create(ctx context.Context, principal *Principal) error
	Update(ctx context.Context, principal *Principal) error
	Delete(ctx context.Context, username string) error
	Find(ctx context.Context, username string) (*Principal, error)
	Exists(ctx context.Context, username string) error

	ChangePassword(ctx context.Context, username string, password string) error
	VerifyResource(ctx context.Context, username string, resource string) error
}

//

type AuthenticationEndpoint interface {
	Authenticate(ctx *gin.Context)
}

type AuthenticationService interface {
	Authenticate(ctx context.Context, principal *Principal) error
	Validate(principal *Principal) []error
}

//

type AuthorizationFilter interface {
	Authorize(ctx *gin.Context)
}

type AuthorizationService interface {
	Authorize(ctx context.Context, tokenString string) (*Principal, error)
}

//

type TokenManager interface {
	Generate(principal *Principal) (*string, error)
	Validate(tokenString string) (*Principal, error)
}

//

const (
	Argon2PrefixKey = "{argon2}"
	BcryptPrefixKey = "{bcrypt}"
	Pbkdf2PrefixKey = "{pbkdf2}"
	ScryptPrefixKey = "{scrypt}"
)

var (
	_ PasswordEncoder   = (*Argon2PasswordEncoder)(nil)
	_ PasswordEncoder   = (*BcryptPasswordEncoder)(nil)
	_ PasswordEncoder   = (*Pbkdf2PasswordEncoder)(nil)
	_ PasswordEncoder   = (*ScryptPasswordEncoder)(nil)
	_ PasswordEncoder   = (*DelegatingPasswordEncoder)(nil)
	_ PasswordEncoder   = (*DefaultPasswordManager)(nil)
	_ PasswordGenerator = (*DefaultPasswordGenerator)(nil)
	_ PasswordGenerator = (*DefaultPasswordManager)(nil)
	_ PasswordManager   = (*DefaultPasswordManager)(nil)
)

type PasswordEncoder interface {
	Encode(rawPassword string) (*string, error)
	Matches(encodedPassword string, rawPassword string) (*bool, error)
	UpgradeEncoding(encodedPassword string) (*bool, error)
}

type PasswordGenerator interface {
	Generate() string
	Validate(rawPassword string) error
}

type PasswordManager interface {
	PasswordEncoder
	PasswordGenerator
}

//

type AuthResource struct {
	Name        *string `gorm:"primaryKey" json:"name,omitempty"`
	Application *string `gorm:"primaryKey" json:"application,omitempty"`
	Enabled     *bool   `gorm:"enabled" json:"enabled,omitempty"`
}

type AuthRole struct {
	Name    *string `gorm:"primaryKey" json:"name,omitempty"`
	Enabled *bool   `gorm:"enabled" json:"enabled,omitempty"`
}

type AuthAccessControlList struct {
	Role       *string `gorm:"primaryKey" json:"role,omitempty"`
	Resource   *string `gorm:"primaryKey" json:"resource,omitempty"`
	Permission *string `gorm:"primaryKey" json:"permission,omitempty"`
	Enabled    *bool   `gorm:"enabled" json:"enabled,omitempty"`
}

type AuthUser struct {
	Username   *string `gorm:"primaryKey" json:"username,omitempty"`
	Role       *string `gorm:"role" json:"role,omitempty"`
	Password   *string `gorm:"password" json:"password,omitempty"`
	Passphrase *string `gorm:"passphrase" json:"passphrase,omitempty"`
	Enabled    *bool   `gorm:"enabled" json:"enabled,omitempty"`
}

type AuthPrincipal struct {
	Username    *string `gorm:"username" json:"username,omitempty"`
	Role        *string `gorm:"role" json:"role,omitempty"`
	Application *string `gorm:"application" json:"application,omitempty"`
	Resource    *string `gorm:"resource" json:"resource,omitempty"`
	Permission  *string `gorm:"permission" json:"permission,omitempty"`
	Password    *string `gorm:"password" json:"password,omitempty"`
	Passphrase  *string `gorm:"passphrase" json:"passphrase,omitempty"`
	Enabled     *bool   `gorm:"enabled" json:"enabled,omitempty"`
}
