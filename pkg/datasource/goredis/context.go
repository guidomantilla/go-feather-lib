package goredis

import (
	"fmt"
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

type context_ struct {
	username string
	password string
	server   []string
}

func NewContext(username string, password string, servers string) Context {
	assert.NotEmpty(username, "starting up - error setting up datasource context: username is empty")
	assert.NotEmpty(password, "starting up - error setting up datasource context: password is empty")
	assert.NotEmpty(servers, "starting up - error setting up datasource context: servers is empty")
	assert.True(len(strings.Split(servers, ";")) > 0, "starting up - error setting up datasource context: servers is empty")

	return &context_{
		username: username,
		password: password,
		server:   strings.Split(servers, ";"),
	}
}

func (context *context_) Url() string {
	return fmt.Sprintf("%s/%s@%s", context.User(), context.Password(), context.Server())
}

func (context *context_) User() string {
	return context.username
}

func (context *context_) Password() string {
	return context.password
}

func (context *context_) Server() any {
	return context.server
}

func (context *context_) Service() string {
	return ""
}
