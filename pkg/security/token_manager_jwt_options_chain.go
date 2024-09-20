package security

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JwtTokenManagerOptionChain struct {
	chain []JwtTokenManagerOption
}

func JwtTokenManagerOptionsChainBuilder() *JwtTokenManagerOptionChain {
	return &JwtTokenManagerOptionChain{
		chain: make([]JwtTokenManagerOption, 0),
	}
}

func (options *JwtTokenManagerOptionChain) Build() JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		for _, option := range options.chain {
			option(tokenManager)
		}
	}
}

func (options *JwtTokenManagerOptionChain) WithIssuer(issuer string) *JwtTokenManagerOptionChain {
	options.chain = append(options.chain, jwtTokenManagerOption.WithIssuer(issuer))
	return options
}

func (options *JwtTokenManagerOptionChain) WithTimeout(timeout time.Duration) *JwtTokenManagerOptionChain {
	options.chain = append(options.chain, jwtTokenManagerOption.WithTimeout(timeout))
	return options
}

func (options *JwtTokenManagerOptionChain) WithSigningMethod(signingMethod jwt.SigningMethod) *JwtTokenManagerOptionChain {
	options.chain = append(options.chain, jwtTokenManagerOption.WithSigningMethod(signingMethod))
	return options
}

func (options *JwtTokenManagerOptionChain) WithSigningKey(signingKey any) *JwtTokenManagerOptionChain {
	options.chain = append(options.chain, jwtTokenManagerOption.WithSigningKey(signingKey))
	return options
}

func (options *JwtTokenManagerOptionChain) WithVerifyingKey(verifyingKey any) *JwtTokenManagerOptionChain {
	options.chain = append(options.chain, jwtTokenManagerOption.WithVerifyingKey(verifyingKey))
	return options
}
