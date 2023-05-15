// Code generated by MockGen. DO NOT EDIT.
// Source: organizations.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	organizations "github.com/aws/aws-sdk-go-v2/service/organizations"
	gomock "github.com/golang/mock/gomock"
)

// MockOrganizationsClient is a mock of OrganizationsClient interface.
type MockOrganizationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsClientMockRecorder
}

// MockOrganizationsClientMockRecorder is the mock recorder for MockOrganizationsClient.
type MockOrganizationsClientMockRecorder struct {
	mock *MockOrganizationsClient
}

// NewMockOrganizationsClient creates a new mock instance.
func NewMockOrganizationsClient(ctrl *gomock.Controller) *MockOrganizationsClient {
	mock := &MockOrganizationsClient{ctrl: ctrl}
	mock.recorder = &MockOrganizationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsClient) EXPECT() *MockOrganizationsClientMockRecorder {
	return m.recorder
}

// DescribeAccount mocks base method.
func (m *MockOrganizationsClient) DescribeAccount(arg0 context.Context, arg1 *organizations.DescribeAccountInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeAccountOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAccount")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccount", varargs...)
	ret0, _ := ret[0].(*organizations.DescribeAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccount indicates an expected call of DescribeAccount.
func (mr *MockOrganizationsClientMockRecorder) DescribeAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccount", reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeAccount), varargs...)
}

