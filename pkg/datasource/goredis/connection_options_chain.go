package goredis

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
