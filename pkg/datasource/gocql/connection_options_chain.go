package gocql

import "github.com/gocql/gocql"

type connectionOptionsChain struct {
	chain []ConnectionOptions
}

func ConnectionOptionsBuilder() *connectionOptionsChain {
	return &connectionOptionsChain{
		chain: make([]ConnectionOptions, 0),
	}
}

func (options *connectionOptionsChain) Build() ConnectionOptions {
	return func(connection Connection) {
		for _, option := range options.chain {
			option(connection)
		}
	}
}

func (options *connectionOptionsChain) WithDialer(dialer gocql.HostDialer) *connectionOptionsChain {
	options.chain = append(options.chain, connectionOptions.WithDialer(dialer))
	return options
}
