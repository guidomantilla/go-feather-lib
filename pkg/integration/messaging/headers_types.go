package messaging

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

const (
	HeaderId           = "id"
	HeaderTimestamp    = "timestamp"
	HeaderReplyChannel = "reply-channel"
	HeaderErrorChannel = "error-channel"
)

type HeadersOptions func(headers Headers)

type HeadersOptionsChain interface {
	Build() HeadersOptions
	Id(id uuid.UUID) HeadersOptionsChain
	Timestamp(timestamp time.Time) HeadersOptionsChain
	ReplyChannel(replyChannel string) HeadersOptionsChain
	ErrorChannel(errorChannel string) HeadersOptionsChain
	Add(property string, value string) HeadersOptionsChain
}

type Headers interface {
	fmt.Stringer
	properties.Properties
	Id() uuid.UUID
	Timestamp() time.Time
	ReplyChannel() string
	ErrorChannel() string
}
