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

type BasicHeaders struct {
	internal     map[string]string
	id           uuid.UUID
	timestamp    time.Time
	replyChannel string
	errorChannel string
}

func NewBasicHeaders(options ...HeadersOptions) *BasicHeaders {
	headers := &BasicHeaders{
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

func NewBasicHeadersFromConfig(config *HeadersConfig) *BasicHeaders {
	return NewBasicHeaders(NewHeadersOptionsFromConfig(config))
}

//

func (headers *BasicHeaders) Id() uuid.UUID {
	return uuid.MustParse(headers.internal[HeaderId])
}

func (headers *BasicHeaders) Timestamp() time.Time {
	value, _ := time.Parse(time.RFC3339, headers.internal[HeaderTimestamp])
	return value
}

func (headers *BasicHeaders) ReplyChannel() string {
	return headers.internal[HeaderReplyChannel]
}

func (headers *BasicHeaders) ErrorChannel() string {
	return headers.internal[HeaderErrorChannel]
}

//

func (headers *BasicHeaders) Add(property string, value string) {

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

func (headers *BasicHeaders) Get(property string) string {
	return headers.internal[property]
}

func (headers *BasicHeaders) AsMap() map[string]string {
	return headers.internal
}

func (headers *BasicHeaders) String() string {
	return fmt.Sprintf("%v", headers.internal)
}
