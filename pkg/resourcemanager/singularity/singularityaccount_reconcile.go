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

// Ensure creates a singularity account
func (sa *azureSingularityManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := sa.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	locations := instance.Spec.Locations
	pollURL := instance.Status.PollingURL

	hash := ""
	stor, err := sa.GetSingularityAccount(ctx, groupName, name)
	if err != nil {
		instance.Status.Message = err.Error()
		instance.Status.State = "NotReady"
	} else {
		instance.Status.State = string(stor.AccountResourceDescriptionProperties.ProvisioningState)

		hash = helpers.Hash256(instance.Spec)
		if instance.Status.SpecHash == hash && (instance.Status.Provisioned || instance.Status.FailedProvisioning) {
			instance.Status.RequestedAt = nil
			return true, nil
		}
	}

	if instance.Status.State == "Succeeded" {

		// everything finished successfully!
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.SpecHash = hash
		instance.Status.ResourceId = *stor.ID
		instance.Status.PollingURL = ""
		return true, nil
	}

	instance.Status.Provisioning = true
	instance.Status.Provisioned = false

	pollURL, _, err = sa.CreateSingularityAccount(
		ctx, groupName, name, location, locations, instance.Spec.Description, instance.Spec.SchedulingPolicy,
		instance.Spec.IdleResourcesHandlingPolicy, instance.Spec.NetworkSettings)

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

// Delete drops a singularity account
func (sa *azureSingularityManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := sa.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	_, err = sa.DeleteSingularityAccount(ctx, groupName, name)
	if err != nil {
		azerr := errhelp.NewAzureError(err)
		if azerr.Code != 404 {
			return true, err
		}
	}

	_, err = sa.GetSingularityAccount(ctx, groupName, name)
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
func (sa *azureSingularityManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
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
	}, nil
}

func (sa *azureSingularityManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := sa.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (sa *azureSingularityManager) convert(obj runtime.Object) (*azurev1alpha1.SingularityAccount, error) {
	local, ok := obj.(*azurev1alpha1.SingularityAccount)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
