// Code generated by MockGen. DO NOT EDIT.
// Source: elastictranscoder.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	elastictranscoder "github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	gomock "github.com/golang/mock/gomock"
)

// MockElastictranscoderClient is a mock of ElastictranscoderClient interface.
type MockElastictranscoderClient struct {
	ctrl     *gomock.Controller
	recorder *MockElastictranscoderClientMockRecorder
}

// MockElastictranscoderClientMockRecorder is the mock recorder for MockElastictranscoderClient.
type MockElastictranscoderClientMockRecorder struct {
	mock *MockElastictranscoderClient
}

// NewMockElastictranscoderClient creates a new mock instance.
func NewMockElastictranscoderClient(ctrl *gomock.Controller) *MockElastictranscoderClient {
	mock := &MockElastictranscoderClient{ctrl: ctrl}
	mock.recorder = &MockElastictranscoderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElastictranscoderClient) EXPECT() *MockElastictranscoderClientMockRecorder {
	return m.recorder
}

// ListJobsByPipeline mocks base method.
func (m *MockElastictranscoderClient) ListJobsByPipeline(arg0 context.Context, arg1 *elastictranscoder.ListJobsByPipelineInput, arg2 ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByPipelineOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elastictranscoder.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListJobsByPipeline")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJobsByPipeline", varargs...)
	ret0, _ := ret[0].(*elastictranscoder.ListJobsByPipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJobsByPipeline indicates an expected call of ListJobsByPipeline.
func (mr *MockElastictranscoderClientMockRecorder) ListJobsByPipeline(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobsByPipeline", reflect.TypeOf((*MockElastictranscoderClient)(nil).ListJobsByPipeline), varargs...)
}

// ListJobsByStatus mocks base method.
func (m *MockElastictranscoderClient) ListJobsByStatus(arg0 context.Context, arg1 *elastictranscoder.ListJobsByStatusInput, arg2 ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elastictranscoder.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListJobsByStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJobsByStatus", varargs...)
	ret0, _ := ret[0].(*elastictranscoder.ListJobsByStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJobsByStatus indicates an expected call of ListJobsByStatus.
func (mr *MockElastictranscoderClientMockRecorder) ListJobsByStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobsByStatus", reflect.TypeOf((*MockElastictranscoderClient)(nil).ListJobsByStatus), varargs...)
}

// ListPipelines mocks base method.
func (m *MockElastictranscoderClient) ListPipelines(arg0 context.Context, arg1 *elastictranscoder.ListPipelinesInput, arg2 ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPipelinesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elastictranscoder.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPipelines")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPipelines", varargs...)
	ret0, _ := ret[0].(*elastictranscoder.ListPipelinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPipelines indicates an expected call of ListPipelines.
func (mr *MockElastictranscoderClientMockRecorder) ListPipelines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelines", reflect.TypeOf((*MockElastictranscoderClient)(nil).ListPipelines), varargs...)
}

// ListPresets mocks base method.
func (m *MockElastictranscoderClient) ListPresets(arg0 context.Context, arg1 *elastictranscoder.ListPresetsInput, arg2 ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPresetsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elastictranscoder.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPresets")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPresets", varargs...)
	ret0, _ := ret[0].(*elastictranscoder.ListPresetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPresets indicates an expected call of ListPresets.
func (mr *MockElastictranscoderClientMockRecorder) ListPresets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPresets", reflect.TypeOf((*MockElastictranscoderClient)(nil).ListPresets), varargs...)
}
