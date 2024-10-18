package config

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
)

func Process(ctx context.Context, environment environment.Environment, cfg *Config) error {
	return envconfig.ProcessWith(ctx, &envconfig.Config{Target: cfg, Lookuper: &EnvironmentLookup{environment: environment}})
}
