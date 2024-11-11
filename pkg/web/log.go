package web

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

func Logger() gin.HandlerFunc {
	return sloggin.New(log.Ref().WithGroup("http"))
}
