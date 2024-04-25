package security

import (
	"crypto/sha512"
	"hash"
	"strings"
)

type HashFunc func() hash.Hash

type Pbkdf2PasswordEncoderOption func(encoder *Pbkdf2PasswordEncoder)

type Pbkdf2PasswordEncoder struct {
	iterations int
	saltLength int
	keyLength  int
	hashFunc   HashFunc
}

func NewPbkdf2PasswordEncoder(options ...Pbkdf2PasswordEncoderOption) *Pbkdf2PasswordEncoder {

	encoder := &Pbkdf2PasswordEncoder{
		iterations: 600_000,
		saltLength: 32,
		keyLength:  64,
		hashFunc:   sha512.New,
	}

	for _, opt := range options {
		opt(encoder)
	}

	return encoder
}

func WithPbkdf2Iterations(iterations int) Pbkdf2PasswordEncoderOption {
	return func(encoder *Pbkdf2PasswordEncoder) {
		encoder.iterations = iterations
	}
}

func WithPbkdf2SaltLength(saltLength int) Pbkdf2PasswordEncoderOption {
	return func(encoder *Pbkdf2PasswordEncoder) {
		encoder.saltLength = saltLength
	}
}

func WithPbkdf2KeyLength(keyLength int) Pbkdf2PasswordEncoderOption {
	return func(encoder *Pbkdf2PasswordEncoder) {
		encoder.keyLength = keyLength
	}
}

func WithHashFunc(hashFunc HashFunc) Pbkdf2PasswordEncoderOption {
	return func(encoder *Pbkdf2PasswordEncoder) {
		encoder.hashFunc = hashFunc
	}
}

func (encoder *Pbkdf2PasswordEncoder) Encode(rawPassword string) (*string, error) {

	var err error
	var salt []byte
	if salt, err = GenerateSalt(encoder.saltLength); err != nil {
		return nil, err
	}

	var value *string
	if value, err = Pbkdf2Encode(rawPassword, salt, encoder.iterations, encoder.keyLength, encoder.hashFunc); err != nil {
		return nil, err
	}

	encodedPassword := *value
	encodedPassword = Pbkdf2PrefixKey + encodedPassword
	return &encodedPassword, nil
}

func (encoder *Pbkdf2PasswordEncoder) Matches(encodedPassword string, rawPassword string) (*bool, error) {

	if rawPassword == "" {
		return nil, ErrRawPasswordIsEmpty
	}

	if !strings.HasPrefix(encodedPassword, Pbkdf2PrefixKey) {
		return nil, ErrEncodedPasswordNotAllowed
	}

	_, iterations, salt, key, err := Pbkdf2Decode(encodedPassword)
	if err != nil {
		return nil, err
	}

	var newEncodedPassword *string
	if newEncodedPassword, err = Pbkdf2Encode(rawPassword, salt, *iterations, len(key), encoder.hashFunc); err != nil {
		return nil, err
	}

	encodedPassword = strings.Replace(encodedPassword, Pbkdf2PrefixKey, "", 1)
	matched := encodedPassword == *(newEncodedPassword)
	return &matched, nil
}

func (encoder *Pbkdf2PasswordEncoder) UpgradeEncoding(encodedPassword string) (*bool, error) {

	if encodedPassword == "" {
		return nil, ErrEncodedPasswordIsEmpty
	}

	if !strings.HasPrefix(encodedPassword, Pbkdf2PrefixKey) {
		return nil, ErrEncodedPasswordNotAllowed
	}

	_, iterations, salt, key, err := Pbkdf2Decode(encodedPassword)
	if err != nil {
		return nil, err
	}

	upgradeNeeded := true
	if encoder.iterations > *(iterations) {
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
