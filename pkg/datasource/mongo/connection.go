package mongo

import (
	"context"
	"fmt"

	retry "github.com/avast/retry-go/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type connection struct {
	context  Context
	database *mongo.Client
	opts     *options.ClientOptionsBuilder
}

func NewConnection(context Context, opts *options.ClientOptionsBuilder) *connection {
	assert.NotNil(context, "starting up - error setting up datasource connection: context is nil")
	assert.NotEmpty(opts, "starting up - error setting up datasource connection: opts is empty")

	return &connection{
		context:  context,
		database: nil,
		opts:     opts,
	}
}

func (datasource *connection) Connect(_ context.Context) (*mongo.Client, error) {

	if datasource.database == nil {

		err := retry.Do(datasource.connect, retry.Attempts(5),
			retry.OnRetry(func(n uint, err error) {
				log.Info("datasource connection - failed to connect")
				log.Info(fmt.Sprintf("datasource connection - trying reconnection to %s", datasource.context.Server()))
				//log.Info(fmt.Sprintf("datasource connection - trying reconnection to %s/%s", datasource.context.Server(), datasource.context.Service()))
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
	if datasource.database, err = mongo.Connect(datasource.opts.ApplyURI(datasource.context.Url())); err != nil {
		log.Error(err.Error())
		return ErrDBConnectionFailed(err)
	}
	log.Info(fmt.Sprintf("datasource connection - connected to %s", datasource.context.Server()))
	//log.Info(fmt.Sprintf("datasource connection - connected to %s/%s", datasource.context.Server(), datasource.context.Service()))
	return nil
}

func (datasource *connection) Close(ctx context.Context) {

	if datasource.database != nil {
		log.Debug("datasource connection - closing connection")
		if err := datasource.database.Disconnect(ctx); err != nil {
			log.Error(fmt.Sprintf("datasource connection - failed to close connection to %s: %s", datasource.context.Server(), err.Error()))
			//log.Error(fmt.Sprintf("datasource connection - failed to close connection to %s/%s: %s", datasource.context.Server(), datasource.context.Service(), err.Error()))
		}
	}
	datasource.database = nil
	log.Debug(fmt.Sprintf("datasource connection - closed connection to %s", datasource.context.Server()))
	//log.Debug(fmt.Sprintf("datasource connection - closed connection to %s/%s", datasource.context.Server(), datasource.context.Service()))
}

func (datasource *connection) Context() Context {
	return datasource.context
}
