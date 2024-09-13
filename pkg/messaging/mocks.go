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

// MockRabbitMQQueueConnection is a mock of RabbitMQQueueConnection interface.
type MockRabbitMQQueueConnection struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQQueueConnectionMockRecorder
}

// MockRabbitMQQueueConnectionMockRecorder is the mock recorder for MockRabbitMQQueueConnection.
type MockRabbitMQQueueConnectionMockRecorder struct {
	mock *MockRabbitMQQueueConnection
}

// NewMockRabbitMQQueueConnection creates a new mock instance.
func NewMockRabbitMQQueueConnection(ctrl *gomock.Controller) *MockRabbitMQQueueConnection {
	mock := &MockRabbitMQQueueConnection{ctrl: ctrl}
	mock.recorder = &MockRabbitMQQueueConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQQueueConnection) EXPECT() *MockRabbitMQQueueConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRabbitMQQueueConnection) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRabbitMQQueueConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRabbitMQQueueConnection)(nil).Close))
}

// Connect mocks base method.
func (m *MockRabbitMQQueueConnection) Connect() (*amqp091.Connection, *amqp091.Channel, *amqp091.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(*amqp091.Connection)
	ret1, _ := ret[1].(*amqp091.Channel)
	ret2, _ := ret[2].(*amqp091.Queue)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Connect indicates an expected call of Connect.
func (mr *MockRabbitMQQueueConnectionMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockRabbitMQQueueConnection)(nil).Connect))
}

// Start mocks base method.
func (m *MockRabbitMQQueueConnection) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockRabbitMQQueueConnectionMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRabbitMQQueueConnection)(nil).Start))
}

// MockRabbitMQQueueMessageListener is a mock of RabbitMQQueueMessageListener interface.
type MockRabbitMQQueueMessageListener struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQQueueMessageListenerMockRecorder
}

// MockRabbitMQQueueMessageListenerMockRecorder is the mock recorder for MockRabbitMQQueueMessageListener.
type MockRabbitMQQueueMessageListenerMockRecorder struct {
	mock *MockRabbitMQQueueMessageListener
}

// NewMockRabbitMQQueueMessageListener creates a new mock instance.
func NewMockRabbitMQQueueMessageListener(ctrl *gomock.Controller) *MockRabbitMQQueueMessageListener {
	mock := &MockRabbitMQQueueMessageListener{ctrl: ctrl}
	mock.recorder = &MockRabbitMQQueueMessageListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQQueueMessageListener) EXPECT() *MockRabbitMQQueueMessageListenerMockRecorder {
	return m.recorder
}

// OnMessage mocks base method.
func (m *MockRabbitMQQueueMessageListener) OnMessage(message *amqp091.Delivery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnMessage", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnMessage indicates an expected call of OnMessage.
func (mr *MockRabbitMQQueueMessageListenerMockRecorder) OnMessage(message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnMessage", reflect.TypeOf((*MockRabbitMQQueueMessageListener)(nil).OnMessage), message)
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
