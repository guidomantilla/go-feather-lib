package gorm

import (
	"context"
	"fmt"

	retry "github.com/avast/retry-go/v4"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type connection struct {
	context  Context
	database *gorm.DB
	openFn   OpenFn
	opts     []gorm.Option
}

func NewConnection(context Context, openFn OpenFn, opts ...gorm.Option) Connection {
	assert.NotNil(context, "starting up - error setting up datasource connection: context is nil")
	assert.NotNil(openFn, "starting up - error setting up datasource connection: open is nil")

	return &connection{
		context:  context,
		database: nil,
		openFn:   openFn,
		opts:     opts,
	}
}

func (datasource *connection) Connect(ctx context.Context) (*gorm.DB, error) {

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

	var err error
	if datasource.database, err = gorm.Open(datasource.openFn(datasource.context.Url()), datasource.opts...); err != nil {
		return ErrDBConnectionFailed(err)
	}

	return nil
}

func (datasource *connection) Close(ctx context.Context) {

	if datasource.database != nil {
		log.Debug(ctx, "datasource connection - closing connection")
		sqlDB, _ := datasource.database.DB()
		if err := sqlDB.Close(); err != nil {
			log.Error(ctx, fmt.Sprintf("datasource connection - failed to close connection to %s/%s: %s", datasource.context.Server(), datasource.context.Service(), err.Error()))
		}
	}
	datasource.database = nil
	log.Debug(ctx, fmt.Sprintf("datasource connection - closed connection to %s/%s", datasource.context.Server(), datasource.context.Service()))
}

func (datasource *connection) Context() Context {
	return datasource.context
}
