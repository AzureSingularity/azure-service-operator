module github.com/Azure/azure-service-operator

go 1.13

replace github.com/Azure/azure-service-operator/pkg/resourcemanager/singularity/microsoftazuremanagementaisupercomputer => ./pkg/resourcemanager/singularity/microsoftazuremanagementaisupercomputer

require (
	github.com/Azure/azure-sdk-for-go v44.0.0+incompatible
	github.com/Azure/azure-service-operator/pkg/resourcemanager/singularity/microsoftazuremanagementaisupercomputer v0.0.0-00010101000000-000000000000
	github.com/Azure/go-autorest/autorest v0.11.10
	github.com/Azure/go-autorest/autorest/adal v0.9.5
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.0
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Azure/go-autorest/autorest/validation v0.3.0
	github.com/Azure/go-autorest/tracing v0.6.0
	github.com/denisenkom/go-mssqldb v0.0.0-20200206145737-bbfc9a55622e
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-logr/logr v0.1.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gobuffalo/envy v1.7.0
	github.com/google/go-cmp v0.3.0
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/google/uuid v1.1.1
	github.com/hashicorp/go-multierror v1.0.0
	github.com/lib/pq v1.6.0
	github.com/marstr/randname v0.0.0-20181206212954-d5b0f288ab8c
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.3.0 // indirect
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/prometheus/client_golang v1.0.0
	github.com/satori/go.uuid v1.2.0
	github.com/sethvargo/go-password v0.1.2
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.5.1
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	golang.org/x/tools v0.0.0-20190920225731-5eefd052ad72
	gopkg.in/ini.v1 v1.55.0 // indirect
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2
	sigs.k8s.io/controller-runtime v0.5.0
	sigs.k8s.io/controller-tools v0.2.5 // indirect
)
