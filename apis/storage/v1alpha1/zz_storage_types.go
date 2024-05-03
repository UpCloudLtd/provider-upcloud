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

type BackupRuleInitParameters struct {

	// (String) The weekday when the backup is created
	// The weekday when the backup is created
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// (Number) The number of days before a backup is automatically deleted
	// The number of days before a backup is automatically deleted
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// (String) The time of day when the backup is created
	// The time of day when the backup is created
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type BackupRuleObservation struct {

	// (String) The weekday when the backup is created
	// The weekday when the backup is created
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// (Number) The number of days before a backup is automatically deleted
	// The number of days before a backup is automatically deleted
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// (String) The time of day when the backup is created
	// The time of day when the backup is created
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type BackupRuleParameters struct {

	// (String) The weekday when the backup is created
	// The weekday when the backup is created
	// +kubebuilder:validation:Optional
	Interval *string `json:"interval" tf:"interval,omitempty"`

	// (Number) The number of days before a backup is automatically deleted
	// The number of days before a backup is automatically deleted
	// +kubebuilder:validation:Optional
	Retention *float64 `json:"retention" tf:"retention,omitempty"`

	// (String) The time of day when the backup is created
	// The time of day when the backup is created
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`
}

type CloneInitParameters struct {

	// (String) The ID of this resource.
	// The unique identifier of the storage/template to clone
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CloneObservation struct {

	// (String) The ID of this resource.
	// The unique identifier of the storage/template to clone
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CloneParameters struct {

	// (String) The ID of this resource.
	// The unique identifier of the storage/template to clone
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`
}

type ImportInitParameters struct {

	// (String) The mode of the import task. One of http_import or direct_upload.
	// The mode of the import task. One of `http_import` or `direct_upload`.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) For direct_upload; an optional hash of the file to upload.
	// For `direct_upload`; an optional hash of the file to upload.
	SourceHash *string `json:"sourceHash,omitempty" tf:"source_hash,omitempty"`

	// (String) The location of the file to import. For http_import an accessible URL for direct_upload a local file.
	// The location of the file to import. For `http_import` an accessible URL for `direct_upload` a local file.
	SourceLocation *string `json:"sourceLocation,omitempty" tf:"source_location,omitempty"`
}

type ImportObservation struct {

	// (String) sha256 sum of the imported data
	// sha256 sum of the imported data
	Sha256Sum *string `json:"sha256sum,omitempty" tf:"sha256sum,omitempty"`

	// (String) The mode of the import task. One of http_import or direct_upload.
	// The mode of the import task. One of `http_import` or `direct_upload`.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) For direct_upload; an optional hash of the file to upload.
	// For `direct_upload`; an optional hash of the file to upload.
	SourceHash *string `json:"sourceHash,omitempty" tf:"source_hash,omitempty"`

	// (String) The location of the file to import. For http_import an accessible URL for direct_upload a local file.
	// The location of the file to import. For `http_import` an accessible URL for `direct_upload` a local file.
	SourceLocation *string `json:"sourceLocation,omitempty" tf:"source_location,omitempty"`

	// (Number) Number of bytes imported
	// Number of bytes imported
	WrittenBytes *float64 `json:"writtenBytes,omitempty" tf:"written_bytes,omitempty"`
}

type ImportParameters struct {

	// (String) The mode of the import task. One of http_import or direct_upload.
	// The mode of the import task. One of `http_import` or `direct_upload`.
	// +kubebuilder:validation:Optional
	Source *string `json:"source" tf:"source,omitempty"`

	// (String) For direct_upload; an optional hash of the file to upload.
	// For `direct_upload`; an optional hash of the file to upload.
	// +kubebuilder:validation:Optional
	SourceHash *string `json:"sourceHash,omitempty" tf:"source_hash,omitempty"`

	// (String) The location of the file to import. For http_import an accessible URL for direct_upload a local file.
	// The location of the file to import. For `http_import` an accessible URL for `direct_upload` a local file.
	// +kubebuilder:validation:Optional
	SourceLocation *string `json:"sourceLocation" tf:"source_location,omitempty"`
}

type StorageInitParameters struct {

	// (see below for nested schema)
	// The criteria to backup the storage
	// Please keep in mind that it's not possible to have a server with backup_rule attached to a server with simple_backup specified.
	// Such configurations will throw errors during execution.
	BackupRule []BackupRuleInitParameters `json:"backupRule,omitempty" tf:"backup_rule,omitempty"`

	// (Block Set, Max: 1) Block defining another storage/template to clone to storage (see below for nested schema)
	// Block defining another storage/template to clone to storage
	Clone []CloneInitParameters `json:"clone,omitempty" tf:"clone,omitempty"`

	// (Boolean) If set to true, the backup taken before the partition and filesystem resize attempt will be deleted immediately after success.
	// If set to true, the backup taken before the partition and filesystem resize attempt will be deleted immediately after success.
	DeleteAutoresizeBackup *bool `json:"deleteAutoresizeBackup,omitempty" tf:"delete_autoresize_backup,omitempty"`

	// (Boolean) Sets if the storage is encrypted at rest
	// Sets if the storage is encrypted at rest
	Encrypt *bool `json:"encrypt,omitempty" tf:"encrypt,omitempty"`

	// (Boolean) If set to true, provider will attempt to resize partition and filesystem when the size of the storage changes.
	// Please note that before the resize attempt is made, backup of the storage will be taken. If the resize attempt fails, the backup will be used
	// to restore the storage and then deleted. If the resize attempt succeeds, backup will be kept (unless delete_autoresize_backup option is set to true).
	// Taking and keeping backups incure costs.
	// If set to true, provider will attempt to resize partition and filesystem when the size of the storage changes.
	// Please note that before the resize attempt is made, backup of the storage will be taken. If the resize attempt fails, the backup will be used
	// to restore the storage and then deleted. If the resize attempt succeeds, backup will be kept (unless delete_autoresize_backup option is set to true).
	// Taking and keeping backups incure costs.
	FilesystemAutoresize *bool `json:"filesystemAutoresize,omitempty" tf:"filesystem_autoresize,omitempty"`

	// (Block Set, Max: 1) Block defining external data to import to storage (see below for nested schema)
	// Block defining external data to import to storage
	Import []ImportInitParameters `json:"import,omitempty" tf:"import,omitempty"`

	// (Number) The size of the storage in gigabytes
	// The size of the storage in gigabytes
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// (String) The storage tier to use
	// The storage tier to use
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// (String) A short, informative description
	// A short, informative description
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// fra1. You can list available zones with upctl zone list.
	// The zone in which the storage will be created, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type StorageObservation struct {

	// (see below for nested schema)
	// The criteria to backup the storage
	// Please keep in mind that it's not possible to have a server with backup_rule attached to a server with simple_backup specified.
	// Such configurations will throw errors during execution.
	BackupRule []BackupRuleObservation `json:"backupRule,omitempty" tf:"backup_rule,omitempty"`

	// (Block Set, Max: 1) Block defining another storage/template to clone to storage (see below for nested schema)
	// Block defining another storage/template to clone to storage
	Clone []CloneObservation `json:"clone,omitempty" tf:"clone,omitempty"`

	// (Boolean) If set to true, the backup taken before the partition and filesystem resize attempt will be deleted immediately after success.
	// If set to true, the backup taken before the partition and filesystem resize attempt will be deleted immediately after success.
	DeleteAutoresizeBackup *bool `json:"deleteAutoresizeBackup,omitempty" tf:"delete_autoresize_backup,omitempty"`

	// (Boolean) Sets if the storage is encrypted at rest
	// Sets if the storage is encrypted at rest
	Encrypt *bool `json:"encrypt,omitempty" tf:"encrypt,omitempty"`

	// (Boolean) If set to true, provider will attempt to resize partition and filesystem when the size of the storage changes.
	// Please note that before the resize attempt is made, backup of the storage will be taken. If the resize attempt fails, the backup will be used
	// to restore the storage and then deleted. If the resize attempt succeeds, backup will be kept (unless delete_autoresize_backup option is set to true).
	// Taking and keeping backups incure costs.
	// If set to true, provider will attempt to resize partition and filesystem when the size of the storage changes.
	// Please note that before the resize attempt is made, backup of the storage will be taken. If the resize attempt fails, the backup will be used
	// to restore the storage and then deleted. If the resize attempt succeeds, backup will be kept (unless delete_autoresize_backup option is set to true).
	// Taking and keeping backups incure costs.
	FilesystemAutoresize *bool `json:"filesystemAutoresize,omitempty" tf:"filesystem_autoresize,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block Set, Max: 1) Block defining external data to import to storage (see below for nested schema)
	// Block defining external data to import to storage
	Import []ImportObservation `json:"import,omitempty" tf:"import,omitempty"`

	// (Number) The size of the storage in gigabytes
	// The size of the storage in gigabytes
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// (String) The storage tier to use
	// The storage tier to use
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// (String) A short, informative description
	// A short, informative description
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// fra1. You can list available zones with upctl zone list.
	// The zone in which the storage will be created, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type StorageParameters struct {

	// (see below for nested schema)
	// The criteria to backup the storage
	// Please keep in mind that it's not possible to have a server with backup_rule attached to a server with simple_backup specified.
	// Such configurations will throw errors during execution.
	// +kubebuilder:validation:Optional
	BackupRule []BackupRuleParameters `json:"backupRule,omitempty" tf:"backup_rule,omitempty"`

	// (Block Set, Max: 1) Block defining another storage/template to clone to storage (see below for nested schema)
	// Block defining another storage/template to clone to storage
	// +kubebuilder:validation:Optional
	Clone []CloneParameters `json:"clone,omitempty" tf:"clone,omitempty"`

	// (Boolean) If set to true, the backup taken before the partition and filesystem resize attempt will be deleted immediately after success.
	// If set to true, the backup taken before the partition and filesystem resize attempt will be deleted immediately after success.
	// +kubebuilder:validation:Optional
	DeleteAutoresizeBackup *bool `json:"deleteAutoresizeBackup,omitempty" tf:"delete_autoresize_backup,omitempty"`

	// (Boolean) Sets if the storage is encrypted at rest
	// Sets if the storage is encrypted at rest
	// +kubebuilder:validation:Optional
	Encrypt *bool `json:"encrypt,omitempty" tf:"encrypt,omitempty"`

	// (Boolean) If set to true, provider will attempt to resize partition and filesystem when the size of the storage changes.
	// Please note that before the resize attempt is made, backup of the storage will be taken. If the resize attempt fails, the backup will be used
	// to restore the storage and then deleted. If the resize attempt succeeds, backup will be kept (unless delete_autoresize_backup option is set to true).
	// Taking and keeping backups incure costs.
	// If set to true, provider will attempt to resize partition and filesystem when the size of the storage changes.
	// Please note that before the resize attempt is made, backup of the storage will be taken. If the resize attempt fails, the backup will be used
	// to restore the storage and then deleted. If the resize attempt succeeds, backup will be kept (unless delete_autoresize_backup option is set to true).
	// Taking and keeping backups incure costs.
	// +kubebuilder:validation:Optional
	FilesystemAutoresize *bool `json:"filesystemAutoresize,omitempty" tf:"filesystem_autoresize,omitempty"`

	// (Block Set, Max: 1) Block defining external data to import to storage (see below for nested schema)
	// Block defining external data to import to storage
	// +kubebuilder:validation:Optional
	Import []ImportParameters `json:"import,omitempty" tf:"import,omitempty"`

	// (Number) The size of the storage in gigabytes
	// The size of the storage in gigabytes
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// (String) The storage tier to use
	// The storage tier to use
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// (String) A short, informative description
	// A short, informative description
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// fra1. You can list available zones with upctl zone list.
	// The zone in which the storage will be created, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// StorageSpec defines the desired state of Storage
type StorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageParameters `json:"forProvider"`
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
	InitProvider StorageInitParameters `json:"initProvider,omitempty"`
}

// StorageStatus defines the observed state of Storage.
type StorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Storage is the Schema for the Storages API. Manages UpCloud Block Storage https://upcloud.com/products/block-storage devices.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upcloud}
type Storage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.title) || (has(self.initProvider) && has(self.initProvider.title))",message="spec.forProvider.title is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   StorageSpec   `json:"spec"`
	Status StorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageList contains a list of Storages
type StorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Storage `json:"items"`
}

// Repository type metadata.
var (
	Storage_Kind             = "Storage"
	Storage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Storage_Kind}.String()
	Storage_KindAPIVersion   = Storage_Kind + "." + CRDGroupVersion.String()
	Storage_GroupVersionKind = CRDGroupVersion.WithKind(Storage_Kind)
)

func init() {
	SchemeBuilder.Register(&Storage{}, &StorageList{})
}
