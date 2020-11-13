// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package singularity

import (
	"context"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	aisc "github.com/Azure/azure-service-operator/pkg/resourcemanager/singularity/microsoftazuremanagementaisupercomputer"
	"github.com/Azure/go-autorest/autorest"
)

type SingularityStorageManager interface {
	CreateSingularityStorageContainer(ctx context.Context,
		groupName string,
		singularityAccountName string,
		storageName string,
		tier string,
		location string,
		description string) (pollingURL string, result aisc.StorageContainerResourceDescription, err error)

	// Get gets the description of the specified singularity account.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// singularityAccountName - the name of the singularity account
	GetSingularityStorageContainer(ctx context.Context, resourceGroupName string, singularityAccountName string, storageName string) (result aisc.StorageContainerResourceDescription, err error)

	// Deletesingularity removes the singularity account
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// singularityAccountName - the name of the singularity account
	DeleteSingularityStorageContainer(ctx context.Context, groupName string, singularityAccountName string, storageName string) (result autorest.Response, err error)

	resourcemanager.ARMClient
}
