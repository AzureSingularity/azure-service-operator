// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package singularity

import (
	"context"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	aisc "github.com/Azure/azure-service-operator/pkg/resourcemanager/singularity/microsoftazuremanagementaisupercomputer"
	"github.com/Azure/go-autorest/autorest"
)

type SingularityJobManager interface {
	CreateSingularityJob(ctx context.Context,
		groupName string,
		singularityAccountName string,
		singularityJobName string,
		frameworkImage azurev1alpha1.JobFrameworkImage,
		placementPolicies *[]azurev1alpha1.JobPlacementPolicy,
		maxJobExecutionTime int32,
		schedulingPriority string,
		groupPolicyName string,
		program string,
		programArgs string,
		description string,
		tensorBoardLogDirectory string,
		codeLocation azurev1alpha1.JobStorageLocation,
		outputLocation azurev1alpha1.JobStorageLocation,
		dataLocation []azurev1alpha1.JobStorageLocation) (pollingURL string, result aisc.JobResourceDescription, err error)

	// Get gets the description of the specified singularity account.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// singularityAccountName - the name of the singularity account
	GetSingularityJob(ctx context.Context, resourceGroupName string, singularityAccountName string, singularityJobName string) (result aisc.JobResourceDescription, err error)

	// Deletesingularity removes the singularity account
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// singularityAccountName - the name of the singularity account
	DeleteSingularityJob(ctx context.Context, groupName string, singularityAccountName string, singularityJobName string) (result autorest.Response, err error)

	resourcemanager.ARMClient
}
