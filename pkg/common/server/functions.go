package server

import (
	"context"
	"syscall"

	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

func Run(name string, version string, fn func(ctx context.Context, application Application) error) {
	assert.NotEmpty(name, "server - error running: name is empty")
	assert.NotEmpty(version, "server - error running: version is empty")
	assert.NotNil(fn, "server - error running: function is nil")

	log.Slog()
	environment.Load()

	app := lifecycle.NewApp(
		lifecycle.WithName(name), lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	app.Attach(server.BuildBaseServer())

	ctx := context.Background()
	if err := fn(ctx, app); err != nil {
		log.Fatal(ctx, err.Error())
	}

	if err := app.Run(); err != nil {
		log.Fatal(ctx, err.Error())
	}
}
