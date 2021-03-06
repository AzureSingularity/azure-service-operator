// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvault

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	kvhelper "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	rghelper "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Keyvault Secrets Client", func() {

	var ctx context.Context
	var err error
	var timeout time.Duration
	var retry time.Duration

	// Define resource group & keyvault constants
	var keyVaultName string
	var resourcegroupName string
	var resourcegroupLocation string
	var userID string

	resourceGroupManager := rghelper.NewAzureResourceGroupManager()

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test

		// Create a context to use in the tests
		ctx = context.Background()

		// Set timeout to 300 seconds
		timeout = 300 * time.Second

		// Set retryinterval to 1 second
		retry = 1 * time.Second

		// Initialize service principal ID to give access to the keyvault
		userID = config.ClientID()

		// Initialize resource names
		keyVaultName = "t-kvtest-kv" + strconv.FormatInt(GinkgoRandomSeed(), 10)
		resourcegroupName = "t-kvtest-rg" + helpers.RandomString(10)
		resourcegroupLocation = config.DefaultLocation()

		// Create a resource group
		log.Println("Creating resource group with name " + resourcegroupName + " in location " + resourcegroupLocation)
		_, err = resourceGroupManager.CreateGroup(ctx, resourcegroupName, resourcegroupLocation)
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() bool {
			result, _ := resourceGroupManager.CheckExistence(ctx, resourcegroupName)
			return result.Response.StatusCode == http.StatusNoContent
		}, timeout, retry,
		).Should(BeTrue())

		// Create a keyvault
		_, err = kvhelper.AzureKeyVaultManager.CreateVaultWithAccessPolicies(ctx, resourcegroupName, keyVaultName, resourcegroupLocation, userID)
		//Expect(err).NotTo(HaveOccurred())

	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		// Delete the keyvault
		kvhelper.AzureKeyVaultManager.DeleteVault(ctx, resourcegroupName, keyVaultName)
		//Expect(err).NotTo(HaveOccurred())

		// Delete the resource group
		_, err = resourceGroupManager.DeleteGroup(context.Background(), resourcegroupName)
		if err != nil {
			azerr := errhelp.NewAzureError(err)
			if azerr.Type == errhelp.AsyncOpIncompleteError {
				err = nil
			}
		}
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() bool {
			result, _ := resourceGroupManager.CheckExistence(ctx, resourcegroupName)
			return result.Response.StatusCode == http.StatusNoContent
		}, timeout, retry,
		).Should(BeFalse())
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete secret in Keyvault", func() {
			secretName := "kvsecret" + strconv.FormatInt(GinkgoRandomSeed(), 10)
			activationDate := time.Date(2018, time.January, 22, 15, 34, 0, 0, time.UTC)
			expiryDate := time.Date(2030, time.February, 1, 12, 22, 0, 0, time.UTC)

			var err error

			data := map[string][]byte{
				"test":  []byte("data"),
				"sweet": []byte("potato"),
			}

			client := New(keyVaultName)

			key := types.NamespacedName{Name: secretName, Namespace: "default"}

			Context("creating secret with KeyVault client", func() {
				err = client.Create(ctx, key, data, secrets.WithActivation(&activationDate), secrets.WithExpiration(&expiryDate))
				Expect(err).To(BeNil())
			})

			Context("ensuring secret exists using keyvault client", func() {
				d, err := client.Get(ctx, key)
				Expect(err).To(BeNil())

				for k, v := range d {
					Expect(data[k]).To(Equal(v))
				}
			})

			datanew := map[string][]byte{
				"french": []byte("fries"),
				"hot":    []byte("dogs"),
			}

			Context("upserting the secret to make sure it can be written", func() {
				err = client.Upsert(ctx, key, datanew, secrets.WithActivation(&activationDate), secrets.WithExpiration(&expiryDate))
				Expect(err).To(BeNil())
			})

			Context("ensuring secret exists using keyvault client", func() {
				d, err := client.Get(ctx, key)
				Expect(err).To(BeNil())

				for k, v := range d {
					Expect(datanew[k]).To(Equal(v))
				}
				Expect(datanew["french"]).To(Equal([]byte("fries")))
			})

			Context("delete secret and ensure it is gone", func() {
				err = client.Delete(ctx, key)
				Expect(err).To(BeNil())

				d, err := client.Get(ctx, key)
				Expect(err).ToNot(BeNil())
				for k, v := range d {
					Expect(data[k]).To(Equal(v))
				}
			})
		})

		It("should create and delete secrets in Keyvault with Flatten enabled", func() {
			secretName := "kvsecret" + strconv.FormatInt(GinkgoRandomSeed(), 10)

			var err error

			data := map[string][]byte{
				"test":  []byte("data"),
				"sweet": []byte("potato"),
			}

			client := New(keyVaultName)

			key := types.NamespacedName{Name: secretName, Namespace: "default"}

			Context("creating flattened secret with KeyVault client", func() {
				err = client.Create(ctx, key, data, secrets.Flatten(true))
				Expect(err).To(BeNil())
			})

			Context("ensuring flattened secrets exist using keyvault client", func() {
				// Look for each originally passed secret item in the keyvault
				for testKey, testValue := range data {
					returnedValue, err := client.Get(
						ctx,
						types.NamespacedName{Namespace: "default", Name: secretName + "-" + testKey},
					)

					Expect(err).To(BeNil())

					expectedReturnSecretKey := "default-" + secretName + "-" + testKey

					Expect(testValue).To(Equal(returnedValue[expectedReturnSecretKey]))
				}
			})

			datanew := map[string][]byte{
				"french": []byte("fries"),
				"hot":    []byte("dogs"),
			}

			Context("upserting the flattened secret to make sure it can be overwritten", func() {
				err = client.Upsert(ctx, key, datanew, secrets.Flatten(true))
				Expect(err).To(BeNil())
			})

			Context("ensuring updated flattened secret exists using keyvault client", func() {
				// Look for each originally passed secret item in the keyvault
				for testKey, testValue := range datanew {
					returnedValue, err := client.Get(
						ctx,
						types.NamespacedName{Namespace: "default", Name: secretName + "-" + testKey},
					)

					Expect(err).To(BeNil())

					expectedReturnSecretKey := "default-" + secretName + "-" + testKey

					Expect(testValue).To(Equal(returnedValue[expectedReturnSecretKey]))
				}
			})

			Context("delete flattened secrets and ensure they're gone", func() {
				for testKey, _ := range datanew {
					err := client.Delete(
						ctx,
						types.NamespacedName{Namespace: "default", Name: secretName + "-" + testKey},
					)

					Expect(err).To(BeNil())

					_, err = client.Get(ctx, key)
					Expect(err).ToNot(BeNil())
				}
			})
		})
	})
})
