package log

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"reflect"
	"testing"
)

func TestNewSlogLogger(t *testing.T) {
	type args struct {
		level   CustomSlogLevel
		writers []io.Writer
	}
	tests := []struct {
		name string
		args args
		want *SlogLogger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSlogLogger(tt.args.level, tt.args.writers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSlogLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlogLogger_Debug(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Debug(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_Info(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Info(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_Warn(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Warn(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_Error(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Error(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_Fatal(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Fatal(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_RetrieveLogger(t *testing.T) {
	tests := []struct {
		name   string
		logger *SlogLogger
		want   any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.logger.RetrieveLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlogLogger.RetrieveLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomSlogLevel_String(t *testing.T) {
	tests := []struct {
		name string
		enum CustomSlogLevel
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.String(); got != tt.want {
				t.Errorf("CustomSlogLevel.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomSlogLevel_ToSlogLevel(t *testing.T) {
	tests := []struct {
		name string
		enum CustomSlogLevel
		want slog.Level
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ToSlogLevel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomSlogLevel.ToSlogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomSlogLevel_ValueFromName(t *testing.T) {
	type args struct {
		slogLevel string
	}
	tests := []struct {
		name string
		enum CustomSlogLevel
		args args
		want CustomSlogLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromName(tt.args.slogLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomSlogLevel.ValueFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomSlogLevel_ValueFromCardinal(t *testing.T) {
	type args struct {
		slogLevel int
	}
	tests := []struct {
		name string
		enum CustomSlogLevel
		args args
		want CustomSlogLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromCardinal(tt.args.slogLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomSlogLevel.ValueFromCardinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomSlogLevel_ValueFromSlogLevel(t *testing.T) {
	type args struct {
		slogLevel slog.Level
	}
	tests := []struct {
		name string
		enum CustomSlogLevel
		args args
		want CustomSlogLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromSlogLevel(tt.args.slogLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomSlogLevel.ValueFromSlogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomSlogFormat_String(t *testing.T) {
	tests := []struct {
		name string
		enum CustomSlogFormat
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.String(); got != tt.want {
				t.Errorf("CustomSlogFormat.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomSlogFormat_ValueFromName(t *testing.T) {
	type args struct {
		loggerFormat string
	}
	tests := []struct {
		name string
		enum CustomSlogFormat
		args args
		want CustomSlogFormat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromName(tt.args.loggerFormat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomSlogFormat.ValueFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomSlogFormat_ValueFromCardinal(t *testing.T) {
	type args struct {
		loggerFormat int
	}
	tests := []struct {
		name string
		enum CustomSlogFormat
		args args
		want CustomSlogFormat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromCardinal(tt.args.loggerFormat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomSlogFormat.ValueFromCardinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomSlogFormat_Handler(t *testing.T) {
	type args struct {
		opts *slog.HandlerOptions
	}
	tests := []struct {
		name  string
		enum  CustomSlogFormat
		args  args
		want  slog.Handler
		wantW string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if got := tt.enum.Handler(w, tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomSlogFormat.Handler() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("CustomSlogFormat.Handler() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
