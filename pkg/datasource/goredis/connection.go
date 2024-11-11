package goredis

import (
	"context"
	"fmt"

	retry "github.com/avast/retry-go/v4"
	redis "github.com/redis/go-redis/v9"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/common/utils"
)

type connection struct {
	context       Context
	database      redis.UniversalClient
	clusterConfig *redis.UniversalOptions
}

func NewConnection(context Context, options ...ConnectionOptions) Connection {
	assert.NotNil(context, "starting up - error setting up datasource connection: context is nil")

	connection := &connection{
		context:  context,
		database: nil,
		clusterConfig: &redis.UniversalOptions{
			Addrs:    context.Server().([]string),
			Username: context.User(),
			Password: context.Password(),
			DB:       0,
			Protocol: 2,
		},
	}

	for _, option := range options {
		option(connection)
	}

	return connection
}

func (datasource *connection) Connect(ctx context.Context) (redis.UniversalClient, error) {

	if datasource.database == nil {

		err := retry.Do(datasource.connect, retry.Attempts(5),
			retry.OnRetry(func(n uint, err error) {
				log.Info(ctx, "datasource connection - failed to connect")
				log.Info(ctx, fmt.Sprintf("datasource connection - trying reconnection to %s/%s", datasource.context.Server(), datasource.context.Service()))
			}),
		)

		if err != nil {
			log.Error(ctx, err.Error())
			return nil, err
		}
	}

	log.Info(ctx, fmt.Sprintf("datasource connection - connected to %s/%s", datasource.context.Server(), datasource.context.Service()))

	return datasource.database, nil
}

func (datasource *connection) connect() error {

	datasource.database = redis.NewUniversalClient(datasource.clusterConfig)
	return nil
}

func (datasource *connection) Close(ctx context.Context) {

	if datasource.database != nil {
		log.Debug(ctx, "datasource connection - closing connection")
		if err := datasource.database.Close(); err != nil {
			log.Error(ctx, fmt.Sprintf("datasource connection - error closing connection: %v", err))
			return
		}
	}
	datasource.database = nil
	log.Debug(ctx, fmt.Sprintf("datasource connection - closed connection to %s/%s", datasource.context.Server(), datasource.context.Service()))
}

func (datasource *connection) Context() Context {
	return datasource.context
}

func (datasource *connection) Set(property string, value any) {
	if utils.IsEmpty(property) || utils.IsEmpty(value) {
		return
	}
}
