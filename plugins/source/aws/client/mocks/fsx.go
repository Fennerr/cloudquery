// Code generated by MockGen. DO NOT EDIT.
// Source: fsx.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	fsx "github.com/aws/aws-sdk-go-v2/service/fsx"
	gomock "github.com/golang/mock/gomock"
)

// MockFsxClient is a mock of FsxClient interface.
type MockFsxClient struct {
	ctrl     *gomock.Controller
	recorder *MockFsxClientMockRecorder
}

// MockFsxClientMockRecorder is the mock recorder for MockFsxClient.
type MockFsxClientMockRecorder struct {
	mock *MockFsxClient
}

// NewMockFsxClient creates a new mock instance.
func NewMockFsxClient(ctrl *gomock.Controller) *MockFsxClient {
	mock := &MockFsxClient{ctrl: ctrl}
	mock.recorder = &MockFsxClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFsxClient) EXPECT() *MockFsxClientMockRecorder {
	return m.recorder
}

// DescribeBackups mocks base method.
func (m *MockFsxClient) DescribeBackups(arg0 context.Context, arg1 *fsx.DescribeBackupsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeBackupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeBackups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackups", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBackups indicates an expected call of DescribeBackups.
func (mr *MockFsxClientMockRecorder) DescribeBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackups", reflect.TypeOf((*MockFsxClient)(nil).DescribeBackups), varargs...)
}

// DescribeDataRepositoryAssociations mocks base method.
func (m *MockFsxClient) DescribeDataRepositoryAssociations(arg0 context.Context, arg1 *fsx.DescribeDataRepositoryAssociationsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryAssociationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDataRepositoryAssociations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDataRepositoryAssociations", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeDataRepositoryAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDataRepositoryAssociations indicates an expected call of DescribeDataRepositoryAssociations.
func (mr *MockFsxClientMockRecorder) DescribeDataRepositoryAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDataRepositoryAssociations", reflect.TypeOf((*MockFsxClient)(nil).DescribeDataRepositoryAssociations), varargs...)
}

// DescribeDataRepositoryTasks mocks base method.
func (m *MockFsxClient) DescribeDataRepositoryTasks(arg0 context.Context, arg1 *fsx.DescribeDataRepositoryTasksInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryTasksOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDataRepositoryTasks")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDataRepositoryTasks", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeDataRepositoryTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDataRepositoryTasks indicates an expected call of DescribeDataRepositoryTasks.
func (mr *MockFsxClientMockRecorder) DescribeDataRepositoryTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDataRepositoryTasks", reflect.TypeOf((*MockFsxClient)(nil).DescribeDataRepositoryTasks), varargs...)
}

// DescribeFileCaches mocks base method.
func (m *MockFsxClient) DescribeFileCaches(arg0 context.Context, arg1 *fsx.DescribeFileCachesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeFileCachesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeFileCaches")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFileCaches", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeFileCachesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileCaches indicates an expected call of DescribeFileCaches.
func (mr *MockFsxClientMockRecorder) DescribeFileCaches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileCaches", reflect.TypeOf((*MockFsxClient)(nil).DescribeFileCaches), varargs...)
}

// DescribeFileSystemAliases mocks base method.
func (m *MockFsxClient) DescribeFileSystemAliases(arg0 context.Context, arg1 *fsx.DescribeFileSystemAliasesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeFileSystemAliasesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeFileSystemAliases")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFileSystemAliases", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeFileSystemAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileSystemAliases indicates an expected call of DescribeFileSystemAliases.
func (mr *MockFsxClientMockRecorder) DescribeFileSystemAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystemAliases", reflect.TypeOf((*MockFsxClient)(nil).DescribeFileSystemAliases), varargs...)
}

// DescribeFileSystems mocks base method.
func (m *MockFsxClient) DescribeFileSystems(arg0 context.Context, arg1 *fsx.DescribeFileSystemsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeFileSystemsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeFileSystems")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFileSystems", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeFileSystemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileSystems indicates an expected call of DescribeFileSystems.
func (mr *MockFsxClientMockRecorder) DescribeFileSystems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystems", reflect.TypeOf((*MockFsxClient)(nil).DescribeFileSystems), varargs...)
}

// DescribeSnapshots mocks base method.
func (m *MockFsxClient) DescribeSnapshots(arg0 context.Context, arg1 *fsx.DescribeSnapshotsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeSnapshotsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeSnapshots")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSnapshots", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSnapshots indicates an expected call of DescribeSnapshots.
func (mr *MockFsxClientMockRecorder) DescribeSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshots", reflect.TypeOf((*MockFsxClient)(nil).DescribeSnapshots), varargs...)
}

// DescribeStorageVirtualMachines mocks base method.
func (m *MockFsxClient) DescribeStorageVirtualMachines(arg0 context.Context, arg1 *fsx.DescribeStorageVirtualMachinesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeStorageVirtualMachinesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeStorageVirtualMachines")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStorageVirtualMachines", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeStorageVirtualMachinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStorageVirtualMachines indicates an expected call of DescribeStorageVirtualMachines.
func (mr *MockFsxClientMockRecorder) DescribeStorageVirtualMachines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStorageVirtualMachines", reflect.TypeOf((*MockFsxClient)(nil).DescribeStorageVirtualMachines), varargs...)
}

// DescribeVolumes mocks base method.
func (m *MockFsxClient) DescribeVolumes(arg0 context.Context, arg1 *fsx.DescribeVolumesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeVolumesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeVolumes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVolumes", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeVolumesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVolumes indicates an expected call of DescribeVolumes.
func (mr *MockFsxClientMockRecorder) DescribeVolumes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVolumes", reflect.TypeOf((*MockFsxClient)(nil).DescribeVolumes), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockFsxClient) ListTagsForResource(arg0 context.Context, arg1 *fsx.ListTagsForResourceInput, arg2 ...func(*fsx.Options)) (*fsx.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &fsx.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*fsx.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockFsxClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockFsxClient)(nil).ListTagsForResource), varargs...)
}
