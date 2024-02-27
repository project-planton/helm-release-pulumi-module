package gcp

import (
	"context"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/helmrelease/stack/kubernetes/model"
)

func Outputs(ctx context.Context, input *model.HelmReleaseKubernetesStackInput) (*model.HelmReleaseKubernetesStackOutputs, error) {
	return &model.HelmReleaseKubernetesStackOutputs{}, nil
}

func OutputMapTransformer(stackOutput map[string]interface{}, input *model.HelmReleaseKubernetesStackInput) *model.HelmReleaseKubernetesStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply || stackOutput == nil {
		return &model.HelmReleaseKubernetesStackOutputs{}
	}
	return &model.HelmReleaseKubernetesStackOutputs{}
}
