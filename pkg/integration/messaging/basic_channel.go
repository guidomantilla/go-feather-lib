package messaging

import (
	"fmt"
	"time"
)

type ChannelConfig struct {
	Timeout time.Duration
}

type BasicChannel struct {
	internal map[string]any
	timeout  time.Duration
	channel  ChannelPipe
}

func NewBasicChannelOptions(options ...ChannelOptions) *BasicChannel {
	channel := &BasicChannel{
		internal: make(map[string]any),
		timeout:  -1,
		channel:  make(ChannelPipe, 100),
	}

	for _, option := range options {
		option(channel)
	}

	channel.internal["timeout"] = channel.timeout
	channel.internal["channel"] = channel.channel.String()

	return channel
}

//

func (channel *BasicChannel) Send(message Message[any], options ...ChannelOptions) error {
	for _, option := range options {
		option(channel)
	}

	select {
	case channel.channel <- message:
		return nil
	case <-time.After(channel.timeout):
		return fmt.Errorf("failed to send message: timeout after %v", channel.timeout)
	}
}

func (channel *BasicChannel) Timeout(timeout time.Duration) {
	channel.internal["timeout"] = channel.timeout
	channel.timeout = timeout
}

func (channel *BasicChannel) Pipe(pipe ChannelPipe) {
	//TODO implement me
	panic("implement me")
}

func (channel *BasicChannel) String() string {
	return fmt.Sprintf("%v", channel.internal)
}
