// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/server/types.go
//
// Generated by this command:
//
//	mockgen -package=server -destination ../pkg/server/mocks.go -source ../pkg/server/types.go github.com/qmdx00/lifecycle Server
//

// Package server is a generated GoMock package.
package server

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDispatcher is a mock of Dispatcher interface.
type MockDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherMockRecorder
}

// MockDispatcherMockRecorder is the mock recorder for MockDispatcher.
type MockDispatcherMockRecorder struct {
	mock *MockDispatcher
}

// NewMockDispatcher creates a new mock instance.
func NewMockDispatcher(ctrl *gomock.Controller) *MockDispatcher {
	mock := &MockDispatcher{ctrl: ctrl}
	mock.recorder = &MockDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDispatcher) EXPECT() *MockDispatcherMockRecorder {
	return m.recorder
}

// Dispatch mocks base method.
func (m *MockDispatcher) Dispatch(message any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dispatch", message)
}

// Dispatch indicates an expected call of Dispatch.
func (mr *MockDispatcherMockRecorder) Dispatch(message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockDispatcher)(nil).Dispatch), message)
}

// ListenAndDispatch mocks base method.
func (m *MockDispatcher) ListenAndDispatch() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAndDispatch")
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenAndDispatch indicates an expected call of ListenAndDispatch.
func (mr *MockDispatcherMockRecorder) ListenAndDispatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndDispatch", reflect.TypeOf((*MockDispatcher)(nil).ListenAndDispatch))
}

// Run mocks base method.
func (m *MockDispatcher) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockDispatcherMockRecorder) Run(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockDispatcher)(nil).Run), ctx)
}

// Stop mocks base method.
func (m *MockDispatcher) Stop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockDispatcherMockRecorder) Stop(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDispatcher)(nil).Stop), ctx)
}
