package security

var passwordGeneratorOptions = NewPasswordGeneratorOptions()

func NewPasswordGeneratorOptions() PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
	}
}

type PasswordGeneratorOptions func(generator PasswordGenerator)

func (option PasswordGeneratorOptions) WithPasswordLength(passwordLength int) PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
		generator.set("passwordLength", passwordLength)
	}
}

func (option PasswordGeneratorOptions) WithMinSpecialChar(minSpecialChar int) PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
		generator.set("minSpecialChar", minSpecialChar)
	}
}

func (option PasswordGeneratorOptions) WithMinNum(minNum int) PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
		generator.set("minNum", minNum)
	}
}

func (option PasswordGeneratorOptions) WithMinUpperCase(minUpperCase int) PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
		generator.set("minUpperCase", minUpperCase)
	}
}
