// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/upgradeseries (interfaces: Facade,UnitDiscovery,Upgrader)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/juju/juju/core/model"
	watcher "github.com/juju/juju/core/watcher"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockFacade is a mock of Facade interface.
type MockFacade struct {
	ctrl     *gomock.Controller
	recorder *MockFacadeMockRecorder
}

// MockFacadeMockRecorder is the mock recorder for MockFacade.
type MockFacadeMockRecorder struct {
	mock *MockFacade
}

// NewMockFacade creates a new mock instance.
func NewMockFacade(ctrl *gomock.Controller) *MockFacade {
	mock := &MockFacade{ctrl: ctrl}
	mock.recorder = &MockFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFacade) EXPECT() *MockFacadeMockRecorder {
	return m.recorder
}

// CurrentSeries mocks base method.
func (m *MockFacade) CurrentSeries() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentSeries")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentSeries indicates an expected call of CurrentSeries.
func (mr *MockFacadeMockRecorder) CurrentSeries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentSeries", reflect.TypeOf((*MockFacade)(nil).CurrentSeries))
}

// FinishUpgradeSeries mocks base method.
func (m *MockFacade) FinishUpgradeSeries(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishUpgradeSeries", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishUpgradeSeries indicates an expected call of FinishUpgradeSeries.
func (mr *MockFacadeMockRecorder) FinishUpgradeSeries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishUpgradeSeries", reflect.TypeOf((*MockFacade)(nil).FinishUpgradeSeries), arg0)
}

// MachineStatus mocks base method.
func (m *MockFacade) MachineStatus() (model.UpgradeSeriesStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineStatus")
	ret0, _ := ret[0].(model.UpgradeSeriesStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MachineStatus indicates an expected call of MachineStatus.
func (mr *MockFacadeMockRecorder) MachineStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineStatus", reflect.TypeOf((*MockFacade)(nil).MachineStatus))
}

// PinMachineApplications mocks base method.
func (m *MockFacade) PinMachineApplications() (map[string]error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PinMachineApplications")
	ret0, _ := ret[0].(map[string]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PinMachineApplications indicates an expected call of PinMachineApplications.
func (mr *MockFacadeMockRecorder) PinMachineApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PinMachineApplications", reflect.TypeOf((*MockFacade)(nil).PinMachineApplications))
}

// SetInstanceStatus mocks base method.
func (m *MockFacade) SetInstanceStatus(arg0 model.UpgradeSeriesStatus, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInstanceStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInstanceStatus indicates an expected call of SetInstanceStatus.
func (mr *MockFacadeMockRecorder) SetInstanceStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInstanceStatus", reflect.TypeOf((*MockFacade)(nil).SetInstanceStatus), arg0, arg1)
}

// SetMachineStatus mocks base method.
func (m *MockFacade) SetMachineStatus(arg0 model.UpgradeSeriesStatus, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMachineStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMachineStatus indicates an expected call of SetMachineStatus.
func (mr *MockFacadeMockRecorder) SetMachineStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMachineStatus", reflect.TypeOf((*MockFacade)(nil).SetMachineStatus), arg0, arg1)
}

// StartUnitCompletion mocks base method.
func (m *MockFacade) StartUnitCompletion(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartUnitCompletion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartUnitCompletion indicates an expected call of StartUnitCompletion.
func (mr *MockFacadeMockRecorder) StartUnitCompletion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartUnitCompletion", reflect.TypeOf((*MockFacade)(nil).StartUnitCompletion), arg0)
}

// TargetSeries mocks base method.
func (m *MockFacade) TargetSeries() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TargetSeries")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TargetSeries indicates an expected call of TargetSeries.
func (mr *MockFacadeMockRecorder) TargetSeries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetSeries", reflect.TypeOf((*MockFacade)(nil).TargetSeries))
}

// UnitsCompleted mocks base method.
func (m *MockFacade) UnitsCompleted() ([]names.UnitTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitsCompleted")
	ret0, _ := ret[0].([]names.UnitTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitsCompleted indicates an expected call of UnitsCompleted.
func (mr *MockFacadeMockRecorder) UnitsCompleted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitsCompleted", reflect.TypeOf((*MockFacade)(nil).UnitsCompleted))
}

// UnitsPrepared mocks base method.
func (m *MockFacade) UnitsPrepared() ([]names.UnitTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitsPrepared")
	ret0, _ := ret[0].([]names.UnitTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitsPrepared indicates an expected call of UnitsPrepared.
func (mr *MockFacadeMockRecorder) UnitsPrepared() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitsPrepared", reflect.TypeOf((*MockFacade)(nil).UnitsPrepared))
}

// UnpinMachineApplications mocks base method.
func (m *MockFacade) UnpinMachineApplications() (map[string]error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpinMachineApplications")
	ret0, _ := ret[0].(map[string]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpinMachineApplications indicates an expected call of UnpinMachineApplications.
func (mr *MockFacadeMockRecorder) UnpinMachineApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpinMachineApplications", reflect.TypeOf((*MockFacade)(nil).UnpinMachineApplications))
}

// WatchUpgradeSeriesNotifications mocks base method.
func (m *MockFacade) WatchUpgradeSeriesNotifications() (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUpgradeSeriesNotifications")
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUpgradeSeriesNotifications indicates an expected call of WatchUpgradeSeriesNotifications.
func (mr *MockFacadeMockRecorder) WatchUpgradeSeriesNotifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpgradeSeriesNotifications", reflect.TypeOf((*MockFacade)(nil).WatchUpgradeSeriesNotifications))
}

// MockUnitDiscovery is a mock of UnitDiscovery interface.
type MockUnitDiscovery struct {
	ctrl     *gomock.Controller
	recorder *MockUnitDiscoveryMockRecorder
}

// MockUnitDiscoveryMockRecorder is the mock recorder for MockUnitDiscovery.
type MockUnitDiscoveryMockRecorder struct {
	mock *MockUnitDiscovery
}

// NewMockUnitDiscovery creates a new mock instance.
func NewMockUnitDiscovery(ctrl *gomock.Controller) *MockUnitDiscovery {
	mock := &MockUnitDiscovery{ctrl: ctrl}
	mock.recorder = &MockUnitDiscoveryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitDiscovery) EXPECT() *MockUnitDiscoveryMockRecorder {
	return m.recorder
}

// Units mocks base method.
func (m *MockUnitDiscovery) Units() ([]names.UnitTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]names.UnitTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units.
func (mr *MockUnitDiscoveryMockRecorder) Units() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockUnitDiscovery)(nil).Units))
}

// MockUpgrader is a mock of Upgrader interface.
type MockUpgrader struct {
	ctrl     *gomock.Controller
	recorder *MockUpgraderMockRecorder
}

// MockUpgraderMockRecorder is the mock recorder for MockUpgrader.
type MockUpgraderMockRecorder struct {
	mock *MockUpgrader
}

// NewMockUpgrader creates a new mock instance.
func NewMockUpgrader(ctrl *gomock.Controller) *MockUpgrader {
	mock := &MockUpgrader{ctrl: ctrl}
	mock.recorder = &MockUpgraderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgrader) EXPECT() *MockUpgraderMockRecorder {
	return m.recorder
}

// PerformUpgrade mocks base method.
func (m *MockUpgrader) PerformUpgrade() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformUpgrade")
	ret0, _ := ret[0].(error)
	return ret0
}

// PerformUpgrade indicates an expected call of PerformUpgrade.
func (mr *MockUpgraderMockRecorder) PerformUpgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformUpgrade", reflect.TypeOf((*MockUpgrader)(nil).PerformUpgrade))
}
