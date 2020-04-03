package clients

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/covidlocaldev/common/utils"
	"net/http"
	"os"
	"time"
)

type SchedulingClient struct {
	HttpServiceClient
}

func NewSchedulingClient() *SchedulingClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	return &SchedulingClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type InternalCreateNewEventRequest struct {
	UserUID          string                   `json:"userUID"`
	Title            string                   `json:"title"`
	Start            time.Time                `json:"start"`
	End              time.Time                `json:"end,omitempty"`
	Availability     string                   `json:"availability"`
	EventType        string                   `json:"eventType"`
	IsAllDay         bool                     `json:"isAllDay"`
	Participants     []*firestore.DocumentRef `json:"participants"`
	Description      string                   `json:"description"`
	Consultation     *firestore.DocumentRef   `json:"consultationRef"`
	ConsultationType string                   `json:"consultationType"`
}

type InternalCreateNewEventResponse struct {
	Event *firestore.DocumentRef `json:"eventRef"`
}

func (client *SchedulingClient) InternalCreateNewEvent(req InternalCreateNewEventRequest) (res InternalCreateNewEventResponse, err error) {
	r, err := client.Post("/scheduling/internal/events", req)
	if err == nil {
		err = utils.ReadBody(r.Body, &res)
	}

	return
}

type InternalUpdateEventRequest struct {
	UserUID      string     `json:"userUID"`
	EventID      string     `json:"eventID"`
	Start        *time.Time `json:"start"`
	End          *time.Time `json:"end"`
	Availability string     `json:"availability"`
}

type InternalUpdateEventResponse struct {
	Event *firestore.DocumentRef `json:"eventRef"`
}

func (client *SchedulingClient) InternalUpdateEvent(req InternalUpdateEventRequest) (res InternalCreateNewEventResponse, err error) {
	r, err := client.Put(fmt.Sprintf("/scheduling/internal/events/%s", req.EventID), req)
	if err == nil {
		err = utils.ReadBody(r.Body, &res)
	}

	return
}

type InternalDeleteEventRequest struct {
	UserUID string `json:"userUID"`
	EventID string `json:"eventID"`
}

func (client *SchedulingClient) InternalDeleteEvent(req InternalUpdateEventRequest) error {
	_, err := client.Delete(fmt.Sprintf("/scheduling/internal/events/%s", req.EventID), req)

	return err
}
