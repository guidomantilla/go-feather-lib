package security

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type jwtTokenManagerOptionsChain struct {
	chain []JwtTokenManagerOptions
}

func JwtTokenManagerOptionsBuilder() *jwtTokenManagerOptionsChain {
	return &jwtTokenManagerOptionsChain{
		chain: make([]JwtTokenManagerOptions, 0),
	}
}

func (options *jwtTokenManagerOptionsChain) Build() JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		for _, option := range options.chain {
			option(tokenManager)
		}
	}
}

func (options *jwtTokenManagerOptionsChain) WithIssuer(issuer string) *jwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithIssuer(issuer))
	return options
}

func (options *jwtTokenManagerOptionsChain) WithTimeout(timeout time.Duration) *jwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithTimeout(timeout))
	return options
}

func (options *jwtTokenManagerOptionsChain) WithSigningMethod(signingMethod jwt.SigningMethod) *jwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithSigningMethod(signingMethod))
	return options
}

func (options *jwtTokenManagerOptionsChain) WithSigningKey(signingKey any) *jwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithSigningKey(signingKey))
	return options
}

func (options *jwtTokenManagerOptionsChain) WithVerifyingKey(verifyingKey any) *jwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithVerifyingKey(verifyingKey))
	return options
}
