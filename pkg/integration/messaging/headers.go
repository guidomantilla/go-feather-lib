package messaging

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

type HeadersConfig struct {
	Id                 uuid.UUID
	MessageType        string
	Timestamp          time.Time
	Expired            bool
	TimeToLive         time.Duration
	ContentType        string
	OriginChannel      string
	DestinationChannel string
	ReplyChannel       string
	ErrorChannel       string
	Headers            map[string]string
}

type BaseHeaders struct {
	internal           map[string]string
	id                 uuid.UUID
	messageType        string
	timestamp          time.Time
	expired            bool
	timeToLive         time.Duration
	contentType        string
	originChannel      string
	destinationChannel string
	replyChannel       string
	errorChannel       string
}

func NewBaseHeaders(options ...HeadersOptions) *BaseHeaders {
	headers := &BaseHeaders{
		internal:           make(map[string]string),
		id:                 uuid.New(),
		messageType:        "",
		timestamp:          time.Now(),
		expired:            false,
		timeToLive:         -1,
		contentType:        "",
		originChannel:      "",
		destinationChannel: "",
		replyChannel:       "",
		errorChannel:       "",
	}

	headers.Add(HeaderErrorChannel, headers.errorChannel)
	headers.Add(HeaderReplyChannel, headers.replyChannel)
	headers.Add(HeaderDestinationChannel, headers.destinationChannel)
	headers.Add(HeaderOriginChannel, headers.originChannel)
	headers.Add(HeaderContentType, headers.contentType)
	headers.Add(HeaderExpiredTimeToLive, headers.timeToLive.String())
	headers.Add(HeaderExpired, strconv.FormatBool(headers.expired))
	headers.Add(HeaderTimestamp, headers.timestamp.Format(time.RFC3339))
	headers.Add(HeaderMessageType, headers.messageType)
	headers.Add(HeaderId, headers.id.String())

	for _, option := range options {
		option(headers)
	}

	return headers
}

func NewBaseHeadersFromConfig(config *HeadersConfig) *BaseHeaders {
	assert.NotNil(config, fmt.Sprintf("integration messaging: %s error - config is required", "base-headers"))
	return NewBaseHeaders(NewHeadersOptionsFromConfig(config))
}

//

func (headers *BaseHeaders) Id() uuid.UUID {
	return uuid.MustParse(headers.internal[HeaderId])
}

func (headers *BaseHeaders) MessageType() string {
	return headers.internal[HeaderMessageType]
}

func (headers *BaseHeaders) Timestamp() time.Time {
	value, _ := time.Parse(time.RFC3339, headers.internal[HeaderTimestamp])
	return value
}

func (headers *BaseHeaders) Expired() bool {
	value, _ := strconv.ParseBool(headers.internal[HeaderExpired])
	return value
}

func (headers *BaseHeaders) TimeToLive() time.Duration {
	value, _ := time.ParseDuration(headers.internal[HeaderExpiredTimeToLive])
	return value
}

func (headers *BaseHeaders) ContentType() string {
	return headers.internal[HeaderContentType]
}

func (headers *BaseHeaders) OriginChannel() string {
	return headers.internal[HeaderOriginChannel]
}

func (headers *BaseHeaders) DestinationChannel() string {
	return headers.internal[HeaderDestinationChannel]
}

func (headers *BaseHeaders) ReplyChannel() string {
	return headers.internal[HeaderReplyChannel]
}

func (headers *BaseHeaders) ErrorChannel() string {
	return headers.internal[HeaderErrorChannel]
}

//

func (headers *BaseHeaders) Add(property string, value string) {

	property, value = strings.TrimSpace(property), strings.TrimSpace(value)
	if property == "" || value == "" {
		return
	}

	property = strings.ToLower(property)
	if property == HeaderId {
		headers.id = uuid.MustParse(value)
	}
	if property == HeaderMessageType {
		headers.messageType = value
	}
	if property == HeaderTimestamp {
		headers.timestamp, _ = time.Parse(time.RFC3339, value)
	}
	if property == HeaderExpired {
		headers.expired, _ = strconv.ParseBool(value)
	}
	if property == HeaderExpiredTimeToLive {
		headers.timeToLive, _ = time.ParseDuration(value)
	}
	if property == HeaderContentType {
		headers.contentType = value
	}
	if property == HeaderOriginChannel {
		headers.originChannel = value
	}
	if property == HeaderDestinationChannel {
		headers.destinationChannel = value
	}
	if property == HeaderReplyChannel {
		headers.replyChannel = value
	}
	if property == HeaderErrorChannel {
		headers.errorChannel = value
	}

	headers.internal[property] = value
}

func (headers *BaseHeaders) Get(property string) string {
	return headers.internal[property]
}

func (headers *BaseHeaders) AsMap() map[string]string {
	return headers.internal
}

func (headers *BaseHeaders) String() string {
	return fmt.Sprintf("%v", headers.internal)
}
