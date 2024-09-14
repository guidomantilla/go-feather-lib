package log

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type SlogLogger struct {
	internal *slog.Logger
}

func NewSlogLogger(level CustomSlogLevel, writers ...io.Writer) *SlogLogger {
	opts := &slog.HandlerOptions{
		Level: level.ToSlogLevel(),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				a.Value = slog.StringValue(SlogLevelOff.ValueFromSlogLevel(level).String())
			}
			return a
		},
	}

	handlers := make([]slog.Handler, 0)
	handlers = append(handlers, SlogTextFormat.Handler(os.Stdout, opts))
	for _, writer := range writers {
		handlers = append(handlers, SlogJsonFormat.Handler(writer, opts))
	}
	internal := slog.New(NewFanoutHandler(handlers...))
	slog.SetDefault(internal)
	slogLogger := &SlogLogger{internal: internal}
	return slogLogger
}

func (logger *SlogLogger) Debug(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelDebug.ToSlogLevel(), msg, args...)
}

func (logger *SlogLogger) Info(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelInfo.ToSlogLevel(), msg, args...)
}

func (logger *SlogLogger) Warn(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelWarning.ToSlogLevel(), msg, args...)
}

func (logger *SlogLogger) Error(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelError.ToSlogLevel(), msg, args...)
}

func (logger *SlogLogger) Fatal(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelFatal.ToSlogLevel(), msg, args...)
	os.Exit(1)
}

func (logger *SlogLogger) Logger() any {
	return logger.internal
}

//

const (
	SlogLevelDebug CustomSlogLevel = iota
	SlogLevelInfo
	SlogLevelWarning
	SlogLevelError
	SlogLevelFatal
	SlogLevelOff
)

type CustomSlogLevel int

func (enum CustomSlogLevel) String() string {

	switch enum {
	case SlogLevelDebug:
		return "DEBUG"
	case SlogLevelInfo:
		return "INFO"
	case SlogLevelWarning:
		return "WARN"
	case SlogLevelError:
		return "ERROR"
	case SlogLevelFatal:
		return "FATAL"
	case SlogLevelOff:
		return "OFF"
	}
	return "OFF"
}

func (enum CustomSlogLevel) ToSlogLevel() slog.Level {
	switch enum {
	case SlogLevelDebug:
		return slog.LevelDebug
	case SlogLevelInfo:
		return slog.LevelInfo
	case SlogLevelWarning:
		return slog.LevelWarn
	case SlogLevelError:
		return slog.LevelError
	case SlogLevelFatal:
		return slog.Level(12)
	case SlogLevelOff:
		return slog.Level(16)
	}
	return slog.Level(16)
}

func (enum CustomSlogLevel) ValueFromName(slogLevel string) CustomSlogLevel {
	switch slogLevel {
	case "DEBUG":
		return SlogLevelDebug
	case "INFO":
		return SlogLevelInfo
	case "WARN":
		return SlogLevelWarning
	case "ERROR":
		return SlogLevelError
	case "FATAL":
		return SlogLevelFatal
	case "OFF":
		return SlogLevelOff
	}
	return SlogLevelOff
}

func (enum CustomSlogLevel) ValueFromCardinal(slogLevel int) CustomSlogLevel {
	switch slogLevel {
	case int(SlogLevelDebug):
		return SlogLevelDebug
	case int(SlogLevelInfo):
		return SlogLevelInfo
	case int(SlogLevelWarning):
		return SlogLevelWarning
	case int(SlogLevelError):
		return SlogLevelError
	case int(SlogLevelFatal):
		return SlogLevelFatal
	case int(SlogLevelOff):
		return SlogLevelOff
	}
	return SlogLevelOff
}

func (enum CustomSlogLevel) ValueFromSlogLevel(slogLevel slog.Level) CustomSlogLevel {
	switch slogLevel {
	case slog.LevelDebug:
		return SlogLevelDebug
	case slog.LevelInfo:
		return SlogLevelInfo
	case slog.LevelWarn:
		return SlogLevelWarning
	case slog.LevelError:
		return SlogLevelError
	case slog.Level(12):
		return SlogLevelFatal
	case slog.Level(16):
		return SlogLevelOff
	}
	return SlogLevelOff
}

//

const (
	SlogTextFormat CustomSlogFormat = iota
	SlogJsonFormat
)

type CustomSlogFormat int

func (enum CustomSlogFormat) String() string {
	switch enum {
	case SlogTextFormat:
		return "TEXT"
	case SlogJsonFormat:
		return "JSON"
	}
	return "TEXT"
}

func (enum CustomSlogFormat) ValueFromName(loggerFormat string) CustomSlogFormat {
	switch loggerFormat {
	case "TEXT":
		return SlogTextFormat
	case "JSON":
		return SlogJsonFormat
	}
	return SlogTextFormat
}

func (enum CustomSlogFormat) ValueFromCardinal(loggerFormat int) CustomSlogFormat {
	switch loggerFormat {
	case int(SlogTextFormat):
		return SlogTextFormat
	case int(SlogJsonFormat):
		return SlogJsonFormat
	}
	return SlogTextFormat
}

func (enum CustomSlogFormat) Handler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	switch enum {
	case SlogTextFormat:
		return slog.NewTextHandler(w, opts)
	case SlogJsonFormat:
		return slog.NewJSONHandler(w, opts)
	}
	return slog.NewTextHandler(w, opts)
}

//

type CustomSlogHandler interface {
}
