// Code generated by MockGen. DO NOT EDIT.
// Source: applicationautoscaling.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	applicationautoscaling "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	gomock "github.com/golang/mock/gomock"
)

// MockApplicationautoscalingClient is a mock of ApplicationautoscalingClient interface.
type MockApplicationautoscalingClient struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationautoscalingClientMockRecorder
}

// MockApplicationautoscalingClientMockRecorder is the mock recorder for MockApplicationautoscalingClient.
type MockApplicationautoscalingClientMockRecorder struct {
	mock *MockApplicationautoscalingClient
}

// NewMockApplicationautoscalingClient creates a new mock instance.
func NewMockApplicationautoscalingClient(ctrl *gomock.Controller) *MockApplicationautoscalingClient {
	mock := &MockApplicationautoscalingClient{ctrl: ctrl}
	mock.recorder = &MockApplicationautoscalingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationautoscalingClient) EXPECT() *MockApplicationautoscalingClientMockRecorder {
	return m.recorder
}

// DescribeScalableTargets mocks base method.
func (m *MockApplicationautoscalingClient) DescribeScalableTargets(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalableTargetsInput, arg2 ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &applicationautoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeScalableTargets")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalableTargets", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalableTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalableTargets indicates an expected call of DescribeScalableTargets.
func (mr *MockApplicationautoscalingClientMockRecorder) DescribeScalableTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalableTargets", reflect.TypeOf((*MockApplicationautoscalingClient)(nil).DescribeScalableTargets), varargs...)
}

// DescribeScalingActivities mocks base method.
func (m *MockApplicationautoscalingClient) DescribeScalingActivities(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalingActivitiesInput, arg2 ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &applicationautoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeScalingActivities")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingActivities", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingActivitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingActivities indicates an expected call of DescribeScalingActivities.
func (mr *MockApplicationautoscalingClientMockRecorder) DescribeScalingActivities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingActivities", reflect.TypeOf((*MockApplicationautoscalingClient)(nil).DescribeScalingActivities), varargs...)
}

// DescribeScalingPolicies mocks base method.
func (m *MockApplicationautoscalingClient) DescribeScalingPolicies(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalingPoliciesInput, arg2 ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &applicationautoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeScalingPolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingPolicies", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPolicies indicates an expected call of DescribeScalingPolicies.
func (mr *MockApplicationautoscalingClientMockRecorder) DescribeScalingPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPolicies", reflect.TypeOf((*MockApplicationautoscalingClient)(nil).DescribeScalingPolicies), varargs...)
}

// DescribeScheduledActions mocks base method.
func (m *MockApplicationautoscalingClient) DescribeScheduledActions(arg0 context.Context, arg1 *applicationautoscaling.DescribeScheduledActionsInput, arg2 ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &applicationautoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeScheduledActions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScheduledActions", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScheduledActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScheduledActions indicates an expected call of DescribeScheduledActions.
func (mr *MockApplicationautoscalingClientMockRecorder) DescribeScheduledActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScheduledActions", reflect.TypeOf((*MockApplicationautoscalingClient)(nil).DescribeScheduledActions), varargs...)
}