// DescribeCreateAccountStatus mocks base method.
func (m *MockOrganizationsClient) DescribeCreateAccountStatus(arg0 context.Context, arg1 *organizations.DescribeCreateAccountStatusInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeCreateAccountStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCreateAccountStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCreateAccountStatus", varargs...)
	ret0, _ := ret[0].(*organizations.DescribeCreateAccountStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCreateAccountStatus indicates an expected call of DescribeCreateAccountStatus.
func (mr *MockOrganizationsClientMockRecorder) DescribeCreateAccountStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCreateAccountStatus", reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeCreateAccountStatus), varargs...)
}

// DescribeEffectivePolicy mocks base method.
func (m *MockOrganizationsClient) DescribeEffectivePolicy(arg0 context.Context, arg1 *organizations.DescribeEffectivePolicyInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeEffectivePolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEffectivePolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEffectivePolicy", varargs...)
	ret0, _ := ret[0].(*organizations.DescribeEffectivePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEffectivePolicy indicates an expected call of DescribeEffectivePolicy.
func (mr *MockOrganizationsClientMockRecorder) DescribeEffectivePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEffectivePolicy", reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeEffectivePolicy), varargs...)
}

// DescribeHandshake mocks base method.
func (m *MockOrganizationsClient) DescribeHandshake(arg0 context.Context, arg1 *organizations.DescribeHandshakeInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeHandshakeOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeHandshake")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeHandshake", varargs...)
	ret0, _ := ret[0].(*organizations.DescribeHandshakeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHandshake indicates an expected call of DescribeHandshake.
func (mr *MockOrganizationsClientMockRecorder) DescribeHandshake(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHandshake", reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeHandshake), varargs...)
}

// DescribeOrganization mocks base method.
func (m *MockOrganizationsClient) DescribeOrganization(arg0 context.Context, arg1 *organizations.DescribeOrganizationInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeOrganizationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeOrganization")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeOrganization", varargs...)
	ret0, _ := ret[0].(*organizations.DescribeOrganizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeOrganization indicates an expected call of DescribeOrganization.
func (mr *MockOrganizationsClientMockRecorder) DescribeOrganization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeOrganization", reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeOrganization), varargs...)
}

// DescribeOrganizationalUnit mocks base method.
func (m *MockOrganizationsClient) DescribeOrganizationalUnit(arg0 context.Context, arg1 *organizations.DescribeOrganizationalUnitInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeOrganizationalUnitOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeOrganizationalUnit")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeOrganizationalUnit", varargs...)
	ret0, _ := ret[0].(*organizations.DescribeOrganizationalUnitOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeOrganizationalUnit indicates an expected call of DescribeOrganizationalUnit.
func (mr *MockOrganizationsClientMockRecorder) DescribeOrganizationalUnit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeOrganizationalUnit", reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeOrganizationalUnit), varargs...)
}

// DescribePolicy mocks base method.
func (m *MockOrganizationsClient) DescribePolicy(arg0 context.Context, arg1 *organizations.DescribePolicyInput, arg2 ...func(*organizations.Options)) (*organizations.DescribePolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribePolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePolicy", varargs...)
	ret0, _ := ret[0].(*organizations.DescribePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePolicy indicates an expected call of DescribePolicy.
func (mr *MockOrganizationsClientMockRecorder) DescribePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePolicy", reflect.TypeOf((*MockOrganizationsClient)(nil).DescribePolicy), varargs...)
}

// DescribeResourcePolicy mocks base method.
func (m *MockOrganizationsClient) DescribeResourcePolicy(arg0 context.Context, arg1 *organizations.DescribeResourcePolicyInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeResourcePolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeResourcePolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeResourcePolicy", varargs...)
	ret0, _ := ret[0].(*organizations.DescribeResourcePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeResourcePolicy indicates an expected call of DescribeResourcePolicy.
func (mr *MockOrganizationsClientMockRecorder) DescribeResourcePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeResourcePolicy", reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeResourcePolicy), varargs...)
}

// ListAWSServiceAccessForOrganization mocks base method.
func (m *MockOrganizationsClient) ListAWSServiceAccessForOrganization(arg0 context.Context, arg1 *organizations.ListAWSServiceAccessForOrganizationInput, arg2 ...func(*organizations.Options)) (*organizations.ListAWSServiceAccessForOrganizationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAWSServiceAccessForOrganization")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAWSServiceAccessForOrganization", varargs...)
	ret0, _ := ret[0].(*organizations.ListAWSServiceAccessForOrganizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAWSServiceAccessForOrganization indicates an expected call of ListAWSServiceAccessForOrganization.
func (mr *MockOrganizationsClientMockRecorder) ListAWSServiceAccessForOrganization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAWSServiceAccessForOrganization", reflect.TypeOf((*MockOrganizationsClient)(nil).ListAWSServiceAccessForOrganization), varargs...)
}

// ListAccounts mocks base method.
func (m *MockOrganizationsClient) ListAccounts(arg0 context.Context, arg1 *organizations.ListAccountsInput, arg2 ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAccounts")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccounts", varargs...)
	ret0, _ := ret[0].(*organizations.ListAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockOrganizationsClientMockRecorder) ListAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockOrganizationsClient)(nil).ListAccounts), varargs...)
}

// ListAccountsForParent mocks base method.
func (m *MockOrganizationsClient) ListAccountsForParent(arg0 context.Context, arg1 *organizations.ListAccountsForParentInput, arg2 ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAccountsForParent")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountsForParent", varargs...)
	ret0, _ := ret[0].(*organizations.ListAccountsForParentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountsForParent indicates an expected call of ListAccountsForParent.
func (mr *MockOrganizationsClientMockRecorder) ListAccountsForParent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsForParent", reflect.TypeOf((*MockOrganizationsClient)(nil).ListAccountsForParent), varargs...)
}

// ListChildren mocks base method.
func (m *MockOrganizationsClient) ListChildren(arg0 context.Context, arg1 *organizations.ListChildrenInput, arg2 ...func(*organizations.Options)) (*organizations.ListChildrenOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListChildren")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListChildren", varargs...)
	ret0, _ := ret[0].(*organizations.ListChildrenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChildren indicates an expected call of ListChildren.
func (mr *MockOrganizationsClientMockRecorder) ListChildren(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChildren", reflect.TypeOf((*MockOrganizationsClient)(nil).ListChildren), varargs...)
}

// ListCreateAccountStatus mocks base method.
func (m *MockOrganizationsClient) ListCreateAccountStatus(arg0 context.Context, arg1 *organizations.ListCreateAccountStatusInput, arg2 ...func(*organizations.Options)) (*organizations.ListCreateAccountStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListCreateAccountStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCreateAccountStatus", varargs...)
	ret0, _ := ret[0].(*organizations.ListCreateAccountStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCreateAccountStatus indicates an expected call of ListCreateAccountStatus.
func (mr *MockOrganizationsClientMockRecorder) ListCreateAccountStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCreateAccountStatus", reflect.TypeOf((*MockOrganizationsClient)(nil).ListCreateAccountStatus), varargs...)
}

// ListDelegatedAdministrators mocks base method.
func (m *MockOrganizationsClient) ListDelegatedAdministrators(arg0 context.Context, arg1 *organizations.ListDelegatedAdministratorsInput, arg2 ...func(*organizations.Options)) (*organizations.ListDelegatedAdministratorsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDelegatedAdministrators")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDelegatedAdministrators", varargs...)
	ret0, _ := ret[0].(*organizations.ListDelegatedAdministratorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDelegatedAdministrators indicates an expected call of ListDelegatedAdministrators.
func (mr *MockOrganizationsClientMockRecorder) ListDelegatedAdministrators(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDelegatedAdministrators", reflect.TypeOf((*MockOrganizationsClient)(nil).ListDelegatedAdministrators), varargs...)
}

// ListDelegatedServicesForAccount mocks base method.
func (m *MockOrganizationsClient) ListDelegatedServicesForAccount(arg0 context.Context, arg1 *organizations.ListDelegatedServicesForAccountInput, arg2 ...func(*organizations.Options)) (*organizations.ListDelegatedServicesForAccountOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDelegatedServicesForAccount")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDelegatedServicesForAccount", varargs...)
	ret0, _ := ret[0].(*organizations.ListDelegatedServicesForAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDelegatedServicesForAccount indicates an expected call of ListDelegatedServicesForAccount.
func (mr *MockOrganizationsClientMockRecorder) ListDelegatedServicesForAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDelegatedServicesForAccount", reflect.TypeOf((*MockOrganizationsClient)(nil).ListDelegatedServicesForAccount), varargs...)
}

// ListHandshakesForAccount mocks base method.
func (m *MockOrganizationsClient) ListHandshakesForAccount(arg0 context.Context, arg1 *organizations.ListHandshakesForAccountInput, arg2 ...func(*organizations.Options)) (*organizations.ListHandshakesForAccountOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListHandshakesForAccount")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHandshakesForAccount", varargs...)
	ret0, _ := ret[0].(*organizations.ListHandshakesForAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHandshakesForAccount indicates an expected call of ListHandshakesForAccount.
func (mr *MockOrganizationsClientMockRecorder) ListHandshakesForAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHandshakesForAccount", reflect.TypeOf((*MockOrganizationsClient)(nil).ListHandshakesForAccount), varargs...)
}

// ListHandshakesForOrganization mocks base method.
func (m *MockOrganizationsClient) ListHandshakesForOrganization(arg0 context.Context, arg1 *organizations.ListHandshakesForOrganizationInput, arg2 ...func(*organizations.Options)) (*organizations.ListHandshakesForOrganizationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListHandshakesForOrganization")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHandshakesForOrganization", varargs...)
	ret0, _ := ret[0].(*organizations.ListHandshakesForOrganizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHandshakesForOrganization indicates an expected call of ListHandshakesForOrganization.
func (mr *MockOrganizationsClientMockRecorder) ListHandshakesForOrganization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHandshakesForOrganization", reflect.TypeOf((*MockOrganizationsClient)(nil).ListHandshakesForOrganization), varargs...)
}

// ListOrganizationalUnitsForParent mocks base method.
func (m *MockOrganizationsClient) ListOrganizationalUnitsForParent(arg0 context.Context, arg1 *organizations.ListOrganizationalUnitsForParentInput, arg2 ...func(*organizations.Options)) (*organizations.ListOrganizationalUnitsForParentOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListOrganizationalUnitsForParent")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizationalUnitsForParent", varargs...)
	ret0, _ := ret[0].(*organizations.ListOrganizationalUnitsForParentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizationalUnitsForParent indicates an expected call of ListOrganizationalUnitsForParent.
func (mr *MockOrganizationsClientMockRecorder) ListOrganizationalUnitsForParent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizationalUnitsForParent", reflect.TypeOf((*MockOrganizationsClient)(nil).ListOrganizationalUnitsForParent), varargs...)
}

// ListParents mocks base method.
func (m *MockOrganizationsClient) ListParents(arg0 context.Context, arg1 *organizations.ListParentsInput, arg2 ...func(*organizations.Options)) (*organizations.ListParentsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListParents")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListParents", varargs...)
	ret0, _ := ret[0].(*organizations.ListParentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListParents indicates an expected call of ListParents.
func (mr *MockOrganizationsClientMockRecorder) ListParents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListParents", reflect.TypeOf((*MockOrganizationsClient)(nil).ListParents), varargs...)
}

// ListPolicies mocks base method.
func (m *MockOrganizationsClient) ListPolicies(arg0 context.Context, arg1 *organizations.ListPoliciesInput, arg2 ...func(*organizations.Options)) (*organizations.ListPoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPolicies", varargs...)
	ret0, _ := ret[0].(*organizations.ListPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicies indicates an expected call of ListPolicies.
func (mr *MockOrganizationsClientMockRecorder) ListPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicies", reflect.TypeOf((*MockOrganizationsClient)(nil).ListPolicies), varargs...)
}

// ListPoliciesForTarget mocks base method.
func (m *MockOrganizationsClient) ListPoliciesForTarget(arg0 context.Context, arg1 *organizations.ListPoliciesForTargetInput, arg2 ...func(*organizations.Options)) (*organizations.ListPoliciesForTargetOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPoliciesForTarget")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPoliciesForTarget", varargs...)
	ret0, _ := ret[0].(*organizations.ListPoliciesForTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPoliciesForTarget indicates an expected call of ListPoliciesForTarget.
func (mr *MockOrganizationsClientMockRecorder) ListPoliciesForTarget(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPoliciesForTarget", reflect.TypeOf((*MockOrganizationsClient)(nil).ListPoliciesForTarget), varargs...)
}

// ListRoots mocks base method.
func (m *MockOrganizationsClient) ListRoots(arg0 context.Context, arg1 *organizations.ListRootsInput, arg2 ...func(*organizations.Options)) (*organizations.ListRootsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListRoots")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRoots", varargs...)
	ret0, _ := ret[0].(*organizations.ListRootsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoots indicates an expected call of ListRoots.
func (mr *MockOrganizationsClientMockRecorder) ListRoots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoots", reflect.TypeOf((*MockOrganizationsClient)(nil).ListRoots), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockOrganizationsClient) ListTagsForResource(arg0 context.Context, arg1 *organizations.ListTagsForResourceInput, arg2 ...func(*organizations.Options)) (*organizations.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
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
	ret0, _ := ret[0].(*organizations.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockOrganizationsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockOrganizationsClient)(nil).ListTagsForResource), varargs...)
}

// ListTargetsForPolicy mocks base method.
func (m *MockOrganizationsClient) ListTargetsForPolicy(arg0 context.Context, arg1 *organizations.ListTargetsForPolicyInput, arg2 ...func(*organizations.Options)) (*organizations.ListTargetsForPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &organizations.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTargetsForPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTargetsForPolicy", varargs...)
	ret0, _ := ret[0].(*organizations.ListTargetsForPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTargetsForPolicy indicates an expected call of ListTargetsForPolicy.
func (mr *MockOrganizationsClientMockRecorder) ListTargetsForPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTargetsForPolicy", reflect.TypeOf((*MockOrganizationsClient)(nil).ListTargetsForPolicy), varargs...)
}
