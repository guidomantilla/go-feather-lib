package mongo

import (
	"context"
	"log/slog"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type loggerWrapper struct {
	logger *slog.Logger
	mu     sync.Mutex
}

func (logger *loggerWrapper) Info(level int, msg string, keysAndValues ...interface{}) {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	if options.LogLevel(level+1) == options.LogLevelDebug {
		logger.logger.Log(context.Background(), log.SlogLevelDebug.ToSlogLevel(), msg, keysAndValues...)
	} else {
		logger.logger.Log(context.Background(), log.SlogLevelInfo.ToSlogLevel(), msg, keysAndValues...)
	}
}

func (logger *loggerWrapper) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.logger.Log(context.Background(), log.SlogLevelError.ToSlogLevel(), msg, keysAndValues...)
}

func Logger() options.LogSink {
	return &loggerWrapper{
		logger: log.Ref(),
	}
}
