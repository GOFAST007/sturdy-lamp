// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: StorageAttachment,StorageInstance,MachinePortRanges,UnitPortRanges,CloudContainer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	network "github.com/juju/juju/core/network"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockStorageAttachment is a mock of StorageAttachment interface.
type MockStorageAttachment struct {
	ctrl     *gomock.Controller
	recorder *MockStorageAttachmentMockRecorder
}

// MockStorageAttachmentMockRecorder is the mock recorder for MockStorageAttachment.
type MockStorageAttachmentMockRecorder struct {
	mock *MockStorageAttachment
}

// NewMockStorageAttachment creates a new mock instance.
func NewMockStorageAttachment(ctrl *gomock.Controller) *MockStorageAttachment {
	mock := &MockStorageAttachment{ctrl: ctrl}
	mock.recorder = &MockStorageAttachmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageAttachment) EXPECT() *MockStorageAttachmentMockRecorder {
	return m.recorder
}

// Life mocks base method.
func (m *MockStorageAttachment) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockStorageAttachmentMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockStorageAttachment)(nil).Life))
}

// StorageInstance mocks base method.
func (m *MockStorageAttachment) StorageInstance() names.StorageTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageInstance")
	ret0, _ := ret[0].(names.StorageTag)
	return ret0
}

// StorageInstance indicates an expected call of StorageInstance.
func (mr *MockStorageAttachmentMockRecorder) StorageInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageInstance", reflect.TypeOf((*MockStorageAttachment)(nil).StorageInstance))
}

// Unit mocks base method.
func (m *MockStorageAttachment) Unit() names.UnitTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit")
	ret0, _ := ret[0].(names.UnitTag)
	return ret0
}

// Unit indicates an expected call of Unit.
func (mr *MockStorageAttachmentMockRecorder) Unit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockStorageAttachment)(nil).Unit))
}

// MockStorageInstance is a mock of StorageInstance interface.
type MockStorageInstance struct {
	ctrl     *gomock.Controller
	recorder *MockStorageInstanceMockRecorder
}

// MockStorageInstanceMockRecorder is the mock recorder for MockStorageInstance.
type MockStorageInstanceMockRecorder struct {
	mock *MockStorageInstance
}

// NewMockStorageInstance creates a new mock instance.
func NewMockStorageInstance(ctrl *gomock.Controller) *MockStorageInstance {
	mock := &MockStorageInstance{ctrl: ctrl}
	mock.recorder = &MockStorageInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageInstance) EXPECT() *MockStorageInstanceMockRecorder {
	return m.recorder
}

// Kind mocks base method.
func (m *MockStorageInstance) Kind() state.StorageKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(state.StorageKind)
	return ret0
}

// Kind indicates an expected call of Kind.
func (mr *MockStorageInstanceMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockStorageInstance)(nil).Kind))
}

// Life mocks base method.
func (m *MockStorageInstance) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockStorageInstanceMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockStorageInstance)(nil).Life))
}

// Owner mocks base method.
func (m *MockStorageInstance) Owner() (names.Tag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner")
	ret0, _ := ret[0].(names.Tag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Owner indicates an expected call of Owner.
func (mr *MockStorageInstanceMockRecorder) Owner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockStorageInstance)(nil).Owner))
}

// Pool mocks base method.
func (m *MockStorageInstance) Pool() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pool")
	ret0, _ := ret[0].(string)
	return ret0
}

// Pool indicates an expected call of Pool.
func (mr *MockStorageInstanceMockRecorder) Pool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pool", reflect.TypeOf((*MockStorageInstance)(nil).Pool))
}

// StorageName mocks base method.
func (m *MockStorageInstance) StorageName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageName")
	ret0, _ := ret[0].(string)
	return ret0
}

// StorageName indicates an expected call of StorageName.
func (mr *MockStorageInstanceMockRecorder) StorageName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageName", reflect.TypeOf((*MockStorageInstance)(nil).StorageName))
}

// StorageTag mocks base method.
func (m *MockStorageInstance) StorageTag() names.StorageTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageTag")
	ret0, _ := ret[0].(names.StorageTag)
	return ret0
}

