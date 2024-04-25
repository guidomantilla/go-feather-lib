package security

import (
	"encoding/json"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/guidomantilla/go-feather-lib/pkg/util"
)

type DefaultClaims struct {
	jwt.RegisteredClaims
	Principal
}

type JwtTokenManagerOption func(tokenManager *JwtTokenManager)

type JwtTokenManager struct {
	issuer        string
	timeout       time.Duration
	signingKey    any
	verifyingKey  any
	signingMethod jwt.SigningMethod
}

func NewJwtTokenManager(options ...JwtTokenManagerOption) *JwtTokenManager {

	tokenManager := &JwtTokenManager{
		issuer:        "",
		timeout:       time.Hour * 24,
		signingKey:    "some_long_signing_key",
		verifyingKey:  "some_long_verifying_key",
		signingMethod: jwt.SigningMethodHS512,
	}

	for _, opt := range options {
		opt(tokenManager)
	}

	return tokenManager
}

func WithIssuer(issuer string) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		tokenManager.issuer = issuer
	}
}

func WithTimeout(timeout time.Duration) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		tokenManager.timeout = timeout
	}
}

func WithSigningMethod(signingMethod jwt.SigningMethod) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		tokenManager.signingMethod = signingMethod
	}
}

func WithSigningKey(signingKey any) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		tokenManager.signingKey = signingKey
	}
}

func WithVerifyingKey(verifyingKey any) JwtTokenManagerOption {
	return func(tokenManager *JwtTokenManager) {
		tokenManager.verifyingKey = verifyingKey
	}
}

func (manager *JwtTokenManager) Generate(principal *Principal) (*string, error) {

	claims := &DefaultClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    manager.issuer,
			Subject:   *principal.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.timeout)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Principal: *principal,
	}

	token := jwt.NewWithClaims(manager.signingMethod, claims)

	var err error
	var tokenString string
	if tokenString, err = token.SignedString(manager.signingKey); err != nil {
		return nil, ErrTokenGenerationFailed(err)
	}

	return &tokenString, nil
}

func (manager *JwtTokenManager) Validate(tokenString string) (*Principal, error) {

	getKeyFunc := func(token *jwt.Token) (any, error) {
		return manager.verifyingKey, nil
	}

	parserOptions := []jwt.ParserOption{
		jwt.WithIssuer(manager.issuer),
		jwt.WithValidMethods([]string{manager.signingMethod.Alg()}),
	}

	var err error
	var token *jwt.Token
	if token, err = jwt.Parse(tokenString, getKeyFunc, parserOptions...); err != nil {
		return nil, ErrTokenValidationFailed(ErrTokenFailedParsing, err)
	}

	if !token.Valid {
		return nil, ErrTokenValidationFailed(ErrTokenInvalid)
	}

	var ok bool
	var mapClaims jwt.MapClaims
	if mapClaims, ok = token.Claims.(jwt.MapClaims); !ok {
		return nil, ErrTokenValidationFailed(ErrTokenEmptyClaims)
	}

	var value any
	if value, ok = mapClaims["username"]; !ok {
		return nil, ErrTokenValidationFailed(ErrTokenEmptyUsernameClaim)
	}

	var username string
	if username, ok = value.(string); !ok {
		return nil, ErrTokenValidationFailed(ErrTokenEmptyUsernameClaim)
	}

	if value, ok = mapClaims["role"]; !ok {
		return nil, ErrTokenValidationFailed(ErrTokenEmptyRoleClaim)
	}

	var role string
	if role, ok = value.(string); !ok {
		return nil, ErrTokenValidationFailed(ErrTokenEmptyRoleClaim)
	}

	if value, ok = mapClaims["resources"]; !ok {
		return nil, ErrTokenValidationFailed(ErrTokenEmptyResourcesClaim)
	}

	var resourcesBytes []byte
	if resourcesBytes, err = json.Marshal(value); err != nil {
		return nil, ErrTokenValidationFailed(ErrTokenInvalidResourcesClaim)
	}

	var resources []string
	if err = json.Unmarshal(resourcesBytes, &resources); err != nil {
		return nil, ErrTokenValidationFailed(ErrTokenInvalidResourcesClaim)
	}

	principal := &Principal{
		Username:  util.ValueToPtr(username),
		Role:      util.ValueToPtr(role),
		Resources: resources,
	}

	return principal, nil
}
