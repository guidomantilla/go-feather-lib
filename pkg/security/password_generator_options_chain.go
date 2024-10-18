package security

type PasswordGeneratorOptionChain struct {
	chain []PasswordGeneratorOption
}

func PasswordGeneratorOptionsChainBuilder() *PasswordGeneratorOptionChain {
	return &PasswordGeneratorOptionChain{
		chain: make([]PasswordGeneratorOption, 0),
	}
}

func (options *PasswordGeneratorOptionChain) Build() PasswordGeneratorOption {
	return func(generator *passwordGenerator) {
		for _, option := range options.chain {
			option(generator)
		}
	}
}

func (options *PasswordGeneratorOptionChain) WithPasswordLength(length int) *PasswordGeneratorOptionChain {
	options.chain = append(options.chain, passwordGeneratorOption.WithPasswordLength(length))
	return options
}

func (options *PasswordGeneratorOptionChain) WithMinSpecialChar(minSpecialChar int) *PasswordGeneratorOptionChain {
	options.chain = append(options.chain, passwordGeneratorOption.WithMinSpecialChar(minSpecialChar))
	return options
}

func (options *PasswordGeneratorOptionChain) WithMinNum(minNum int) *PasswordGeneratorOptionChain {
	options.chain = append(options.chain, passwordGeneratorOption.WithMinNum(minNum))
	return options
}

func (options *PasswordGeneratorOptionChain) WithMinUpperCase(minUpperCase int) *PasswordGeneratorOptionChain {
	options.chain = append(options.chain, passwordGeneratorOption.WithMinUpperCase(minUpperCase))
	return options
}
