package environment

import "github.com/guidomantilla/go-feather-lib/pkg/common/properties"

type OptionsChain struct {
	chain []Options
}

func OptionsChainBuilder() *OptionsChain {
	return &OptionsChain{
		chain: make([]Options, 0),
	}
}

func (options *OptionsChain) Build() Options {
	return func(environment Environment) {
		for _, option := range options.chain {
			option(environment)
		}
	}
}

func (options *OptionsChain) WithCmd(cmdArgs []string) *OptionsChain {
	options.chain = append(options.chain, options_.WithCmd(cmdArgs))
	return options
}

func (options *OptionsChain) WithSSL() *OptionsChain {
	options.chain = append(options.chain, options_.WithSSL())
	return options
}

func (options *OptionsChain) WithOs() *OptionsChain {
	options.chain = append(options.chain, options_.WithOs())
	return options
}

func (options *OptionsChain) WithArraySource(name string, array []string) *OptionsChain {
	options.chain = append(options.chain, options_.WithArraySource(name, array))
	return options
}

func (options *OptionsChain) WithPropertySources(propertySources ...properties.PropertiesSource) *OptionsChain {
	options.chain = append(options.chain, options_.WithPropertySources(propertySources...))
	return options
}
