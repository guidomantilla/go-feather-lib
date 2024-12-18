// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/common/log/types.go
//
// Generated by this command:
//
//	mockgen -package=log -destination ../pkg/common/log/mocks.go -source ../pkg/common/log/types.go
//

// Package log is a generated GoMock package.
package log

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder[T]
	isgomock struct{}
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder[T any] struct {
	mock *MockLogger[T]
}

// NewMockLogger creates a new mock instance.
func NewMockLogger[T any](ctrl *gomock.Controller) *MockLogger[T] {
	mock := &MockLogger[T]{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger[T]) EXPECT() *MockLoggerMockRecorder[T] {
	return m.recorder
}

// Debug mocks base method.
func (m *MockLogger[T]) Debug(ctx context.Context, msg string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder[T]) Debug(ctx, msg any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger[T])(nil).Debug), varargs...)
}

// Error mocks base method.
func (m *MockLogger[T]) Error(ctx context.Context, msg string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder[T]) Error(ctx, msg any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger[T])(nil).Error), varargs...)
}

// Fatal mocks base method.
func (m *MockLogger[T]) Fatal(ctx context.Context, msg string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockLoggerMockRecorder[T]) Fatal(ctx, msg any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger[T])(nil).Fatal), varargs...)
}

// Info mocks base method.
func (m *MockLogger[T]) Info(ctx context.Context, msg string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder[T]) Info(ctx, msg any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger[T])(nil).Info), varargs...)
}

// Logger mocks base method.
func (m *MockLogger[T]) Logger() T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(T)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockLoggerMockRecorder[T]) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockLogger[T])(nil).Logger))
}

// Trace mocks base method.
func (m *MockLogger[T]) Trace(ctx context.Context, msg string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Trace", varargs...)
}

// Trace indicates an expected call of Trace.
func (mr *MockLoggerMockRecorder[T]) Trace(ctx, msg any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trace", reflect.TypeOf((*MockLogger[T])(nil).Trace), varargs...)
}

// Warn mocks base method.
func (m *MockLogger[T]) Warn(ctx context.Context, msg string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerMockRecorder[T]) Warn(ctx, msg any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLogger[T])(nil).Warn), varargs...)
}
