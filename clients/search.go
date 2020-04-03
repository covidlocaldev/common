package clients

import (
	"net/http"
	"os"
)

type SearchClient struct {
	HttpServiceClient
}

func NewSearchClient() *SearchClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	return &SearchClient{
		HttpServiceClient{
			Client: http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}
