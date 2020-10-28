package microsoftazuremanagementaisupercomputerapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "./"
)

        // AccountQuotaPolicyClientAPI contains the set of methods on the AccountQuotaPolicyClient type.
        type AccountQuotaPolicyClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, policyName string, body microsoftazuremanagementaisupercomputer.AccountQuotaPolicyResourceDescription) (result microsoftazuremanagementaisupercomputer.AccountQuotaPolicyCreateOrUpdateFuture, err error)
            Get(ctx context.Context, resourceGroupName string, accountName string, policyName string) (result microsoftazuremanagementaisupercomputer.AccountQuotaPolicyResourceDescription, err error)
            ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.AccountQuotaPolicyResourceDescriptionListPage, err error)
                ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.AccountQuotaPolicyResourceDescriptionListIterator, err error)
            Update(ctx context.Context, resourceGroupName string, accountName string, policyName string, body microsoftazuremanagementaisupercomputer.AccountQuotaPolicyResourcePatchDescription) (result microsoftazuremanagementaisupercomputer.AccountQuotaPolicyUpdateFuture, err error)
        }

        var _ AccountQuotaPolicyClientAPI = (*microsoftazuremanagementaisupercomputer.AccountQuotaPolicyClient)(nil)
        // AccountClientAPI contains the set of methods on the AccountClient type.
        type AccountClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, body microsoftazuremanagementaisupercomputer.AccountResourceDescription) (result microsoftazuremanagementaisupercomputer.AccountCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.AccountDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.AccountResourceDescription, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result microsoftazuremanagementaisupercomputer.AccountResourceDescriptionListPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result microsoftazuremanagementaisupercomputer.AccountResourceDescriptionListIterator, err error)
            ListBySubscription(ctx context.Context) (result microsoftazuremanagementaisupercomputer.AccountResourceDescriptionListPage, err error)
                ListBySubscriptionComplete(ctx context.Context) (result microsoftazuremanagementaisupercomputer.AccountResourceDescriptionListIterator, err error)
            Update(ctx context.Context, resourceGroupName string, accountName string, body microsoftazuremanagementaisupercomputer.AccountResourcePatchDescription) (result microsoftazuremanagementaisupercomputer.AccountUpdateFuture, err error)
        }

        var _ AccountClientAPI = (*microsoftazuremanagementaisupercomputer.AccountClient)(nil)
        // GroupPolicyClientAPI contains the set of methods on the GroupPolicyClient type.
        type GroupPolicyClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, policyName string, body microsoftazuremanagementaisupercomputer.GroupPolicyResourceDescription) (result microsoftazuremanagementaisupercomputer.GroupPolicyCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, accountName string, policyName string) (result microsoftazuremanagementaisupercomputer.GroupPolicyDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, accountName string, policyName string) (result microsoftazuremanagementaisupercomputer.GroupPolicyResourceDescription, err error)
            ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.GroupPolicyResourceDescriptionListPage, err error)
                ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.GroupPolicyResourceDescriptionListIterator, err error)
            Update(ctx context.Context, resourceGroupName string, accountName string, policyName string, body microsoftazuremanagementaisupercomputer.GroupPolicyResourcePatchDescription) (result microsoftazuremanagementaisupercomputer.GroupPolicyUpdateFuture, err error)
        }

        var _ GroupPolicyClientAPI = (*microsoftazuremanagementaisupercomputer.GroupPolicyClient)(nil)
        // InstanceTypeSeriesClientAPI contains the set of methods on the InstanceTypeSeriesClient type.
        type InstanceTypeSeriesClientAPI interface {
            List(ctx context.Context, location string) (result microsoftazuremanagementaisupercomputer.AvailableInstanceTypeSeriesDescriptionListPage, err error)
                ListComplete(ctx context.Context, location string) (result microsoftazuremanagementaisupercomputer.AvailableInstanceTypeSeriesDescriptionListIterator, err error)
            ListInstanceType(ctx context.Context, location string, instanceTypeSeriesID string) (result microsoftazuremanagementaisupercomputer.AvailableInstanceTypeDescriptionListPage, err error)
                ListInstanceTypeComplete(ctx context.Context, location string, instanceTypeSeriesID string) (result microsoftazuremanagementaisupercomputer.AvailableInstanceTypeDescriptionListIterator, err error)
        }

        var _ InstanceTypeSeriesClientAPI = (*microsoftazuremanagementaisupercomputer.InstanceTypeSeriesClient)(nil)
        // JobClientAPI contains the set of methods on the JobClient type.
        type JobClientAPI interface {
            Approve(ctx context.Context, resourceGroupName string, accountName string, jobName string, body *microsoftazuremanagementaisupercomputer.JobUserActionDescription) (result microsoftazuremanagementaisupercomputer.JobApproveFuture, err error)
            CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, jobName string, body microsoftazuremanagementaisupercomputer.JobResourceDescription) (result microsoftazuremanagementaisupercomputer.JobCreateOrUpdateFuture, err error)
            CreateSasToken(ctx context.Context, resourceGroupName string, accountName string, jobName string, body *microsoftazuremanagementaisupercomputer.JobSasTokenDescription) (result microsoftazuremanagementaisupercomputer.JobCreateSasTokenFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, accountName string, jobName string) (result microsoftazuremanagementaisupercomputer.JobDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, accountName string, jobName string) (result microsoftazuremanagementaisupercomputer.JobResourceDescription, err error)
            ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.JobResourceDescriptionListPage, err error)
                ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.JobResourceDescriptionListIterator, err error)
            Remove(ctx context.Context, resourceGroupName string, accountName string, jobName string, body *microsoftazuremanagementaisupercomputer.JobUserActionDescription) (result microsoftazuremanagementaisupercomputer.JobRemoveFuture, err error)
            RenewKey(ctx context.Context, resourceGroupName string, accountName string, jobName string) (result microsoftazuremanagementaisupercomputer.JobSasTokenDescription, err error)
            Resume(ctx context.Context, resourceGroupName string, accountName string, jobName string, body *microsoftazuremanagementaisupercomputer.JobUserActionDescription) (result microsoftazuremanagementaisupercomputer.JobResumeFuture, err error)
            Suspend(ctx context.Context, resourceGroupName string, accountName string, jobName string, body *microsoftazuremanagementaisupercomputer.JobUserActionDescription) (result microsoftazuremanagementaisupercomputer.JobSuspendFuture, err error)
            Update(ctx context.Context, resourceGroupName string, accountName string, jobName string, body microsoftazuremanagementaisupercomputer.JobResourcePatchDescription) (result microsoftazuremanagementaisupercomputer.JobUpdateFuture, err error)
        }

        var _ JobClientAPI = (*microsoftazuremanagementaisupercomputer.JobClient)(nil)
        // StorageContainerClientAPI contains the set of methods on the StorageContainerClient type.
        type StorageContainerClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, storageContainerResourceName string, body microsoftazuremanagementaisupercomputer.StorageContainerResourceDescription) (result microsoftazuremanagementaisupercomputer.StorageContainerCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, accountName string, storageContainerResourceName string) (result microsoftazuremanagementaisupercomputer.StorageContainerDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, accountName string, storageContainerResourceName string) (result microsoftazuremanagementaisupercomputer.StorageContainerResourceDescription, err error)
            ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.StorageContainerResourceDescriptionListPage, err error)
                ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result microsoftazuremanagementaisupercomputer.StorageContainerResourceDescriptionListIterator, err error)
            Resume(ctx context.Context, resourceGroupName string, accountName string, storageContainerResourceName string) (result microsoftazuremanagementaisupercomputer.StorageContainerResumeFuture, err error)
            Suspend(ctx context.Context, resourceGroupName string, accountName string, storageContainerResourceName string) (result microsoftazuremanagementaisupercomputer.StorageContainerSuspendFuture, err error)
            Update(ctx context.Context, resourceGroupName string, accountName string, storageContainerResourceName string, body microsoftazuremanagementaisupercomputer.StorageContainerResourcePatchDescription) (result microsoftazuremanagementaisupercomputer.StorageContainerUpdateFuture, err error)
        }

        var _ StorageContainerClientAPI = (*microsoftazuremanagementaisupercomputer.StorageContainerClient)(nil)
