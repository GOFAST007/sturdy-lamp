// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/mongo (interfaces: SearchTools)

// Package mongo is a generated GoMock package.
package mongo

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSearchTools is a mock of SearchTools interface.
type MockSearchTools struct {
	ctrl     *gomock.Controller
	recorder *MockSearchToolsMockRecorder
}

// MockSearchToolsMockRecorder is the mock recorder for MockSearchTools.
type MockSearchToolsMockRecorder struct {
	mock *MockSearchTools
}

// NewMockSearchTools creates a new mock instance.
func NewMockSearchTools(ctrl *gomock.Controller) *MockSearchTools {
	mock := &MockSearchTools{ctrl: ctrl}
	mock.recorder = &MockSearchToolsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchTools) EXPECT() *MockSearchToolsMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockSearchTools) Exists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockSearchToolsMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSearchTools)(nil).Exists), arg0)
}

// GetCommandOutput mocks base method.
func (m *MockSearchTools) GetCommandOutput(arg0 string, arg1 ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCommandOutput", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommandOutput indicates an expected call of GetCommandOutput.
func (mr *MockSearchToolsMockRecorder) GetCommandOutput(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommandOutput", reflect.TypeOf((*MockSearchTools)(nil).GetCommandOutput), varargs...)
}
