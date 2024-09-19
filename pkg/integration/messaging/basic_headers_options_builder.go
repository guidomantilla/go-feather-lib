package messaging

import (
	"time"

	"github.com/google/uuid"
)

type BasicHeadersOptionsChain struct {
	chain []HeadersOptions
}

func HeadersOptionsChainBuilder() HeadersOptionsChain {
	return &BasicHeadersOptionsChain{
		chain: make([]HeadersOptions, 0),
	}
}

func (options *BasicHeadersOptionsChain) Build() HeadersOptions {
	return func(headers Headers) {
		for _, option := range options.chain {
			option(headers)
		}
	}
}

func (options *BasicHeadersOptionsChain) Id(id uuid.UUID) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Id(id))
	return options
}

func (options *BasicHeadersOptionsChain) Timestamp(timestamp time.Time) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Timestamp(timestamp))
	return options
}

func (options *BasicHeadersOptionsChain) ReplyChannel(replyChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.ReplyChannel(replyChannel))
	return options
}

func (options *BasicHeadersOptionsChain) ErrorChannel(errorChannel string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.ErrorChannel(errorChannel))
	return options
}

func (options *BasicHeadersOptionsChain) Add(property string, value string) HeadersOptionsChain {
	options.chain = append(options.chain, headersOptions.Add(property, value))
	return options
}
