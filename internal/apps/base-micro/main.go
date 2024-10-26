package main

import (
	"os"

	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
)

func main() {

	_ = os.Setenv("LOG_LEVEL", "TRACE")
	cserver.Run("base-micro", "1.0.0", func(application cserver.Application) error {

		return nil
	})
}
