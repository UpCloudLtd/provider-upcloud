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

type ManagedObjectStoragePolicyInitParameters struct {

	// (String) Description of the policy.
	// Description of the policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// encoded compliant with RFC 3986.
	// Policy document, URL-encoded compliant with RFC 3986.
	Document *string `json:"document,omitempty" tf:"document,omitempty"`
}

type ManagedObjectStoragePolicyObservation struct {

	// (String) Policy ARN.
	// Policy ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// (Number) Attachment count.
	// Attachment count.
	AttachmentCount *float64 `json:"attachmentCount,omitempty" tf:"attachment_count,omitempty"`

	// (String) Creation time.
	// Creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) Default version id.
	// Default version id.
	DefaultVersionID *string `json:"defaultVersionId,omitempty" tf:"default_version_id,omitempty"`

	// (String) Description of the policy.
	// Description of the policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// encoded compliant with RFC 3986.
	// Policy document, URL-encoded compliant with RFC 3986.
	Document *string `json:"document,omitempty" tf:"document,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Managed Object Storage service UUID.
	// Managed Object Storage service UUID.
	ServiceUUID *string `json:"serviceUuid,omitempty" tf:"service_uuid,omitempty"`

	// (Boolean) Defines whether the policy was set up by the system.
	// Defines whether the policy was set up by the system.
	System *bool `json:"system,omitempty" tf:"system,omitempty"`

	// (String) Update time.
	// Update time.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ManagedObjectStoragePolicyParameters struct {

	// (String) Description of the policy.
	// Description of the policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// encoded compliant with RFC 3986.
	// Policy document, URL-encoded compliant with RFC 3986.
	// +kubebuilder:validation:Optional
	Document *string `json:"document,omitempty" tf:"document,omitempty"`

	// (String) Managed Object Storage service UUID.
	// Managed Object Storage service UUID.
	// +crossplane:generate:reference:type=ManagedObjectStorage
	// +kubebuilder:validation:Optional
	ServiceUUID *string `json:"serviceUuid,omitempty" tf:"service_uuid,omitempty"`

	// Reference to a ManagedObjectStorage to populate serviceUuid.
	// +kubebuilder:validation:Optional
	ServiceUUIDRef *v1.Reference `json:"serviceUuidRef,omitempty" tf:"-"`

	// Selector for a ManagedObjectStorage to populate serviceUuid.
	// +kubebuilder:validation:Optional
	ServiceUUIDSelector *v1.Selector `json:"serviceUuidSelector,omitempty" tf:"-"`
}

// ManagedObjectStoragePolicySpec defines the desired state of ManagedObjectStoragePolicy
type ManagedObjectStoragePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedObjectStoragePolicyParameters `json:"forProvider"`
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
	InitProvider ManagedObjectStoragePolicyInitParameters `json:"initProvider,omitempty"`
}

// ManagedObjectStoragePolicyStatus defines the observed state of ManagedObjectStoragePolicy.
type ManagedObjectStoragePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedObjectStoragePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ManagedObjectStoragePolicy is the Schema for the ManagedObjectStoragePolicys API. This resource represents an UpCloud Managed Object Storage policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upcloud}
type ManagedObjectStoragePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.document) || (has(self.initProvider) && has(self.initProvider.document))",message="spec.forProvider.document is a required parameter"
	Spec   ManagedObjectStoragePolicySpec   `json:"spec"`
	Status ManagedObjectStoragePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedObjectStoragePolicyList contains a list of ManagedObjectStoragePolicys
type ManagedObjectStoragePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedObjectStoragePolicy `json:"items"`
}

// Repository type metadata.
var (
	ManagedObjectStoragePolicy_Kind             = "ManagedObjectStoragePolicy"
	ManagedObjectStoragePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedObjectStoragePolicy_Kind}.String()
	ManagedObjectStoragePolicy_KindAPIVersion   = ManagedObjectStoragePolicy_Kind + "." + CRDGroupVersion.String()
	ManagedObjectStoragePolicy_GroupVersionKind = CRDGroupVersion.WithKind(ManagedObjectStoragePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedObjectStoragePolicy{}, &ManagedObjectStoragePolicyList{})
}