/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Network.
func (mg *Network) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Router),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RouterRef,
		Selector:     mg.Spec.ForProvider.RouterSelector,
		To: reference.To{
			List:    &RouterList{},
			Managed: &Router{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Router")
	}
	mg.Spec.ForProvider.Router = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouterRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Router),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RouterRef,
		Selector:     mg.Spec.InitProvider.RouterSelector,
		To: reference.To{
			List:    &RouterList{},
			Managed: &Router{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Router")
	}
	mg.Spec.InitProvider.Router = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RouterRef = rsp.ResolvedReference

	return nil
}
