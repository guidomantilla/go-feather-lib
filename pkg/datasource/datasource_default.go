package datasource

import (
	"fmt"
	"github.com/avast/retry-go/v4"

	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultDatasource struct {
	url       string
	server    string
	service   string
	database  *gorm.DB
	dialector gorm.Dialector
	opts      []gorm.Option
}

func NewDefaultDatasource(datasourceContext DatasourceContext, dialector gorm.Dialector, opts ...gorm.Option) *DefaultDatasource {

	if datasourceContext == nil {
		log.Fatal("starting up - error setting up datasource: datasourceContext is nil")
	}

	return &DefaultDatasource{
		url:       datasourceContext.GetUrl(),
		server:    datasourceContext.GetServer(),
		service:   datasourceContext.GetService(),
		database:  nil,
		dialector: dialector,
		opts:      opts,
	}
}

func (datasource *DefaultDatasource) GetDatabase() (*gorm.DB, error) {

	if datasource.database == nil {

		err := retry.Do(datasource.Connect, retry.Attempts(5),
			retry.OnRetry(func(n uint, err error) {
				log.Info("connection - failed to connect")
				log.Info(fmt.Sprintf("connection - retrying connection to %s/%s", datasource.server, datasource.service))
			}),
		)

		if err != nil {
			return nil, err
		}
	}

	return datasource.database, nil
}

func (datasource *DefaultDatasource) Connect() error {

	var err error
	if datasource.database, err = gorm.Open(datasource.dialector, datasource.opts...); err != nil {
		log.Error(err.Error())
		return ErrDBConnectionFailed(err)
	}
	log.Debug(fmt.Sprintf("connection - connected to %s/%s", datasource.server, datasource.service))

	return nil
}
