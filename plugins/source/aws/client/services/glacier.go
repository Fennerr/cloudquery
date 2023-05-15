// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/glacier.go -source=glacier.go GlacierClient
type GlacierClient interface {
	DescribeJob(context.Context, *glacier.DescribeJobInput, ...func(*glacier.Options)) (*glacier.DescribeJobOutput, error)
	DescribeVault(context.Context, *glacier.DescribeVaultInput, ...func(*glacier.Options)) (*glacier.DescribeVaultOutput, error)
	GetDataRetrievalPolicy(context.Context, *glacier.GetDataRetrievalPolicyInput, ...func(*glacier.Options)) (*glacier.GetDataRetrievalPolicyOutput, error)
	GetJobOutput(context.Context, *glacier.GetJobOutputInput, ...func(*glacier.Options)) (*glacier.GetJobOutputOutput, error)
	GetVaultAccessPolicy(context.Context, *glacier.GetVaultAccessPolicyInput, ...func(*glacier.Options)) (*glacier.GetVaultAccessPolicyOutput, error)
	GetVaultLock(context.Context, *glacier.GetVaultLockInput, ...func(*glacier.Options)) (*glacier.GetVaultLockOutput, error)
	GetVaultNotifications(context.Context, *glacier.GetVaultNotificationsInput, ...func(*glacier.Options)) (*glacier.GetVaultNotificationsOutput, error)
	ListJobs(context.Context, *glacier.ListJobsInput, ...func(*glacier.Options)) (*glacier.ListJobsOutput, error)
	ListMultipartUploads(context.Context, *glacier.ListMultipartUploadsInput, ...func(*glacier.Options)) (*glacier.ListMultipartUploadsOutput, error)
	ListParts(context.Context, *glacier.ListPartsInput, ...func(*glacier.Options)) (*glacier.ListPartsOutput, error)
	ListProvisionedCapacity(context.Context, *glacier.ListProvisionedCapacityInput, ...func(*glacier.Options)) (*glacier.ListProvisionedCapacityOutput, error)
	ListTagsForVault(context.Context, *glacier.ListTagsForVaultInput, ...func(*glacier.Options)) (*glacier.ListTagsForVaultOutput, error)
	ListVaults(context.Context, *glacier.ListVaultsInput, ...func(*glacier.Options)) (*glacier.ListVaultsOutput, error)
}
