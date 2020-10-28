// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SingularityJobSpec defines the desired state of SingularityJob
type SingularityJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +kubebuilder:validation:Required
	ResourceGroup  string            `json:"resourceGroup"`
	AccountName    string            `json:"accountName"`
	FrameworkImage JobFrameworkImage `json:"frameworkImage"`
	// +kubebuilder:validation:MinLength:1
	PlacementPolicies   *[]JobPlacementPolicy `json:"placementPolicies"`
	MaxJobExecutionTime int32                 `json:"maxJobExecutionTime"`

	// optional
	SchedulingPriority      string               `json:"schedulingPriority,omitempty"`
	GroupPolicyName         string               `json:"groupPolicyName,omitempty"`
	Program                 string               `json:"program,omitempty"`
	ProgramArgs             string               `json:"programArgs,omitempty"`
	Description             string               `json:"description,omitempty"`
	TensorBoardLogDirectory string               `json:"tensorBoardLogDirectory,omitempty"`
	CodeLocation            JobStorageLocation   `json:"codeLocation,omitempty"`
	OutputLocation          JobStorageLocation   `json:"outputLocation,omitempty"`
	DataLocation            []JobStorageLocation `json:"dataLocation,omitempty"`
}

// +kubebuilder:object:root=true

// SingularityJob is the Schema for the singularityjobs API
type SingularityJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SingularityJobSpec `json:"spec,omitempty"`
	Status ASOStatus          `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SingularityJobList contains a list of SingularityJob
type SingularityJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SingularityJob `json:"items"`
}

type JobFrameworkImage struct {
	Version string `json:"version,omitempty"`

	// +kubebuilder:validation:Required
	Kind string `json:"kind"`
}

type JobPlacementPolicy struct {
	// +kubebuilder:validation:Required
	Location      string             `json:"location"`
	InstanceTypes *[]JobInstanceType `json:"instanceTypes"`
}

type JobInstanceType struct {
	// +kubebuilder:validation:Required
	InstanceType string          `json:"instanceType"`
	ScalePolicy  *JobScalePolicy `json:"scalePolicy"`
}

type JobScalePolicy struct {
	MinInstanceTypeCount int32 `json:"minInstanceTypeCount,omitempty"`
	MaxInstanceTypeCount int32 `json:"maxInstanceTypeCount,omitempty"`
	// +kubebuilder:validation:Required
	AutoScale                bool  `json:"autoScale"`
	CurrentInstanceTypeCount int32 `json:"currentInstanceTypeCount"`
}

type JobStorageLocation struct {
	// +kubebuilder:validation:Required
	Source JobStorageSource `json:"source,omitempty"`
	Mount  JobStorageMount  `json:"mount,omitempty"`
}

type JobStorageMount struct {
	// +kubebuilder:validation:Required
	Path string `json:"path,omitempty"`
}

type JobStorageSource struct {
	FileEndpoint   string `json:"fileEndpoint,omitempty"`
	StorageAccount string `json:"storageAccount,omitempty"`
	AccountKey     string `json:"accountKey,omitempty"`

	BlobEndpoint string `json:"blobEndpoint,omitempty"`

	StorageContainerName string `json:"storageContainerName,omitempty"`

	// +kubebuilder:validation:Required
	Kind string `json:"kind,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SingularityJob{}, &SingularityJobList{})
}
