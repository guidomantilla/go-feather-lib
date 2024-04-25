package security

import (
	"strings"
)

type ScryptPasswordEncoderOption func(encoder *ScryptPasswordEncoder)

type ScryptPasswordEncoder struct {
	N          int
	r          int
	p          int
	saltLength int
	keyLength  int
}

func NewScryptPasswordEncoder(options ...ScryptPasswordEncoderOption) *ScryptPasswordEncoder {

	encoder := &ScryptPasswordEncoder{
		N:          32768,
		r:          8,
		p:          1,
		saltLength: 16,
		keyLength:  32,
	}

	for _, opt := range options {
		opt(encoder)
	}

	return encoder
}

func WithScryptN(N int) ScryptPasswordEncoderOption {
	return func(encoder *ScryptPasswordEncoder) {
		encoder.N = N
	}
}

func WithScryptR(r int) ScryptPasswordEncoderOption {
	return func(encoder *ScryptPasswordEncoder) {
		encoder.r = r
	}
}

func WithScryptP(p int) ScryptPasswordEncoderOption {
	return func(encoder *ScryptPasswordEncoder) {
		encoder.p = p
	}
}

func WithScryptSaltLength(saltLength int) ScryptPasswordEncoderOption {
	return func(encoder *ScryptPasswordEncoder) {
		encoder.saltLength = saltLength
	}
}

func WithScryptKeyLength(keyLength int) ScryptPasswordEncoderOption {
	return func(encoder *ScryptPasswordEncoder) {
		encoder.keyLength = keyLength
	}
}

func (encoder *ScryptPasswordEncoder) Encode(rawPassword string) (*string, error) {

	var err error
	var salt []byte
	if salt, err = GenerateSalt(encoder.saltLength); err != nil {
		return nil, err
	}

	var value *string
	if value, err = ScryptEncode(rawPassword, salt, encoder.N, encoder.r, encoder.p, encoder.keyLength); err != nil {
		return nil, err
	}

	encodedPassword := *value
	encodedPassword = ScryptPrefixKey + encodedPassword
	return &encodedPassword, nil
}

func (encoder *ScryptPasswordEncoder) Matches(encodedPassword string, rawPassword string) (*bool, error) {

	if rawPassword == "" {
		return nil, ErrRawPasswordIsEmpty
	}

	if !strings.HasPrefix(encodedPassword, ScryptPrefixKey) {
		return nil, ErrEncodedPasswordNotAllowed
	}

	_, N, r, p, salt, key, err := ScryptDecode(encodedPassword)
	if err != nil {
		return nil, err
	}

	var newEncodedPassword *string
	if newEncodedPassword, err = ScryptEncode(rawPassword, salt, *N, *r, *p, len(key)); err != nil {
		return nil, err
	}

	encodedPassword = strings.Replace(encodedPassword, ScryptPrefixKey, "", 1)
	matched := encodedPassword == *(newEncodedPassword)
	return &matched, nil
}

func (encoder *ScryptPasswordEncoder) UpgradeEncoding(encodedPassword string) (*bool, error) {

	if encodedPassword == "" {
		return nil, ErrEncodedPasswordIsEmpty
	}

	if !strings.HasPrefix(encodedPassword, ScryptPrefixKey) {
		return nil, ErrEncodedPasswordNotAllowed
	}

	_, N, r, p, salt, key, err := ScryptDecode(encodedPassword)
	if err != nil {
		return nil, err
	}

	upgradeNeeded := true
	if encoder.N > *(N) {
		return &upgradeNeeded, nil
	}

	if encoder.r > *(r) {
		return &upgradeNeeded, nil
	}

	if encoder.p > *(p) {
		return &upgradeNeeded, nil
	}

	if encoder.saltLength > len(salt) {
		return &upgradeNeeded, nil
	}

	if encoder.keyLength > len(key) {
		return &upgradeNeeded, nil
	}

	upgradeNeeded = false
	return &upgradeNeeded, nil
}
