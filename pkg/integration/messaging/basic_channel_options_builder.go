package messaging

import "time"

type BasicChannelOptionsChain struct {
	chain []ChannelOptions
}

func ChannelOptionsChainBuilder() ChannelOptionsChain {
	return &BasicChannelOptionsChain{
		chain: make([]ChannelOptions, 0),
	}
}

func (options *BasicChannelOptionsChain) Build() ChannelOptions {
	return func(channel Channel) {
		for _, option := range options.chain {
			option(channel)
		}
	}
}

func (options *BasicChannelOptionsChain) Timeout(timeout time.Duration) ChannelOptionsChain {
	options.chain = append(options.chain, channelOptions.Timeout(timeout))
	return options
}
