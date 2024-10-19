package security

import (
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtTokenManagerOptions = NewJwtTokenManagerOptions()

func NewJwtTokenManagerOptions() JwtTokenManagerOptions {
	return func(tokenManager *JwtTokenManager) {
	}
}

type JwtTokenManagerOptions func(tokenManager *JwtTokenManager)

func (option JwtTokenManagerOptions) WithIssuer(issuer string) JwtTokenManagerOptions {
	return func(tokenManager *JwtTokenManager) {
		if strings.TrimSpace(issuer) != "" {
			tokenManager.issuer = issuer
		}
	}
}

func (option JwtTokenManagerOptions) WithTimeout(timeout time.Duration) JwtTokenManagerOptions {
	return func(tokenManager *JwtTokenManager) {
		if timeout > 0 {
			tokenManager.timeout = timeout
		}
	}
}

func (option JwtTokenManagerOptions) WithSigningMethod(signingMethod jwt.SigningMethod) JwtTokenManagerOptions {
	return func(tokenManager *JwtTokenManager) {
		if signingMethod != nil {
			tokenManager.signingMethod = signingMethod
		}
	}
}

func (option JwtTokenManagerOptions) WithSigningKey(signingKey any) JwtTokenManagerOptions {
	return func(tokenManager *JwtTokenManager) {
		if signingKey != nil {
			tokenManager.signingKey = signingKey
		}
	}
}

func (option JwtTokenManagerOptions) WithVerifyingKey(verifyingKey any) JwtTokenManagerOptions {
	return func(tokenManager *JwtTokenManager) {
		if verifyingKey != nil {
			tokenManager.verifyingKey = verifyingKey
		}
	}
}
