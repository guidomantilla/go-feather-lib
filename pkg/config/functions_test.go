package config

import (
	"context"
	"testing"

	envconfig "github.com/sethvargo/go-envconfig"

	"github.com/guidomantilla/go-feather-lib/pkg/environment"
)

func TestProcess(t *testing.T) {
	env := environment.Default()
	config := &envconfig.Config{
		Target: &struct {
			SomeEnvVar string `envconfig:"SOME_ENV_VAR"`
		}{
			SomeEnvVar: "some-value",
		},
	}

	type args struct {
		ctx         context.Context
		environment environment.Environment
		config      *envconfig.Config
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
				config:      config,
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
