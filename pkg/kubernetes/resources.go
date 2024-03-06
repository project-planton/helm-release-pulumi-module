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

	print("added kubernetes provider %#v", kubernetesProvider)

	//// Convert map[string]string to pulumi.Map
	//valuesPulumiMap := pulumi.Map{}
	//for k, v := range s.Input.HelmRelease.Spec.Values {
	//	valuesPulumiMap[k] = pulumi.String(v)
	//}

	_, err = helmv3.NewRelease(ctx, s.Input.HelmRelease.Metadata.Name, &helmv3.ReleaseArgs{
		RepositoryOpts: &helmv3.RepositoryOptsArgs{
			Repo: pulumi.String(s.Input.HelmRelease.Spec.ChartRepo),
		},
		Namespace: pulumi.String(s.Input.HelmRelease.Metadata.Id),
		Chart:     pulumi.String(s.Input.HelmRelease.Spec.ChartName),
		Version:   pulumi.String(s.Input.HelmRelease.Spec.ChartVersion),
		//Values:    pulumi.MapInput(valuesPulumiMap),
	}, pulumi.Provider(kubernetesProvider))

	if err != nil {
		return errors.Wrap(err, "failed to create new helm-release")
	}
	return nil
}
