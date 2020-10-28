// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package singularity

type SingularityManagers struct {
	SingularityAccount SingularityAccountManager
	SingularityJob     SingularityJobManager
	SingularityStorage SingularityStorageManager
}

var AzureSingularityManagers = SingularityManagers{
	SingularityAccount: &azureSingularityManager{},
	SingularityJob:     &azureSingularityJobManager{},
	SingularityStorage: &azureSingularityStorageManager{},
}
