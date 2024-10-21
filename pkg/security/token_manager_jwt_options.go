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

func (options JwtTokenManagerOptions) WithIssuer(issuer string) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("issuer", issuer)
	}
}

func (options JwtTokenManagerOptions) WithTimeout(timeout time.Duration) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("timeout", timeout)
	}
}

func (options JwtTokenManagerOptions) WithSigningMethod(signingMethod jwt.SigningMethod) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("signingMethod", signingMethod)
	}
}

func (options JwtTokenManagerOptions) WithSigningKey(signingKey any) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("signingKey", signingKey)
	}
}

func (options JwtTokenManagerOptions) WithVerifyingKey(verifyingKey any) JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		tokenManager.set("verifyingKey", verifyingKey)
	}
}
