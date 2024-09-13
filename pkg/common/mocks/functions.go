package mocks

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/datasource"
)

func BuildMockGormTransactionHandler() (datasource.TransactionHandler, sqlmock.Sqlmock) {
	db, mock := BuildMockGormDatasource()
	return datasource.NewTransactionHandler(db), mock
}

func BuildMockGormDatasource() (datasource.Datasource, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mock",
		SkipInitializeWithVersion: true,
	})
	datasourceContext := datasource.NewDefaultDatasourceContext("some url", "some username", "some password", "some server", "some service")
	datasrc := datasource.NewDefaultDatasource(datasourceContext, dialector, &gorm.Config{})
	return datasrc, mock
}
