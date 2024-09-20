package security

var (
	_ TokenManager = (*MockTokenManager)(nil)
	_ TokenManager = (*JwtTokenManager)(nil)
)

type TokenManager interface {
	Generate(principal *Principal) (*string, error)
	Validate(tokenString string) (*Principal, error)
}
