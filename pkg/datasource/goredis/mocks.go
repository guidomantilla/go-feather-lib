// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/datasource/goredis/types.go
//
// Generated by this command:
//
//	mockgen -package=goredis -destination ../pkg/datasource/goredis/mocks.go -source ../pkg/datasource/goredis/types.go
//

// Package goredis is a generated GoMock package.
package goredis

import (
	context "context"
	reflect "reflect"

	redis "github.com/redis/go-redis/v9"
	gomock "go.uber.org/mock/gomock"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
	isgomock struct{}
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Password mocks base method.
func (m *MockContext) Password() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Password")
	ret0, _ := ret[0].(string)
	return ret0
}

// Password indicates an expected call of Password.
func (mr *MockContextMockRecorder) Password() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Password", reflect.TypeOf((*MockContext)(nil).Password))
}

// Server mocks base method.
func (m *MockContext) Server() any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Server")
	ret0, _ := ret[0].(any)
	return ret0
}

// Server indicates an expected call of Server.
func (mr *MockContextMockRecorder) Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Server", reflect.TypeOf((*MockContext)(nil).Server))
}

// Service mocks base method.
func (m *MockContext) Service() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Service")
	ret0, _ := ret[0].(string)
	return ret0
}

// Service indicates an expected call of Service.
func (mr *MockContextMockRecorder) Service() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Service", reflect.TypeOf((*MockContext)(nil).Service))
}

// Url mocks base method.
func (m *MockContext) Url() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Url")
	ret0, _ := ret[0].(string)
	return ret0
}

// Url indicates an expected call of Url.
func (mr *MockContextMockRecorder) Url() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Url", reflect.TypeOf((*MockContext)(nil).Url))
}

// User mocks base method.
func (m *MockContext) User() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(string)
	return ret0
}

// User indicates an expected call of User.
func (mr *MockContextMockRecorder) User() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockContext)(nil).User))
}

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
	isgomock struct{}
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConnection) Close(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close", ctx)
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close), ctx)
}

// Connect mocks base method.
func (m *MockConnection) Connect(ctx context.Context) (redis.UniversalClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx)
	ret0, _ := ret[0].(redis.UniversalClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockConnectionMockRecorder) Connect(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockConnection)(nil).Connect), ctx)
}

// Context mocks base method.
func (m *MockConnection) Context() Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockConnectionMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockConnection)(nil).Context))
}

// Set mocks base method.
func (m *MockConnection) Set(key string, value any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set.
func (mr *MockConnectionMockRecorder) Set(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockConnection)(nil).Set), key, value)
}