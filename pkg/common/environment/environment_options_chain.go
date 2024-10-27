package environment

import "github.com/guidomantilla/go-feather-lib/pkg/common/properties"

type optionsChain struct {
	chain []Options
}

func OptionsBuilder() *optionsChain {
	return &optionsChain{
		chain: make([]Options, 0),
	}
}

func (options *optionsChain) Build() Options {
	return func(environment Environment) {
		for _, option := range options.chain {
			option(environment)
		}
	}
}

func (options *optionsChain) WithCmd(cmdArgs []string) *optionsChain {
	options.chain = append(options.chain, options_.WithCmd(cmdArgs))
	return options
}

func (options *optionsChain) WithSSL() *optionsChain {
	options.chain = append(options.chain, options_.WithSSL())
	return options
}

func (options *optionsChain) WithOs() *optionsChain {
	options.chain = append(options.chain, options_.WithOs())
	return options
}

func (options *optionsChain) WithArraySource(name string, array []string) *optionsChain {
	options.chain = append(options.chain, options_.WithArraySource(name, array))
	return options
}

func (options *optionsChain) WithPropertySources(propertySources ...properties.PropertiesSource) *optionsChain {
	options.chain = append(options.chain, options_.WithPropertySources(propertySources...))
	return options
}
