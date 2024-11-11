package security

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

var SupportedDecoders = map[string]PasswordEncoder{
	Argon2PrefixKey: NewArgon2PasswordEncoder(),
	BcryptPrefixKey: NewBcryptPasswordEncoder(),
	Pbkdf2PrefixKey: NewPbkdf2PasswordEncoder(),
	ScryptPrefixKey: NewScryptPasswordEncoder(),
}

type DelegatingPasswordEncoderOption func(encoder *DelegatingPasswordEncoder)

type DelegatingPasswordEncoder struct {
	encoder  PasswordEncoder
	decoders map[string]PasswordEncoder
}

func NewDelegatingPasswordEncoder(encoder PasswordEncoder, options ...DelegatingPasswordEncoderOption) *DelegatingPasswordEncoder {

	assert.NotNil(encoder, "starting up - error setting up delegating password encoder: encoder is nil")

	delegator := &DelegatingPasswordEncoder{
		encoder:  encoder,
		decoders: SupportedDecoders,
	}

	for _, opt := range options {
		opt(delegator)
	}

	return delegator
}

func WithSupportedDecoders(decoders map[string]PasswordEncoder) DelegatingPasswordEncoderOption {
	return func(encoder *DelegatingPasswordEncoder) {
		encoder.decoders = decoders
	}
}

func (delegate *DelegatingPasswordEncoder) Encode(rawPassword string) (*string, error) {
	return delegate.encoder.Encode(rawPassword)
}

func (delegate *DelegatingPasswordEncoder) Matches(encodedPassword string, rawPassword string) (*bool, error) {

	for prefix, decoder := range delegate.decoders {
		if strings.HasPrefix(encodedPassword, prefix) {
			return decoder.Matches(encodedPassword, rawPassword)
		}
	}

	return nil, ErrPasswordEncoderNotFound
}

func (delegate *DelegatingPasswordEncoder) UpgradeEncoding(encodedPassword string) (*bool, error) {

	for prefix, decoder := range delegate.decoders {
		if strings.HasPrefix(encodedPassword, prefix) {
			return decoder.UpgradeEncoding(encodedPassword)
		}
	}

	return nil, ErrPasswordEncoderNotFound
}
