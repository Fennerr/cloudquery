// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/efs.go -source=efs.go EfsClient
type EfsClient interface {
	DescribeAccessPoints(context.Context, *efs.DescribeAccessPointsInput, ...func(*efs.Options)) (*efs.DescribeAccessPointsOutput, error)
	DescribeAccountPreferences(context.Context, *efs.DescribeAccountPreferencesInput, ...func(*efs.Options)) (*efs.DescribeAccountPreferencesOutput, error)
	DescribeBackupPolicy(context.Context, *efs.DescribeBackupPolicyInput, ...func(*efs.Options)) (*efs.DescribeBackupPolicyOutput, error)
	DescribeFileSystemPolicy(context.Context, *efs.DescribeFileSystemPolicyInput, ...func(*efs.Options)) (*efs.DescribeFileSystemPolicyOutput, error)
	DescribeFileSystems(context.Context, *efs.DescribeFileSystemsInput, ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error)
	DescribeLifecycleConfiguration(context.Context, *efs.DescribeLifecycleConfigurationInput, ...func(*efs.Options)) (*efs.DescribeLifecycleConfigurationOutput, error)
	DescribeMountTargetSecurityGroups(context.Context, *efs.DescribeMountTargetSecurityGroupsInput, ...func(*efs.Options)) (*efs.DescribeMountTargetSecurityGroupsOutput, error)
	DescribeMountTargets(context.Context, *efs.DescribeMountTargetsInput, ...func(*efs.Options)) (*efs.DescribeMountTargetsOutput, error)
	DescribeReplicationConfigurations(context.Context, *efs.DescribeReplicationConfigurationsInput, ...func(*efs.Options)) (*efs.DescribeReplicationConfigurationsOutput, error)
	DescribeTags(context.Context, *efs.DescribeTagsInput, ...func(*efs.Options)) (*efs.DescribeTagsOutput, error)
	ListTagsForResource(context.Context, *efs.ListTagsForResourceInput, ...func(*efs.Options)) (*efs.ListTagsForResourceOutput, error)
}
