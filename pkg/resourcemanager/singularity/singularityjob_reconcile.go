// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package singularity

import (
	"context"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a singularity job
func (sa *azureSingularityJobManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := sa.convert(obj)
	if err != nil {
		return false, err
	}

	accountName := instance.Spec.AccountName
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	pollURL := instance.Status.PollingURL

	hash := ""
	stor, err := sa.GetSingularityJob(ctx, groupName, accountName, name)
	if err != nil {
		instance.Status.Message = err.Error()
		instance.Status.State = "NotReady"
	} else {
		instance.Status.State = string(stor.JobResourceDescriptionProperties.Status)

		hash = helpers.Hash256(instance.Spec)
		if instance.Status.SpecHash == hash && (instance.Status.Provisioned || instance.Status.FailedProvisioning) {
			instance.Status.RequestedAt = nil
			return true, nil
		}
	}

	if instance.Status.State == "Running" || instance.Status.State == "Completed" {
		// everything finished successfully!
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.SpecHash = hash
		instance.Status.ResourceId = *stor.ID
		instance.Status.PollingURL = ""
		instance.Endpoints = azurev1alpha1.JobEndpoints{
			NotebookEndpoint:    *stor.NotebookEndpoint,
			GrafanaEndpoint:     *stor.GrafanaEndpoint,
			TensorBoardEndpoint: *stor.TensorBoardEndpoint,
			JobLogsEndpoint:     *stor.JobLogsEndpoint,
		}
		return true, nil
	}

	if instance.Status.State != "NotReady" {
		return false, nil
	}

	instance.Status.Provisioning = true
	instance.Status.Provisioned = false

	pollURL, _, err = sa.CreateSingularityJob(
		ctx, groupName, accountName, name, instance.Spec.FrameworkImage, instance.Spec.PlacementPolicies, instance.Spec.MaxJobExecutionTime, instance.Spec.SchedulingPriority, instance.Spec.GroupPolicyName, instance.Spec.Program, instance.Spec.ProgramArgs, instance.Spec.Description, instance.Spec.TensorBoardLogDirectory, instance.Spec.CodeLocation, instance.Spec.OutputLocation, instance.Spec.DataLocation)

	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)
		instance.Status.Provisioning = false

		ignore := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		if helpers.ContainsString(ignore, azerr.Type) {
			instance.Status.Provisioning = false
			return false, nil
		}

		wait := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.AlreadyExists,
		}
		if helpers.ContainsString(wait, azerr.Type) {

			if azerr.Type == errhelp.AsyncOpIncompleteError {
				instance.Status.Provisioning = true
				instance.Status.PollingURL = pollURL
			}
			return false, nil
		}

		return false, err
	}

	return false, nil
}

// Delete drops a singularity job
func (sa *azureSingularityJobManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := sa.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	accountName := instance.Spec.AccountName
	groupName := instance.Spec.ResourceGroup

	_, err = sa.DeleteSingularityJob(ctx, groupName, accountName, name)
	if err != nil {
		azerr := errhelp.NewAzureError(err)
		if azerr.Code != 404 {
			return true, err
		}
	}

	_, err = sa.GetSingularityJob(ctx, groupName, accountName, name)
	if err != nil {
		catch := []string{
			errhelp.AsyncOpIncompleteError,
		}
		gone := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		} else if helpers.ContainsString(gone, azerr.Type) {
			return false, nil
		}
		return true, err
	}
	return true, nil
}

// GetParents returns the parents of a singularity account
func (sa *azureSingularityJobManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := sa.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Name:      instance.Spec.ResourceGroup,
				Namespace: instance.Namespace,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
		{
			Key: types.NamespacedName{
				Name:      instance.Spec.AccountName,
				Namespace: instance.Namespace,
			},
			Target: &azurev1alpha1.SingularityAccount{},
		},
	}, nil
}

func (sa *azureSingularityJobManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := sa.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (sa *azureSingularityJobManager) convert(obj runtime.Object) (*azurev1alpha1.SingularityJob, error) {
	local, ok := obj.(*azurev1alpha1.SingularityJob)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
