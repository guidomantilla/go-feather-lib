package security

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtTokenManagerOptions = NewJwtTokenManagerOptions()

func NewJwtTokenManagerOptions() JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
	}
}

type JwtTokenManagerOptions func(tokenManager TokenManager)

func (option JwtTokenManagerOptions) WithIssuer(issuer string) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("issuer", issuer)
	}
}

func (option JwtTokenManagerOptions) WithTimeout(timeout time.Duration) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("timeout", timeout)
	}
}

func (option JwtTokenManagerOptions) WithSigningMethod(signingMethod jwt.SigningMethod) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("signingMethod", signingMethod)
	}
}

func (option JwtTokenManagerOptions) WithSigningKey(signingKey any) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("signingKey", signingKey)
	}
}

func (option JwtTokenManagerOptions) WithVerifyingKey(verifyingKey any) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("verifyingKey", verifyingKey)
	}
}
