// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/networking/v1beta1 (interfaces: NetworkingV1beta1Interface,IngressInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1beta1 "k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v1beta10 "k8s.io/client-go/applyconfigurations/networking/v1beta1"
	v1beta11 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	rest "k8s.io/client-go/rest"
)

// MockNetworkingV1beta1Interface is a mock of NetworkingV1beta1Interface interface.
type MockNetworkingV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkingV1beta1InterfaceMockRecorder
}

// MockNetworkingV1beta1InterfaceMockRecorder is the mock recorder for MockNetworkingV1beta1Interface.
type MockNetworkingV1beta1InterfaceMockRecorder struct {
	mock *MockNetworkingV1beta1Interface
}

// NewMockNetworkingV1beta1Interface creates a new mock instance.
func NewMockNetworkingV1beta1Interface(ctrl *gomock.Controller) *MockNetworkingV1beta1Interface {
	mock := &MockNetworkingV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockNetworkingV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkingV1beta1Interface) EXPECT() *MockNetworkingV1beta1InterfaceMockRecorder {
	return m.recorder
}

// IngressClasses mocks base method.
func (m *MockNetworkingV1beta1Interface) IngressClasses() v1beta11.IngressClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngressClasses")
	ret0, _ := ret[0].(v1beta11.IngressClassInterface)
	return ret0
}

// IngressClasses indicates an expected call of IngressClasses.
func (mr *MockNetworkingV1beta1InterfaceMockRecorder) IngressClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngressClasses", reflect.TypeOf((*MockNetworkingV1beta1Interface)(nil).IngressClasses))
}

// Ingresses mocks base method.
func (m *MockNetworkingV1beta1Interface) Ingresses(arg0 string) v1beta11.IngressInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ingresses", arg0)
	ret0, _ := ret[0].(v1beta11.IngressInterface)
	return ret0
}

// Ingresses indicates an expected call of Ingresses.
func (mr *MockNetworkingV1beta1InterfaceMockRecorder) Ingresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ingresses", reflect.TypeOf((*MockNetworkingV1beta1Interface)(nil).Ingresses), arg0)
}

// RESTClient mocks base method.
func (m *MockNetworkingV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockNetworkingV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockNetworkingV1beta1Interface)(nil).RESTClient))
}

// MockIngressV1Beta1Interface is a mock of IngressInterface interface.
type MockIngressV1Beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockIngressV1Beta1InterfaceMockRecorder
}

// MockIngressV1Beta1InterfaceMockRecorder is the mock recorder for MockIngressV1Beta1Interface.
type MockIngressV1Beta1InterfaceMockRecorder struct {
	mock *MockIngressV1Beta1Interface
}

// NewMockIngressV1Beta1Interface creates a new mock instance.
func NewMockIngressV1Beta1Interface(ctrl *gomock.Controller) *MockIngressV1Beta1Interface {
	mock := &MockIngressV1Beta1Interface{ctrl: ctrl}
	mock.recorder = &MockIngressV1Beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngressV1Beta1Interface) EXPECT() *MockIngressV1Beta1InterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockIngressV1Beta1Interface) Apply(arg0 context.Context, arg1 *v1beta10.IngressApplyConfiguration, arg2 v1.ApplyOptions) (*v1beta1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).Apply), arg0, arg1, arg2)
}

// ApplyStatus mocks base method.
func (m *MockIngressV1Beta1Interface) ApplyStatus(arg0 context.Context, arg1 *v1beta10.IngressApplyConfiguration, arg2 v1.ApplyOptions) (*v1beta1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyStatus indicates an expected call of ApplyStatus.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatus", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).ApplyStatus), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockIngressV1Beta1Interface) Create(arg0 context.Context, arg1 *v1beta1.Ingress, arg2 v1.CreateOptions) (*v1beta1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockIngressV1Beta1Interface) Delete(arg0 context.Context, arg1 string, arg2 v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockIngressV1Beta1Interface) DeleteCollection(arg0 context.Context, arg1 v1.DeleteOptions, arg2 v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockIngressV1Beta1Interface) Get(arg0 context.Context, arg1 string, arg2 v1.GetOptions) (*v1beta1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockIngressV1Beta1Interface) List(arg0 context.Context, arg1 v1.ListOptions) (*v1beta1.IngressList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.IngressList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockIngressV1Beta1Interface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v1.PatchOptions, arg5 ...string) (*v1beta1.Ingress, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockIngressV1Beta1Interface) Update(arg0 context.Context, arg1 *v1beta1.Ingress, arg2 v1.UpdateOptions) (*v1beta1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockIngressV1Beta1Interface) UpdateStatus(arg0 context.Context, arg1 *v1beta1.Ingress, arg2 v1.UpdateOptions) (*v1beta1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockIngressV1Beta1Interface) Watch(arg0 context.Context, arg1 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockIngressV1Beta1InterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockIngressV1Beta1Interface)(nil).Watch), arg0, arg1)
}