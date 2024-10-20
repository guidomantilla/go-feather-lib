package datasource

import (
	"fmt"

	retry "github.com/avast/retry-go/v4"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type connection struct {
	context   Context
	database  *gorm.DB
	dialector gorm.Dialector
	opts      []gorm.Option
}

func NewConnection(context Context, dialector gorm.Dialector, opts ...gorm.Option) Connection[*gorm.DB] {
	assert.NotNil(context, "starting up - error setting up datasource connection: context is nil")
	assert.NotNil(dialector, "starting up - error setting up datasource connection: dialector is nil")
	//assert.NotEmpty(opts, "starting up - error setting up datasource connection: opts is empty")

	return &connection{
		context:   context,
		database:  nil,
		dialector: dialector,
		opts:      opts,
	}
}

func (datasource *connection) Connect() (*gorm.DB, error) {

	if datasource.database == nil {

		err := retry.Do(datasource.connect, retry.Attempts(5),
			retry.OnRetry(func(n uint, err error) {
				log.Info("datasource connection - failed to connect")
				log.Info(fmt.Sprintf("datasource connection - trying reconnection to %s/%s", datasource.context.Server(), datasource.context.Service()))
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
	if datasource.database, err = gorm.Open(datasource.dialector, datasource.opts...); err != nil {
		log.Error(err.Error())
		return ErrDBConnectionFailed(err)
	}
	log.Info(fmt.Sprintf("datasource connection - connected to %s/%s", datasource.context.Server(), datasource.context.Service()))

	return nil
}

func (datasource *connection) Close() {

	if datasource.database != nil {
		log.Debug("datasource connection - closing connection")
		sqlDB, _ := datasource.database.DB()
		if err := sqlDB.Close(); err != nil {
			log.Error(fmt.Sprintf("datasource connection - failed to close connection to %s/%s: %s", datasource.context.Server(), datasource.context.Service(), err.Error()))
		}
	}
	datasource.database = nil
	log.Debug(fmt.Sprintf("datasource connection - closed connection to to %s/%s", datasource.context.Server(), datasource.context.Service()))
}

func (datasource *connection) Context() Context {
	return datasource.context
}
