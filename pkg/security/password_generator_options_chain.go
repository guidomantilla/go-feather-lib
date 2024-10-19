package security

type PasswordGeneratorOptionsChain struct {
	chain []PasswordGeneratorOptions
}

func PasswordGeneratorOptionsChainBuilder() *PasswordGeneratorOptionsChain {
	return &PasswordGeneratorOptionsChain{
		chain: make([]PasswordGeneratorOptions, 0),
	}
}

func (options *PasswordGeneratorOptionsChain) Build() PasswordGeneratorOptions {
	return func(generator *passwordGenerator) {
		for _, option := range options.chain {
			option(generator)
		}
	}
}

func (options *PasswordGeneratorOptionsChain) WithPasswordLength(length int) *PasswordGeneratorOptionsChain {
	options.chain = append(options.chain, passwordGeneratorOptions.WithPasswordLength(length))
	return options
}

func (options *PasswordGeneratorOptionsChain) WithMinSpecialChar(minSpecialChar int) *PasswordGeneratorOptionsChain {
	options.chain = append(options.chain, passwordGeneratorOptions.WithMinSpecialChar(minSpecialChar))
	return options
}

func (options *PasswordGeneratorOptionsChain) WithMinNum(minNum int) *PasswordGeneratorOptionsChain {
	options.chain = append(options.chain, passwordGeneratorOptions.WithMinNum(minNum))
	return options
}

func (options *PasswordGeneratorOptionsChain) WithMinUpperCase(minUpperCase int) *PasswordGeneratorOptionsChain {
	options.chain = append(options.chain, passwordGeneratorOptions.WithMinUpperCase(minUpperCase))
	return options
}
