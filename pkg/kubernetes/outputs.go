package gcp

import (
	"context"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/helmrelease/stack/kubernetes/model"
)

func Outputs(ctx context.Context, input *model.HelmReleaseKubernetesStackInput) (*model.HelmReleaseKubernetesStackOutputs, error) {
	return &model.HelmReleaseKubernetesStackOutputs{}, nil
}

func OutputMapTransformer(stackOutput map[string]interface{}, input *model.HelmReleaseKubernetesStackInput) *model.HelmReleaseKubernetesStackOutputs {
	return &model.HelmReleaseKubernetesStackOutputs{}
}
