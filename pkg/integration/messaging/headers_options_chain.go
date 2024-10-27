package messaging

import (
	"time"

	"github.com/google/uuid"
)

type headersOptionsChain struct {
	chain []HeadersOptions
}

func HeadersOptionsBuilder() HeadersOptionsChain {
	return &headersOptionsChain{
		chain: make([]HeadersOptions, 0),
	}
}

func (options *headersOptionsChain) Build() HeadersOptions {
	return func(headers Headers) {
		for _, option := range options.chain {
			option(headers)
		}
	}
}

func (options *headersOptionsChain) Id(id uuid.UUID) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Id(id))
	return options
}

func (options *headersOptionsChain) MessageType(messageType string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.MessageType(messageType))
	return options
}

func (options *headersOptionsChain) Timestamp(timestamp time.Time) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Timestamp(timestamp))
	return options
}

func (options *headersOptionsChain) Expired(expired bool) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Expired(expired))
	return options
}

func (options *headersOptionsChain) TimeToLive(timeToLive time.Duration) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.TimeToLive(timeToLive))
	return options
}

func (options *headersOptionsChain) ContentType(contentType string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.ContentType(contentType))
	return options
}

func (options *headersOptionsChain) OriginChannel(originChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.OriginChannel(originChannel))
	return options
}

func (options *headersOptionsChain) DestinationChannel(destinationChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.DestinationChannel(destinationChannel))
	return options
}

func (options *headersOptionsChain) ReplyChannel(replyChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.ReplyChannel(replyChannel))
	return options
}

func (options *headersOptionsChain) ErrorChannel(errorChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.ErrorChannel(errorChannel))
	return options
}

func (options *headersOptionsChain) Add(property string, value string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Add(property, value))
	return options
}
