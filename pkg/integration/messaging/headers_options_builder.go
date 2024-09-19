package messaging

import (
	"time"

	"github.com/google/uuid"
)

type BaseHeadersOptionsChain struct {
	chain []HeadersOptions
}

func HeadersOptionsChainBuilder() HeadersOptionsChain {
	return &BaseHeadersOptionsChain{
		chain: make([]HeadersOptions, 0),
	}
}

func (options *BaseHeadersOptionsChain) Build() HeadersOptions {
	return func(headers Headers) {
		for _, option := range options.chain {
			option(headers)
		}
	}
}

func (options *BaseHeadersOptionsChain) Id(id uuid.UUID) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Id(id))
	return options
}

func (options *BaseHeadersOptionsChain) MessageType(messageType string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.MessageType(messageType))
	return options
}

func (options *BaseHeadersOptionsChain) Timestamp(timestamp time.Time) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Timestamp(timestamp))
	return options
}

func (options *BaseHeadersOptionsChain) Expired(expired bool) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Expired(expired))
	return options
}

func (options *BaseHeadersOptionsChain) TimeToLive(timeToLive time.Duration) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.TimeToLive(timeToLive))
	return options
}

func (options *BaseHeadersOptionsChain) ContentType(contentType string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.ContentType(contentType))
	return options
}

func (options *BaseHeadersOptionsChain) OriginChannel(originChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.OriginChannel(originChannel))
	return options
}

func (options *BaseHeadersOptionsChain) DestinationChannel(destinationChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.DestinationChannel(destinationChannel))
	return options
}

func (options *BaseHeadersOptionsChain) ReplyChannel(replyChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.ReplyChannel(replyChannel))
	return options
}

func (options *BaseHeadersOptionsChain) ErrorChannel(errorChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.ErrorChannel(errorChannel))
	return options
}

func (options *BaseHeadersOptionsChain) Add(property string, value string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Add(property, value))
	return options
}
