package security

var passwordGeneratorOptions = NewPasswordGeneratorOptions()

func NewPasswordGeneratorOptions() PasswordGeneratorOptions {
	return func(generator *passwordGenerator) {
	}
}

type PasswordGeneratorOptions func(generator *passwordGenerator)

func (option PasswordGeneratorOptions) WithPasswordLength(length int) PasswordGeneratorOptions {
	return func(generator *passwordGenerator) {
		if length >= 16 {
			generator.passwordLength = length
		}
	}
}

func (option PasswordGeneratorOptions) WithMinSpecialChar(minSpecialChar int) PasswordGeneratorOptions {
	return func(generator *passwordGenerator) {
		if minSpecialChar >= 2 {
			generator.minSpecialChar = minSpecialChar
		}
	}
}

func (option PasswordGeneratorOptions) WithMinNum(minNum int) PasswordGeneratorOptions {
	return func(generator *passwordGenerator) {
		if minNum >= 2 {
			generator.minNum = minNum
		}
	}
}

func (option PasswordGeneratorOptions) WithMinUpperCase(minUpperCase int) PasswordGeneratorOptions {
	return func(generator *passwordGenerator) {
		if minUpperCase >= 2 {
			generator.minUpperCase = minUpperCase
		}
	}
}
