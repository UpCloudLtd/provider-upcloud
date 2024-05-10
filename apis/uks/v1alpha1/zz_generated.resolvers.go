/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/UpCloudLtd/provider-upcloud/apis/network/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this KubernetesCluster.
func (mg *KubernetesCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkRef,
		Selector:     mg.Spec.ForProvider.NetworkSelector,
		To: reference.To{
			List:    &v1alpha1.NetworkList{},
			Managed: &v1alpha1.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Network")
	}
	mg.Spec.ForProvider.Network = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Network),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkRef,
		Selector:     mg.Spec.InitProvider.NetworkSelector,
		To: reference.To{
			List:    &v1alpha1.NetworkList{},
			Managed: &v1alpha1.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Network")
	}
	mg.Spec.InitProvider.Network = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KubernetesNodeGroup.
func (mg *KubernetesNodeGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Cluster),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterRef,
		Selector:     mg.Spec.ForProvider.ClusterSelector,
		To: reference.To{
			List:    &KubernetesClusterList{},
			Managed: &KubernetesCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Cluster")
	}
	mg.Spec.ForProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterRef = rsp.ResolvedReference

	return nil
}