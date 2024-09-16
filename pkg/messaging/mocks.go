// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/messaging/types.go
//
// Generated by this command:
//
//	mockgen -package=messaging -destination ../pkg/messaging/mocks.go -source ../pkg/messaging/types.go
//

// Package messaging is a generated GoMock package.
package messaging

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMessagingContext is a mock of MessagingContext interface.
type MockMessagingContext struct {
	ctrl     *gomock.Controller
	recorder *MockMessagingContextMockRecorder
}

// MockMessagingContextMockRecorder is the mock recorder for MockMessagingContext.
type MockMessagingContextMockRecorder struct {
	mock *MockMessagingContext
}

// NewMockMessagingContext creates a new mock instance.
func NewMockMessagingContext(ctrl *gomock.Controller) *MockMessagingContext {
	mock := &MockMessagingContext{ctrl: ctrl}
	mock.recorder = &MockMessagingContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessagingContext) EXPECT() *MockMessagingContextMockRecorder {
	return m.recorder
}

// Server mocks base method.
func (m *MockMessagingContext) Server() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Server")
	ret0, _ := ret[0].(string)
	return ret0
}

// Server indicates an expected call of Server.
func (mr *MockMessagingContextMockRecorder) Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Server", reflect.TypeOf((*MockMessagingContext)(nil).Server))
}

// Url mocks base method.
func (m *MockMessagingContext) Url() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Url")
	ret0, _ := ret[0].(string)
	return ret0
}

// Url indicates an expected call of Url.
func (mr *MockMessagingContextMockRecorder) Url() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Url", reflect.TypeOf((*MockMessagingContext)(nil).Url))
}

// MockMessagingConnection is a mock of MessagingConnection interface.
type MockMessagingConnection[T MessagingConnectionTypes] struct {
	ctrl     *gomock.Controller
	recorder *MockMessagingConnectionMockRecorder[T]
}

// MockMessagingConnectionMockRecorder is the mock recorder for MockMessagingConnection.
type MockMessagingConnectionMockRecorder[T MessagingConnectionTypes] struct {
	mock *MockMessagingConnection[T]
}

// NewMockMessagingConnection creates a new mock instance.
func NewMockMessagingConnection[T MessagingConnectionTypes](ctrl *gomock.Controller) *MockMessagingConnection[T] {
	mock := &MockMessagingConnection[T]{ctrl: ctrl}
	mock.recorder = &MockMessagingConnectionMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessagingConnection[T]) EXPECT() *MockMessagingConnectionMockRecorder[T] {
	return m.recorder
}

// Close mocks base method.
func (m *MockMessagingConnection[T]) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockMessagingConnectionMockRecorder[T]) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMessagingConnection[T])(nil).Close))
}

// Connect mocks base method.
func (m *MockMessagingConnection[T]) Connect() (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockMessagingConnectionMockRecorder[T]) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockMessagingConnection[T])(nil).Connect))
}

// MessagingContext mocks base method.
func (m *MockMessagingConnection[T]) MessagingContext() MessagingContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessagingContext")
	ret0, _ := ret[0].(MessagingContext)
	return ret0
}

// MessagingContext indicates an expected call of MessagingContext.
func (mr *MockMessagingConnectionMockRecorder[T]) MessagingContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessagingContext", reflect.TypeOf((*MockMessagingConnection[T])(nil).MessagingContext))
}

// MockMessagingListener is a mock of MessagingListener interface.
type MockMessagingListener[T MessagingListenerTypes] struct {
	ctrl     *gomock.Controller
	recorder *MockMessagingListenerMockRecorder[T]
}

// MockMessagingListenerMockRecorder is the mock recorder for MockMessagingListener.
type MockMessagingListenerMockRecorder[T MessagingListenerTypes] struct {
	mock *MockMessagingListener[T]
}

// NewMockMessagingListener creates a new mock instance.
func NewMockMessagingListener[T MessagingListenerTypes](ctrl *gomock.Controller) *MockMessagingListener[T] {
	mock := &MockMessagingListener[T]{ctrl: ctrl}
	mock.recorder = &MockMessagingListenerMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessagingListener[T]) EXPECT() *MockMessagingListenerMockRecorder[T] {
	return m.recorder
}

// OnMessage mocks base method.
func (m *MockMessagingListener[T]) OnMessage(message T) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnMessage", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnMessage indicates an expected call of OnMessage.
func (mr *MockMessagingListenerMockRecorder[T]) OnMessage(message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnMessage", reflect.TypeOf((*MockMessagingListener[T])(nil).OnMessage), message)
}

// MockMessagingConsumer is a mock of MessagingConsumer interface.
type MockMessagingConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockMessagingConsumerMockRecorder
}

// MockMessagingConsumerMockRecorder is the mock recorder for MockMessagingConsumer.
type MockMessagingConsumerMockRecorder struct {
	mock *MockMessagingConsumer
}

// NewMockMessagingConsumer creates a new mock instance.
func NewMockMessagingConsumer(ctrl *gomock.Controller) *MockMessagingConsumer {
	mock := &MockMessagingConsumer{ctrl: ctrl}
	mock.recorder = &MockMessagingConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessagingConsumer) EXPECT() *MockMessagingConsumerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMessagingConsumer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockMessagingConsumerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMessagingConsumer)(nil).Close))
}

// Consume mocks base method.
func (m *MockMessagingConsumer) Consume() (MessagingEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume")
	ret0, _ := ret[0].(MessagingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockMessagingConsumerMockRecorder) Consume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockMessagingConsumer)(nil).Consume))
}

// MessagingContext mocks base method.
func (m *MockMessagingConsumer) MessagingContext() MessagingContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessagingContext")
	ret0, _ := ret[0].(MessagingContext)
	return ret0
}

// MessagingContext indicates an expected call of MessagingContext.
func (mr *MockMessagingConsumerMockRecorder) MessagingContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessagingContext", reflect.TypeOf((*MockMessagingConsumer)(nil).MessagingContext))
}
