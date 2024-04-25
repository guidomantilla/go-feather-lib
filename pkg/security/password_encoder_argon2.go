package security

import (
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2PasswordEncoderOption func(encoder *Argon2PasswordEncoder)

type Argon2PasswordEncoder struct {
	iterations int
	memory     int
	threads    int
	saltLength int
	keyLength  int
}

func NewArgon2PasswordEncoder(options ...Argon2PasswordEncoderOption) *Argon2PasswordEncoder {

	encoder := &Argon2PasswordEncoder{
		iterations: 1,
		memory:     64 * 1024,
		threads:    2,
		saltLength: 16,
		keyLength:  32,
	}

	for _, opt := range options {
		opt(encoder)
	}

	return encoder
}

func WithArgon2Iterations(iterations int) Argon2PasswordEncoderOption {
	return func(encoder *Argon2PasswordEncoder) {
		encoder.iterations = iterations
	}
}

func WithArgon2Memory(memory int) Argon2PasswordEncoderOption {
	return func(encoder *Argon2PasswordEncoder) {
		encoder.memory = memory
	}
}

func WithArgon2Threads(threads int) Argon2PasswordEncoderOption {
	return func(encoder *Argon2PasswordEncoder) {
		encoder.threads = threads
	}
}

func WithArgon2SaltLength(saltLength int) Argon2PasswordEncoderOption {
	return func(encoder *Argon2PasswordEncoder) {
		encoder.saltLength = saltLength
	}
}

func WithArgon2KeyLength(keyLength int) Argon2PasswordEncoderOption {
	return func(encoder *Argon2PasswordEncoder) {
		encoder.keyLength = keyLength
	}
}

func (encoder *Argon2PasswordEncoder) Encode(rawPassword string) (*string, error) {

	var err error
	var salt []byte
	if salt, err = GenerateSalt(encoder.saltLength); err != nil {
		return nil, err
	}

	var value *string
	if value, err = Argon2Encode(rawPassword, salt, encoder.iterations, encoder.memory, encoder.threads, encoder.keyLength); err != nil {
		return nil, err
	}

	encodedPassword := *value
	encodedPassword = Argon2PrefixKey + encodedPassword
	return &encodedPassword, nil
}

func (encoder *Argon2PasswordEncoder) Matches(encodedPassword string, rawPassword string) (*bool, error) {

	if rawPassword == "" {
		return nil, ErrRawPasswordIsEmpty
	}

	if !strings.HasPrefix(encodedPassword, Argon2PrefixKey) {
		return nil, ErrEncodedPasswordNotAllowed
	}

	_, _, iterations, memory, threads, salt, key, err := Argon2Decode(encodedPassword)
	if err != nil {
		return nil, err
	}

	var newEncodedPassword *string
	if newEncodedPassword, err = Argon2Encode(rawPassword, salt, *iterations, *memory, *threads, len(key)); err != nil {
		return nil, err
	}

	encodedPassword = strings.Replace(encodedPassword, Argon2PrefixKey, "", 1)
	matched := encodedPassword == *(newEncodedPassword)
	return &matched, nil
}

func (encoder *Argon2PasswordEncoder) UpgradeEncoding(encodedPassword string) (*bool, error) {

	if encodedPassword == "" {
		return nil, ErrEncodedPasswordIsEmpty
	}

	if !strings.HasPrefix(encodedPassword, Argon2PrefixKey) {
		return nil, ErrEncodedPasswordNotAllowed
	}

	_, version, iterations, memory, threads, salt, key, err := Argon2Decode(encodedPassword)
	if err != nil {
		return nil, err
	}

	upgradeNeeded := true
	if argon2.Version > *(version) {
		return &upgradeNeeded, nil
	}

	if encoder.iterations > *(iterations) {
		return &upgradeNeeded, nil
	}

	if encoder.memory > *(memory) {
		return &upgradeNeeded, nil
	}

	if encoder.threads > *(threads) {
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
