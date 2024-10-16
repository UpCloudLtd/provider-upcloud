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

type KubernetesClusterInitParameters struct {

	// (Set of String) IP addresses or IP ranges in CIDR format which are allowed to access the cluster control plane. To allow access from any source, use ["0.0.0.0/0"]. To deny access from all sources, use []. Values set here do not restrict access to node groups or exposed Kubernetes services.
	// IP addresses or IP ranges in CIDR format which are allowed to access the cluster control plane. To allow access from any source, use `["0.0.0.0/0"]`. To deny access from all sources, use `[]`. Values set here do not restrict access to node groups or exposed Kubernetes services.
	// +listType=set
	ControlPlaneIPFilter []*string `json:"controlPlaneIpFilter,omitempty" tf:"control_plane_ip_filter,omitempty"`

	// value pairs to classify the cluster.
	// User defined key-value pairs to classify the cluster.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) Cluster name. Needs to be unique within the account.
	// Cluster name. Needs to be unique within the account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Network ID for the cluster to run in.
	// Network ID for the cluster to run in.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/crossplane-provider-upcloud/apis/network/v1alpha1.Network
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// (String) The pricing plan used for the cluster. Default plan is development. You can list available plans with upctl kubernetes plans.
	// The pricing plan used for the cluster. Default plan is `development`. You can list available plans with `upctl kubernetes plans`.
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// (Boolean) Enable private node groups. Private node groups requires a network that is routed through NAT gateway.
	// Enable private node groups. Private node groups requires a network that is routed through NAT gateway.
	PrivateNodeGroups *bool `json:"privateNodeGroups,omitempty" tf:"private_node_groups,omitempty"`

	// (String) Set default storage encryption strategy for all nodes in the cluster.
	// Set default storage encryption strategy for all nodes in the cluster.
	StorageEncryption *string `json:"storageEncryption,omitempty" tf:"storage_encryption,omitempty"`

	// (String) Kubernetes version ID, e.g. 1.29. You can list available version IDs with upctl kubernetes versions.
	// Kubernetes version ID, e.g. `1.29`. You can list available version IDs with `upctl kubernetes versions`.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// fra1. You can list available zones with upctl zone list.
	// Zone in which the Kubernetes cluster will be hosted, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type KubernetesClusterObservation struct {

	// (Set of String) IP addresses or IP ranges in CIDR format which are allowed to access the cluster control plane. To allow access from any source, use ["0.0.0.0/0"]. To deny access from all sources, use []. Values set here do not restrict access to node groups or exposed Kubernetes services.
	// IP addresses or IP ranges in CIDR format which are allowed to access the cluster control plane. To allow access from any source, use `["0.0.0.0/0"]`. To deny access from all sources, use `[]`. Values set here do not restrict access to node groups or exposed Kubernetes services.
	// +listType=set
	ControlPlaneIPFilter []*string `json:"controlPlaneIpFilter,omitempty" tf:"control_plane_ip_filter,omitempty"`

	// (String) UUID of the cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// value pairs to classify the cluster.
	// User defined key-value pairs to classify the cluster.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) Cluster name. Needs to be unique within the account.
	// Cluster name. Needs to be unique within the account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Network ID for the cluster to run in.
	// Network ID for the cluster to run in.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// (String) Network CIDR for the given network. Computed automatically.
	// Network CIDR for the given network. Computed automatically.
	NetworkCidr *string `json:"networkCidr,omitempty" tf:"network_cidr,omitempty"`

	// (List of String) Names of the node groups configured to cluster
	// Names of the node groups configured to cluster
	NodeGroups []*string `json:"nodeGroups,omitempty" tf:"node_groups,omitempty"`

	// (String) The pricing plan used for the cluster. Default plan is development. You can list available plans with upctl kubernetes plans.
	// The pricing plan used for the cluster. Default plan is `development`. You can list available plans with `upctl kubernetes plans`.
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// (Boolean) Enable private node groups. Private node groups requires a network that is routed through NAT gateway.
	// Enable private node groups. Private node groups requires a network that is routed through NAT gateway.
	PrivateNodeGroups *bool `json:"privateNodeGroups,omitempty" tf:"private_node_groups,omitempty"`

	// (String) Operational state of the cluster.
	// Operational state of the cluster.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// (String) Set default storage encryption strategy for all nodes in the cluster.
	// Set default storage encryption strategy for all nodes in the cluster.
	StorageEncryption *string `json:"storageEncryption,omitempty" tf:"storage_encryption,omitempty"`

	// (String) Kubernetes version ID, e.g. 1.29. You can list available version IDs with upctl kubernetes versions.
	// Kubernetes version ID, e.g. `1.29`. You can list available version IDs with `upctl kubernetes versions`.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// fra1. You can list available zones with upctl zone list.
	// Zone in which the Kubernetes cluster will be hosted, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type KubernetesClusterParameters struct {

	// (Set of String) IP addresses or IP ranges in CIDR format which are allowed to access the cluster control plane. To allow access from any source, use ["0.0.0.0/0"]. To deny access from all sources, use []. Values set here do not restrict access to node groups or exposed Kubernetes services.
	// IP addresses or IP ranges in CIDR format which are allowed to access the cluster control plane. To allow access from any source, use `["0.0.0.0/0"]`. To deny access from all sources, use `[]`. Values set here do not restrict access to node groups or exposed Kubernetes services.
	// +kubebuilder:validation:Optional
	// +listType=set
	ControlPlaneIPFilter []*string `json:"controlPlaneIpFilter" tf:"control_plane_ip_filter,omitempty"`

	// value pairs to classify the cluster.
	// User defined key-value pairs to classify the cluster.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels" tf:"labels,omitempty"`

	// (String) Cluster name. Needs to be unique within the account.
	// Cluster name. Needs to be unique within the account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Network ID for the cluster to run in.
	// Network ID for the cluster to run in.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/crossplane-provider-upcloud/apis/network/v1alpha1.Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network" tf:"network,omitempty"`

	// Reference to a Network in network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// (String) The pricing plan used for the cluster. Default plan is development. You can list available plans with upctl kubernetes plans.
	// The pricing plan used for the cluster. Default plan is `development`. You can list available plans with `upctl kubernetes plans`.
	// +kubebuilder:validation:Optional
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// (Boolean) Enable private node groups. Private node groups requires a network that is routed through NAT gateway.
	// Enable private node groups. Private node groups requires a network that is routed through NAT gateway.
	// +kubebuilder:validation:Optional
	PrivateNodeGroups *bool `json:"privateNodeGroups,omitempty" tf:"private_node_groups,omitempty"`

	// (String) Set default storage encryption strategy for all nodes in the cluster.
	// Set default storage encryption strategy for all nodes in the cluster.
	// +kubebuilder:validation:Optional
	StorageEncryption *string `json:"storageEncryption,omitempty" tf:"storage_encryption,omitempty"`

	// (String) Kubernetes version ID, e.g. 1.29. You can list available version IDs with upctl kubernetes versions.
	// Kubernetes version ID, e.g. `1.29`. You can list available version IDs with `upctl kubernetes versions`.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// fra1. You can list available zones with upctl zone list.
	// Zone in which the Kubernetes cluster will be hosted, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

// KubernetesClusterSpec defines the desired state of KubernetesCluster
type KubernetesClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KubernetesClusterParameters `json:"forProvider"`
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
	InitProvider KubernetesClusterInitParameters `json:"initProvider,omitempty"`
}

// KubernetesClusterStatus defines the observed state of KubernetesCluster.
type KubernetesClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KubernetesClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KubernetesCluster is the Schema for the KubernetesClusters API. This resource represents a Managed Kubernetes https://upcloud.com/products/managed-kubernetes cluster.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upcloud}
type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.controlPlaneIpFilter) || (has(self.initProvider) && has(self.initProvider.controlPlaneIpFilter))",message="spec.forProvider.controlPlaneIpFilter is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.labels) || (has(self.initProvider) && has(self.initProvider.labels))",message="spec.forProvider.labels is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   KubernetesClusterSpec   `json:"spec"`
	Status KubernetesClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesClusterList contains a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesCluster `json:"items"`
}

// Repository type metadata.
var (
	KubernetesCluster_Kind             = "KubernetesCluster"
	KubernetesCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KubernetesCluster_Kind}.String()
	KubernetesCluster_KindAPIVersion   = KubernetesCluster_Kind + "." + CRDGroupVersion.String()
	KubernetesCluster_GroupVersionKind = CRDGroupVersion.WithKind(KubernetesCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&KubernetesCluster{}, &KubernetesClusterList{})
}
