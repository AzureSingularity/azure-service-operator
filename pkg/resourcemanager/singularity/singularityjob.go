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

type azureSingularityJobManager struct {
}

func NewSingulartiyJoblient() *azureSingularityJobManager {
	return &azureSingularityJobManager{}
}

func getSingularityJobClient() (aisc.JobClient, error) {
	singularityClient := aisc.NewJobClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return aisc.JobClient{}, err
	}
	singularityClient.Authorizer = a
	singularityClient.AddToUserAgent(config.UserAgent())
	return singularityClient, nil
}

// CreateSingularityJob creates a new singularity Job
func (_ *azureSingularityJobManager) CreateSingularityJob(ctx context.Context,
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
	dataLocation []azurev1alpha1.JobStorageLocation) (pollingURL string, result aisc.JobResourceDescription, err error) {
	singularityClient, err := getSingularityJobClient()
	if err != nil {
		return "", aisc.JobResourceDescription{}, err
	}

	var policies []aisc.PlacementPolicy
	for _, pol := range *placementPolicies {
		policy := aisc.PlacementPolicy{
			Location: to.StringPtr(pol.Location),
		}

		var instanceTypes []aisc.InstanceTypeSettings
		for _, inst := range *pol.InstanceTypes {
			instance := aisc.InstanceTypeSettings{
				InstanceType: to.StringPtr(inst.InstanceType),
				ScalePolicy: &aisc.ScalePolicy{
					AutoScale:                to.BoolPtr(inst.ScalePolicy.AutoScale),
					CurrentInstanceTypeCount: to.Int32Ptr(inst.ScalePolicy.CurrentInstanceTypeCount),
				},
			}

			instanceTypes = append(instanceTypes, instance)
		}

		policy.InstanceTypes = &instanceTypes
		policies = append(policies, policy)
	}

	jobResourceProperties := &aisc.JobResourceDescriptionProperties{
		PlacementPolicies:   &policies,
		MaxJobExecutionTime: to.Float64Ptr(float64(maxJobExecutionTime)),
		GroupPolicyName:     to.StringPtr(groupPolicyName),
		SchedulingPriority:  aisc.SchedulingPriority(schedulingPriority),
		Program:             to.StringPtr(program),
		ProgramArgs:         to.StringPtr(programArgs),
		Description:         to.StringPtr(description),
	}

	if len(tensorBoardLogDirectory) > 0 {
		jobResourceProperties.TensorBoardLogDirectory = to.StringPtr(tensorBoardLogDirectory)
	}

	if frameworkImage.Kind == "PyTorch" {
		jobResourceProperties.FrameworkImage = &aisc.PyTorchFrameworkImage{
			Version: &frameworkImage.Version,
		}
	}

	jobResourceProperties.CodeLocation = GetStorageLocation(codeLocation)
	jobResourceProperties.OutputLocation = GetStorageLocation(outputLocation)

	var dataLocs []aisc.StorageLocation

	for _, tempLoc := range dataLocation {
		dataLocs = append(dataLocs, *GetStorageLocation(tempLoc))
	}

	jobResourceProperties.DataLocation = &dataLocs

	if frameworkImage.Kind == "TensorFlow" {
		jobResourceProperties.FrameworkImage = &aisc.TensorFlowFrameworkImage{
			Version: &frameworkImage.Version,
		}
	}

	jobResource := aisc.JobResourceDescription{
		JobResourceDescriptionProperties: jobResourceProperties,
	}

	future, err := singularityClient.CreateOrUpdate(ctx, groupName, singularityAccountName, singularityJobName, jobResource)
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
func (_ *azureSingularityJobManager) GetSingularityJob(ctx context.Context, resourceGroupName string, singularityAccountName string, singularityJobName string) (result aisc.JobResourceDescription, err error) {
	singularityClient, err := getSingularityJobClient()
	if err != nil {
		return aisc.JobResourceDescription{}, err
	}

	return singularityClient.Get(ctx, resourceGroupName, singularityAccountName, singularityJobName)
}

func GetStorageLocation(storageLocation azurev1alpha1.JobStorageLocation) (result *aisc.StorageLocation) {

	if storageLocation.Source.Kind == "AzureAISupercomputerStorage" {
		return &aisc.StorageLocation{
			Mount: &aisc.StorageMount{
				Path: &storageLocation.Mount.Path,
			},
			Source: &aisc.StorageSourceAzureAISupercomputerStorage{
				StorageContainerName: &storageLocation.Source.StorageContainerName,
			},
		}
	}

	if storageLocation.Source.Kind == "AzureBlobStorage" {
		return &aisc.StorageLocation{
			Mount: &aisc.StorageMount{
				Path: &storageLocation.Mount.Path,
			},
			Source: &aisc.StorageSourceAzureBlob{
				BlobEndpoint: &storageLocation.Source.BlobEndpoint,
			},
		}
	}
	if storageLocation.Source.Kind == "AzureFileStorage" {
		return &aisc.StorageLocation{
			Mount: &aisc.StorageMount{
				Path: &storageLocation.Mount.Path,
			},
			Source: &aisc.StorageSourceAzureFile{
				FileEndpoint:   &storageLocation.Source.FileEndpoint,
				StorageAccount: &storageLocation.Source.StorageAccount,
				AccountKey:     &storageLocation.Source.AccountKey,
			},
		}
	}

	return nil
}

// Deletesingularity removes the resource group named by env var
func (_ *azureSingularityJobManager) DeleteSingularityJob(ctx context.Context, groupName string, singularityAccountName string, singularityJobName string) (result autorest.Response, err error) {
	singularityClient, err := getSingularityJobClient()
	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}

	future, err := singularityClient.Delete(ctx, groupName, singularityAccountName, singularityJobName)
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
