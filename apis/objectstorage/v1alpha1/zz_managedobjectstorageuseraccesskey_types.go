/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ManagedObjectStorageUserAccessKeyInitParameters struct {

	// (String) Status of the key. Valid values: Active|Inactive
	// Status of the key. Valid values: `Active`|`Inactive`
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ManagedObjectStorageUserAccessKeyObservation struct {

	// (String) Access key id.
	// Access key id.
	AccessKeyID *string `json:"accessKeyId,omitempty" tf:"access_key_id,omitempty"`

	// (String) Creation time.
	// Creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Last used.
	// Last used.
	LastUsedAt *string `json:"lastUsedAt,omitempty" tf:"last_used_at,omitempty"`

	// (String) Managed Object Storage service UUID.
	// Managed Object Storage service UUID.
	ServiceUUID *string `json:"serviceUuid,omitempty" tf:"service_uuid,omitempty"`

	// (String) Status of the key. Valid values: Active|Inactive
	// Status of the key. Valid values: `Active`|`Inactive`
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (String) Username.
	// Username.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ManagedObjectStorageUserAccessKeyParameters struct {

	// (String) Managed Object Storage service UUID.
	// Managed Object Storage service UUID.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/provider-upcloud/apis/objectstorage/v1alpha1.ManagedObjectStorage
	// +kubebuilder:validation:Optional
	ServiceUUID *string `json:"serviceUuid" tf:"service_uuid,omitempty"`

	// Reference to a ManagedObjectStorage in objectstorage to populate serviceUuid.
	// +kubebuilder:validation:Optional
	ServiceUUIDRef *v1.Reference `json:"serviceUuidRef,omitempty" tf:"-"`

	// Selector for a ManagedObjectStorage in objectstorage to populate serviceUuid.
	// +kubebuilder:validation:Optional
	ServiceUUIDSelector *v1.Selector `json:"serviceUuidSelector,omitempty" tf:"-"`

	// (String) Status of the key. Valid values: Active|Inactive
	// Status of the key. Valid values: `Active`|`Inactive`
	// +kubebuilder:validation:Optional
	Status *string `json:"status" tf:"status,omitempty"`

	// (String) Username.
	// Username.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/provider-upcloud/apis/objectstorage/v1alpha1.ManagedObjectStorageUser
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`

	// Reference to a ManagedObjectStorageUser in objectstorage to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a ManagedObjectStorageUser in objectstorage to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

// ManagedObjectStorageUserAccessKeySpec defines the desired state of ManagedObjectStorageUserAccessKey
type ManagedObjectStorageUserAccessKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedObjectStorageUserAccessKeyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ManagedObjectStorageUserAccessKeyInitParameters `json:"initProvider,omitempty"`
}

// ManagedObjectStorageUserAccessKeyStatus defines the observed state of ManagedObjectStorageUserAccessKey.
type ManagedObjectStorageUserAccessKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedObjectStorageUserAccessKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ManagedObjectStorageUserAccessKey is the Schema for the ManagedObjectStorageUserAccessKeys API. This resource represents an UpCloud Managed Object Storage user access key.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upcloud}
type ManagedObjectStorageUserAccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.status) || (has(self.initProvider) && has(self.initProvider.status))",message="spec.forProvider.status is a required parameter"
	Spec   ManagedObjectStorageUserAccessKeySpec   `json:"spec"`
	Status ManagedObjectStorageUserAccessKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedObjectStorageUserAccessKeyList contains a list of ManagedObjectStorageUserAccessKeys
type ManagedObjectStorageUserAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedObjectStorageUserAccessKey `json:"items"`
}

// Repository type metadata.
var (
	ManagedObjectStorageUserAccessKey_Kind             = "ManagedObjectStorageUserAccessKey"
	ManagedObjectStorageUserAccessKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedObjectStorageUserAccessKey_Kind}.String()
	ManagedObjectStorageUserAccessKey_KindAPIVersion   = ManagedObjectStorageUserAccessKey_Kind + "." + CRDGroupVersion.String()
	ManagedObjectStorageUserAccessKey_GroupVersionKind = CRDGroupVersion.WithKind(ManagedObjectStorageUserAccessKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedObjectStorageUserAccessKey{}, &ManagedObjectStorageUserAccessKeyList{})
}
