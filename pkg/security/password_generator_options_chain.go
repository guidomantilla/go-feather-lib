package security

type passwordGeneratorOptionsChain struct {
	chain []PasswordGeneratorOptions
}

func PasswordGeneratorOptionsBuilder() *passwordGeneratorOptionsChain {
	return &passwordGeneratorOptionsChain{
		chain: make([]PasswordGeneratorOptions, 0),
	}
}

func (options *passwordGeneratorOptionsChain) Build() PasswordGeneratorOptions {
	return func(generator PasswordGenerator) {
		for _, option := range options.chain {
			option(generator)
		}
	}
}

func (options *passwordGeneratorOptionsChain) WithPasswordLength(length int) *passwordGeneratorOptionsChain {
	options.chain = append(options.chain, passwordGeneratorOptions.WithPasswordLength(length))
	return options
}

func (options *passwordGeneratorOptionsChain) WithMinSpecialChar(minSpecialChar int) *passwordGeneratorOptionsChain {
	options.chain = append(options.chain, passwordGeneratorOptions.WithMinSpecialChar(minSpecialChar))
	return options
}

func (options *passwordGeneratorOptionsChain) WithMinNum(minNum int) *passwordGeneratorOptionsChain {
	options.chain = append(options.chain, passwordGeneratorOptions.WithMinNum(minNum))
	return options
}

func (options *passwordGeneratorOptionsChain) WithMinUpperCase(minUpperCase int) *passwordGeneratorOptionsChain {
	options.chain = append(options.chain, passwordGeneratorOptions.WithMinUpperCase(minUpperCase))
	return options
}
