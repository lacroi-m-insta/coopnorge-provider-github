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

type GroupInitParameters struct {

	// The description of the IdP group.
	// The description of the IdP group.
	GroupDescription *string `json:"groupDescription,omitempty" tf:"group_description,omitempty"`

	// The ID of the IdP group.
	// The ID of the IdP group.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The name of the IdP group.
	// The name of the IdP group.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`
}

type GroupObservation struct {

	// The description of the IdP group.
	// The description of the IdP group.
	GroupDescription *string `json:"groupDescription,omitempty" tf:"group_description,omitempty"`

	// The ID of the IdP group.
	// The ID of the IdP group.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The name of the IdP group.
	// The name of the IdP group.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`
}

type GroupParameters struct {

	// The description of the IdP group.
	// The description of the IdP group.
	// +kubebuilder:validation:Optional
	GroupDescription *string `json:"groupDescription" tf:"group_description,omitempty"`

	// The ID of the IdP group.
	// The ID of the IdP group.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId" tf:"group_id,omitempty"`

	// The name of the IdP group.
	// The name of the IdP group.
	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName" tf:"group_name,omitempty"`
}

type TeamSyncGroupMappingInitParameters struct {

	// An Array of GitHub Identity Provider Groups (or empty []).  Each group block consists of the fields documented below.
	// An Array of GitHub Identity Provider Groups (or empty []).
	Group []GroupInitParameters `json:"group,omitempty" tf:"group,omitempty"`

	// Slug of the team
	// Slug of the team.
	TeamSlug *string `json:"teamSlug,omitempty" tf:"team_slug,omitempty"`
}

type TeamSyncGroupMappingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// An Array of GitHub Identity Provider Groups (or empty []).  Each group block consists of the fields documented below.
	// An Array of GitHub Identity Provider Groups (or empty []).
	Group []GroupObservation `json:"group,omitempty" tf:"group,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Slug of the team
	// Slug of the team.
	TeamSlug *string `json:"teamSlug,omitempty" tf:"team_slug,omitempty"`
}

type TeamSyncGroupMappingParameters struct {

	// An Array of GitHub Identity Provider Groups (or empty []).  Each group block consists of the fields documented below.
	// An Array of GitHub Identity Provider Groups (or empty []).
	// +kubebuilder:validation:Optional
	Group []GroupParameters `json:"group,omitempty" tf:"group,omitempty"`

	// Slug of the team
	// Slug of the team.
	// +kubebuilder:validation:Optional
	TeamSlug *string `json:"teamSlug,omitempty" tf:"team_slug,omitempty"`
}

// TeamSyncGroupMappingSpec defines the desired state of TeamSyncGroupMapping
type TeamSyncGroupMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TeamSyncGroupMappingParameters `json:"forProvider"`
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
	InitProvider TeamSyncGroupMappingInitParameters `json:"initProvider,omitempty"`
}

// TeamSyncGroupMappingStatus defines the observed state of TeamSyncGroupMapping.
type TeamSyncGroupMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TeamSyncGroupMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TeamSyncGroupMapping is the Schema for the TeamSyncGroupMappings API. Creates and manages the connections between a team and its IdP group(s).
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type TeamSyncGroupMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.teamSlug) || (has(self.initProvider) && has(self.initProvider.teamSlug))",message="spec.forProvider.teamSlug is a required parameter"
	Spec   TeamSyncGroupMappingSpec   `json:"spec"`
	Status TeamSyncGroupMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamSyncGroupMappingList contains a list of TeamSyncGroupMappings
type TeamSyncGroupMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamSyncGroupMapping `json:"items"`
}

// Repository type metadata.
var (
	TeamSyncGroupMapping_Kind             = "TeamSyncGroupMapping"
	TeamSyncGroupMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TeamSyncGroupMapping_Kind}.String()
	TeamSyncGroupMapping_KindAPIVersion   = TeamSyncGroupMapping_Kind + "." + CRDGroupVersion.String()
	TeamSyncGroupMapping_GroupVersionKind = CRDGroupVersion.WithKind(TeamSyncGroupMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&TeamSyncGroupMapping{}, &TeamSyncGroupMappingList{})
}
