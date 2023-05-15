// Code generated by MockGen. DO NOT EDIT.
// Source: efs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	efs "github.com/aws/aws-sdk-go-v2/service/efs"
	gomock "github.com/golang/mock/gomock"
)

// MockEfsClient is a mock of EfsClient interface.
type MockEfsClient struct {
	ctrl     *gomock.Controller
	recorder *MockEfsClientMockRecorder
}

// MockEfsClientMockRecorder is the mock recorder for MockEfsClient.
type MockEfsClientMockRecorder struct {
	mock *MockEfsClient
}

// NewMockEfsClient creates a new mock instance.
func NewMockEfsClient(ctrl *gomock.Controller) *MockEfsClient {
	mock := &MockEfsClient{ctrl: ctrl}
	mock.recorder = &MockEfsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEfsClient) EXPECT() *MockEfsClientMockRecorder {
	return m.recorder
}

// DescribeAccessPoints mocks base method.
func (m *MockEfsClient) DescribeAccessPoints(arg0 context.Context, arg1 *efs.DescribeAccessPointsInput, arg2 ...func(*efs.Options)) (*efs.DescribeAccessPointsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAccessPoints")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccessPoints", varargs...)
	ret0, _ := ret[0].(*efs.DescribeAccessPointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccessPoints indicates an expected call of DescribeAccessPoints.
func (mr *MockEfsClientMockRecorder) DescribeAccessPoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccessPoints", reflect.TypeOf((*MockEfsClient)(nil).DescribeAccessPoints), varargs...)
}

// DescribeAccountPreferences mocks base method.
func (m *MockEfsClient) DescribeAccountPreferences(arg0 context.Context, arg1 *efs.DescribeAccountPreferencesInput, arg2 ...func(*efs.Options)) (*efs.DescribeAccountPreferencesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAccountPreferences")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccountPreferences", varargs...)
	ret0, _ := ret[0].(*efs.DescribeAccountPreferencesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccountPreferences indicates an expected call of DescribeAccountPreferences.
func (mr *MockEfsClientMockRecorder) DescribeAccountPreferences(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccountPreferences", reflect.TypeOf((*MockEfsClient)(nil).DescribeAccountPreferences), varargs...)
}

// DescribeBackupPolicy mocks base method.
func (m *MockEfsClient) DescribeBackupPolicy(arg0 context.Context, arg1 *efs.DescribeBackupPolicyInput, arg2 ...func(*efs.Options)) (*efs.DescribeBackupPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeBackupPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackupPolicy", varargs...)
	ret0, _ := ret[0].(*efs.DescribeBackupPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBackupPolicy indicates an expected call of DescribeBackupPolicy.
func (mr *MockEfsClientMockRecorder) DescribeBackupPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackupPolicy", reflect.TypeOf((*MockEfsClient)(nil).DescribeBackupPolicy), varargs...)
}

// DescribeFileSystemPolicy mocks base method.
func (m *MockEfsClient) DescribeFileSystemPolicy(arg0 context.Context, arg1 *efs.DescribeFileSystemPolicyInput, arg2 ...func(*efs.Options)) (*efs.DescribeFileSystemPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeFileSystemPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFileSystemPolicy", varargs...)
	ret0, _ := ret[0].(*efs.DescribeFileSystemPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileSystemPolicy indicates an expected call of DescribeFileSystemPolicy.
func (mr *MockEfsClientMockRecorder) DescribeFileSystemPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystemPolicy", reflect.TypeOf((*MockEfsClient)(nil).DescribeFileSystemPolicy), varargs...)
}

// DescribeFileSystems mocks base method.
func (m *MockEfsClient) DescribeFileSystems(arg0 context.Context, arg1 *efs.DescribeFileSystemsInput, arg2 ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
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
	ret0, _ := ret[0].(*efs.DescribeFileSystemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileSystems indicates an expected call of DescribeFileSystems.
func (mr *MockEfsClientMockRecorder) DescribeFileSystems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystems", reflect.TypeOf((*MockEfsClient)(nil).DescribeFileSystems), varargs...)
}

// DescribeLifecycleConfiguration mocks base method.
func (m *MockEfsClient) DescribeLifecycleConfiguration(arg0 context.Context, arg1 *efs.DescribeLifecycleConfigurationInput, arg2 ...func(*efs.Options)) (*efs.DescribeLifecycleConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeLifecycleConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLifecycleConfiguration", varargs...)
	ret0, _ := ret[0].(*efs.DescribeLifecycleConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLifecycleConfiguration indicates an expected call of DescribeLifecycleConfiguration.
func (mr *MockEfsClientMockRecorder) DescribeLifecycleConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLifecycleConfiguration", reflect.TypeOf((*MockEfsClient)(nil).DescribeLifecycleConfiguration), varargs...)
}

// DescribeMountTargetSecurityGroups mocks base method.
func (m *MockEfsClient) DescribeMountTargetSecurityGroups(arg0 context.Context, arg1 *efs.DescribeMountTargetSecurityGroupsInput, arg2 ...func(*efs.Options)) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeMountTargetSecurityGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMountTargetSecurityGroups", varargs...)
	ret0, _ := ret[0].(*efs.DescribeMountTargetSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMountTargetSecurityGroups indicates an expected call of DescribeMountTargetSecurityGroups.
func (mr *MockEfsClientMockRecorder) DescribeMountTargetSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargetSecurityGroups", reflect.TypeOf((*MockEfsClient)(nil).DescribeMountTargetSecurityGroups), varargs...)
}

// DescribeMountTargets mocks base method.
func (m *MockEfsClient) DescribeMountTargets(arg0 context.Context, arg1 *efs.DescribeMountTargetsInput, arg2 ...func(*efs.Options)) (*efs.DescribeMountTargetsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeMountTargets")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMountTargets", varargs...)
	ret0, _ := ret[0].(*efs.DescribeMountTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMountTargets indicates an expected call of DescribeMountTargets.
func (mr *MockEfsClientMockRecorder) DescribeMountTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargets", reflect.TypeOf((*MockEfsClient)(nil).DescribeMountTargets), varargs...)
}

// DescribeReplicationConfigurations mocks base method.
func (m *MockEfsClient) DescribeReplicationConfigurations(arg0 context.Context, arg1 *efs.DescribeReplicationConfigurationsInput, arg2 ...func(*efs.Options)) (*efs.DescribeReplicationConfigurationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeReplicationConfigurations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReplicationConfigurations", varargs...)
	ret0, _ := ret[0].(*efs.DescribeReplicationConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReplicationConfigurations indicates an expected call of DescribeReplicationConfigurations.
func (mr *MockEfsClientMockRecorder) DescribeReplicationConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReplicationConfigurations", reflect.TypeOf((*MockEfsClient)(nil).DescribeReplicationConfigurations), varargs...)
}

// DescribeTags mocks base method.
func (m *MockEfsClient) DescribeTags(arg0 context.Context, arg1 *efs.DescribeTagsInput, arg2 ...func(*efs.Options)) (*efs.DescribeTagsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeTags")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTags", varargs...)
	ret0, _ := ret[0].(*efs.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTags indicates an expected call of DescribeTags.
func (mr *MockEfsClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTags", reflect.TypeOf((*MockEfsClient)(nil).DescribeTags), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockEfsClient) ListTagsForResource(arg0 context.Context, arg1 *efs.ListTagsForResourceInput, arg2 ...func(*efs.Options)) (*efs.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &efs.Options{}
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
	ret0, _ := ret[0].(*efs.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockEfsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockEfsClient)(nil).ListTagsForResource), varargs...)
}
