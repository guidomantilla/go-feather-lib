package astra

import (
	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/datasource/gocql"
)

type context_ struct {
	url      string
	database string
	token    string
}

func NewContext(url string, database string, token string) gocql.Context {
	assert.NotEmpty(url, "starting up - error setting up datasource context: url is empty")
	assert.NotEmpty(database, "starting up - error setting up datasource context: database is empty")
	assert.NotEmpty(token, "starting up - error setting up datasource context: token is empty")

	return &context_{
		url:      url,
		database: database,
		token:    token,
	}
}

func (context *context_) Server() []string {
	return []string{context.url}
}
