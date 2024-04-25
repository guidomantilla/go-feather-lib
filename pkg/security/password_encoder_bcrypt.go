package security

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordEncoderOption func(encoder *BcryptPasswordEncoder)

type BcryptPasswordEncoder struct {
	cost int
}

func NewBcryptPasswordEncoder(options ...BcryptPasswordEncoderOption) *BcryptPasswordEncoder {

	encoder := &BcryptPasswordEncoder{cost: bcrypt.DefaultCost}

	for _, opt := range options {
		opt(encoder)
	}

	return encoder
}

func WithBcryptCost(cost int) BcryptPasswordEncoderOption {
	return func(encoder *BcryptPasswordEncoder) {
		encoder.cost = cost
	}
}

func (encoder *BcryptPasswordEncoder) Encode(rawPassword string) (*string, error) {

	if encoder.cost < bcrypt.MinCost || encoder.cost > bcrypt.MaxCost {
		return nil, ErrBcryptCostNotAllowed
	}

	if rawPassword == "" {
		return nil, ErrRawPasswordIsEmpty
	}

	var err error
	var bytes []byte
	if bytes, err = bcrypt.GenerateFromPassword([]byte(rawPassword), encoder.cost); err != nil {
		return nil, err
	}

	encodedPassword := BcryptPrefixKey + string(bytes)
	return &encodedPassword, nil
}

func (encoder *BcryptPasswordEncoder) Matches(encodedPassword string, rawPassword string) (*bool, error) {

	if rawPassword == "" {
		return nil, ErrRawPasswordIsEmpty
	}

	if encodedPassword == "" {
		return nil, ErrEncodedPasswordIsEmpty
	}

	if !strings.HasPrefix(encodedPassword, BcryptPrefixKey) {
		return nil, ErrEncodedPasswordNotAllowed
	}

	matched := true
	encodedPassword = strings.Replace(encodedPassword, BcryptPrefixKey, "", 1)
	if err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword)); err != nil {
		matched = false
	}

	return &matched, nil
}

func (encoder *BcryptPasswordEncoder) UpgradeEncoding(encodedPassword string) (*bool, error) {

	if encodedPassword == "" {
		return nil, ErrEncodedPasswordIsEmpty
	}

	if !strings.HasPrefix(encodedPassword, BcryptPrefixKey) {
		return nil, ErrEncodedPasswordNotAllowed
	}

	encodedPassword = strings.Replace(encodedPassword, BcryptPrefixKey, "", 1)

	cost, _ := bcrypt.Cost([]byte(encodedPassword))
	upgradeNeeded := cost < encoder.cost

	return &upgradeNeeded, nil
}
