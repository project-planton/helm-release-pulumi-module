package gcp

import (
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/helmrelease/stack/kubernetes/model"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	Input *model.HelmReleaseKubernetesStackInput
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	return nil
}
