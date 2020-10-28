// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SingularityAccountSpec defines the desired state of SingularityAccount
type SingularityAccountSpec struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Location string `json:"location"`
	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength:1
	ResourceGroup string `json:"resourceGroup"`
	// +kubebuilder:validation:MinLength:1
	Locations *[]SingularityAccountLocation `json:"locations"`

	// optional
	Description                 string                      `json:"description,omitempty"`
	SchedulingPolicy            SchedulingPolicy            `json:"schedulingPolicy,omitempty"`
	IdleResourcesHandlingPolicy IdleResourcesHandlingPolicy `json:"idleResourcesHandlingPolicy,omitempty"`
	NetworkSettings             NetworkSettings             `json:"networkSettings,omitempty"`
}

// +kubebuilder:object:root=true

// SingularityAccount is the Schema for the singularityaccounts API
type SingularityAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SingularityAccountSpec `json:"spec,omitempty"`
	Status ASOStatus              `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SingularityAccountList contains a list of SingularityAccount
type SingularityAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SingularityAccount `json:"items"`
}

type SingularityAccountLocation struct {
	// +kubebuilder:validation:Required
	// Specifies the location where jobs can be created
	Name string `json:"name"`

	// +kubebuilder:validation:Required
	// Specifies the failover priority for the location
	FailoverPriority int32 `json:"failoverPriority"`
}

type SchedulingPolicy struct {
	SchedulingMode string `json:"schedulingMode,omitempty"`
}

type IdleResourcesHandlingPolicy struct {
	IdleResourcesHandlingMode string `json:"idleResourcesHandlingMode,omitempty"`
}

type NetworkSettings struct {
	SubnetId string `json:"subnetId,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SingularityAccount{}, &SingularityAccountList{})
}

func (singularityAccount *SingularityAccount) IsSubmitted() bool {
	return singularityAccount.Status.Provisioning || singularityAccount.Status.Provisioned
}
