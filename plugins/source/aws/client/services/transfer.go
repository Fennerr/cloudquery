// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/transfer.go -source=transfer.go TransferClient
type TransferClient interface {
	DescribeAccess(context.Context, *transfer.DescribeAccessInput, ...func(*transfer.Options)) (*transfer.DescribeAccessOutput, error)
	DescribeAgreement(context.Context, *transfer.DescribeAgreementInput, ...func(*transfer.Options)) (*transfer.DescribeAgreementOutput, error)
	DescribeCertificate(context.Context, *transfer.DescribeCertificateInput, ...func(*transfer.Options)) (*transfer.DescribeCertificateOutput, error)
	DescribeConnector(context.Context, *transfer.DescribeConnectorInput, ...func(*transfer.Options)) (*transfer.DescribeConnectorOutput, error)
	DescribeExecution(context.Context, *transfer.DescribeExecutionInput, ...func(*transfer.Options)) (*transfer.DescribeExecutionOutput, error)
	DescribeHostKey(context.Context, *transfer.DescribeHostKeyInput, ...func(*transfer.Options)) (*transfer.DescribeHostKeyOutput, error)
	DescribeProfile(context.Context, *transfer.DescribeProfileInput, ...func(*transfer.Options)) (*transfer.DescribeProfileOutput, error)
	DescribeSecurityPolicy(context.Context, *transfer.DescribeSecurityPolicyInput, ...func(*transfer.Options)) (*transfer.DescribeSecurityPolicyOutput, error)
	DescribeServer(context.Context, *transfer.DescribeServerInput, ...func(*transfer.Options)) (*transfer.DescribeServerOutput, error)
	DescribeUser(context.Context, *transfer.DescribeUserInput, ...func(*transfer.Options)) (*transfer.DescribeUserOutput, error)
	DescribeWorkflow(context.Context, *transfer.DescribeWorkflowInput, ...func(*transfer.Options)) (*transfer.DescribeWorkflowOutput, error)
	ListAccesses(context.Context, *transfer.ListAccessesInput, ...func(*transfer.Options)) (*transfer.ListAccessesOutput, error)
	ListAgreements(context.Context, *transfer.ListAgreementsInput, ...func(*transfer.Options)) (*transfer.ListAgreementsOutput, error)
	ListCertificates(context.Context, *transfer.ListCertificatesInput, ...func(*transfer.Options)) (*transfer.ListCertificatesOutput, error)
	ListConnectors(context.Context, *transfer.ListConnectorsInput, ...func(*transfer.Options)) (*transfer.ListConnectorsOutput, error)
	ListExecutions(context.Context, *transfer.ListExecutionsInput, ...func(*transfer.Options)) (*transfer.ListExecutionsOutput, error)
	ListHostKeys(context.Context, *transfer.ListHostKeysInput, ...func(*transfer.Options)) (*transfer.ListHostKeysOutput, error)
	ListProfiles(context.Context, *transfer.ListProfilesInput, ...func(*transfer.Options)) (*transfer.ListProfilesOutput, error)
	ListSecurityPolicies(context.Context, *transfer.ListSecurityPoliciesInput, ...func(*transfer.Options)) (*transfer.ListSecurityPoliciesOutput, error)
	ListServers(context.Context, *transfer.ListServersInput, ...func(*transfer.Options)) (*transfer.ListServersOutput, error)
	ListTagsForResource(context.Context, *transfer.ListTagsForResourceInput, ...func(*transfer.Options)) (*transfer.ListTagsForResourceOutput, error)
	ListUsers(context.Context, *transfer.ListUsersInput, ...func(*transfer.Options)) (*transfer.ListUsersOutput, error)
	ListWorkflows(context.Context, *transfer.ListWorkflowsInput, ...func(*transfer.Options)) (*transfer.ListWorkflowsOutput, error)
}
