// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SingularityStorageContainerSpec defines the desired state of SingularityStorageContainer
type SingularityStorageContainerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`
	AccountName   string `json:"accountName"`
	Tier          string `json:"tier"`
	Location      string `json:"location"`

	// optional
	Description string `json:"description,omitempty"`
}

// +kubebuilder:object:root=true

// SingularityStorageContainer is the Schema for the singularitystoragecontainers API
type SingularityStorageContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SingularityStorageContainerSpec `json:"spec,omitempty"`
	Status ASOStatus                       `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SingularityStorageContainerList contains a list of SingularityStorageContainer
type SingularityStorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SingularityStorageContainer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SingularityStorageContainer{}, &SingularityStorageContainerList{})
}
