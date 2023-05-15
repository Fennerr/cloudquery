// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/support"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/support.go -source=support.go SupportClient
type SupportClient interface {
	DescribeAttachment(context.Context, *support.DescribeAttachmentInput, ...func(*support.Options)) (*support.DescribeAttachmentOutput, error)
	DescribeCases(context.Context, *support.DescribeCasesInput, ...func(*support.Options)) (*support.DescribeCasesOutput, error)
	DescribeCommunications(context.Context, *support.DescribeCommunicationsInput, ...func(*support.Options)) (*support.DescribeCommunicationsOutput, error)
	DescribeServices(context.Context, *support.DescribeServicesInput, ...func(*support.Options)) (*support.DescribeServicesOutput, error)
	DescribeSeverityLevels(context.Context, *support.DescribeSeverityLevelsInput, ...func(*support.Options)) (*support.DescribeSeverityLevelsOutput, error)
	DescribeTrustedAdvisorCheckRefreshStatuses(context.Context, *support.DescribeTrustedAdvisorCheckRefreshStatusesInput, ...func(*support.Options)) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error)
	DescribeTrustedAdvisorCheckResult(context.Context, *support.DescribeTrustedAdvisorCheckResultInput, ...func(*support.Options)) (*support.DescribeTrustedAdvisorCheckResultOutput, error)
	DescribeTrustedAdvisorCheckSummaries(context.Context, *support.DescribeTrustedAdvisorCheckSummariesInput, ...func(*support.Options)) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error)
	DescribeTrustedAdvisorChecks(context.Context, *support.DescribeTrustedAdvisorChecksInput, ...func(*support.Options)) (*support.DescribeTrustedAdvisorChecksOutput, error)
}
