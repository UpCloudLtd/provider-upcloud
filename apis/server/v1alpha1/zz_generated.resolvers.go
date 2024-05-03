/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/UpCloudLtd/provider-upcloud/apis/network/v1alpha1"
	v1alpha11 "github.com/UpCloudLtd/provider-upcloud/apis/storage/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this FirewallRules.
func (mg *FirewallRules) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerIDRef,
		Selector:     mg.Spec.ForProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServerIDRef,
		Selector:     mg.Spec.InitProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerID")
	}
	mg.Spec.InitProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkInterface); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkInterface[i3].Network),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NetworkInterface[i3].NetworkRef,
			Selector:     mg.Spec.ForProvider.NetworkInterface[i3].NetworkSelector,
			To: reference.To{
				List:    &v1alpha1.NetworkList{},
				Managed: &v1alpha1.Network{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterface[i3].Network")
		}
		mg.Spec.ForProvider.NetworkInterface[i3].Network = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkInterface[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageDevices); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageDevices[i3].Storage),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.StorageDevices[i3].StorageRef,
			Selector:     mg.Spec.ForProvider.StorageDevices[i3].StorageSelector,
			To: reference.To{
				List:    &v1alpha11.StorageList{},
				Managed: &v1alpha11.Storage{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageDevices[i3].Storage")
		}
		mg.Spec.ForProvider.StorageDevices[i3].Storage = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageDevices[i3].StorageRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.NetworkInterface); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkInterface[i3].Network),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NetworkInterface[i3].NetworkRef,
			Selector:     mg.Spec.InitProvider.NetworkInterface[i3].NetworkSelector,
			To: reference.To{
				List:    &v1alpha1.NetworkList{},
				Managed: &v1alpha1.Network{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.NetworkInterface[i3].Network")
		}
		mg.Spec.InitProvider.NetworkInterface[i3].Network = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.NetworkInterface[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageDevices); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageDevices[i3].Storage),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.StorageDevices[i3].StorageRef,
			Selector:     mg.Spec.InitProvider.StorageDevices[i3].StorageSelector,
			To: reference.To{
				List:    &v1alpha11.StorageList{},
				Managed: &v1alpha11.Storage{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageDevices[i3].Storage")
		}
		mg.Spec.InitProvider.StorageDevices[i3].Storage = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageDevices[i3].StorageRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ServerGroup.
func (mg *ServerGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Members),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.MembersRefs,
		Selector:      mg.Spec.ForProvider.MembersSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Members")
	}
	mg.Spec.ForProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.MembersRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Members),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.MembersRefs,
		Selector:      mg.Spec.InitProvider.MembersSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Members")
	}
	mg.Spec.InitProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.MembersRefs = mrsp.ResolvedReferences

	return nil
}
