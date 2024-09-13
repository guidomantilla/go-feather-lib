package config

import (
	"context"
	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"

	envconfig "github.com/sethvargo/go-envconfig"
)

func Process(ctx context.Context, environment environment.Environment, cfg *Config) error {

	internalConfig := &envconfig.Config{Target: cfg, Lookuper: &EnvironmentLookup{environment: environment}}
	if err := envconfig.ProcessWith(ctx, internalConfig); err != nil {
		return err
	}

	return nil
}
