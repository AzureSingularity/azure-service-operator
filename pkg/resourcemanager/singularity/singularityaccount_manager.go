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

type SingularityAccountManager interface {
	CreateSingularityAccount(ctx context.Context,
		groupName string,
		singularityAccountName string,
		location string,
		locations *[]azurev1alpha1.SingularityAccountLocation,
		description string,
		schedulingPolicy azurev1alpha1.SchedulingPolicy,
		idleResourcesHandlingPolicy azurev1alpha1.IdleResourcesHandlingPolicy,
		networkSettings azurev1alpha1.NetworkSettings) (pollingURL string, result aisc.AccountResourceDescription, err error)

	// Get gets the description of the specified singularity account.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// singularityAccountName - the name of the singularity account
	GetSingularityAccount(ctx context.Context, resourceGroupName string, singularityAccountName string) (result aisc.AccountResourceDescription, err error)

	// Deletesingularity removes the singularity account
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// singularityAccountName - the name of the singularity account
	DeleteSingularityAccount(ctx context.Context, groupName string, singularityAccountName string) (result autorest.Response, err error)

	resourcemanager.ARMClient
}
