package config

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"

	"github.com/guidomantilla/go-feather-lib/pkg/environment"
)

func Process(ctx context.Context, environment environment.Environment, config *envconfig.Config) error {
	config.Lookuper = &EnvironmentLookup{
		environment: environment,
	}
	return envconfig.ProcessWith(ctx, config)
}
