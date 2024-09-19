package messaging

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

var headersOptions = NewHeadersOptions()

func NewHeadersOptions() HeadersOptions {
	return func(headers Headers) {
	}
}

func NewHeadersOptionsFromConfig(config *HeadersConfig) HeadersOptions {
	builder := HeadersOptionsChainBuilder().Id(config.Id).Timestamp(config.Timestamp).ReplyChannel(config.ReplyChannel).ErrorChannel(config.ErrorChannel)
	for key, value := range config.Headers {
		builder = builder.Add(key, value)
	}
	return builder.Build()
}

func (options HeadersOptions) Id(id uuid.UUID) HeadersOptions {
	return func(headers Headers) {
		if id != uuid.Nil {
			headers.Add(HeaderId, id.String())
		}
	}
}

func (options HeadersOptions) Timestamp(timestamp time.Time) HeadersOptions {
	return func(headers Headers) {
		if !timestamp.IsZero() {
			headers.Add(HeaderTimestamp, timestamp.Format(time.RFC3339))
		}
	}
}

func (options HeadersOptions) ReplyChannel(replyChannel string) HeadersOptions {
	return func(headers Headers) {
		if replyChannel != "" {
			headers.Add(HeaderReplyChannel, replyChannel)
		}
	}
}

func (options HeadersOptions) ErrorChannel(errorChannel string) HeadersOptions {
	return func(headers Headers) {
		if errorChannel != "" {
			headers.Add(HeaderErrorChannel, errorChannel)
		}
	}
}

func (options HeadersOptions) Add(property string, value string) HeadersOptions {
	return func(headers Headers) {
		if strings.TrimSpace(property) != "" && strings.TrimSpace(value) != "" {
			headers.Add(strings.ToLower(property), value)
		}
	}
}
