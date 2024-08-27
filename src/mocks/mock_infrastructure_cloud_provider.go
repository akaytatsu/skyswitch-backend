// Code generated by MockGen. DO NOT EDIT.
// Source: app/infrastructure/cloud_provider (interfaces: ICloudProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "app/entity"
	infrastructure_cloud_provider "app/infrastructure/cloud_provider"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICloudProvider is a mock of ICloudProvider interface.
type MockICloudProvider struct {
	ctrl     *gomock.Controller
	recorder *MockICloudProviderMockRecorder
}

// MockICloudProviderMockRecorder is the mock recorder for MockICloudProvider.
type MockICloudProviderMockRecorder struct {
	mock *MockICloudProvider
}

// NewMockICloudProvider creates a new mock instance.
func NewMockICloudProvider(ctrl *gomock.Controller) *MockICloudProvider {
	mock := &MockICloudProvider{ctrl: ctrl}
	mock.recorder = &MockICloudProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICloudProvider) EXPECT() *MockICloudProviderMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockICloudProvider) Connect(arg0 entity.EntityCloudAccount) (infrastructure_cloud_provider.ICloudProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0)
	ret0, _ := ret[0].(infrastructure_cloud_provider.ICloudProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockICloudProviderMockRecorder) Connect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockICloudProvider)(nil).Connect), arg0)
}

// GetAutoScalingGroupByID mocks base method.
func (m *MockICloudProvider) GetAutoScalingGroupByID(arg0 string) (*entity.EntityAutoScalingGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutoScalingGroupByID", arg0)
	ret0, _ := ret[0].(*entity.EntityAutoScalingGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutoScalingGroupByID indicates an expected call of GetAutoScalingGroupByID.
func (mr *MockICloudProviderMockRecorder) GetAutoScalingGroupByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoScalingGroupByID", reflect.TypeOf((*MockICloudProvider)(nil).GetAutoScalingGroupByID), arg0)
}

// GetAutoScalingGroups mocks base method.
func (m *MockICloudProvider) GetAutoScalingGroups() ([]*entity.EntityAutoScalingGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutoScalingGroups")
	ret0, _ := ret[0].([]*entity.EntityAutoScalingGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutoScalingGroups indicates an expected call of GetAutoScalingGroups.
func (mr *MockICloudProviderMockRecorder) GetAutoScalingGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoScalingGroups", reflect.TypeOf((*MockICloudProvider)(nil).GetAutoScalingGroups))
}

// GetDBInstanceByID mocks base method.
func (m *MockICloudProvider) GetDBInstanceByID(arg0 string) (*entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBInstanceByID", arg0)
	ret0, _ := ret[0].(*entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDBInstanceByID indicates an expected call of GetDBInstanceByID.
func (mr *MockICloudProviderMockRecorder) GetDBInstanceByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBInstanceByID", reflect.TypeOf((*MockICloudProvider)(nil).GetDBInstanceByID), arg0)
}

// GetDBInstances mocks base method.
func (m *MockICloudProvider) GetDBInstances() ([]*entity.EntityDbinstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBInstances")
	ret0, _ := ret[0].([]*entity.EntityDbinstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDBInstances indicates an expected call of GetDBInstances.
func (mr *MockICloudProviderMockRecorder) GetDBInstances() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBInstances", reflect.TypeOf((*MockICloudProvider)(nil).GetDBInstances))
}

// GetInstanceByID mocks base method.
func (m *MockICloudProvider) GetInstanceByID(arg0 string) (*entity.EntityInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceByID", arg0)
	ret0, _ := ret[0].(*entity.EntityInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceByID indicates an expected call of GetInstanceByID.
func (mr *MockICloudProviderMockRecorder) GetInstanceByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceByID", reflect.TypeOf((*MockICloudProvider)(nil).GetInstanceByID), arg0)
}

// GetInstances mocks base method.
func (m *MockICloudProvider) GetInstances() ([]*entity.EntityInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstances")
	ret0, _ := ret[0].([]*entity.EntityInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstances indicates an expected call of GetInstances.
func (mr *MockICloudProviderMockRecorder) GetInstances() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstances", reflect.TypeOf((*MockICloudProvider)(nil).GetInstances))
}

// StartAutoScalingGroup mocks base method.
func (m *MockICloudProvider) StartAutoScalingGroup(arg0 *entity.EntityAutoScalingGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartAutoScalingGroup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartAutoScalingGroup indicates an expected call of StartAutoScalingGroup.
func (mr *MockICloudProviderMockRecorder) StartAutoScalingGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAutoScalingGroup", reflect.TypeOf((*MockICloudProvider)(nil).StartAutoScalingGroup), arg0)
}

// StartDBInstance mocks base method.
func (m *MockICloudProvider) StartDBInstance(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDBInstance", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartDBInstance indicates an expected call of StartDBInstance.
func (mr *MockICloudProviderMockRecorder) StartDBInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDBInstance", reflect.TypeOf((*MockICloudProvider)(nil).StartDBInstance), arg0)
}

// StartInstance mocks base method.
func (m *MockICloudProvider) StartInstance(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartInstance", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartInstance indicates an expected call of StartInstance.
func (mr *MockICloudProviderMockRecorder) StartInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartInstance", reflect.TypeOf((*MockICloudProvider)(nil).StartInstance), arg0)
}

// StopAutoScalingGroup mocks base method.
func (m *MockICloudProvider) StopAutoScalingGroup(arg0 *entity.EntityAutoScalingGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopAutoScalingGroup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopAutoScalingGroup indicates an expected call of StopAutoScalingGroup.
func (mr *MockICloudProviderMockRecorder) StopAutoScalingGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAutoScalingGroup", reflect.TypeOf((*MockICloudProvider)(nil).StopAutoScalingGroup), arg0)
}

// StopDBInstance mocks base method.
func (m *MockICloudProvider) StopDBInstance(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopDBInstance", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopDBInstance indicates an expected call of StopDBInstance.
func (mr *MockICloudProviderMockRecorder) StopDBInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopDBInstance", reflect.TypeOf((*MockICloudProvider)(nil).StopDBInstance), arg0)
}

// StopInstance mocks base method.
func (m *MockICloudProvider) StopInstance(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopInstance", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopInstance indicates an expected call of StopInstance.
func (mr *MockICloudProviderMockRecorder) StopInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopInstance", reflect.TypeOf((*MockICloudProvider)(nil).StopInstance), arg0)
}
