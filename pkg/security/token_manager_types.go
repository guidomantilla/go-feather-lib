package security

var (
	_ TokenManager = (*JwtTokenManager)(nil)
	_ TokenManager = (*MockTokenManager)(nil)
)

type TokenManager interface {
	Generate(principal *Principal) (*string, error)
	Validate(tokenString string) (*Principal, error)
	set(property string, value any)
}