// StorageTag indicates an expected call of StorageTag.
func (mr *MockStorageInstanceMockRecorder) StorageTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageTag", reflect.TypeOf((*MockStorageInstance)(nil).StorageTag))
}

// Tag mocks base method.
func (m *MockStorageInstance) Tag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockStorageInstanceMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockStorageInstance)(nil).Tag))
}

// MockMachinePortRanges is a mock of MachinePortRanges interface.
type MockMachinePortRanges struct {
	ctrl     *gomock.Controller
	recorder *MockMachinePortRangesMockRecorder
}

// MockMachinePortRangesMockRecorder is the mock recorder for MockMachinePortRanges.
type MockMachinePortRangesMockRecorder struct {
	mock *MockMachinePortRanges
}

// NewMockMachinePortRanges creates a new mock instance.
func NewMockMachinePortRanges(ctrl *gomock.Controller) *MockMachinePortRanges {
	mock := &MockMachinePortRanges{ctrl: ctrl}
	mock.recorder = &MockMachinePortRangesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachinePortRanges) EXPECT() *MockMachinePortRangesMockRecorder {
	return m.recorder
}

// ByUnit mocks base method.
func (m *MockMachinePortRanges) ByUnit() map[string]state.UnitPortRanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByUnit")
	ret0, _ := ret[0].(map[string]state.UnitPortRanges)
	return ret0
}

// ByUnit indicates an expected call of ByUnit.
func (mr *MockMachinePortRangesMockRecorder) ByUnit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByUnit", reflect.TypeOf((*MockMachinePortRanges)(nil).ByUnit))
}

// Changes mocks base method.
func (m *MockMachinePortRanges) Changes() state.ModelOperation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(state.ModelOperation)
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockMachinePortRangesMockRecorder) Changes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockMachinePortRanges)(nil).Changes))
}

// ForUnit mocks base method.
func (m *MockMachinePortRanges) ForUnit(arg0 string) state.UnitPortRanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForUnit", arg0)
	ret0, _ := ret[0].(state.UnitPortRanges)
	return ret0
}

// ForUnit indicates an expected call of ForUnit.
func (mr *MockMachinePortRangesMockRecorder) ForUnit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForUnit", reflect.TypeOf((*MockMachinePortRanges)(nil).ForUnit), arg0)
}

// MachineID mocks base method.
func (m *MockMachinePortRanges) MachineID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineID")
	ret0, _ := ret[0].(string)
	return ret0
}

// MachineID indicates an expected call of MachineID.
func (mr *MockMachinePortRangesMockRecorder) MachineID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineID", reflect.TypeOf((*MockMachinePortRanges)(nil).MachineID))
}

// UniquePortRanges mocks base method.
func (m *MockMachinePortRanges) UniquePortRanges() []network.PortRange {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UniquePortRanges")
	ret0, _ := ret[0].([]network.PortRange)
	return ret0
}

// UniquePortRanges indicates an expected call of UniquePortRanges.
func (mr *MockMachinePortRangesMockRecorder) UniquePortRanges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UniquePortRanges", reflect.TypeOf((*MockMachinePortRanges)(nil).UniquePortRanges))
}

// MockUnitPortRanges is a mock of UnitPortRanges interface.
type MockUnitPortRanges struct {
	ctrl     *gomock.Controller
	recorder *MockUnitPortRangesMockRecorder
}

// MockUnitPortRangesMockRecorder is the mock recorder for MockUnitPortRanges.
type MockUnitPortRangesMockRecorder struct {
	mock *MockUnitPortRanges
}

// NewMockUnitPortRanges creates a new mock instance.
func NewMockUnitPortRanges(ctrl *gomock.Controller) *MockUnitPortRanges {
	mock := &MockUnitPortRanges{ctrl: ctrl}
	mock.recorder = &MockUnitPortRangesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitPortRanges) EXPECT() *MockUnitPortRangesMockRecorder {
	return m.recorder
}

// ByEndpoint mocks base method.
func (m *MockUnitPortRanges) ByEndpoint() network.GroupedPortRanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByEndpoint")
	ret0, _ := ret[0].(network.GroupedPortRanges)
	return ret0
}

// ByEndpoint indicates an expected call of ByEndpoint.
func (mr *MockUnitPortRangesMockRecorder) ByEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByEndpoint", reflect.TypeOf((*MockUnitPortRanges)(nil).ByEndpoint))
}

