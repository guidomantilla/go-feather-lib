package security

import (
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtTokenManagerOption = NewJwtTokenManagerOption()

func NewJwtTokenManagerOption() JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
	}
}

type JwtTokenManagerOption func(tokenManager *JwtTokenManager)

func (option JwtTokenManagerOption) WithIssuer(issuer string) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		if strings.TrimSpace(issuer) != "" {
			tokenManager.issuer = issuer
		}
	}
}

func (option JwtTokenManagerOption) WithTimeout(timeout time.Duration) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		if timeout > 0 {
			tokenManager.timeout = timeout
		}
	}
}

func (option JwtTokenManagerOption) WithSigningMethod(signingMethod jwt.SigningMethod) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		if signingMethod != nil {
			tokenManager.signingMethod = signingMethod
		}
	}
}

func (option JwtTokenManagerOption) WithSigningKey(signingKey any) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		if signingKey != nil {
			tokenManager.signingKey = signingKey
		}
	}
}

func (option JwtTokenManagerOption) WithVerifyingKey(verifyingKey any) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		if verifyingKey != nil {
			tokenManager.verifyingKey = verifyingKey
		}
	}
}
