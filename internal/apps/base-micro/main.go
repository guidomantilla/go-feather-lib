package main

import (
	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	"os"
)

func main() {

	_ = os.Setenv("LOG_LEVEL", "TRACE")
	cserver.Run("nats-micro", "1.0.0", func(application cserver.Application) error {

		return nil
	})
}
