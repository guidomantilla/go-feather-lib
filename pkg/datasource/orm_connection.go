package datasource

import (
	"fmt"

	retry "github.com/avast/retry-go/v4"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type OrmConnection struct {
	context   StoreContext
	database  *gorm.DB
	dialector gorm.Dialector
	opts      []gorm.Option
}

func NewOrmConnection(context StoreContext, dialector gorm.Dialector, opts ...gorm.Option) *OrmConnection {

	if context == nil {
		log.Fatal("starting up - error setting up datasource connection: context is nil")
	}

	return &OrmConnection{
		context:   context,
		database:  nil,
		dialector: dialector,
		opts:      opts,
	}
}

func (datasource *OrmConnection) Connect() (*gorm.DB, error) {

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

func (datasource *OrmConnection) connect() error {

	var err error
	if datasource.database, err = gorm.Open(datasource.dialector, datasource.opts...); err != nil {
		log.Error(err.Error())
		return ErrDBConnectionFailed(err)
	}
	log.Info(fmt.Sprintf("datasource connection - connected to %s/%s", datasource.context.Server(), datasource.context.Service()))

	return nil
}

func (datasource *OrmConnection) Close() {

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

func (datasource *OrmConnection) Context() StoreContext {
	return datasource.context
}
