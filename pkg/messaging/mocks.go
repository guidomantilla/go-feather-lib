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

	nats "github.com/nats-io/nats.go"
	amqp091 "github.com/rabbitmq/amqp091-go"
	gomock "go.uber.org/mock/gomock"
)

// MockRabbitMQContext is a mock of RabbitMQContext interface.
type MockRabbitMQContext struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQContextMockRecorder
}

// MockRabbitMQContextMockRecorder is the mock recorder for MockRabbitMQContext.
type MockRabbitMQContextMockRecorder struct {
	mock *MockRabbitMQContext
}

// NewMockRabbitMQContext creates a new mock instance.
func NewMockRabbitMQContext(ctrl *gomock.Controller) *MockRabbitMQContext {
	mock := &MockRabbitMQContext{ctrl: ctrl}
	mock.recorder = &MockRabbitMQContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQContext) EXPECT() *MockRabbitMQContextMockRecorder {
	return m.recorder
}

// Server mocks base method.
func (m *MockRabbitMQContext) Server() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Server")
	ret0, _ := ret[0].(string)
	return ret0
}

// Server indicates an expected call of Server.
func (mr *MockRabbitMQContextMockRecorder) Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Server", reflect.TypeOf((*MockRabbitMQContext)(nil).Server))
}

// Url mocks base method.
func (m *MockRabbitMQContext) Url() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Url")
	ret0, _ := ret[0].(string)
	return ret0
}

// Url indicates an expected call of Url.
func (mr *MockRabbitMQContextMockRecorder) Url() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Url", reflect.TypeOf((*MockRabbitMQContext)(nil).Url))
}

// VHost mocks base method.
func (m *MockRabbitMQContext) VHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// VHost indicates an expected call of VHost.
func (mr *MockRabbitMQContextMockRecorder) VHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VHost", reflect.TypeOf((*MockRabbitMQContext)(nil).VHost))
}

// MockRabbitMQConnection is a mock of RabbitMQConnection interface.
type MockRabbitMQConnection[T RabbitMQConnectionTypes] struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQConnectionMockRecorder[T]
}

// MockRabbitMQConnectionMockRecorder is the mock recorder for MockRabbitMQConnection.
type MockRabbitMQConnectionMockRecorder[T RabbitMQConnectionTypes] struct {
	mock *MockRabbitMQConnection[T]
}

// NewMockRabbitMQConnection creates a new mock instance.
func NewMockRabbitMQConnection[T RabbitMQConnectionTypes](ctrl *gomock.Controller) *MockRabbitMQConnection[T] {
	mock := &MockRabbitMQConnection[T]{ctrl: ctrl}
	mock.recorder = &MockRabbitMQConnectionMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQConnection[T]) EXPECT() *MockRabbitMQConnectionMockRecorder[T] {
	return m.recorder
}

// Close mocks base method.
func (m *MockRabbitMQConnection[T]) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRabbitMQConnectionMockRecorder[T]) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRabbitMQConnection[T])(nil).Close))
}

// Connect mocks base method.
func (m *MockRabbitMQConnection[T]) Connect() (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockRabbitMQConnectionMockRecorder[T]) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockRabbitMQConnection[T])(nil).Connect))
}

// RabbitMQContext mocks base method.
func (m *MockRabbitMQConnection[T]) RabbitMQContext() RabbitMQContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RabbitMQContext")
	ret0, _ := ret[0].(RabbitMQContext)
	return ret0
}

// RabbitMQContext indicates an expected call of RabbitMQContext.
func (mr *MockRabbitMQConnectionMockRecorder[T]) RabbitMQContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RabbitMQContext", reflect.TypeOf((*MockRabbitMQConnection[T])(nil).RabbitMQContext))
}

// MockRabbitMQMessageListener is a mock of RabbitMQMessageListener interface.
type MockRabbitMQMessageListener[T RabbitMQMessageListenerTypes] struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQMessageListenerMockRecorder[T]
}

// MockRabbitMQMessageListenerMockRecorder is the mock recorder for MockRabbitMQMessageListener.
type MockRabbitMQMessageListenerMockRecorder[T RabbitMQMessageListenerTypes] struct {
	mock *MockRabbitMQMessageListener[T]
}

// NewMockRabbitMQMessageListener creates a new mock instance.
func NewMockRabbitMQMessageListener[T RabbitMQMessageListenerTypes](ctrl *gomock.Controller) *MockRabbitMQMessageListener[T] {
	mock := &MockRabbitMQMessageListener[T]{ctrl: ctrl}
	mock.recorder = &MockRabbitMQMessageListenerMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQMessageListener[T]) EXPECT() *MockRabbitMQMessageListenerMockRecorder[T] {
	return m.recorder
}

