package log

import (
	"context"
	"log/slog"
	"reflect"
	"testing"
)

func TestNewFanoutHandler(t *testing.T) {
	type args struct {
		handlers []slog.Handler
	}
	tests := []struct {
		name string
		args args
		want slog.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFanoutHandler(tt.args.handlers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFanoutHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFanoutHandler_Enabled(t *testing.T) {
	type args struct {
		ctx context.Context
		l   slog.Level
	}
	tests := []struct {
		name string
		h    *FanoutHandler
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Enabled(tt.args.ctx, tt.args.l); got != tt.want {
				t.Errorf("FanoutHandler.Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFanoutHandler_Handle(t *testing.T) {
	type args struct {
		ctx context.Context
		r   slog.Record
	}
	tests := []struct {
		name    string
		h       *FanoutHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Handle(tt.args.ctx, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("FanoutHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFanoutHandler_WithAttrs(t *testing.T) {
	type args struct {
		attrs []slog.Attr
	}
	tests := []struct {
		name string
		h    *FanoutHandler
		args args
		want slog.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.WithAttrs(tt.args.attrs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FanoutHandler.WithAttrs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFanoutHandler_WithGroup(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		h    *FanoutHandler
		args args
		want slog.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.WithGroup(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FanoutHandler.WithGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
