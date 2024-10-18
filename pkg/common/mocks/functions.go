package mocks

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/datasource"
)

func BuildMockGormTransactionHandler() (datasource.TransactionHandler[*gorm.DB], sqlmock.Sqlmock) {
	db, mock := BuildMockGormDatasource()
	return datasource.NewOrmTransactionHandler(db), mock
}

func BuildMockGormDatasource() (datasource.Connection[*gorm.DB], sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mock",
		SkipInitializeWithVersion: true,
	})
	context := datasource.NewContext("some url", "some username", "some password", "some server", "some service")
	connection := datasource.NewConnection(context, dialector, &gorm.Config{})
	return connection, mock
}
