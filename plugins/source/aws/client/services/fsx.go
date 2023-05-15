// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/fsx.go -source=fsx.go FsxClient
type FsxClient interface {
	DescribeBackups(context.Context, *fsx.DescribeBackupsInput, ...func(*fsx.Options)) (*fsx.DescribeBackupsOutput, error)
	DescribeDataRepositoryAssociations(context.Context, *fsx.DescribeDataRepositoryAssociationsInput, ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryAssociationsOutput, error)
	DescribeDataRepositoryTasks(context.Context, *fsx.DescribeDataRepositoryTasksInput, ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryTasksOutput, error)
	DescribeFileCaches(context.Context, *fsx.DescribeFileCachesInput, ...func(*fsx.Options)) (*fsx.DescribeFileCachesOutput, error)
	DescribeFileSystemAliases(context.Context, *fsx.DescribeFileSystemAliasesInput, ...func(*fsx.Options)) (*fsx.DescribeFileSystemAliasesOutput, error)
	DescribeFileSystems(context.Context, *fsx.DescribeFileSystemsInput, ...func(*fsx.Options)) (*fsx.DescribeFileSystemsOutput, error)
	DescribeSnapshots(context.Context, *fsx.DescribeSnapshotsInput, ...func(*fsx.Options)) (*fsx.DescribeSnapshotsOutput, error)
	DescribeStorageVirtualMachines(context.Context, *fsx.DescribeStorageVirtualMachinesInput, ...func(*fsx.Options)) (*fsx.DescribeStorageVirtualMachinesOutput, error)
	DescribeVolumes(context.Context, *fsx.DescribeVolumesInput, ...func(*fsx.Options)) (*fsx.DescribeVolumesOutput, error)
	ListTagsForResource(context.Context, *fsx.ListTagsForResourceInput, ...func(*fsx.Options)) (*fsx.ListTagsForResourceOutput, error)
}
