// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package singularity

import (
	"context"
	"net/http"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	aisc "github.com/Azure/azure-service-operator/pkg/resourcemanager/singularity/microsoftazuremanagementaisupercomputer"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type azureSingularityManager struct {
}

func NewSingulartiyAccountlient() *azureSingularityManager {
	return &azureSingularityManager{}
}

func getSingularityClient() (aisc.AccountClient, error) {
	singularityClient := aisc.NewAccountClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return aisc.AccountClient{}, err
	}
	singularityClient.Authorizer = a
	singularityClient.AddToUserAgent(config.UserAgent())
	return singularityClient, nil
}

// CreateSingularityAccount creates a new singularity account
func (_ *azureSingularityManager) CreateSingularityAccount(ctx context.Context,
	groupName string,
	singularityAccountName string,
	location string,
	locations *[]azurev1alpha1.SingularityAccountLocation,
	description string,
	schedulingPolicy azurev1alpha1.SchedulingPolicy,
	idleResourcesHandlingPolicy azurev1alpha1.IdleResourcesHandlingPolicy,
	networkSettings azurev1alpha1.NetworkSettings) (pollingURL string, result aisc.AccountResourceDescription, err error) {
	singularityClient, err := getSingularityClient()
	if err != nil {
		return "", aisc.AccountResourceDescription{}, err
	}

	var resourceLocations []aisc.AccountLocation
	for _, loc := range *locations {
		loc := aisc.AccountLocation{
			Name:             to.StringPtr(loc.Name),
			FailoverPriority: to.Int32Ptr(loc.FailoverPriority),
		}
		resourceLocations = append(resourceLocations, loc)
	}

	accResourceProperties := &aisc.AccountResourceDescriptionProperties{
		Locations: &resourceLocations,
	}

	if len(description) > 0 {
		accResourceProperties.Description = &description
	}

	if len(schedulingPolicy.SchedulingMode) > 0 {
		accResourceProperties.SchedulingPolicy = &aisc.SchedulingPolicy{
			SchedulingMode: aisc.SchedulingMode(schedulingPolicy.SchedulingMode),
		}
	}

	if len(idleResourcesHandlingPolicy.IdleResourcesHandlingMode) > 0 {
		accResourceProperties.IdleResourcesHandlingPolicy = &aisc.IdleResourcesHandlingPolicy{
			IdleResourcesHandlingMode: aisc.IdleResourcesHandlingMode(idleResourcesHandlingPolicy.IdleResourcesHandlingMode),
		}
	}

	if len(networkSettings.SubnetId) > 0 {
		accResourceProperties.NetworkSettings = &aisc.NetworkSettings{
			SubnetID: &networkSettings.SubnetId,
		}
	}

	accResource := aisc.AccountResourceDescription{
		AccountResourceDescriptionProperties: accResourceProperties,
		Location:                             &location,
	}

	future, err := singularityClient.CreateOrUpdate(ctx, groupName, singularityAccountName, accResource)
	if err != nil {
		return "", result, err
	}

	result, err = future.Result(singularityClient)

	return future.PollingURL(), result, err

}

// Get gets the description of the specified singularity account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// singularityAccountName - the name of the singularity account
func (_ *azureSingularityManager) GetSingularityAccount(ctx context.Context, resourceGroupName string, singularityAccountName string) (result aisc.AccountResourceDescription, err error) {
	singularityClient, err := getSingularityClient()
	if err != nil {
		return aisc.AccountResourceDescription{}, err
	}

	return singularityClient.Get(ctx, resourceGroupName, singularityAccountName)
}

// Deletesingularity removes the resource group named by env var
func (_ *azureSingularityManager) DeleteSingularityAccount(ctx context.Context, groupName string, singularityAccountName string) (result autorest.Response, err error) {
	singularityClient, err := getSingularityClient()
	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}

	future, err := singularityClient.Delete(ctx, groupName, singularityAccountName)
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
