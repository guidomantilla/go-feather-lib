package mocks

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	datasource2 "github.com/guidomantilla/go-feather-lib/pkg/common/datasource"
)

func BuildMockGormTransactionHandler() (datasource2.TransactionHandler, sqlmock.Sqlmock) {
	db, mock := BuildMockGormDatasource()
	return datasource2.NewTransactionHandler(db), mock
}

func BuildMockGormDatasource() (datasource2.Datasource, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mock",
		SkipInitializeWithVersion: true,
	})
	datasourceContext := datasource2.NewDefaultDatasourceContext("some url", "some username", "some password", "some server", "some service")
	datasrc := datasource2.NewDefaultDatasource(datasourceContext, dialector, &gorm.Config{})
	return datasrc, mock
}
