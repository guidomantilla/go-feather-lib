// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/datasource/types.go
//
// Generated by this command:
//
//	mockgen -package datasource -destination ../pkg/datasource/mocks.go -source ../pkg/datasource/types.go
//

// Package datasource is a generated GoMock package.
package datasource

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDatasourceContext is a mock of DatasourceContext interface.
type MockDatasourceContext struct {
	ctrl     *gomock.Controller
	recorder *MockDatasourceContextMockRecorder
}

// MockDatasourceContextMockRecorder is the mock recorder for MockDatasourceContext.
type MockDatasourceContextMockRecorder struct {
	mock *MockDatasourceContext
}

// NewMockDatasourceContext creates a new mock instance.
func NewMockDatasourceContext(ctrl *gomock.Controller) *MockDatasourceContext {
	mock := &MockDatasourceContext{ctrl: ctrl}
	mock.recorder = &MockDatasourceContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatasourceContext) EXPECT() *MockDatasourceContextMockRecorder {
	return m.recorder
}

// GetServer mocks base method.
func (m *MockDatasourceContext) GetServer() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServer")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServer indicates an expected call of GetServer.
func (mr *MockDatasourceContextMockRecorder) GetServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServer", reflect.TypeOf((*MockDatasourceContext)(nil).GetServer))
}

// GetService mocks base method.
func (m *MockDatasourceContext) GetService() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetService indicates an expected call of GetService.
func (mr *MockDatasourceContextMockRecorder) GetService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockDatasourceContext)(nil).GetService))
}

// GetUrl mocks base method.
func (m *MockDatasourceContext) GetUrl() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUrl indicates an expected call of GetUrl.
func (mr *MockDatasourceContextMockRecorder) GetUrl() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockDatasourceContext)(nil).GetUrl))
}

// MockDatasource is a mock of Datasource interface.
type MockDatasource struct {
	ctrl     *gomock.Controller
	recorder *MockDatasourceMockRecorder
}

// MockDatasourceMockRecorder is the mock recorder for MockDatasource.
type MockDatasourceMockRecorder struct {
	mock *MockDatasource
}

// NewMockDatasource creates a new mock instance.
func NewMockDatasource(ctrl *gomock.Controller) *MockDatasource {
	mock := &MockDatasource{ctrl: ctrl}
	mock.recorder = &MockDatasourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatasource) EXPECT() *MockDatasourceMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockDatasource) Connect() (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockDatasourceMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockDatasource)(nil).Connect))
}

// MockTransactionHandler is a mock of TransactionHandler interface.
type MockTransactionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionHandlerMockRecorder
}

// MockTransactionHandlerMockRecorder is the mock recorder for MockTransactionHandler.
type MockTransactionHandlerMockRecorder struct {
	mock *MockTransactionHandler
}

// NewMockTransactionHandler creates a new mock instance.
func NewMockTransactionHandler(ctrl *gomock.Controller) *MockTransactionHandler {
	mock := &MockTransactionHandler{ctrl: ctrl}
	mock.recorder = &MockTransactionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionHandler) EXPECT() *MockTransactionHandlerMockRecorder {
	return m.recorder
}

// HandleTransaction mocks base method.
func (m *MockTransactionHandler) HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTransaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleTransaction indicates an expected call of HandleTransaction.
func (mr *MockTransactionHandlerMockRecorder) HandleTransaction(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTransaction", reflect.TypeOf((*MockTransactionHandler)(nil).HandleTransaction), ctx, fn)
}
