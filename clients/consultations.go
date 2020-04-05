package clients

import (
	"cloud.google.com/go/firestore"
	"github.com/covidlocaldev/common/emr/resources"
	"log"
	"net/http"
	"os"
	"time"
)

type ConsultationsClient struct {
	HttpServiceClient
}

func NewConsultationsClient() *ConsultationsClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	log.Printf("service host: %s", serviceHost)

	return &ConsultationsClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateConsultationRequest struct {
	PatientUserUID      string             `json:"patientUserUID"`
	PractitionerUserUID string             `json:"practitionerUserUID"`
	ConsultationType    string             `json:"consultationType"`   // see consultations.types for enum values
	ConsultationStatus  string             `json:"consultationStatus"` // see consultations.status for enum values
	ScheduledStart      time.Time          `json:"requestedStart"`
	ScheduledEnd        time.Time          `json:"requestedEnd"`
	Reasons             []resources.Coding `json:"reasons"`
}

type CreateConsultationResponse struct {
	Chat         *firestore.DocumentRef `json:"chatRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

func (client *ConsultationsClient) CreateConsultation(req CreateConsultationRequest) (*http.Response, error) {
	return client.Post("/consultations/internal/consultations", req)
}

type DeliverChatConsultationKeyRequest struct {
	RecipientUID string `json:"recipientUID"`
}

type StartChatConsultationRequest struct {
	Practitioner *firestore.DocumentRef `json:"practitionerRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

type StartChatConsultationResponse struct {
	Chat         *firestore.DocumentRef `json:"chatRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

func (client *ConsultationsClient) StartChatConsultation(req StartChatConsultationRequest) (*http.Response, error) {
	return client.Post("/consultations/internal/consultations/start-chat", req)
}
