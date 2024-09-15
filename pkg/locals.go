package pkg

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/helmrelease"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcekind"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	Labels map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *helmrelease.HelmReleaseStackInput) *Locals {
	locals := &Locals{}

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Environment:  stackInput.Target.Spec.EnvironmentInfo.EnvId,
		kuberneteslabelkeys.Organization: stackInput.Target.Spec.EnvironmentInfo.OrgId,
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceId:   stackInput.Target.Metadata.Id,
		kuberneteslabelkeys.ResourceKind: apiresourcekind.ApiResourceKind_helm_release.String(),
	}
	return locals
}
