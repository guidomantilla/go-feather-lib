package security

import (
	"github.com/guidomantilla/go-feather-lib/pkg/log"
)

type DefaultPasswordManager struct {
	passwordEncoder   PasswordEncoder
	passwordGenerator PasswordGenerator
}

func NewDefaultPasswordManager(passwordEncoder PasswordEncoder, passwordGenerator PasswordGenerator) *DefaultPasswordManager {

	if passwordEncoder == nil {
		log.Fatal("starting up - error setting up passwordManager: passwordEncoder is nil")
	}

	if passwordGenerator == nil {
		log.Fatal("starting up - error setting up passwordManager: passwordGenerator is nil")
	}

	return &DefaultPasswordManager{
		passwordEncoder:   passwordEncoder,
		passwordGenerator: passwordGenerator,
	}
}

func (manager *DefaultPasswordManager) Encode(rawPassword string) (*string, error) {

	var err error
	var password *string
	if password, err = manager.passwordEncoder.Encode(rawPassword); err != nil {
		return nil, ErrPasswordEncodingFailed(err)
	}

	return password, nil
}

func (manager *DefaultPasswordManager) Matches(encodedPassword string, rawPassword string) (*bool, error) {

	var err error
	var ok *bool
	if ok, err = manager.passwordEncoder.Matches(encodedPassword, rawPassword); err != nil {
		return nil, ErrPasswordMatchingFailed(err)
	}

	return ok, nil
}

func (manager *DefaultPasswordManager) UpgradeEncoding(encodedPassword string) (*bool, error) {

	var err error
	var ok *bool
	if ok, err = manager.passwordEncoder.UpgradeEncoding(encodedPassword); err != nil {
		return nil, ErrPasswordUpgradeEncodingValidationFailed(err)
	}

	return ok, nil
}

func (manager *DefaultPasswordManager) Generate() string {
	return manager.passwordGenerator.Generate()
}

func (manager *DefaultPasswordManager) Validate(rawPassword string) error {

	var err error
	if err = manager.passwordGenerator.Validate(rawPassword); err != nil {
		return ErrPasswordValidationFailed(err)
	}

	return nil
}
