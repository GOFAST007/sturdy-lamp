// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/caasapplicationprovisioner (interfaces: CAASUnitProvisionerFacade)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	watcher "github.com/juju/juju/core/watcher"
	params "github.com/juju/juju/rpc/params"
	gomock "go.uber.org/mock/gomock"
)

// MockCAASUnitProvisionerFacade is a mock of CAASUnitProvisionerFacade interface.
type MockCAASUnitProvisionerFacade struct {
	ctrl     *gomock.Controller
	recorder *MockCAASUnitProvisionerFacadeMockRecorder
}

// MockCAASUnitProvisionerFacadeMockRecorder is the mock recorder for MockCAASUnitProvisionerFacade.
type MockCAASUnitProvisionerFacadeMockRecorder struct {
	mock *MockCAASUnitProvisionerFacade
}

// NewMockCAASUnitProvisionerFacade creates a new mock instance.
func NewMockCAASUnitProvisionerFacade(ctrl *gomock.Controller) *MockCAASUnitProvisionerFacade {
	mock := &MockCAASUnitProvisionerFacade{ctrl: ctrl}
	mock.recorder = &MockCAASUnitProvisionerFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCAASUnitProvisionerFacade) EXPECT() *MockCAASUnitProvisionerFacadeMockRecorder {
	return m.recorder
}

// ApplicationScale mocks base method.
func (m *MockCAASUnitProvisionerFacade) ApplicationScale(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationScale", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationScale indicates an expected call of ApplicationScale.
func (mr *MockCAASUnitProvisionerFacadeMockRecorder) ApplicationScale(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationScale", reflect.TypeOf((*MockCAASUnitProvisionerFacade)(nil).ApplicationScale), arg0)
}

// ApplicationTrust mocks base method.
func (m *MockCAASUnitProvisionerFacade) ApplicationTrust(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationTrust", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationTrust indicates an expected call of ApplicationTrust.
func (mr *MockCAASUnitProvisionerFacadeMockRecorder) ApplicationTrust(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationTrust", reflect.TypeOf((*MockCAASUnitProvisionerFacade)(nil).ApplicationTrust), arg0)
}

// UpdateApplicationService mocks base method.
func (m *MockCAASUnitProvisionerFacade) UpdateApplicationService(arg0 params.UpdateApplicationServiceArg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplicationService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplicationService indicates an expected call of UpdateApplicationService.
func (mr *MockCAASUnitProvisionerFacadeMockRecorder) UpdateApplicationService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplicationService", reflect.TypeOf((*MockCAASUnitProvisionerFacade)(nil).UpdateApplicationService), arg0)
}

// WatchApplicationScale mocks base method.
func (m *MockCAASUnitProvisionerFacade) WatchApplicationScale(arg0 string) (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplicationScale", arg0)
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplicationScale indicates an expected call of WatchApplicationScale.
func (mr *MockCAASUnitProvisionerFacadeMockRecorder) WatchApplicationScale(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplicationScale", reflect.TypeOf((*MockCAASUnitProvisionerFacade)(nil).WatchApplicationScale), arg0)
}

// WatchApplicationTrustHash mocks base method.
func (m *MockCAASUnitProvisionerFacade) WatchApplicationTrustHash(arg0 string) (watcher.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplicationTrustHash", arg0)
	ret0, _ := ret[0].(watcher.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplicationTrustHash indicates an expected call of WatchApplicationTrustHash.
func (mr *MockCAASUnitProvisionerFacadeMockRecorder) WatchApplicationTrustHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplicationTrustHash", reflect.TypeOf((*MockCAASUnitProvisionerFacade)(nil).WatchApplicationTrustHash), arg0)
}
