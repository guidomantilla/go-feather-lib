package datasource

import (
	"database/sql"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewDefaultDatasource(t *testing.T) {
	var err error
	var db *sql.DB
	if db, _, err = sqlmock.New(sqlmock.MonitorPingsOption(true)); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	datasourceCtx := &DefaultDatasourceContext{
		url: "some_usersome_passsome_serversome_service",
	}
	datasource := &DefaultDatasource{
		url:      datasourceCtx.url,
		database: nil,
		dialector: mysql.New(mysql.Config{
			Conn:                      db,
			DriverName:                "mock",
			SkipInitializeWithVersion: true,
		}),
		opts: []gorm.Option{&gorm.Config{}},
	}

	type args struct {
		datasourceContext DatasourceContext
		dialector         gorm.Dialector
		config            *gorm.Config
	}
	tests := []struct {
		name string
		args args
		want *DefaultDatasource
	}{
		{
			name: "Happy Path",
			args: args{
				datasourceContext: datasourceCtx,
				config:            &gorm.Config{},
			},
			want: datasource,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDefaultDatasource(tt.args.datasourceContext, tt.args.dialector, tt.args.config)
			if !reflect.DeepEqual(got.url, tt.want.url) {
				t.Errorf("NewDefaultDatasource() = %v, want %v", got.url, tt.want.url)
			}
			if !reflect.DeepEqual(got.database, tt.want.database) {
				t.Errorf("NewDefaultDatasource() = %v, want %v", got.database, tt.want.database)
			}
		})
	}
}

func TestDefaultDatasource_GetDatabase(t *testing.T) {

}