// Changes mocks base method.
func (m *MockUnitPortRanges) Changes() state.ModelOperation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(state.ModelOperation)
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockUnitPortRangesMockRecorder) Changes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockUnitPortRanges)(nil).Changes))
}

// Close mocks base method.
func (m *MockUnitPortRanges) Close(arg0 string, arg1 network.PortRange) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close", arg0, arg1)
}

// Close indicates an expected call of Close.
func (mr *MockUnitPortRangesMockRecorder) Close(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUnitPortRanges)(nil).Close), arg0, arg1)
}

// ForEndpoint mocks base method.
func (m *MockUnitPortRanges) ForEndpoint(arg0 string) []network.PortRange {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEndpoint", arg0)
	ret0, _ := ret[0].([]network.PortRange)
	return ret0
}

// ForEndpoint indicates an expected call of ForEndpoint.
func (mr *MockUnitPortRangesMockRecorder) ForEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEndpoint", reflect.TypeOf((*MockUnitPortRanges)(nil).ForEndpoint), arg0)
}

// Open mocks base method.
func (m *MockUnitPortRanges) Open(arg0 string, arg1 network.PortRange) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Open", arg0, arg1)
}

// Open indicates an expected call of Open.
func (mr *MockUnitPortRangesMockRecorder) Open(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockUnitPortRanges)(nil).Open), arg0, arg1)
}

// UniquePortRanges mocks base method.
func (m *MockUnitPortRanges) UniquePortRanges() []network.PortRange {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UniquePortRanges")
	ret0, _ := ret[0].([]network.PortRange)
	return ret0
}

// UniquePortRanges indicates an expected call of UniquePortRanges.
func (mr *MockUnitPortRangesMockRecorder) UniquePortRanges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UniquePortRanges", reflect.TypeOf((*MockUnitPortRanges)(nil).UniquePortRanges))
}

// UnitName mocks base method.
func (m *MockUnitPortRanges) UnitName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitName")
	ret0, _ := ret[0].(string)
	return ret0
}

// UnitName indicates an expected call of UnitName.
func (mr *MockUnitPortRangesMockRecorder) UnitName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitName", reflect.TypeOf((*MockUnitPortRanges)(nil).UnitName))
}

// MockCloudContainer is a mock of CloudContainer interface.
type MockCloudContainer struct {
	ctrl     *gomock.Controller
	recorder *MockCloudContainerMockRecorder
}

// MockCloudContainerMockRecorder is the mock recorder for MockCloudContainer.
type MockCloudContainerMockRecorder struct {
	mock *MockCloudContainer
}

// NewMockCloudContainer creates a new mock instance.
func NewMockCloudContainer(ctrl *gomock.Controller) *MockCloudContainer {
	mock := &MockCloudContainer{ctrl: ctrl}
	mock.recorder = &MockCloudContainerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudContainer) EXPECT() *MockCloudContainerMockRecorder {
	return m.recorder
}

// Address mocks base method.
func (m *MockCloudContainer) Address() *network.SpaceAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(*network.SpaceAddress)
	return ret0
}

// Address indicates an expected call of Address.
func (mr *MockCloudContainerMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockCloudContainer)(nil).Address))
}

// Ports mocks base method.
func (m *MockCloudContainer) Ports() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ports")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Ports indicates an expected call of Ports.
func (mr *MockCloudContainerMockRecorder) Ports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ports", reflect.TypeOf((*MockCloudContainer)(nil).Ports))
}

// ProviderId mocks base method.
func (m *MockCloudContainer) ProviderId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProviderId")
	ret0, _ := ret[0].(string)
	return ret0
}

// ProviderId indicates an expected call of ProviderId.
func (mr *MockCloudContainerMockRecorder) ProviderId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProviderId", reflect.TypeOf((*MockCloudContainer)(nil).ProviderId))
}

// Unit mocks base method.
func (m *MockCloudContainer) Unit() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit")
	ret0, _ := ret[0].(string)
	return ret0
}

// Unit indicates an expected call of Unit.
func (mr *MockCloudContainerMockRecorder) Unit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockCloudContainer)(nil).Unit))
}
