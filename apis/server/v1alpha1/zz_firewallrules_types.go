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

type FirewallRuleInitParameters struct {

	// (String) Action to take if the rule conditions are met
	// Action to take if the rule conditions are met
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// (String) Freeform comment string for the rule
	// Freeform comment string for the rule
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (String) The destination address range ends from this address
	// The destination address range ends from this address
	DestinationAddressEnd *string `json:"destinationAddressEnd,omitempty" tf:"destination_address_end,omitempty"`

	// (String) The destination address range starts from this address
	// The destination address range starts from this address
	DestinationAddressStart *string `json:"destinationAddressStart,omitempty" tf:"destination_address_start,omitempty"`

	// (String) The destination port range ends from this port number
	// The destination port range ends from this port number
	DestinationPortEnd *string `json:"destinationPortEnd,omitempty" tf:"destination_port_end,omitempty"`

	// (String) The destination port range starts from this port number
	// The destination port range starts from this port number
	DestinationPortStart *string `json:"destinationPortStart,omitempty" tf:"destination_port_start,omitempty"`

	// (String) The direction of network traffic this rule will be applied to
	// The direction of network traffic this rule will be applied to
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// (String) The address family of new firewall rule
	// The address family of new firewall rule
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// (String) The ICMP type
	// The ICMP type
	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// (String) The protocol this rule will be applied to
	// The protocol this rule will be applied to
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (String) The source address range ends from this address
	// The source address range ends from this address
	SourceAddressEnd *string `json:"sourceAddressEnd,omitempty" tf:"source_address_end,omitempty"`

	// (String) The source address range starts from this address
	// The source address range starts from this address
	SourceAddressStart *string `json:"sourceAddressStart,omitempty" tf:"source_address_start,omitempty"`

	// (String) The source port range ends from this port number
	// The source port range ends from this port number
	SourcePortEnd *string `json:"sourcePortEnd,omitempty" tf:"source_port_end,omitempty"`

	// (String) The source port range starts from this port number
	// The source port range starts from this port number
	SourcePortStart *string `json:"sourcePortStart,omitempty" tf:"source_port_start,omitempty"`
}

type FirewallRuleObservation struct {

	// (String) Action to take if the rule conditions are met
	// Action to take if the rule conditions are met
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// (String) Freeform comment string for the rule
	// Freeform comment string for the rule
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (String) The destination address range ends from this address
	// The destination address range ends from this address
	DestinationAddressEnd *string `json:"destinationAddressEnd,omitempty" tf:"destination_address_end,omitempty"`

	// (String) The destination address range starts from this address
	// The destination address range starts from this address
	DestinationAddressStart *string `json:"destinationAddressStart,omitempty" tf:"destination_address_start,omitempty"`

	// (String) The destination port range ends from this port number
	// The destination port range ends from this port number
	DestinationPortEnd *string `json:"destinationPortEnd,omitempty" tf:"destination_port_end,omitempty"`

	// (String) The destination port range starts from this port number
	// The destination port range starts from this port number
	DestinationPortStart *string `json:"destinationPortStart,omitempty" tf:"destination_port_start,omitempty"`

	// (String) The direction of network traffic this rule will be applied to
	// The direction of network traffic this rule will be applied to
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// (String) The address family of new firewall rule
	// The address family of new firewall rule
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// (String) The ICMP type
	// The ICMP type
	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// (String) The protocol this rule will be applied to
	// The protocol this rule will be applied to
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (String) The source address range ends from this address
	// The source address range ends from this address
	SourceAddressEnd *string `json:"sourceAddressEnd,omitempty" tf:"source_address_end,omitempty"`

	// (String) The source address range starts from this address
	// The source address range starts from this address
	SourceAddressStart *string `json:"sourceAddressStart,omitempty" tf:"source_address_start,omitempty"`

	// (String) The source port range ends from this port number
	// The source port range ends from this port number
	SourcePortEnd *string `json:"sourcePortEnd,omitempty" tf:"source_port_end,omitempty"`

	// (String) The source port range starts from this port number
	// The source port range starts from this port number
	SourcePortStart *string `json:"sourcePortStart,omitempty" tf:"source_port_start,omitempty"`
}

