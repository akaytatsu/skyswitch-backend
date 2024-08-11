// Code generated by MockGen. DO NOT EDIT.
// Source: app/usecase/dbinstance (interfaces: IUsecaseDbinstance)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "app/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUsecaseDbinstance is a mock of IUsecaseDbinstance interface.
type MockIUsecaseDbinstance struct {
	ctrl     *gomock.Controller
	recorder *MockIUsecaseDbinstanceMockRecorder
}

// MockIUsecaseDbinstanceMockRecorder is the mock recorder for MockIUsecaseDbinstance.
type MockIUsecaseDbinstanceMockRecorder struct {
	mock *MockIUsecaseDbinstance
}

// NewMockIUsecaseDbinstance creates a new mock instance.
func NewMockIUsecaseDbinstance(ctrl *gomock.Controller) *MockIUsecaseDbinstance {
	mock := &MockIUsecaseDbinstance{ctrl: ctrl}
	mock.recorder = &MockIUsecaseDbinstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUsecaseDbinstance) EXPECT() *MockIUsecaseDbinstanceMockRecorder {
	return m.recorder
}

// ActiveDeactiveInstance mocks base method.
func (m *MockIUsecaseDbinstance) ActiveDeactiveInstance(arg0 int64, arg1 bool) (*entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveDeactiveInstance", arg0, arg1)
	ret0, _ := ret[0].(*entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveDeactiveInstance indicates an expected call of ActiveDeactiveInstance.
func (mr *MockIUsecaseDbinstanceMockRecorder) ActiveDeactiveInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveDeactiveInstance", reflect.TypeOf((*MockIUsecaseDbinstance)(nil).ActiveDeactiveInstance), arg0, arg1)
}

// Create mocks base method.
func (m *MockIUsecaseDbinstance) Create(arg0 *entity.EntityDbinstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIUsecaseDbinstanceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUsecaseDbinstance)(nil).Create), arg0)
}

// CreateOrUpdateDbInstance mocks base method.
func (m *MockIUsecaseDbinstance) CreateOrUpdateDbInstance(arg0 *entity.EntityDbinstance, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateDbInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateDbInstance indicates an expected call of CreateOrUpdateDbInstance.
func (mr *MockIUsecaseDbinstanceMockRecorder) CreateOrUpdateDbInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateDbInstance", reflect.TypeOf((*MockIUsecaseDbinstance)(nil).CreateOrUpdateDbInstance), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIUsecaseDbinstance) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUsecaseDbinstanceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUsecaseDbinstance)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockIUsecaseDbinstance) Get(arg0 int) (*entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIUsecaseDbinstanceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIUsecaseDbinstance)(nil).Get), arg0)
}

// GetAll mocks base method.
func (m *MockIUsecaseDbinstance) GetAll(arg0 entity.SearchEntityDbinstanceParams) ([]entity.EntityDbinstance, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]entity.EntityDbinstance)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIUsecaseDbinstanceMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIUsecaseDbinstance)(nil).GetAll), arg0)
}

// GetAllOFCalendar mocks base method.
func (m *MockIUsecaseDbinstance) GetAllOFCalendar(arg0 int) ([]entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOFCalendar", arg0)
	ret0, _ := ret[0].([]entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOFCalendar indicates an expected call of GetAllOFCalendar.
func (mr *MockIUsecaseDbinstanceMockRecorder) GetAllOFCalendar(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOFCalendar", reflect.TypeOf((*MockIUsecaseDbinstance)(nil).GetAllOFCalendar), arg0)
}

// GetByInstanceID mocks base method.
func (m *MockIUsecaseDbinstance) GetByInstanceID(arg0 string) (*entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByInstanceID", arg0)
	ret0, _ := ret[0].(*entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByInstanceID indicates an expected call of GetByInstanceID.
func (mr *MockIUsecaseDbinstanceMockRecorder) GetByInstanceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByInstanceID", reflect.TypeOf((*MockIUsecaseDbinstance)(nil).GetByInstanceID), arg0)
}

// Update mocks base method.
func (m *MockIUsecaseDbinstance) Update(arg0 *entity.EntityDbinstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIUsecaseDbinstanceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUsecaseDbinstance)(nil).Update), arg0)
}