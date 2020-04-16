// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all cosmos

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCosmosDBHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	cosmosDBAccountName := GenerateTestResourceNameWithRandom("cosmosdb", 8)
	cosmosDBNamespace := "default"

	dbInstance := &v1alpha1.CosmosDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cosmosDBAccountName,
			Namespace: cosmosDBNamespace,
		},
		Spec: v1alpha1.CosmosDBSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			Kind:          v1alpha1.CosmosDBKindGlobalDocumentDB,
			Properties: v1alpha1.CosmosDBProperties{
				DatabaseAccountOfferType: v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard,
			},
		},
	}

	EnsureInstance(ctx, t, tc, dbInstance)

	EnsureDelete(ctx, t, tc, dbInstance)

}

func TestCosmosDBControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgLocation := tc.resourceGroupLocation
	//wrong resource group name
	resourceGroupName := "gone"

	cosmosDBAccountName := GenerateTestResourceNameWithRandom("cosmosdb", 8)
	cosmosDBNamespace := "default"

	dbInstance1 := &v1alpha1.CosmosDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cosmosDBAccountName,
			Namespace: cosmosDBNamespace,
		},
		Spec: v1alpha1.CosmosDBSpec{
			Location:      rgLocation,
			ResourceGroup: resourceGroupName,
			Kind:          v1alpha1.CosmosDBKindGlobalDocumentDB,
			Properties: v1alpha1.CosmosDBProperties{
				DatabaseAccountOfferType: v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard,
			},
		},
	}
	//the expected error meessage to be shown
	errMessage := "Waiting for resource group '" + resourceGroupName + "' to be available"

	EnsureInstanceWithResult(ctx, t, tc, dbInstance1, errMessage, false)
	EnsureDelete(ctx, t, tc, dbInstance1)
}

func TestCosmosDBControllerInvalidLocation(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	resourceGroupName := tc.resourceGroupName
	//rglocation doesnot exist
	rgLocation := GenerateTestResourceNameWithRandom("cosmos-lo", 10)

	cosmosDBAccountName := GenerateTestResourceNameWithRandom("cosmos-db", 8)
	cosmosDBNamespace := "default"

	dbInstance2 := &v1alpha1.CosmosDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cosmosDBAccountName,
			Namespace: cosmosDBNamespace,
		},
		Spec: v1alpha1.CosmosDBSpec{
			Location:      rgLocation,
			ResourceGroup: resourceGroupName,
			Kind:          v1alpha1.CosmosDBKindGlobalDocumentDB,
			Properties: v1alpha1.CosmosDBProperties{
				DatabaseAccountOfferType: v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard,
			},
		},
	}

	//error meessage to be expected
	errMessage := "The specified location '" + rgLocation + "' is invalid"

	EnsureInstanceWithResult(ctx, t, tc, dbInstance2, errMessage, false)
	EnsureDelete(ctx, t, tc, dbInstance2)
}
