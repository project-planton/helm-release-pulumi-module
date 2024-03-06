package gcp

import (
	"github.com/pkg/errors"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/helmrelease/stack/kubernetes/model"
	pulumikubernetesprovider "github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/automation/provider/kubernetes"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	Input *model.HelmReleaseKubernetesStackInput
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	kubernetesProvider, err := pulumikubernetesprovider.GetWithStackCredentials(ctx, s.Input.CredentialsInput.Kubernetes)
	if err != nil {
		return errors.Wrap(err, "failed to setup kubernetes provider")
	}

	//// Convert map[string]string to pulumi.Map
	//valuesPulumiMap := pulumi.Map{}
	//for k, v := range s.Input.HelmRelease.Spec.Values {
	//	valuesPulumiMap[k] = pulumi.String(v)
	//}

	if _, err := helmv3.NewRelease(ctx, s.Input.HelmRelease.Metadata.Name, &helmv3.ReleaseArgs{
		Name:      pulumi.String(s.Input.HelmRelease.Metadata.Name),
		Namespace: pulumi.String(s.Input.HelmRelease.Metadata.Id),
		RepositoryOpts: &helmv3.RepositoryOptsArgs{
			Repo: pulumi.String(s.Input.HelmRelease.Spec.ChartRepo),
		},
		Chart:           pulumi.String(s.Input.HelmRelease.Spec.ChartName),
		Version:         pulumi.String(s.Input.HelmRelease.Spec.ChartVersion),
		CreateNamespace: pulumi.Bool(true),
		CleanupOnFail:   pulumi.Bool(false),
		ForceUpdate:     pulumi.Bool(true),
		AllowNullValues: pulumi.Bool(true),
		Atomic:          pulumi.Bool(false),
		Values:          pulumi.Map{},
	}, pulumi.Provider(kubernetesProvider)); err != nil {
		return errors.Wrap(err, "failed to create new helm-release")
	}
	return nil
}
