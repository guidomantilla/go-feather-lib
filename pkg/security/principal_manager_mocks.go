// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/security/principal_manager_types.go
//
// Generated by this command:
//
//	mockgen -package=security -destination ../pkg/security/principal_manager_mocks.go -source ../pkg/security/principal_manager_types.go
//

// Package security is a generated GoMock package.
package security

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPrincipalManager is a mock of PrincipalManager interface.
type MockPrincipalManager struct {
	ctrl     *gomock.Controller
	recorder *MockPrincipalManagerMockRecorder
	isgomock struct{}
}

// MockPrincipalManagerMockRecorder is the mock recorder for MockPrincipalManager.
type MockPrincipalManagerMockRecorder struct {
	mock *MockPrincipalManager
}

// NewMockPrincipalManager creates a new mock instance.
func NewMockPrincipalManager(ctrl *gomock.Controller) *MockPrincipalManager {
	mock := &MockPrincipalManager{ctrl: ctrl}
	mock.recorder = &MockPrincipalManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrincipalManager) EXPECT() *MockPrincipalManagerMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockPrincipalManager) ChangePassword(ctx context.Context, username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockPrincipalManagerMockRecorder) ChangePassword(ctx, username, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockPrincipalManager)(nil).ChangePassword), ctx, username, password)
}

// Create mocks base method.
func (m *MockPrincipalManager) Create(ctx context.Context, principal *Principal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, principal)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPrincipalManagerMockRecorder) Create(ctx, principal any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPrincipalManager)(nil).Create), ctx, principal)
}

// Delete mocks base method.
func (m *MockPrincipalManager) Delete(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPrincipalManagerMockRecorder) Delete(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPrincipalManager)(nil).Delete), ctx, username)
}

// Exists mocks base method.
func (m *MockPrincipalManager) Exists(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockPrincipalManagerMockRecorder) Exists(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockPrincipalManager)(nil).Exists), ctx, username)
}

// Find mocks base method.
func (m *MockPrincipalManager) Find(ctx context.Context, username string) (*Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, username)
	ret0, _ := ret[0].(*Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockPrincipalManagerMockRecorder) Find(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPrincipalManager)(nil).Find), ctx, username)
}

// Update mocks base method.
func (m *MockPrincipalManager) Update(ctx context.Context, principal *Principal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, principal)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPrincipalManagerMockRecorder) Update(ctx, principal any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPrincipalManager)(nil).Update), ctx, principal)
}

// VerifyResource mocks base method.
func (m *MockPrincipalManager) VerifyResource(ctx context.Context, username, resource string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyResource", ctx, username, resource)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyResource indicates an expected call of VerifyResource.
func (mr *MockPrincipalManagerMockRecorder) VerifyResource(ctx, username, resource any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyResource", reflect.TypeOf((*MockPrincipalManager)(nil).VerifyResource), ctx, username, resource)
}
