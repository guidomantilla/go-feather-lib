package messaging

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type HeadersConfig struct {
	Id           uuid.UUID
	Timestamp    time.Time
	ReplyChannel string
	ErrorChannel string
	Headers      map[string]string
}

type BaseHeaders struct {
	internal     map[string]string
	id           uuid.UUID
	timestamp    time.Time
	replyChannel string
	errorChannel string
}

func NewBaseHeaders(options ...HeadersOptions) *BaseHeaders {
	headers := &BaseHeaders{
		internal:     make(map[string]string),
		id:           uuid.New(),
		timestamp:    time.Now(),
		replyChannel: "",
		errorChannel: "",
	}

	headers.Add(HeaderErrorChannel, headers.errorChannel)
	headers.Add(HeaderReplyChannel, headers.replyChannel)
	headers.Add(HeaderTimestamp, headers.timestamp.Format(time.RFC3339))
	headers.Add(HeaderId, headers.id.String())

	for _, option := range options {
		option(headers)
	}

	return headers
}

func NewBasicHeadersFromConfig(config *HeadersConfig) *BaseHeaders {
	return NewBaseHeaders(NewHeadersOptionsFromConfig(config))
}

//

func (headers *BaseHeaders) Id() uuid.UUID {
	return uuid.MustParse(headers.internal[HeaderId])
}

func (headers *BaseHeaders) Timestamp() time.Time {
	value, _ := time.Parse(time.RFC3339, headers.internal[HeaderTimestamp])
	return value
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
	if property == HeaderTimestamp {
		headers.timestamp, _ = time.Parse(time.RFC3339, value)
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
