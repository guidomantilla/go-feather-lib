package security

var passwordGeneratorOptions = NewPasswordGeneratorOptions()

func NewPasswordGeneratorOptions() PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
	}
}

type PasswordGeneratorOptions func(generator PasswordGenerator)

func (options PasswordGeneratorOptions) WithPasswordLength(passwordLength int) PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
		generator.set("passwordLength", passwordLength)
	}
}

func (options PasswordGeneratorOptions) WithMinSpecialChar(minSpecialChar int) PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
		generator.set("minSpecialChar", minSpecialChar)
	}
}

func (options PasswordGeneratorOptions) WithMinNum(minNum int) PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
		generator.set("minNum", minNum)
	}
}

func (options PasswordGeneratorOptions) WithMinUpperCase(minUpperCase int) PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
		generator.set("minUpperCase", minUpperCase)
	}
}
