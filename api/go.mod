module github.com/keptn/keptn/api

go 1.13

require (
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/go-openapi/errors v0.19.9
	github.com/go-openapi/loads v0.20.2
	github.com/go-openapi/runtime v0.19.24
	github.com/go-openapi/spec v0.20.3
	github.com/go-openapi/strfmt v0.20.0
	github.com/go-openapi/swag v0.19.14
	github.com/go-openapi/validate v0.20.2
	github.com/go-test/deep v1.0.7
	github.com/google/uuid v1.2.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/keptn/go-utils v0.8.1-0.20210308154435-bdd99f3d0e5f
	github.com/keptn/kubernetes-utils v0.8.1-0.20210308145316-785b480a4e52
	golang.org/x/net v0.0.0-20210224082022-3d97a244fca7
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible
