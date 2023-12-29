// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/ssh (interfaces: Context,LeaderAPI,SSHClientAPI,SSHControllerAPI,CloudCredentialAPI,ApplicationAPI,CharmsAPI,ModelCommand)

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	os "os"
	reflect "reflect"

	api "github.com/juju/juju/api"
	application "github.com/juju/juju/api/client/application"
	client "github.com/juju/juju/api/client/client"
	charms "github.com/juju/juju/api/common/charms"
	cloud "github.com/juju/juju/cloud"
	controller "github.com/juju/juju/controller"
	cloudspec "github.com/juju/juju/environs/cloudspec"
	jujuclient "github.com/juju/juju/jujuclient"
	params "github.com/juju/juju/rpc/params"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// GetStderr mocks base method.
func (m *MockContext) GetStderr() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStderr")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// GetStderr indicates an expected call of GetStderr.
func (mr *MockContextMockRecorder) GetStderr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStderr", reflect.TypeOf((*MockContext)(nil).GetStderr))
}

// GetStdin mocks base method.
func (m *MockContext) GetStdin() io.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStdin")
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

// GetStdin indicates an expected call of GetStdin.
func (mr *MockContextMockRecorder) GetStdin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStdin", reflect.TypeOf((*MockContext)(nil).GetStdin))
}

// GetStdout mocks base method.
func (m *MockContext) GetStdout() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStdout")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// GetStdout indicates an expected call of GetStdout.
func (mr *MockContextMockRecorder) GetStdout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStdout", reflect.TypeOf((*MockContext)(nil).GetStdout))
}

// InterruptNotify mocks base method.
func (m *MockContext) InterruptNotify(arg0 chan<- os.Signal) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InterruptNotify", arg0)
}

// InterruptNotify indicates an expected call of InterruptNotify.
func (mr *MockContextMockRecorder) InterruptNotify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterruptNotify", reflect.TypeOf((*MockContext)(nil).InterruptNotify), arg0)
}

// StopInterruptNotify mocks base method.
func (m *MockContext) StopInterruptNotify(arg0 chan<- os.Signal) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopInterruptNotify", arg0)
}

// StopInterruptNotify indicates an expected call of StopInterruptNotify.
func (mr *MockContextMockRecorder) StopInterruptNotify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopInterruptNotify", reflect.TypeOf((*MockContext)(nil).StopInterruptNotify), arg0)
}

// MockLeaderAPI is a mock of LeaderAPI interface.
type MockLeaderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockLeaderAPIMockRecorder
}

// MockLeaderAPIMockRecorder is the mock recorder for MockLeaderAPI.
type MockLeaderAPIMockRecorder struct {
	mock *MockLeaderAPI
}

// NewMockLeaderAPI creates a new mock instance.
func NewMockLeaderAPI(ctrl *gomock.Controller) *MockLeaderAPI {
	mock := &MockLeaderAPI{ctrl: ctrl}
	mock.recorder = &MockLeaderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaderAPI) EXPECT() *MockLeaderAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockLeaderAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockLeaderAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLeaderAPI)(nil).Close))
}

// Leader mocks base method.
func (m *MockLeaderAPI) Leader(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leader", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leader indicates an expected call of Leader.
func (mr *MockLeaderAPIMockRecorder) Leader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leader", reflect.TypeOf((*MockLeaderAPI)(nil).Leader), arg0)
}

// MockSSHClientAPI is a mock of SSHClientAPI interface.
type MockSSHClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSSHClientAPIMockRecorder
}

// MockSSHClientAPIMockRecorder is the mock recorder for MockSSHClientAPI.
type MockSSHClientAPIMockRecorder struct {
	mock *MockSSHClientAPI
}

// NewMockSSHClientAPI creates a new mock instance.
func NewMockSSHClientAPI(ctrl *gomock.Controller) *MockSSHClientAPI {
	mock := &MockSSHClientAPI{ctrl: ctrl}
	mock.recorder = &MockSSHClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSHClientAPI) EXPECT() *MockSSHClientAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSSHClientAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSSHClientAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSSHClientAPI)(nil).Close))
}

// ModelCredentialForSSH mocks base method.
func (m *MockSSHClientAPI) ModelCredentialForSSH() (cloudspec.CloudSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelCredentialForSSH")
	ret0, _ := ret[0].(cloudspec.CloudSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelCredentialForSSH indicates an expected call of ModelCredentialForSSH.
func (mr *MockSSHClientAPIMockRecorder) ModelCredentialForSSH() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelCredentialForSSH", reflect.TypeOf((*MockSSHClientAPI)(nil).ModelCredentialForSSH))
}

// MockSSHControllerAPI is a mock of SSHControllerAPI interface.
type MockSSHControllerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSSHControllerAPIMockRecorder
}

// MockSSHControllerAPIMockRecorder is the mock recorder for MockSSHControllerAPI.
type MockSSHControllerAPIMockRecorder struct {
	mock *MockSSHControllerAPI
}

// NewMockSSHControllerAPI creates a new mock instance.
func NewMockSSHControllerAPI(ctrl *gomock.Controller) *MockSSHControllerAPI {
	mock := &MockSSHControllerAPI{ctrl: ctrl}
	mock.recorder = &MockSSHControllerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSHControllerAPI) EXPECT() *MockSSHControllerAPIMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockSSHControllerAPI) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockSSHControllerAPIMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockSSHControllerAPI)(nil).ControllerConfig))
}

// MockCloudCredentialAPI is a mock of CloudCredentialAPI interface.
type MockCloudCredentialAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCloudCredentialAPIMockRecorder
}

// MockCloudCredentialAPIMockRecorder is the mock recorder for MockCloudCredentialAPI.
type MockCloudCredentialAPIMockRecorder struct {
	mock *MockCloudCredentialAPI
}

// NewMockCloudCredentialAPI creates a new mock instance.
func NewMockCloudCredentialAPI(ctrl *gomock.Controller) *MockCloudCredentialAPI {
	mock := &MockCloudCredentialAPI{ctrl: ctrl}
	mock.recorder = &MockCloudCredentialAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudCredentialAPI) EXPECT() *MockCloudCredentialAPIMockRecorder {
	return m.recorder
}

// BestAPIVersion mocks base method.
func (m *MockCloudCredentialAPI) BestAPIVersion() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestAPIVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// BestAPIVersion indicates an expected call of BestAPIVersion.
func (mr *MockCloudCredentialAPIMockRecorder) BestAPIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestAPIVersion", reflect.TypeOf((*MockCloudCredentialAPI)(nil).BestAPIVersion))
}

// Close mocks base method.
func (m *MockCloudCredentialAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCloudCredentialAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloudCredentialAPI)(nil).Close))
}

// Cloud mocks base method.
func (m *MockCloudCredentialAPI) Cloud(arg0 names.CloudTag) (cloud.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud", arg0)
	ret0, _ := ret[0].(cloud.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cloud indicates an expected call of Cloud.
func (mr *MockCloudCredentialAPIMockRecorder) Cloud(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockCloudCredentialAPI)(nil).Cloud), arg0)
}

// CredentialContents mocks base method.
func (m *MockCloudCredentialAPI) CredentialContents(arg0, arg1 string, arg2 bool) ([]params.CredentialContentResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialContents", arg0, arg1, arg2)
	ret0, _ := ret[0].([]params.CredentialContentResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredentialContents indicates an expected call of CredentialContents.
func (mr *MockCloudCredentialAPIMockRecorder) CredentialContents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialContents", reflect.TypeOf((*MockCloudCredentialAPI)(nil).CredentialContents), arg0, arg1, arg2)
}

// MockApplicationAPI is a mock of ApplicationAPI interface.
type MockApplicationAPI struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationAPIMockRecorder
}

// MockApplicationAPIMockRecorder is the mock recorder for MockApplicationAPI.
type MockApplicationAPIMockRecorder struct {
	mock *MockApplicationAPI
}

// NewMockApplicationAPI creates a new mock instance.
func NewMockApplicationAPI(ctrl *gomock.Controller) *MockApplicationAPI {
	mock := &MockApplicationAPI{ctrl: ctrl}
	mock.recorder = &MockApplicationAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationAPI) EXPECT() *MockApplicationAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockApplicationAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockApplicationAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockApplicationAPI)(nil).Close))
}

// Leader mocks base method.
func (m *MockApplicationAPI) Leader(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leader", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leader indicates an expected call of Leader.
func (mr *MockApplicationAPIMockRecorder) Leader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leader", reflect.TypeOf((*MockApplicationAPI)(nil).Leader), arg0)
}

// UnitsInfo mocks base method.
func (m *MockApplicationAPI) UnitsInfo(arg0 []names.UnitTag) ([]application.UnitInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitsInfo", arg0)
	ret0, _ := ret[0].([]application.UnitInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitsInfo indicates an expected call of UnitsInfo.
func (mr *MockApplicationAPIMockRecorder) UnitsInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitsInfo", reflect.TypeOf((*MockApplicationAPI)(nil).UnitsInfo), arg0)
}

// MockCharmsAPI is a mock of CharmsAPI interface.
type MockCharmsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCharmsAPIMockRecorder
}

// MockCharmsAPIMockRecorder is the mock recorder for MockCharmsAPI.
type MockCharmsAPIMockRecorder struct {
	mock *MockCharmsAPI
}

// NewMockCharmsAPI creates a new mock instance.
func NewMockCharmsAPI(ctrl *gomock.Controller) *MockCharmsAPI {
	mock := &MockCharmsAPI{ctrl: ctrl}
	mock.recorder = &MockCharmsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmsAPI) EXPECT() *MockCharmsAPIMockRecorder {
	return m.recorder
}

// CharmInfo mocks base method.
func (m *MockCharmsAPI) CharmInfo(arg0 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmInfo", arg0)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmInfo indicates an expected call of CharmInfo.
func (mr *MockCharmsAPIMockRecorder) CharmInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmInfo", reflect.TypeOf((*MockCharmsAPI)(nil).CharmInfo), arg0)
}

// Close mocks base method.
func (m *MockCharmsAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCharmsAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCharmsAPI)(nil).Close))
}

// MockModelCommand is a mock of ModelCommand interface.
type MockModelCommand struct {
	ctrl     *gomock.Controller
	recorder *MockModelCommandMockRecorder
}

// MockModelCommandMockRecorder is the mock recorder for MockModelCommand.
type MockModelCommandMockRecorder struct {
	mock *MockModelCommand
}

// NewMockModelCommand creates a new mock instance.
func NewMockModelCommand(ctrl *gomock.Controller) *MockModelCommand {
	mock := &MockModelCommand{ctrl: ctrl}
	mock.recorder = &MockModelCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelCommand) EXPECT() *MockModelCommandMockRecorder {
	return m.recorder
}

// ModelDetails mocks base method.
func (m *MockModelCommand) ModelDetails() (string, *jujuclient.ModelDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelDetails")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*jujuclient.ModelDetails)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ModelDetails indicates an expected call of ModelDetails.
func (mr *MockModelCommandMockRecorder) ModelDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelDetails", reflect.TypeOf((*MockModelCommand)(nil).ModelDetails))
}

// ModelIdentifier mocks base method.
func (m *MockModelCommand) ModelIdentifier() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelIdentifier")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelIdentifier indicates an expected call of ModelIdentifier.
func (mr *MockModelCommandMockRecorder) ModelIdentifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelIdentifier", reflect.TypeOf((*MockModelCommand)(nil).ModelIdentifier))
}

// NewAPIClient mocks base method.
func (m *MockModelCommand) NewAPIClient() (*client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAPIClient")
	ret0, _ := ret[0].(*client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAPIClient indicates an expected call of NewAPIClient.
func (mr *MockModelCommandMockRecorder) NewAPIClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAPIClient", reflect.TypeOf((*MockModelCommand)(nil).NewAPIClient))
}

// NewAPIRoot mocks base method.
func (m *MockModelCommand) NewAPIRoot() (api.Connection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAPIRoot")
	ret0, _ := ret[0].(api.Connection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAPIRoot indicates an expected call of NewAPIRoot.
func (mr *MockModelCommandMockRecorder) NewAPIRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAPIRoot", reflect.TypeOf((*MockModelCommand)(nil).NewAPIRoot))
}

// NewControllerAPIRoot mocks base method.
func (m *MockModelCommand) NewControllerAPIRoot() (api.Connection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewControllerAPIRoot")
	ret0, _ := ret[0].(api.Connection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewControllerAPIRoot indicates an expected call of NewControllerAPIRoot.
func (mr *MockModelCommandMockRecorder) NewControllerAPIRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewControllerAPIRoot", reflect.TypeOf((*MockModelCommand)(nil).NewControllerAPIRoot))
}
