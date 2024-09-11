package config

import (
	"context"
	"testing"

	"github.com/xorcare/pointer"

	"github.com/guidomantilla/go-feather-lib/pkg/environment"
)

func TestProcess(t *testing.T) {
	env := environment.Default()

	type args struct {
		ctx         context.Context
		environment environment.Environment
		config      *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Path",
			args: args{
				ctx:         context.TODO(),
				environment: env,
				config: &Config{
					Host:                 pointer.Of("localhost"),
					HttpPort:             nil,
					GrpcPort:             nil,
					TokenSignatureKey:    nil,
					TokenVerificationKey: nil,
					TokenTimeout:         nil,
					DatasourceDriver:     nil,
					DatasourceUsername:   nil,
					DatasourcePassword:   nil,
					DatasourceServer:     nil,
					DatasourceService:    nil,
					DatasourceUrl:        nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Process(tt.args.ctx, tt.args.environment, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
