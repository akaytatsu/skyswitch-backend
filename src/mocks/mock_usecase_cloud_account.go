// Code generated by MockGen. DO NOT EDIT.
// Source: app/usecase/cloud_account (interfaces: IUsecaseCloudAccount)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "app/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUsecaseCloudAccount is a mock of IUsecaseCloudAccount interface.
type MockIUsecaseCloudAccount struct {
	ctrl     *gomock.Controller
	recorder *MockIUsecaseCloudAccountMockRecorder
}

// MockIUsecaseCloudAccountMockRecorder is the mock recorder for MockIUsecaseCloudAccount.
type MockIUsecaseCloudAccountMockRecorder struct {
	mock *MockIUsecaseCloudAccount
}

// NewMockIUsecaseCloudAccount creates a new mock instance.
func NewMockIUsecaseCloudAccount(ctrl *gomock.Controller) *MockIUsecaseCloudAccount {
	mock := &MockIUsecaseCloudAccount{ctrl: ctrl}
	mock.recorder = &MockIUsecaseCloudAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUsecaseCloudAccount) EXPECT() *MockIUsecaseCloudAccountMockRecorder {
	return m.recorder
}

// ActiveDeactiveCloudAccount mocks base method.
func (m *MockIUsecaseCloudAccount) ActiveDeactiveCloudAccount(arg0 int64, arg1 bool) (*entity.EntityCloudAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveDeactiveCloudAccount", arg0, arg1)
	ret0, _ := ret[0].(*entity.EntityCloudAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveDeactiveCloudAccount indicates an expected call of ActiveDeactiveCloudAccount.
func (mr *MockIUsecaseCloudAccountMockRecorder) ActiveDeactiveCloudAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveDeactiveCloudAccount", reflect.TypeOf((*MockIUsecaseCloudAccount)(nil).ActiveDeactiveCloudAccount), arg0, arg1)
}

// CreateCloudAccount mocks base method.
func (m *MockIUsecaseCloudAccount) CreateCloudAccount(arg0 *entity.EntityCloudAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCloudAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCloudAccount indicates an expected call of CreateCloudAccount.
func (mr *MockIUsecaseCloudAccountMockRecorder) CreateCloudAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCloudAccount", reflect.TypeOf((*MockIUsecaseCloudAccount)(nil).CreateCloudAccount), arg0)
}

// DeleteCloudAccount mocks base method.
func (m *MockIUsecaseCloudAccount) DeleteCloudAccount(arg0 *entity.EntityCloudAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCloudAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCloudAccount indicates an expected call of DeleteCloudAccount.
func (mr *MockIUsecaseCloudAccountMockRecorder) DeleteCloudAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCloudAccount", reflect.TypeOf((*MockIUsecaseCloudAccount)(nil).DeleteCloudAccount), arg0)
}

// GetAll mocks base method.
func (m *MockIUsecaseCloudAccount) GetAll(arg0 entity.SearchEntityCloudAccountParams) ([]entity.EntityCloudAccount, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]entity.EntityCloudAccount)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIUsecaseCloudAccountMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIUsecaseCloudAccount)(nil).GetAll), arg0)
}

// GetByID mocks base method.
func (m *MockIUsecaseCloudAccount) GetByID(arg0 int64) (*entity.EntityCloudAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*entity.EntityCloudAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIUsecaseCloudAccountMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIUsecaseCloudAccount)(nil).GetByID), arg0)
}

// UpdateAllInstancesOnAllCloudAccountProvider mocks base method.
func (m *MockIUsecaseCloudAccount) UpdateAllInstancesOnAllCloudAccountProvider() ([]*entity.EntityInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllInstancesOnAllCloudAccountProvider")
	ret0, _ := ret[0].([]*entity.EntityInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllInstancesOnAllCloudAccountProvider indicates an expected call of UpdateAllInstancesOnAllCloudAccountProvider.
func (mr *MockIUsecaseCloudAccountMockRecorder) UpdateAllInstancesOnAllCloudAccountProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllInstancesOnAllCloudAccountProvider", reflect.TypeOf((*MockIUsecaseCloudAccount)(nil).UpdateAllInstancesOnAllCloudAccountProvider))
}

// UpdateAllInstancesOnCloudAccountProvider mocks base method.
func (m *MockIUsecaseCloudAccount) UpdateAllInstancesOnCloudAccountProvider(arg0 *entity.EntityCloudAccount) ([]*entity.EntityInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllInstancesOnCloudAccountProvider", arg0)
	ret0, _ := ret[0].([]*entity.EntityInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllInstancesOnCloudAccountProvider indicates an expected call of UpdateAllInstancesOnCloudAccountProvider.
func (mr *MockIUsecaseCloudAccountMockRecorder) UpdateAllInstancesOnCloudAccountProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllInstancesOnCloudAccountProvider", reflect.TypeOf((*MockIUsecaseCloudAccount)(nil).UpdateAllInstancesOnCloudAccountProvider), arg0)
}

// UpdateAllInstancesOnCloudAccountProviderFromID mocks base method.
func (m *MockIUsecaseCloudAccount) UpdateAllInstancesOnCloudAccountProviderFromID(arg0 int) ([]*entity.EntityInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllInstancesOnCloudAccountProviderFromID", arg0)
	ret0, _ := ret[0].([]*entity.EntityInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllInstancesOnCloudAccountProviderFromID indicates an expected call of UpdateAllInstancesOnCloudAccountProviderFromID.
func (mr *MockIUsecaseCloudAccountMockRecorder) UpdateAllInstancesOnCloudAccountProviderFromID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllInstancesOnCloudAccountProviderFromID", reflect.TypeOf((*MockIUsecaseCloudAccount)(nil).UpdateAllInstancesOnCloudAccountProviderFromID), arg0)
}

// UpdateCloudAccount mocks base method.
func (m *MockIUsecaseCloudAccount) UpdateCloudAccount(arg0 *entity.EntityCloudAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCloudAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCloudAccount indicates an expected call of UpdateCloudAccount.
func (mr *MockIUsecaseCloudAccountMockRecorder) UpdateCloudAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCloudAccount", reflect.TypeOf((*MockIUsecaseCloudAccount)(nil).UpdateCloudAccount), arg0)
}
