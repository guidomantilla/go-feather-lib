package datasource

import (
	"reflect"
	"testing"
)

func TestNewDefaultDatasourceContext(t *testing.T) {

	datasourceCtx := &DefaultDatasourceContext{
		url:     "some_usersome_passsome_serversome_service",
		server:  "some_server",
		service: "some_service",
	}
	type args struct {
		url      string
		username string
		password string
		server   string
		service  string
	}
	tests := []struct {
		name string
		args args
		want *DefaultDatasourceContext
	}{
		{
			name: "Happy Path",
			args: args{
				url:      ":username:password:server:service",
				username: "some_user",
				password: "some_pass",
				server:   "some_server",
				service:  "some_service",
			},
			want: datasourceCtx,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDefaultDatasourceContext(tt.args.url, tt.args.username, tt.args.password, tt.args.server, tt.args.service)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultDatasourceContext() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got.GetUrl(), tt.want.GetUrl()) {
				t.Errorf("NewDefaultDatasourceContext() = %v, want %v", got.GetUrl(), tt.want.GetUrl())
			}
		})
	}
}
