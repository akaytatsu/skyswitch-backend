// Code generated by MockGen. DO NOT EDIT.
// Source: app/usecase/dbinstance (interfaces: IRepositoryDbinstance)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "app/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepositoryDbinstance is a mock of IRepositoryDbinstance interface.
type MockIRepositoryDbinstance struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryDbinstanceMockRecorder
}

// MockIRepositoryDbinstanceMockRecorder is the mock recorder for MockIRepositoryDbinstance.
type MockIRepositoryDbinstanceMockRecorder struct {
	mock *MockIRepositoryDbinstance
}

// NewMockIRepositoryDbinstance creates a new mock instance.
func NewMockIRepositoryDbinstance(ctrl *gomock.Controller) *MockIRepositoryDbinstance {
	mock := &MockIRepositoryDbinstance{ctrl: ctrl}
	mock.recorder = &MockIRepositoryDbinstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepositoryDbinstance) EXPECT() *MockIRepositoryDbinstanceMockRecorder {
	return m.recorder
}

// ActiveDeactiveInstance mocks base method.
func (m *MockIRepositoryDbinstance) ActiveDeactiveInstance(arg0 int64, arg1 bool) (*entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveDeactiveInstance", arg0, arg1)
	ret0, _ := ret[0].(*entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveDeactiveInstance indicates an expected call of ActiveDeactiveInstance.
func (mr *MockIRepositoryDbinstanceMockRecorder) ActiveDeactiveInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveDeactiveInstance", reflect.TypeOf((*MockIRepositoryDbinstance)(nil).ActiveDeactiveInstance), arg0, arg1)
}

// Create mocks base method.
func (m *MockIRepositoryDbinstance) Create(arg0 *entity.EntityDbinstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIRepositoryDbinstanceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRepositoryDbinstance)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockIRepositoryDbinstance) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIRepositoryDbinstanceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRepositoryDbinstance)(nil).Delete), arg0)
}

// FromCalendar mocks base method.
func (m *MockIRepositoryDbinstance) FromCalendar(arg0 int) ([]entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromCalendar", arg0)
	ret0, _ := ret[0].([]entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromCalendar indicates an expected call of FromCalendar.
func (mr *MockIRepositoryDbinstanceMockRecorder) FromCalendar(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromCalendar", reflect.TypeOf((*MockIRepositoryDbinstance)(nil).FromCalendar), arg0)
}

// GetAll mocks base method.
func (m *MockIRepositoryDbinstance) GetAll(arg0 entity.SearchEntityDbinstanceParams) ([]entity.EntityDbinstance, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]entity.EntityDbinstance)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIRepositoryDbinstanceMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIRepositoryDbinstance)(nil).GetAll), arg0)
}

// GetByInstanceID mocks base method.
func (m *MockIRepositoryDbinstance) GetByInstanceID(arg0 string) (*entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByInstanceID", arg0)
	ret0, _ := ret[0].(*entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByInstanceID indicates an expected call of GetByInstanceID.
func (mr *MockIRepositoryDbinstanceMockRecorder) GetByInstanceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByInstanceID", reflect.TypeOf((*MockIRepositoryDbinstance)(nil).GetByInstanceID), arg0)
}

// GetFromID mocks base method.
func (m *MockIRepositoryDbinstance) GetFromID(arg0 int) (*entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromID", arg0)
	ret0, _ := ret[0].(*entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromID indicates an expected call of GetFromID.
func (mr *MockIRepositoryDbinstanceMockRecorder) GetFromID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromID", reflect.TypeOf((*MockIRepositoryDbinstance)(nil).GetFromID), arg0)
}

// Update mocks base method.
func (m *MockIRepositoryDbinstance) Update(arg0 *entity.EntityDbinstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIRepositoryDbinstanceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRepositoryDbinstance)(nil).Update), arg0)
}
