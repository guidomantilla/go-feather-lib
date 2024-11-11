package gorm

import (
	"time"

	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/gorm/logger"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

func Logger() logger.Interface {
	return slogGorm.New(slogGorm.WithHandler(log.Handler()), slogGorm.WithTraceAll(), slogGorm.WithRecordNotFoundError())
}

func LoggerWithSlowThreshold(threshold time.Duration) logger.Interface {
	return slogGorm.New(slogGorm.WithHandler(log.Handler()), slogGorm.WithTraceAll(), slogGorm.WithRecordNotFoundError(),
		slogGorm.WithSlowThreshold(threshold))
}
