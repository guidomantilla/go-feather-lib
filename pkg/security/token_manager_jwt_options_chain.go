package security

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JwtTokenManagerOptionsChain struct {
	chain []JwtTokenManagerOptions
}

func JwtTokenManagerOptionsChainBuilder() *JwtTokenManagerOptionsChain {
	return &JwtTokenManagerOptionsChain{
		chain: make([]JwtTokenManagerOptions, 0),
	}
}

func (options *JwtTokenManagerOptionsChain) Build() JwtTokenManagerOptions {
	return func(tokenManager TokenManager) {
		for _, option := range options.chain {
			option(tokenManager)
		}
	}
}

func (options *JwtTokenManagerOptionsChain) WithIssuer(issuer string) *JwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithIssuer(issuer))
	return options
}

func (options *JwtTokenManagerOptionsChain) WithTimeout(timeout time.Duration) *JwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithTimeout(timeout))
	return options
}

func (options *JwtTokenManagerOptionsChain) WithSigningMethod(signingMethod jwt.SigningMethod) *JwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithSigningMethod(signingMethod))
	return options
}

func (options *JwtTokenManagerOptionsChain) WithSigningKey(signingKey any) *JwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithSigningKey(signingKey))
	return options
}

func (options *JwtTokenManagerOptionsChain) WithVerifyingKey(verifyingKey any) *JwtTokenManagerOptionsChain {
	options.chain = append(options.chain, jwtTokenManagerOptions.WithVerifyingKey(verifyingKey))
	return options
}
