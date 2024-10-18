package log

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type slogLogger struct {
	internal *slog.Logger
}

func New(level SlogLevel, writers ...io.Writer) Logger[*slog.Logger] {
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
	internal := slog.New(NewSlogFanoutHandler(handlers...))
	slog.SetDefault(internal)

	return &slogLogger{internal: internal}
}

func (logger *slogLogger) Trace(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelTrace.ToSlogLevel(), msg, args...)
}

func (logger *slogLogger) Debug(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelDebug.ToSlogLevel(), msg, args...)
}

func (logger *slogLogger) Info(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelInfo.ToSlogLevel(), msg, args...)
}

func (logger *slogLogger) Warn(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelWarning.ToSlogLevel(), msg, args...)
}

func (logger *slogLogger) Error(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelError.ToSlogLevel(), msg, args...)
}

func (logger *slogLogger) Fatal(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelFatal.ToSlogLevel(), msg, args...)
	os.Exit(1)
}

func (logger *slogLogger) Logger() *slog.Logger {
	return logger.internal
}

//

const (
	SlogLevelTrace SlogLevel = iota
	SlogLevelDebug
	SlogLevelInfo
	SlogLevelWarning
	SlogLevelError
	SlogLevelFatal
	SlogLevelOff
)

type SlogLevel int

func (enum SlogLevel) String() string {

	switch enum {
	case SlogLevelTrace:
		return "TRACE"
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

func (enum SlogLevel) ToSlogLevel() slog.Level {
	switch enum {
	case SlogLevelTrace:
		return slog.Level(-8)
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

func (enum SlogLevel) ValueFromName(slogLevel string) SlogLevel {
	switch slogLevel {
	case "TRACE":
		return SlogLevelTrace
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

func (enum SlogLevel) ValueFromCardinal(slogLevel int) SlogLevel {
	switch slogLevel {
	case int(SlogLevelTrace):
		return SlogLevelTrace
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

func (enum SlogLevel) ValueFromSlogLevel(slogLevel slog.Level) SlogLevel {
	switch slogLevel {
	case slog.Level(-8):
		return SlogLevelTrace
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
	SlogTextFormat SlogFormat = iota
	SlogJsonFormat
)

type SlogFormat int

func (enum SlogFormat) String() string {
	switch enum {
	case SlogTextFormat:
		return "TEXT"
	case SlogJsonFormat:
		return "JSON"
	}
	return "TEXT"
}

func (enum SlogFormat) ValueFromName(loggerFormat string) SlogFormat {
	switch loggerFormat {
	case "TEXT":
		return SlogTextFormat
	case "JSON":
		return SlogJsonFormat
	}
	return SlogTextFormat
}

func (enum SlogFormat) ValueFromCardinal(loggerFormat int) SlogFormat {
	switch loggerFormat {
	case int(SlogTextFormat):
		return SlogTextFormat
	case int(SlogJsonFormat):
		return SlogJsonFormat
	}
	return SlogTextFormat
}

func (enum SlogFormat) Handler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	switch enum {
	case SlogTextFormat:
		return slog.NewTextHandler(w, opts)
	case SlogJsonFormat:
		return slog.NewJSONHandler(w, opts)
	}
	return slog.NewTextHandler(w, opts)
}
