package messaging

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

var headersOptions = NewHeadersOptions()

func NewHeadersOptions() HeadersOptions {
	return func(headers Headers) {
	}
}

func NewHeadersOptionsFromConfig(config *HeadersConfig) HeadersOptions {
	assert.NotNil(config, fmt.Sprintf("integration messaging: %s error - config is required", "headers-options"))
	builder := HeadersOptionsChainBuilder().Id(config.Id).MessageType(config.MessageType).Timestamp(config.Timestamp).
		Expired(config.Expired).TimeToLive(config.TimeToLive).ContentType(config.ContentType).
		OriginChannel(config.OriginChannel).DestinationChannel(config.DestinationChannel).
		ReplyChannel(config.ReplyChannel).ErrorChannel(config.ErrorChannel)
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

func (options HeadersOptions) MessageType(messageType string) HeadersOptions {
	return func(headers Headers) {
		if strings.TrimSpace(messageType) != "" {
			headers.Add(HeaderMessageType, messageType)
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

func (options HeadersOptions) Expired(expired bool) HeadersOptions {
	return func(headers Headers) {
		headers.Add(HeaderExpired, strconv.FormatBool(expired))
	}
}

func (options HeadersOptions) TimeToLive(timeToLive time.Duration) HeadersOptions {
	return func(headers Headers) {
		if timeToLive > 0 {
			headers.Add(HeaderTimeToLive, timeToLive.String())
		}
	}
}

func (options HeadersOptions) ContentType(contentType string) HeadersOptions {
	return func(headers Headers) {
		if strings.TrimSpace(contentType) != "" {
			headers.Add(HeaderContentType, contentType)
		}
	}
}

func (options HeadersOptions) OriginChannel(originChannel string) HeadersOptions {
	return func(headers Headers) {
		if strings.TrimSpace(originChannel) != "" {
			headers.Add(HeaderOriginChannel, originChannel)
		}
	}
}

func (options HeadersOptions) DestinationChannel(destinationChannel string) HeadersOptions {
	return func(headers Headers) {
		if strings.TrimSpace(destinationChannel) != "" {
			headers.Add(HeaderDestinationChannel, destinationChannel)
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
