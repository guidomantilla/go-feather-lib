package server

import (
	"syscall"

	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

func Run(name string, version string, fn func(application Application) error) {
	log.Slog()
	environment.Load()

	app := lifecycle.NewApp(
		lifecycle.WithName(name), lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	app.Attach(server.BuildDefaultServer())
	if err := fn(app); err != nil {
		log.Fatal(err.Error())
	}

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
