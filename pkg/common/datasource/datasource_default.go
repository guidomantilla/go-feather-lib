package datasource

import (
	"fmt"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"

	"gorm.io/gorm"
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

	var err error

	if datasource.database == nil {
		if datasource.database, err = gorm.Open(datasource.dialector, datasource.opts...); err != nil {
			log.Error(err.Error())
			return nil, ErrDBConnectionFailed(err)
		}
		log.Debug(fmt.Sprintf("connection - connected to %s/%s", datasource.server, datasource.service))
	}

	return datasource.database, nil
}
