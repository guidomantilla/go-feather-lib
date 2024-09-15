package datasource

import (
	"fmt"

	retry "github.com/avast/retry-go/v4"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultDatasource struct {
	datasourceContext DatasourceContext
	database          *gorm.DB
	dialector         gorm.Dialector
	opts              []gorm.Option
}

func NewDefaultDatasource(datasourceContext DatasourceContext, dialector gorm.Dialector, opts ...gorm.Option) *DefaultDatasource {

	if datasourceContext == nil {
		log.Fatal("starting up - error setting up datasource: datasourceContext is nil")
	}

	return &DefaultDatasource{
		datasourceContext: datasourceContext,
		database:          nil,
		dialector:         dialector,
		opts:              opts,
	}
}

func (datasource *DefaultDatasource) Connect() (*gorm.DB, error) {

	if datasource.database == nil {

		err := retry.Do(datasource.connect, retry.Attempts(5),
			retry.OnRetry(func(n uint, err error) {
				log.Info("datasource connection - failed to connect")
				log.Info(fmt.Sprintf("datasource connection - trying reconnection to %s/%s", datasource.datasourceContext.Server(), datasource.datasourceContext.Service()))
			}),
		)

		if err != nil {
			return nil, err
		}
	}

	return datasource.database, nil
}

func (datasource *DefaultDatasource) connect() error {

	var err error
	if datasource.database, err = gorm.Open(datasource.dialector, datasource.opts...); err != nil {
		log.Error(err.Error())
		return ErrDBConnectionFailed(err)
	}
	log.Info(fmt.Sprintf("datasource connection - connected to %s/%s", datasource.datasourceContext.Server(), datasource.datasourceContext.Service()))

	return nil
}

func (datasource *DefaultDatasource) Close() {

	if datasource.database != nil {
		log.Debug("datasource connection - closing connection")
		sqlDB, _ := datasource.database.DB()
		if err := sqlDB.Close(); err != nil {
			log.Error(fmt.Sprintf("datasource connection - failed to close connection to %s/%s: %s", datasource.datasourceContext.Server(), datasource.datasourceContext.Service(), err.Error()))
		}
	}
	datasource.database = nil
	log.Debug(fmt.Sprintf("datasource connection - closed connection to to %s/%s", datasource.datasourceContext.Server(), datasource.datasourceContext.Service()))
}

func (datasource *DefaultDatasource) DatasourceContext() DatasourceContext {
	return datasource.datasourceContext
}
