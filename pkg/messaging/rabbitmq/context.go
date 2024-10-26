package rabbitmq

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/utils"
)

type context_ struct {
	url     string
	server  string
	service string
	vhost   string
}

func NewContext(url string, username string, password string, server string, options ...ContextOptions) Context {
	assert.NotEmpty(url, "starting up - error setting up messaging context: url is empty")
	assert.NotEmpty(username, "starting up - error setting up messaging context: username is empty")
	assert.NotEmpty(password, "starting up - error setting up messaging context: password is empty")
	assert.NotEmpty(server, "starting up - error setting up messaging context: server is empty")

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)

	context := &context_{
		url:    url,
		server: server,
	}

	for _, option := range options {
		option(context)
	}

	if context.service != "" && context.vhost == "" {
		context.url = strings.Replace(url, ":service", context.service, 1)
	}

	if context.service == "" && context.vhost != "" {
		context.url = strings.Replace(url, ":vhost", context.vhost, 1)
	}

	return context
}

func (context *context_) Url() string {
	return context.url
}

func (context *context_) Server() string {

	if context.service != "" && context.vhost == "" {
		return context.server + context.service
	}

	if context.service == "" && context.vhost != "" {
		return context.server + context.vhost
	}

	return context.server
}

func (context *context_) Set(property string, value string) {
	if utils.IsEmpty(property) || utils.IsEmpty(value) {
		return
	}

	switch property {
	case "service":
		context.service = strings.TrimSpace(value)
	case "vhost":
		context.vhost = strings.TrimSpace(value)
	}
}
