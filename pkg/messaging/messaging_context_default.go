package messaging

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultMessagingContextOption func(defaultMessagingContext *DefaultMessagingContext)

func WithService(service string) DefaultMessagingContextOption {
	return func(defaultMessagingContext *DefaultMessagingContext) {
		defaultMessagingContext.service = service
	}
}

func WithVhost(vhost string) DefaultMessagingContextOption {
	return func(defaultMessagingContext *DefaultMessagingContext) {
		defaultMessagingContext.vhost = vhost
	}
}

type DefaultMessagingContext struct {
	url     string
	server  string
	service string
	vhost   string
}

func NewDefaultMessagingContext(url string, username string, password string, server string, options ...DefaultMessagingContextOption) *DefaultMessagingContext {

	if strings.TrimSpace(url) == "" {
		log.Fatal("starting up - error setting up messaging context: url is empty")
	}

	if strings.TrimSpace(username) == "" {
		log.Fatal("starting up - error setting up messaging context: username is empty")
	}

	if strings.TrimSpace(password) == "" {
		log.Fatal("starting up - error setting up messaging context: password is empty")
	}

	if strings.TrimSpace(server) == "" {
		log.Fatal("starting up - error setting up messaging context: server is empty")
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)

	defaultMessagingContext := &DefaultMessagingContext{
		url:    url,
		server: server,
	}

	for _, option := range options {
		option(defaultMessagingContext)
	}

	if defaultMessagingContext.service != "" && defaultMessagingContext.vhost == "" {
		defaultMessagingContext.url = strings.Replace(url, ":service", defaultMessagingContext.service, 1)
	}

	if defaultMessagingContext.service == "" && defaultMessagingContext.vhost != "" {
		defaultMessagingContext.url = strings.Replace(url, ":vhost", defaultMessagingContext.vhost, 1)
	}

	return defaultMessagingContext
}

func (context *DefaultMessagingContext) Url() string {
	return context.url
}

func (context *DefaultMessagingContext) Server() string {

	if context.service != "" && context.vhost == "" {
		return context.server + context.service
	}

	if context.service == "" && context.vhost != "" {
		return context.server + context.vhost
	}

	return context.server
}
