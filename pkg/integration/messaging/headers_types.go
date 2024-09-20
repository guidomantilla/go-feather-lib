package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

const (
	HeaderId                 = "x-id"
	HeaderCorrelationId      = "x-correlation-id"
	HeaderMessageType        = "x-message-type"
	HeaderTimestamp          = "x-timestamp"
	HeaderContentType        = "x-content-type"
	HeaderExpired            = "x-expired"
	HeaderExpiredTimeToLive  = "x-expired-ttl"
	HeaderExpiredTimestamp   = "x-expired-timestamp"
	HeaderExpiredChannel     = "x-expired-channel"
	HeaderInvalid            = "x-invalid"
	HeaderInvalidCause       = "x-invalid-cause"
	HeaderInvalidChannel     = "x-invalid-channel"
	HeaderError              = "x-error"
	HeaderErrorCause         = "x-error-cause"
	HeaderErrorChannel       = "x-error-channel"
	HeaderOriginChannel      = "x-origin-channel"
	HeaderDestinationChannel = "x-destination-channel"
	HeaderReplyChannel       = "x-reply-channel"
)

type HeadersOptions func(headers Headers)

type HeadersOptionsChain interface {
	Build() HeadersOptions
	Id(id uuid.UUID) HeadersOptionsChain
	MessageType(messageType string) HeadersOptionsChain
	Timestamp(timestamp time.Time) HeadersOptionsChain
	Expired(expired bool) HeadersOptionsChain
	TimeToLive(timeToLive time.Duration) HeadersOptionsChain
	ContentType(contentType string) HeadersOptionsChain
	OriginChannel(originChannel string) HeadersOptionsChain
	DestinationChannel(destinationChannel string) HeadersOptionsChain
	ReplyChannel(replyChannel string) HeadersOptionsChain
	ErrorChannel(errorChannel string) HeadersOptionsChain
	Add(property string, value string) HeadersOptionsChain
}

type Headers interface {
	fmt.Stringer
	properties.Properties
	Id() uuid.UUID
	MessageType() string
	Timestamp() time.Time
	Expired() bool
	TimeToLive() time.Duration
	ContentType() string
	OriginChannel() string
	DestinationChannel() string
	ReplyChannel() string
	ErrorChannel() string
}

type HeadersValidatorHandler func(ctx context.Context, headers Headers) error

type HeadersValidator interface {
	Validate(ctx context.Context, headers Headers) error
}
