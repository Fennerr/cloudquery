// Code generated by MockGen. DO NOT EDIT.
// Source: ecr.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	ecr "github.com/aws/aws-sdk-go-v2/service/ecr"
	gomock "github.com/golang/mock/gomock"
)

// MockEcrClient is a mock of EcrClient interface.
type MockEcrClient struct {
	ctrl     *gomock.Controller
	recorder *MockEcrClientMockRecorder
}

// MockEcrClientMockRecorder is the mock recorder for MockEcrClient.
type MockEcrClientMockRecorder struct {
	mock *MockEcrClient
}

// NewMockEcrClient creates a new mock instance.
func NewMockEcrClient(ctrl *gomock.Controller) *MockEcrClient {
	mock := &MockEcrClient{ctrl: ctrl}
	mock.recorder = &MockEcrClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEcrClient) EXPECT() *MockEcrClientMockRecorder {
	return m.recorder
}

// BatchGetImage mocks base method.
func (m *MockEcrClient) BatchGetImage(arg0 context.Context, arg1 *ecr.BatchGetImageInput, arg2 ...func(*ecr.Options)) (*ecr.BatchGetImageOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to BatchGetImage")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetImage", varargs...)
	ret0, _ := ret[0].(*ecr.BatchGetImageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetImage indicates an expected call of BatchGetImage.
func (mr *MockEcrClientMockRecorder) BatchGetImage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetImage", reflect.TypeOf((*MockEcrClient)(nil).BatchGetImage), varargs...)
}

