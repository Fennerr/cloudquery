package cloudformation

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildStacksWithTemplate(tmpl string) func(t *testing.T, ctrl *gomock.Controller) client.Services {
	return func(t *testing.T, ctrl *gomock.Controller) client.Services {
		mock := mocks.NewMockCloudformationClient(ctrl)

		var stack types.Stack
		if err := faker.FakeObject(&stack); err != nil {
			t.Fatal(err)
		}
		mock.EXPECT().DescribeStacks(
			gomock.Any(),
			&cloudformation.DescribeStacksInput{},
			gomock.Any(),
		).Return(
			&cloudformation.DescribeStacksOutput{Stacks: []types.Stack{stack}},
			nil,
		)

		var resource types.StackResourceSummary
		if err := faker.FakeObject(&resource); err != nil {
			t.Fatal(err)
		}
		mock.EXPECT().ListStackResources(
			gomock.Any(),
			&cloudformation.ListStackResourcesInput{StackName: stack.StackName},
			gomock.Any(),
		).Return(
			&cloudformation.ListStackResourcesOutput{StackResourceSummaries: []types.StackResourceSummary{resource}},
			nil,
		)

		var summary cloudformation.GetTemplateSummaryOutput
		if err := faker.FakeObject(&summary); err != nil {
			t.Fatal(err)
		}
		summary.Metadata = aws.String(`{ "some": "metadata" }`) // Required as faker doesn't handle this.

		mock.EXPECT().GetTemplateSummary(
			gomock.Any(),
			&cloudformation.GetTemplateSummaryInput{StackName: stack.StackName},
			gomock.Any(),
		).Return(
			&summary,
			nil,
		)

		var template cloudformation.GetTemplateOutput
		if err := faker.FakeObject(&template); err != nil {
			t.Fatal(err)
		}
		template.TemplateBody = aws.String(tmpl) // Required as faker doesn't handle this.

		mock.EXPECT().GetTemplate(
			gomock.Any(),
			&cloudformation.GetTemplateInput{StackName: stack.StackName},
			gomock.Any(),
		).Return(
			&template,
			nil,
		)

		return client.Services{Cloudformation: mock}
	}
}

func TestCloudformationStacksWithYAML(t *testing.T) {
	// check that cloudformation templates are parsed correctly for both YAML and JSON
	client.AwsMockTestHelper(t, Stacks(), buildStacksWithTemplate(`---
AWSTemplateFormatVersion: "version date"`), client.TestOptions{})
}

func TestCloudformationStacksWithJSON(t *testing.T) {
	client.AwsMockTestHelper(t, Stacks(), buildStacksWithTemplate(`{"test": "key"}`), client.TestOptions{})
}