type FirewallRuleParameters struct {

	// (String) Action to take if the rule conditions are met
	// Action to take if the rule conditions are met
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// (String) Freeform comment string for the rule
	// Freeform comment string for the rule
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (String) The destination address range ends from this address
	// The destination address range ends from this address
	// +kubebuilder:validation:Optional
	DestinationAddressEnd *string `json:"destinationAddressEnd,omitempty" tf:"destination_address_end,omitempty"`

	// (String) The destination address range starts from this address
	// The destination address range starts from this address
	// +kubebuilder:validation:Optional
	DestinationAddressStart *string `json:"destinationAddressStart,omitempty" tf:"destination_address_start,omitempty"`

	// (String) The destination port range ends from this port number
	// The destination port range ends from this port number
	// +kubebuilder:validation:Optional
	DestinationPortEnd *string `json:"destinationPortEnd,omitempty" tf:"destination_port_end,omitempty"`

	// (String) The destination port range starts from this port number
	// The destination port range starts from this port number
	// +kubebuilder:validation:Optional
	DestinationPortStart *string `json:"destinationPortStart,omitempty" tf:"destination_port_start,omitempty"`

	// (String) The direction of network traffic this rule will be applied to
	// The direction of network traffic this rule will be applied to
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// (String) The address family of new firewall rule
	// The address family of new firewall rule
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// (String) The ICMP type
	// The ICMP type
	// +kubebuilder:validation:Optional
	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// (String) The protocol this rule will be applied to
	// The protocol this rule will be applied to
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (String) The source address range ends from this address
	// The source address range ends from this address
	// +kubebuilder:validation:Optional
	SourceAddressEnd *string `json:"sourceAddressEnd,omitempty" tf:"source_address_end,omitempty"`

	// (String) The source address range starts from this address
	// The source address range starts from this address
	// +kubebuilder:validation:Optional
	SourceAddressStart *string `json:"sourceAddressStart,omitempty" tf:"source_address_start,omitempty"`

	// (String) The source port range ends from this port number
	// The source port range ends from this port number
	// +kubebuilder:validation:Optional
	SourcePortEnd *string `json:"sourcePortEnd,omitempty" tf:"source_port_end,omitempty"`

	// (String) The source port range starts from this port number
	// The source port range starts from this port number
	// +kubebuilder:validation:Optional
	SourcePortStart *string `json:"sourcePortStart,omitempty" tf:"source_port_start,omitempty"`
}

type FirewallRulesInitParameters struct {

	// address/port range.
	// The default rule can be created by providing only "action" and "direction" attributes. Default rule should be defined last. (see below for nested schema)
	// A single firewall rule.
	// If used, IP address and port ranges must have both start and end values specified. These can be the same value if only one IP address or port number is specified.
	// Source and destination port numbers can only be set if the protocol is TCP or UDP.
	// The ICMP type may only be set if the protocol is ICMP.
	// Typical firewall rule should have "action", "direction", "protocol", "family" and at least one destination/source-address/port range.
	// The default rule can be created by providing only "action" and "direction" attributes. Default rule should be defined last.
	FirewallRule []FirewallRuleInitParameters `json:"firewallRule,omitempty" tf:"firewall_rule,omitempty"`

	// (String) The unique id of the server to be protected the firewall rules
	// The unique id of the server to be protected the firewall rules
	// +crossplane:generate:reference:type=Server
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

type FirewallRulesObservation struct {

	// address/port range.
	// The default rule can be created by providing only "action" and "direction" attributes. Default rule should be defined last. (see below for nested schema)
	// A single firewall rule.
	// If used, IP address and port ranges must have both start and end values specified. These can be the same value if only one IP address or port number is specified.
	// Source and destination port numbers can only be set if the protocol is TCP or UDP.
	// The ICMP type may only be set if the protocol is ICMP.
	// Typical firewall rule should have "action", "direction", "protocol", "family" and at least one destination/source-address/port range.
	// The default rule can be created by providing only "action" and "direction" attributes. Default rule should be defined last.
	FirewallRule []FirewallRuleObservation `json:"firewallRule,omitempty" tf:"firewall_rule,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The unique id of the server to be protected the firewall rules
	// The unique id of the server to be protected the firewall rules
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

type FirewallRulesParameters struct {

	// address/port range.
	// The default rule can be created by providing only "action" and "direction" attributes. Default rule should be defined last. (see below for nested schema)
	// A single firewall rule.
	// If used, IP address and port ranges must have both start and end values specified. These can be the same value if only one IP address or port number is specified.
	// Source and destination port numbers can only be set if the protocol is TCP or UDP.
	// The ICMP type may only be set if the protocol is ICMP.
	// Typical firewall rule should have "action", "direction", "protocol", "family" and at least one destination/source-address/port range.
	// The default rule can be created by providing only "action" and "direction" attributes. Default rule should be defined last.
	// +kubebuilder:validation:Optional
	FirewallRule []FirewallRuleParameters `json:"firewallRule,omitempty" tf:"firewall_rule,omitempty"`

	// (String) The unique id of the server to be protected the firewall rules
	// The unique id of the server to be protected the firewall rules
	// +crossplane:generate:reference:type=Server
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

// FirewallRulesSpec defines the desired state of FirewallRules
type FirewallRulesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallRulesParameters `json:"forProvider"`
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
	InitProvider FirewallRulesInitParameters `json:"initProvider,omitempty"`
}

// FirewallRulesStatus defines the observed state of FirewallRules.
type FirewallRulesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallRulesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FirewallRules is the Schema for the FirewallRuless API. This resource represents a generated list of UpCloud firewall rules. Firewall rules are used in conjunction with UpCloud servers. Each server has its own firewall rules. The firewall is enabled on all network interfaces except ones attached to private virtual networks. The maximum number of firewall rules per server is 1000.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upcloud}
type FirewallRules struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.firewallRule) || (has(self.initProvider) && has(self.initProvider.firewallRule))",message="spec.forProvider.firewallRule is a required parameter"
	Spec   FirewallRulesSpec   `json:"spec"`
	Status FirewallRulesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRulesList contains a list of FirewallRuless
type FirewallRulesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallRules `json:"items"`
}

// Repository type metadata.
var (
	FirewallRules_Kind             = "FirewallRules"
	FirewallRules_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallRules_Kind}.String()
	FirewallRules_KindAPIVersion   = FirewallRules_Kind + "." + CRDGroupVersion.String()
	FirewallRules_GroupVersionKind = CRDGroupVersion.WithKind(FirewallRules_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallRules{}, &FirewallRulesList{})
}
