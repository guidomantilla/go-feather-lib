package security

var passwordGeneratorOption = NewPasswordGeneratorOption()

func NewPasswordGeneratorOption() PasswordGeneratorOption {
	return func(generator *passwordGenerator) {
	}
}

type PasswordGeneratorOption func(generator *passwordGenerator)

func (option PasswordGeneratorOption) WithPasswordLength(length int) PasswordGeneratorOption {
	return func(generator *passwordGenerator) {
		if length >= 16 {
			generator.passwordLength = length
		}
	}
}

func (option PasswordGeneratorOption) WithMinSpecialChar(minSpecialChar int) PasswordGeneratorOption {
	return func(generator *passwordGenerator) {
		if minSpecialChar >= 2 {
			generator.minSpecialChar = minSpecialChar
		}
	}
}

func (option PasswordGeneratorOption) WithMinNum(minNum int) PasswordGeneratorOption {
	return func(generator *passwordGenerator) {
		if minNum >= 2 {
			generator.minNum = minNum
		}
	}
}

func (option PasswordGeneratorOption) WithMinUpperCase(minUpperCase int) PasswordGeneratorOption {
	return func(generator *passwordGenerator) {
		if minUpperCase >= 2 {
			generator.minUpperCase = minUpperCase
		}
	}
}