// OnMessage mocks base method.
func (m *MockRabbitMQMessageListener[T]) OnMessage(message T) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnMessage", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnMessage indicates an expected call of OnMessage.
func (mr *MockRabbitMQMessageListenerMockRecorder[T]) OnMessage(message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnMessage", reflect.TypeOf((*MockRabbitMQMessageListener[T])(nil).OnMessage), message)
}

// MockRabbitMQChannel is a mock of RabbitMQChannel interface.
type MockRabbitMQChannel struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQChannelMockRecorder
}

// MockRabbitMQChannelMockRecorder is the mock recorder for MockRabbitMQChannel.
type MockRabbitMQChannelMockRecorder struct {
	mock *MockRabbitMQChannel
}

// NewMockRabbitMQChannel creates a new mock instance.
func NewMockRabbitMQChannel(ctrl *gomock.Controller) *MockRabbitMQChannel {
	mock := &MockRabbitMQChannel{ctrl: ctrl}
	mock.recorder = &MockRabbitMQChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQChannel) EXPECT() *MockRabbitMQChannelMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRabbitMQChannel) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRabbitMQChannelMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRabbitMQChannel)(nil).Close))
}

// Connect mocks base method.
func (m *MockRabbitMQChannel) Connect() (*amqp091.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(*amqp091.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockRabbitMQChannelMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockRabbitMQChannel)(nil).Connect))
}

// RabbitMQContext mocks base method.
func (m *MockRabbitMQChannel) RabbitMQContext() RabbitMQContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RabbitMQContext")
	ret0, _ := ret[0].(RabbitMQContext)
	return ret0
}

// RabbitMQContext indicates an expected call of RabbitMQContext.
func (mr *MockRabbitMQChannelMockRecorder) RabbitMQContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RabbitMQContext", reflect.TypeOf((*MockRabbitMQChannel)(nil).RabbitMQContext))
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

// RabbitMQContext mocks base method.
func (m *MockRabbitMQQueue) RabbitMQContext() RabbitMQContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RabbitMQContext")
	ret0, _ := ret[0].(RabbitMQContext)
	return ret0
}

// RabbitMQContext indicates an expected call of RabbitMQContext.
func (mr *MockRabbitMQQueueMockRecorder) RabbitMQContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RabbitMQContext", reflect.TypeOf((*MockRabbitMQQueue)(nil).RabbitMQContext))
}

// MockNatsSubjectConnection is a mock of NatsSubjectConnection interface.
type MockNatsSubjectConnection struct {
	ctrl     *gomock.Controller
	recorder *MockNatsSubjectConnectionMockRecorder
}

// MockNatsSubjectConnectionMockRecorder is the mock recorder for MockNatsSubjectConnection.
type MockNatsSubjectConnectionMockRecorder struct {
	mock *MockNatsSubjectConnection
}

// NewMockNatsSubjectConnection creates a new mock instance.
func NewMockNatsSubjectConnection(ctrl *gomock.Controller) *MockNatsSubjectConnection {
	mock := &MockNatsSubjectConnection{ctrl: ctrl}
	mock.recorder = &MockNatsSubjectConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNatsSubjectConnection) EXPECT() *MockNatsSubjectConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockNatsSubjectConnection) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockNatsSubjectConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNatsSubjectConnection)(nil).Close))
}

// Connect mocks base method.
func (m *MockNatsSubjectConnection) Connect() (*nats.Conn, *nats.Subscription, chan *nats.Msg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(*nats.Conn)
	ret1, _ := ret[1].(*nats.Subscription)
	ret2, _ := ret[2].(chan *nats.Msg)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Connect indicates an expected call of Connect.
func (mr *MockNatsSubjectConnectionMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockNatsSubjectConnection)(nil).Connect))
}

// Start mocks base method.
func (m *MockNatsSubjectConnection) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockNatsSubjectConnectionMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockNatsSubjectConnection)(nil).Start))
}

// MockNatsMessageListener is a mock of NatsMessageListener interface.
type MockNatsMessageListener struct {
	ctrl     *gomock.Controller
	recorder *MockNatsMessageListenerMockRecorder
}

// MockNatsMessageListenerMockRecorder is the mock recorder for MockNatsMessageListener.
type MockNatsMessageListenerMockRecorder struct {
	mock *MockNatsMessageListener
}

// NewMockNatsMessageListener creates a new mock instance.
func NewMockNatsMessageListener(ctrl *gomock.Controller) *MockNatsMessageListener {
	mock := &MockNatsMessageListener{ctrl: ctrl}
	mock.recorder = &MockNatsMessageListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNatsMessageListener) EXPECT() *MockNatsMessageListenerMockRecorder {
	return m.recorder
}

// OnMessage mocks base method.
func (m *MockNatsMessageListener) OnMessage(message *nats.Msg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnMessage", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnMessage indicates an expected call of OnMessage.
func (mr *MockNatsMessageListenerMockRecorder) OnMessage(message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnMessage", reflect.TypeOf((*MockNatsMessageListener)(nil).OnMessage), message)
}
