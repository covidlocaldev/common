package clients

import (
	"fmt"
	"github.com/covidlocaldev/common/emr/resources"
	"log"
	"net/http"
	"os"
)

type EMRClient struct {
	HttpServiceClient
}

func NewEMRClient() *EMRClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	log.Printf("service host: %s", serviceHost)

	return &EMRClient{
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

func (client *EMRClient) CreateResource(req CreateResourceRequest) (resp *http.Response, err error) {
	return client.Post(fmt.Sprintf("/emr/resources/%s", req.Resource.GetResourceType()), req.Resource)
}
