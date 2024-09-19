package messaging

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

const (
	HeaderId                 = "x-id"
	HeaderMessageType        = "x-message-type"
	HeaderTimestamp          = "x-timestamp"
	HeaderExpired            = "x-expired"
	HeaderTimeToLive         = "x-ttl"
	HeaderContentType        = "x-content-type"
	HeaderOriginChannel      = "x-origin-channel"
	HeaderDestinationChannel = "x-destination-channel"
	HeaderReplyChannel       = "x-reply-channel"
	HeaderErrorChannel       = "x-error-channel"
	/*
		HeaderCorrelationId = "x-correlation-id"
		HeaderPriority      = "x-priority"
		HeaderUserId        = "x-user-id"
		HeaderSessionId     = "x-session-id"
		HeaderSource        = "x-source"
		HeaderDestination   = "x-destination"
	*/
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
