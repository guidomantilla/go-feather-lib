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

	amqp091 "github.com/rabbitmq/amqp091-go"
	stream "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
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

// MockRabbitMQQueue is a mock of RabbitMQQueue interface.
type MockRabbitMQQueue struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQQueueMockRecorder
}

// MockRabbitMQQueueMockRecorder is the mock recorder for MockRabbitMQQueue.
type MockRabbitMQQueueMockRecorder struct {
	mock *MockRabbitMQQueue
}

// NewMockRabbitMQQueue creates a new mock instance.
func NewMockRabbitMQQueue(ctrl *gomock.Controller) *MockRabbitMQQueue {
	mock := &MockRabbitMQQueue{ctrl: ctrl}
	mock.recorder = &MockRabbitMQQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQQueue) EXPECT() *MockRabbitMQQueueMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRabbitMQQueue) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRabbitMQQueueMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRabbitMQQueue)(nil).Close))
}

// Connect mocks base method.
func (m *MockRabbitMQQueue) Connect() (*amqp091.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(*amqp091.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockRabbitMQQueueMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockRabbitMQQueue)(nil).Connect))
}

// Consumer mocks base method.
func (m *MockRabbitMQQueue) Consumer() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consumer")
	ret0, _ := ret[0].(string)
	return ret0
}

// Consumer indicates an expected call of Consumer.
func (mr *MockRabbitMQQueueMockRecorder) Consumer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consumer", reflect.TypeOf((*MockRabbitMQQueue)(nil).Consumer))
}

// MessagingContext mocks base method.
func (m *MockRabbitMQQueue) MessagingContext() MessagingContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessagingContext")
	ret0, _ := ret[0].(MessagingContext)
	return ret0
}

// MessagingContext indicates an expected call of MessagingContext.
func (mr *MockRabbitMQQueueMockRecorder) MessagingContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessagingContext", reflect.TypeOf((*MockRabbitMQQueue)(nil).MessagingContext))
}

// Name mocks base method.
func (m *MockRabbitMQQueue) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockRabbitMQQueueMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockRabbitMQQueue)(nil).Name))
}

// MockRabbitMQStreams is a mock of RabbitMQStreams interface.
type MockRabbitMQStreams struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQStreamsMockRecorder
}

// MockRabbitMQStreamsMockRecorder is the mock recorder for MockRabbitMQStreams.
type MockRabbitMQStreamsMockRecorder struct {
	mock *MockRabbitMQStreams
}

// NewMockRabbitMQStreams creates a new mock instance.
func NewMockRabbitMQStreams(ctrl *gomock.Controller) *MockRabbitMQStreams {
	mock := &MockRabbitMQStreams{ctrl: ctrl}
	mock.recorder = &MockRabbitMQStreamsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQStreams) EXPECT() *MockRabbitMQStreamsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRabbitMQStreams) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRabbitMQStreamsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRabbitMQStreams)(nil).Close))
}

// Connect mocks base method.
func (m *MockRabbitMQStreams) Connect() (*stream.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(*stream.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockRabbitMQStreamsMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockRabbitMQStreams)(nil).Connect))
}

// Consumer mocks base method.
func (m *MockRabbitMQStreams) Consumer() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consumer")
	ret0, _ := ret[0].(string)
	return ret0
}

// Consumer indicates an expected call of Consumer.
func (mr *MockRabbitMQStreamsMockRecorder) Consumer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consumer", reflect.TypeOf((*MockRabbitMQStreams)(nil).Consumer))
}

// MessagingContext mocks base method.
func (m *MockRabbitMQStreams) MessagingContext() MessagingContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessagingContext")
	ret0, _ := ret[0].(MessagingContext)
	return ret0
}

// MessagingContext indicates an expected call of MessagingContext.
func (mr *MockRabbitMQStreamsMockRecorder) MessagingContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessagingContext", reflect.TypeOf((*MockRabbitMQStreams)(nil).MessagingContext))
}

// Name mocks base method.
func (m *MockRabbitMQStreams) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockRabbitMQStreamsMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockRabbitMQStreams)(nil).Name))
}
