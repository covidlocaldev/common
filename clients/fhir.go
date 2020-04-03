package clients

import (
	"fmt"
	"github.com/covidlocaldev/common/fhir/resources"
	"log"
	"net/http"
	"os"
)

type FHIRClient struct {
	HttpServiceClient
}

func NewFHIRClient() *FHIRClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	log.Printf("service host: %s", serviceHost)

	return &FHIRClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateResourceRequest struct {
	Resource resources.Resource
}

type CreateResourceResponse struct {
	Resource resources.Resource
}

func (client *FHIRClient) CreateResource(req CreateResourceRequest) (resp *http.Response, err error) {
	return client.Post(fmt.Sprintf("/fhir/resources/%s", req.Resource.GetResourceType()), req.Resource)
}
