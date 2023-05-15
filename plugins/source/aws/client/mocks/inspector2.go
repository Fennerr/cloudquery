// Code generated by MockGen. DO NOT EDIT.
// Source: inspector2.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	inspector2 "github.com/aws/aws-sdk-go-v2/service/inspector2"
	gomock "github.com/golang/mock/gomock"
)

// MockInspector2Client is a mock of Inspector2Client interface.
type MockInspector2Client struct {
	ctrl     *gomock.Controller
	recorder *MockInspector2ClientMockRecorder
}

// MockInspector2ClientMockRecorder is the mock recorder for MockInspector2Client.
type MockInspector2ClientMockRecorder struct {
	mock *MockInspector2Client
}

// NewMockInspector2Client creates a new mock instance.
func NewMockInspector2Client(ctrl *gomock.Controller) *MockInspector2Client {
	mock := &MockInspector2Client{ctrl: ctrl}
	mock.recorder = &MockInspector2ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInspector2Client) EXPECT() *MockInspector2ClientMockRecorder {
	return m.recorder
}

// BatchGetAccountStatus mocks base method.
func (m *MockInspector2Client) BatchGetAccountStatus(arg0 context.Context, arg1 *inspector2.BatchGetAccountStatusInput, arg2 ...func(*inspector2.Options)) (*inspector2.BatchGetAccountStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to BatchGetAccountStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetAccountStatus", varargs...)
	ret0, _ := ret[0].(*inspector2.BatchGetAccountStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetAccountStatus indicates an expected call of BatchGetAccountStatus.
func (mr *MockInspector2ClientMockRecorder) BatchGetAccountStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetAccountStatus", reflect.TypeOf((*MockInspector2Client)(nil).BatchGetAccountStatus), varargs...)
}

// BatchGetFreeTrialInfo mocks base method.
func (m *MockInspector2Client) BatchGetFreeTrialInfo(arg0 context.Context, arg1 *inspector2.BatchGetFreeTrialInfoInput, arg2 ...func(*inspector2.Options)) (*inspector2.BatchGetFreeTrialInfoOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to BatchGetFreeTrialInfo")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetFreeTrialInfo", varargs...)
	ret0, _ := ret[0].(*inspector2.BatchGetFreeTrialInfoOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetFreeTrialInfo indicates an expected call of BatchGetFreeTrialInfo.
func (mr *MockInspector2ClientMockRecorder) BatchGetFreeTrialInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetFreeTrialInfo", reflect.TypeOf((*MockInspector2Client)(nil).BatchGetFreeTrialInfo), varargs...)
}

// DescribeOrganizationConfiguration mocks base method.
func (m *MockInspector2Client) DescribeOrganizationConfiguration(arg0 context.Context, arg1 *inspector2.DescribeOrganizationConfigurationInput, arg2 ...func(*inspector2.Options)) (*inspector2.DescribeOrganizationConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeOrganizationConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeOrganizationConfiguration", varargs...)
	ret0, _ := ret[0].(*inspector2.DescribeOrganizationConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeOrganizationConfiguration indicates an expected call of DescribeOrganizationConfiguration.
func (mr *MockInspector2ClientMockRecorder) DescribeOrganizationConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeOrganizationConfiguration", reflect.TypeOf((*MockInspector2Client)(nil).DescribeOrganizationConfiguration), varargs...)
}

// GetConfiguration mocks base method.
func (m *MockInspector2Client) GetConfiguration(arg0 context.Context, arg1 *inspector2.GetConfigurationInput, arg2 ...func(*inspector2.Options)) (*inspector2.GetConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfiguration", varargs...)
	ret0, _ := ret[0].(*inspector2.GetConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *MockInspector2ClientMockRecorder) GetConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockInspector2Client)(nil).GetConfiguration), varargs...)
}

// GetDelegatedAdminAccount mocks base method.
func (m *MockInspector2Client) GetDelegatedAdminAccount(arg0 context.Context, arg1 *inspector2.GetDelegatedAdminAccountInput, arg2 ...func(*inspector2.Options)) (*inspector2.GetDelegatedAdminAccountOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDelegatedAdminAccount")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDelegatedAdminAccount", varargs...)
	ret0, _ := ret[0].(*inspector2.GetDelegatedAdminAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDelegatedAdminAccount indicates an expected call of GetDelegatedAdminAccount.
func (mr *MockInspector2ClientMockRecorder) GetDelegatedAdminAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDelegatedAdminAccount", reflect.TypeOf((*MockInspector2Client)(nil).GetDelegatedAdminAccount), varargs...)
}

