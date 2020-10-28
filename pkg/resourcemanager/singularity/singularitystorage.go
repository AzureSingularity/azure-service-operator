// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package singularity

import (
	"context"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	aisc "github.com/Azure/azure-service-operator/pkg/resourcemanager/singularity/microsoftazuremanagementaisupercomputer"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type azureSingularityStorageManager struct {
}

func NewSingulartiyStorageClient() *azureSingularityStorageManager {
	return &azureSingularityStorageManager{}
}

func getSingularityStorageClient() (aisc.StorageContainerClient, error) {
	singularityClient := aisc.NewStorageContainerClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return aisc.StorageContainerClient{}, err
	}
	singularityClient.Authorizer = a
	singularityClient.AddToUserAgent(config.UserAgent())
	return singularityClient, nil
}

// CreateSingularityJob creates a new singularity Job
func (_ *azureSingularityStorageManager) CreateSingularityStorageContainer(ctx context.Context,
	groupName string,
	singularityAccountName string,
	storageName string,
	tier string,
	location string,
	description string) (pollingURL string, result aisc.StorageContainerResourceDescription, err error) {
	singularityClient, err := getSingularityStorageClient()
	if err != nil {
		return "", aisc.StorageContainerResourceDescription{}, err
	}

	storageResourceProperties := &aisc.StorageContainerResourceDescriptionProperties{
		Tier:        aisc.StorageContainerTier(tier),
		Description: to.StringPtr(description),
		Location:    to.StringPtr(location),
	}

	storageResource := aisc.StorageContainerResourceDescription{
		StorageContainerResourceDescriptionProperties: storageResourceProperties,
	}

	future, err := singularityClient.CreateOrUpdate(ctx, groupName, singularityAccountName, storageName, storageResource)
	if err != nil {
		return "", result, err
	}

	result, err = future.Result(singularityClient)

	return future.PollingURL(), result, err

}

// Get gets the description of the specified singularity Job.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// singularityJobName - the name of the singularity Job
func (_ *azureSingularityStorageManager) GetSingularityStorageContainer(ctx context.Context, resourceGroupName string, singularityAccountName string, storageName string) (result aisc.StorageContainerResourceDescription, err error) {
	singularityClient, err := getSingularityStorageClient()
	if err != nil {
		return aisc.StorageContainerResourceDescription{}, err
	}

	return singularityClient.Get(ctx, resourceGroupName, singularityAccountName, storageName)
}

// Deletesingularity removes the resource group named by env var
func (_ *azureSingularityStorageManager) DeleteSingularityStorageContainer(ctx context.Context, groupName string, singularityAccountName string, storageName string) (result autorest.Response, err error) {
	singularityClient, err := getSingularityStorageClient()
	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}

	future, err := singularityClient.Delete(ctx, groupName, singularityAccountName, storageName)
	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}

	result, err = future.Result(singularityClient)

	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	} else {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 200,
			},
		}, err
	}
}