// BatchGetRepositoryScanningConfiguration mocks base method.
func (m *MockEcrClient) BatchGetRepositoryScanningConfiguration(arg0 context.Context, arg1 *ecr.BatchGetRepositoryScanningConfigurationInput, arg2 ...func(*ecr.Options)) (*ecr.BatchGetRepositoryScanningConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to BatchGetRepositoryScanningConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetRepositoryScanningConfiguration", varargs...)
	ret0, _ := ret[0].(*ecr.BatchGetRepositoryScanningConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetRepositoryScanningConfiguration indicates an expected call of BatchGetRepositoryScanningConfiguration.
func (mr *MockEcrClientMockRecorder) BatchGetRepositoryScanningConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetRepositoryScanningConfiguration", reflect.TypeOf((*MockEcrClient)(nil).BatchGetRepositoryScanningConfiguration), varargs...)
}

// DescribeImageReplicationStatus mocks base method.
func (m *MockEcrClient) DescribeImageReplicationStatus(arg0 context.Context, arg1 *ecr.DescribeImageReplicationStatusInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeImageReplicationStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeImageReplicationStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImageReplicationStatus", varargs...)
	ret0, _ := ret[0].(*ecr.DescribeImageReplicationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeImageReplicationStatus indicates an expected call of DescribeImageReplicationStatus.
func (mr *MockEcrClientMockRecorder) DescribeImageReplicationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImageReplicationStatus", reflect.TypeOf((*MockEcrClient)(nil).DescribeImageReplicationStatus), varargs...)
}

// DescribeImageScanFindings mocks base method.
func (m *MockEcrClient) DescribeImageScanFindings(arg0 context.Context, arg1 *ecr.DescribeImageScanFindingsInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeImageScanFindingsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeImageScanFindings")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImageScanFindings", varargs...)
	ret0, _ := ret[0].(*ecr.DescribeImageScanFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeImageScanFindings indicates an expected call of DescribeImageScanFindings.
func (mr *MockEcrClientMockRecorder) DescribeImageScanFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImageScanFindings", reflect.TypeOf((*MockEcrClient)(nil).DescribeImageScanFindings), varargs...)
}

// DescribeImages mocks base method.
func (m *MockEcrClient) DescribeImages(arg0 context.Context, arg1 *ecr.DescribeImagesInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeImagesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeImages")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImages", varargs...)
	ret0, _ := ret[0].(*ecr.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeImages indicates an expected call of DescribeImages.
func (mr *MockEcrClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImages", reflect.TypeOf((*MockEcrClient)(nil).DescribeImages), varargs...)
}

// DescribePullThroughCacheRules mocks base method.
func (m *MockEcrClient) DescribePullThroughCacheRules(arg0 context.Context, arg1 *ecr.DescribePullThroughCacheRulesInput, arg2 ...func(*ecr.Options)) (*ecr.DescribePullThroughCacheRulesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribePullThroughCacheRules")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePullThroughCacheRules", varargs...)
	ret0, _ := ret[0].(*ecr.DescribePullThroughCacheRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePullThroughCacheRules indicates an expected call of DescribePullThroughCacheRules.
func (mr *MockEcrClientMockRecorder) DescribePullThroughCacheRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePullThroughCacheRules", reflect.TypeOf((*MockEcrClient)(nil).DescribePullThroughCacheRules), varargs...)
}

// DescribeRegistry mocks base method.
func (m *MockEcrClient) DescribeRegistry(arg0 context.Context, arg1 *ecr.DescribeRegistryInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeRegistryOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeRegistry")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRegistry", varargs...)
	ret0, _ := ret[0].(*ecr.DescribeRegistryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRegistry indicates an expected call of DescribeRegistry.
func (mr *MockEcrClientMockRecorder) DescribeRegistry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRegistry", reflect.TypeOf((*MockEcrClient)(nil).DescribeRegistry), varargs...)
}

// DescribeRepositories mocks base method.
func (m *MockEcrClient) DescribeRepositories(arg0 context.Context, arg1 *ecr.DescribeRepositoriesInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeRepositoriesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeRepositories")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRepositories", varargs...)
	ret0, _ := ret[0].(*ecr.DescribeRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRepositories indicates an expected call of DescribeRepositories.
func (mr *MockEcrClientMockRecorder) DescribeRepositories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRepositories", reflect.TypeOf((*MockEcrClient)(nil).DescribeRepositories), varargs...)
}

// GetAuthorizationToken mocks base method.
func (m *MockEcrClient) GetAuthorizationToken(arg0 context.Context, arg1 *ecr.GetAuthorizationTokenInput, arg2 ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAuthorizationToken")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuthorizationToken", varargs...)
	ret0, _ := ret[0].(*ecr.GetAuthorizationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizationToken indicates an expected call of GetAuthorizationToken.
func (mr *MockEcrClientMockRecorder) GetAuthorizationToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizationToken", reflect.TypeOf((*MockEcrClient)(nil).GetAuthorizationToken), varargs...)
}

// GetDownloadUrlForLayer mocks base method.
func (m *MockEcrClient) GetDownloadUrlForLayer(arg0 context.Context, arg1 *ecr.GetDownloadUrlForLayerInput, arg2 ...func(*ecr.Options)) (*ecr.GetDownloadUrlForLayerOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDownloadUrlForLayer")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDownloadUrlForLayer", varargs...)
	ret0, _ := ret[0].(*ecr.GetDownloadUrlForLayerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDownloadUrlForLayer indicates an expected call of GetDownloadUrlForLayer.
func (mr *MockEcrClientMockRecorder) GetDownloadUrlForLayer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadUrlForLayer", reflect.TypeOf((*MockEcrClient)(nil).GetDownloadUrlForLayer), varargs...)
}

// GetLifecyclePolicy mocks base method.
func (m *MockEcrClient) GetLifecyclePolicy(arg0 context.Context, arg1 *ecr.GetLifecyclePolicyInput, arg2 ...func(*ecr.Options)) (*ecr.GetLifecyclePolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetLifecyclePolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLifecyclePolicy", varargs...)
	ret0, _ := ret[0].(*ecr.GetLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifecyclePolicy indicates an expected call of GetLifecyclePolicy.
func (mr *MockEcrClientMockRecorder) GetLifecyclePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePolicy", reflect.TypeOf((*MockEcrClient)(nil).GetLifecyclePolicy), varargs...)
}

// GetLifecyclePolicyPreview mocks base method.
func (m *MockEcrClient) GetLifecyclePolicyPreview(arg0 context.Context, arg1 *ecr.GetLifecyclePolicyPreviewInput, arg2 ...func(*ecr.Options)) (*ecr.GetLifecyclePolicyPreviewOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetLifecyclePolicyPreview")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLifecyclePolicyPreview", varargs...)
	ret0, _ := ret[0].(*ecr.GetLifecyclePolicyPreviewOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifecyclePolicyPreview indicates an expected call of GetLifecyclePolicyPreview.
func (mr *MockEcrClientMockRecorder) GetLifecyclePolicyPreview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePolicyPreview", reflect.TypeOf((*MockEcrClient)(nil).GetLifecyclePolicyPreview), varargs...)
}

// GetRegistryPolicy mocks base method.
func (m *MockEcrClient) GetRegistryPolicy(arg0 context.Context, arg1 *ecr.GetRegistryPolicyInput, arg2 ...func(*ecr.Options)) (*ecr.GetRegistryPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetRegistryPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegistryPolicy", varargs...)
	ret0, _ := ret[0].(*ecr.GetRegistryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegistryPolicy indicates an expected call of GetRegistryPolicy.
func (mr *MockEcrClientMockRecorder) GetRegistryPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistryPolicy", reflect.TypeOf((*MockEcrClient)(nil).GetRegistryPolicy), varargs...)
}

// GetRegistryScanningConfiguration mocks base method.
func (m *MockEcrClient) GetRegistryScanningConfiguration(arg0 context.Context, arg1 *ecr.GetRegistryScanningConfigurationInput, arg2 ...func(*ecr.Options)) (*ecr.GetRegistryScanningConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetRegistryScanningConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegistryScanningConfiguration", varargs...)
	ret0, _ := ret[0].(*ecr.GetRegistryScanningConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegistryScanningConfiguration indicates an expected call of GetRegistryScanningConfiguration.
func (mr *MockEcrClientMockRecorder) GetRegistryScanningConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistryScanningConfiguration", reflect.TypeOf((*MockEcrClient)(nil).GetRegistryScanningConfiguration), varargs...)
}

// GetRepositoryPolicy mocks base method.
func (m *MockEcrClient) GetRepositoryPolicy(arg0 context.Context, arg1 *ecr.GetRepositoryPolicyInput, arg2 ...func(*ecr.Options)) (*ecr.GetRepositoryPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetRepositoryPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRepositoryPolicy", varargs...)
	ret0, _ := ret[0].(*ecr.GetRepositoryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryPolicy indicates an expected call of GetRepositoryPolicy.
func (mr *MockEcrClientMockRecorder) GetRepositoryPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryPolicy", reflect.TypeOf((*MockEcrClient)(nil).GetRepositoryPolicy), varargs...)
}

// ListImages mocks base method.
func (m *MockEcrClient) ListImages(arg0 context.Context, arg1 *ecr.ListImagesInput, arg2 ...func(*ecr.Options)) (*ecr.ListImagesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListImages")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListImages", varargs...)
	ret0, _ := ret[0].(*ecr.ListImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockEcrClientMockRecorder) ListImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockEcrClient)(nil).ListImages), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockEcrClient) ListTagsForResource(arg0 context.Context, arg1 *ecr.ListTagsForResourceInput, arg2 ...func(*ecr.Options)) (*ecr.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ecr.Options{}
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
	ret0, _ := ret[0].(*ecr.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockEcrClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockEcrClient)(nil).ListTagsForResource), varargs...)
}