// GetFindingsReportStatus mocks base method.
func (m *MockInspector2Client) GetFindingsReportStatus(arg0 context.Context, arg1 *inspector2.GetFindingsReportStatusInput, arg2 ...func(*inspector2.Options)) (*inspector2.GetFindingsReportStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetFindingsReportStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFindingsReportStatus", varargs...)
	ret0, _ := ret[0].(*inspector2.GetFindingsReportStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFindingsReportStatus indicates an expected call of GetFindingsReportStatus.
func (mr *MockInspector2ClientMockRecorder) GetFindingsReportStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFindingsReportStatus", reflect.TypeOf((*MockInspector2Client)(nil).GetFindingsReportStatus), varargs...)
}

// GetMember mocks base method.
func (m *MockInspector2Client) GetMember(arg0 context.Context, arg1 *inspector2.GetMemberInput, arg2 ...func(*inspector2.Options)) (*inspector2.GetMemberOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetMember")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMember", varargs...)
	ret0, _ := ret[0].(*inspector2.GetMemberOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMember indicates an expected call of GetMember.
func (mr *MockInspector2ClientMockRecorder) GetMember(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMember", reflect.TypeOf((*MockInspector2Client)(nil).GetMember), varargs...)
}

// ListAccountPermissions mocks base method.
func (m *MockInspector2Client) ListAccountPermissions(arg0 context.Context, arg1 *inspector2.ListAccountPermissionsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListAccountPermissionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAccountPermissions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountPermissions", varargs...)
	ret0, _ := ret[0].(*inspector2.ListAccountPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountPermissions indicates an expected call of ListAccountPermissions.
func (mr *MockInspector2ClientMockRecorder) ListAccountPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountPermissions", reflect.TypeOf((*MockInspector2Client)(nil).ListAccountPermissions), varargs...)
}

// ListCoverage mocks base method.
func (m *MockInspector2Client) ListCoverage(arg0 context.Context, arg1 *inspector2.ListCoverageInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListCoverageOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListCoverage")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCoverage", varargs...)
	ret0, _ := ret[0].(*inspector2.ListCoverageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCoverage indicates an expected call of ListCoverage.
func (mr *MockInspector2ClientMockRecorder) ListCoverage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCoverage", reflect.TypeOf((*MockInspector2Client)(nil).ListCoverage), varargs...)
}

// ListCoverageStatistics mocks base method.
func (m *MockInspector2Client) ListCoverageStatistics(arg0 context.Context, arg1 *inspector2.ListCoverageStatisticsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListCoverageStatisticsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListCoverageStatistics")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCoverageStatistics", varargs...)
	ret0, _ := ret[0].(*inspector2.ListCoverageStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCoverageStatistics indicates an expected call of ListCoverageStatistics.
func (mr *MockInspector2ClientMockRecorder) ListCoverageStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCoverageStatistics", reflect.TypeOf((*MockInspector2Client)(nil).ListCoverageStatistics), varargs...)
}

// ListDelegatedAdminAccounts mocks base method.
func (m *MockInspector2Client) ListDelegatedAdminAccounts(arg0 context.Context, arg1 *inspector2.ListDelegatedAdminAccountsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListDelegatedAdminAccountsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDelegatedAdminAccounts")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDelegatedAdminAccounts", varargs...)
	ret0, _ := ret[0].(*inspector2.ListDelegatedAdminAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDelegatedAdminAccounts indicates an expected call of ListDelegatedAdminAccounts.
func (mr *MockInspector2ClientMockRecorder) ListDelegatedAdminAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDelegatedAdminAccounts", reflect.TypeOf((*MockInspector2Client)(nil).ListDelegatedAdminAccounts), varargs...)
}

// ListFilters mocks base method.
func (m *MockInspector2Client) ListFilters(arg0 context.Context, arg1 *inspector2.ListFiltersInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListFiltersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListFilters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFilters", varargs...)
	ret0, _ := ret[0].(*inspector2.ListFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFilters indicates an expected call of ListFilters.
func (mr *MockInspector2ClientMockRecorder) ListFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFilters", reflect.TypeOf((*MockInspector2Client)(nil).ListFilters), varargs...)
}

// ListFindingAggregations mocks base method.
func (m *MockInspector2Client) ListFindingAggregations(arg0 context.Context, arg1 *inspector2.ListFindingAggregationsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListFindingAggregationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListFindingAggregations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFindingAggregations", varargs...)
	ret0, _ := ret[0].(*inspector2.ListFindingAggregationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFindingAggregations indicates an expected call of ListFindingAggregations.
func (mr *MockInspector2ClientMockRecorder) ListFindingAggregations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFindingAggregations", reflect.TypeOf((*MockInspector2Client)(nil).ListFindingAggregations), varargs...)
}

// ListFindings mocks base method.
func (m *MockInspector2Client) ListFindings(arg0 context.Context, arg1 *inspector2.ListFindingsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListFindingsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListFindings")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFindings", varargs...)
	ret0, _ := ret[0].(*inspector2.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFindings indicates an expected call of ListFindings.
func (mr *MockInspector2ClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFindings", reflect.TypeOf((*MockInspector2Client)(nil).ListFindings), varargs...)
}

// ListMembers mocks base method.
func (m *MockInspector2Client) ListMembers(arg0 context.Context, arg1 *inspector2.ListMembersInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListMembersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListMembers")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMembers", varargs...)
	ret0, _ := ret[0].(*inspector2.ListMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembers indicates an expected call of ListMembers.
func (mr *MockInspector2ClientMockRecorder) ListMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockInspector2Client)(nil).ListMembers), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockInspector2Client) ListTagsForResource(arg0 context.Context, arg1 *inspector2.ListTagsForResourceInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
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
	ret0, _ := ret[0].(*inspector2.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockInspector2ClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockInspector2Client)(nil).ListTagsForResource), varargs...)
}

// ListUsageTotals mocks base method.
func (m *MockInspector2Client) ListUsageTotals(arg0 context.Context, arg1 *inspector2.ListUsageTotalsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListUsageTotalsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListUsageTotals")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsageTotals", varargs...)
	ret0, _ := ret[0].(*inspector2.ListUsageTotalsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsageTotals indicates an expected call of ListUsageTotals.
func (mr *MockInspector2ClientMockRecorder) ListUsageTotals(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsageTotals", reflect.TypeOf((*MockInspector2Client)(nil).ListUsageTotals), varargs...)
}
