package security

import (
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type passwordManager struct {
	passwordEncoder   PasswordEncoder
	passwordGenerator PasswordGenerator
}

func NewPasswordManager(passwordEncoder PasswordEncoder, passwordGenerator PasswordGenerator) PasswordManager {

	if passwordEncoder == nil {
		log.Fatal("starting up - error setting up passwordManager: passwordEncoder is nil")
	}

	if passwordGenerator == nil {
		log.Fatal("starting up - error setting up passwordManager: passwordGenerator is nil")
	}

	return &passwordManager{
		passwordEncoder:   passwordEncoder,
		passwordGenerator: passwordGenerator,
	}
}

func (manager *passwordManager) Encode(rawPassword string) (*string, error) {

	var err error
	var password *string
	if password, err = manager.passwordEncoder.Encode(rawPassword); err != nil {
		return nil, ErrPasswordEncodingFailed(err)
	}

	return password, nil
}

func (manager *passwordManager) Matches(encodedPassword string, rawPassword string) (*bool, error) {

	var err error
	var ok *bool
	if ok, err = manager.passwordEncoder.Matches(encodedPassword, rawPassword); err != nil {
		return nil, ErrPasswordMatchingFailed(err)
	}

	return ok, nil
}

func (manager *passwordManager) UpgradeEncoding(encodedPassword string) (*bool, error) {

	var err error
	var ok *bool
	if ok, err = manager.passwordEncoder.UpgradeEncoding(encodedPassword); err != nil {
		return nil, ErrPasswordUpgradeEncodingValidationFailed(err)
	}

	return ok, nil
}

func (manager *passwordManager) Generate() string {
	return manager.passwordGenerator.Generate()
}

func (manager *passwordManager) Validate(rawPassword string) error {

	var err error
	if err = manager.passwordGenerator.Validate(rawPassword); err != nil {
		return ErrPasswordValidationFailed(err)
	}

	return nil
}
