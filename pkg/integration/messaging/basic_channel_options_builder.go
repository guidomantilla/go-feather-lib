package messaging

import "time"

type BasicChannelOptionsChain struct {
	Chain []ChannelOptions
}

func ChannelOptionsChainBuilder() ChannelOptionsChain {
	return &BasicChannelOptionsChain{
		Chain: make([]ChannelOptions, 0),
	}
}

func (options *BasicChannelOptionsChain) Build() ChannelOptions {
	return func(channel Channel) {
		for _, option := range options.Chain {
			option(channel)
		}
	}
}

func (options *BasicChannelOptionsChain) Timeout(timeout time.Duration) ChannelOptionsChain {
	options.Chain = append(options.Chain, channelOptions.Timeout(timeout))
	return options
}
