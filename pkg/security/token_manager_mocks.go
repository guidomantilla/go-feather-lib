// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/security/token_manager_types.go
//
// Generated by this command:
//
//	mockgen -package=security -destination ../pkg/security/token_manager_mocks.go -source ../pkg/security/token_manager_types.go
//

// Package security is a generated GoMock package.
package security

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTokenManager is a mock of TokenManager interface.
type MockTokenManager struct {
	ctrl     *gomock.Controller
	recorder *MockTokenManagerMockRecorder
	isgomock struct{}
}

// MockTokenManagerMockRecorder is the mock recorder for MockTokenManager.
type MockTokenManagerMockRecorder struct {
	mock *MockTokenManager
}

// NewMockTokenManager creates a new mock instance.
func NewMockTokenManager(ctrl *gomock.Controller) *MockTokenManager {
	mock := &MockTokenManager{ctrl: ctrl}
	mock.recorder = &MockTokenManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenManager) EXPECT() *MockTokenManagerMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockTokenManager) Generate(principal *Principal) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", principal)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockTokenManagerMockRecorder) Generate(principal any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockTokenManager)(nil).Generate), principal)
}

// Validate mocks base method.
func (m *MockTokenManager) Validate(tokenString string) (*Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", tokenString)
	ret0, _ := ret[0].(*Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockTokenManagerMockRecorder) Validate(tokenString any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokenManager)(nil).Validate), tokenString)
}

// set mocks base method.
func (m *MockTokenManager) set(property string, value any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "set", property, value)
}

// set indicates an expected call of set.
func (mr *MockTokenManagerMockRecorder) set(property, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "set", reflect.TypeOf((*MockTokenManager)(nil).set), property, value)
}
