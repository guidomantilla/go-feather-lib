package gocql

import (
	"context"
	"fmt"

	retry "github.com/avast/retry-go/v4"
	"github.com/gocql/gocql"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/common/utils"
)

type connection struct {
	context  Context
	database *gocql.Session
	dialer   gocql.HostDialer
}

func NewConnection(context Context, options ...ConnectionOptions) Connection {
	assert.NotNil(context, "starting up - error setting up datasource connection: context is nil")
	assert.NotEmpty(options, "starting up - error setting up datasource connection: options is empty")

	connection := &connection{
		context:  context,
		database: nil,
		dialer:   nil,
	}

	for _, option := range options {
		option(connection)
	}

	return connection
}

func (datasource *connection) Connect(_ context.Context) (*gocql.Session, error) {

	if datasource.database == nil {

		err := retry.Do(datasource.connect, retry.Attempts(5),
			retry.OnRetry(func(n uint, err error) {
				log.Info("datasource connection - failed to connect")
				log.Info(fmt.Sprintf("datasource connection - trying reconnection to %s", datasource.context.Server()))
			}),
		)

		if err != nil {
			return nil, err
		}
	}

	return datasource.database, nil
}

func (datasource *connection) connect() error {

	var err error
	if datasource.database, err = gocql.NewCluster(datasource.context.Server()...).CreateSession(); err != nil {
		log.Error(err.Error())
		return ErrDBConnectionFailed(err)
	}
	log.Info(fmt.Sprintf("datasource connection - connected to %s", datasource.context.Server()))

	return nil
}

func (datasource *connection) Close(_ context.Context) {

	if datasource.database != nil && !datasource.database.Closed() {
		log.Debug("datasource connection - closing connection")
		datasource.database.Close()
	}
	datasource.database = nil
	log.Debug(fmt.Sprintf("datasource connection - closed connection to %s", datasource.context.Server()))
}

func (datasource *connection) Context() Context {
	return datasource.context
}

func (datasource *connection) Set(property string, value any) {
	if utils.IsEmpty(property) || utils.IsEmpty(value) {
		return
	}

	switch property {
	case "dialer":
		datasource.dialer = utils.ToType[gocql.HostDialer](value)
	}
}
