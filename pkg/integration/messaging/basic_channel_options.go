package messaging

import "time"

var channelOptions = NewChannelOptions()

func NewChannelOptions() ChannelOptions {
	return func(channel Channel) {
	}
}

func NewChannelOptionsFromConfig(config *ChannelConfig) HeadersOptions {

	return nil
}

func (options ChannelOptions) Timeout(timeout time.Duration) ChannelOptions {
	return func(channel Channel) {
		if timeout > 0 {
			channel.Timeout(timeout)
		}
	}
}
