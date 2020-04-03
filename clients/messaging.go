package clients

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
	"os"
)

type MessagingClient struct {
	HttpServiceClient
}

func NewMessagingClient() *MessagingClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	return &MessagingClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateChatRequest struct {
	PatientUserUID      string `json:"patientUserUID"`
	PractitionerUserUID string `json:"practitionerUserUID"`
}

type CreateChatResponse struct {
	ChatID string `json:"chatID"`
}

func (client *MessagingClient) CreateChat(patientUID, practitionerUserUID string) (*http.Response, error) {
	req := CreateChatRequest{
		PatientUserUID:      patientUID,
		PractitionerUserUID: practitionerUserUID,
	}

	return client.Post("/messaging/internal/chats", req)
}

type SendAutomatedChatMessageRequest struct {
	ConsultationID string `json:"consultationID"`
	MessageContent string `json:"messageContent"`
	ContentType    string `json:"contentType"`
}

func (client *MessagingClient) SendAutomatedChatMessage(chatID, consultationID, messageContent, contentType, messageType string) (*http.Response, error) {
	req := SendAutomatedChatMessageRequest{
		ConsultationID: consultationID,
		MessageContent: messageContent,
		ContentType:    contentType,
	}

	return client.Post(fmt.Sprintf("/messaging/internal/chats/%s", chatID), req)
}

type SendChatMessageRequest struct {
	ChatID         string `json:"chatID"`
	ConsultationID string `json:"consultationID"`
	SenderID       string `json:"senderID"`
	Message        string `json:"message"`
	ContentType    string `json:"contentType"`
}

func (client *MessagingClient) SendChatMessage(chatID, consultationID, senderID, message, contentType, messageType string) (*http.Response, error) {
	req := SendChatMessageRequest{
		SenderID:       senderID,
		ConsultationID: consultationID,
		Message:        message,
		ContentType:    contentType,
	}

	return client.Post(fmt.Sprintf("/messaging/chats/%s/messages", chatID), req)
}

type SendEmailRequest struct {
	From             *mail.Email `json:"from"`
	To               *mail.Email `json:"to"`
	Subject          string      `json:"subject"`
	PlainTextContent string      `json:"plainTextContent"`
	HTMLContent      string      `json:"htmlContent"`
}

func (client *MessagingClient) SendEmail(req SendEmailRequest) (*http.Response, error) {
	return client.Post("/messaging/emails", req)
}

type SendTemplateEmailRequest struct {
	From       *mail.Email            `json:"from"`
	To         *mail.Email            `json:"to"`
	Subject    string                 `json:"subject"`
	TemplateID string                 `json:"templateID"`
	Args       map[string]interface{} `json:"args"`
}

func (client *MessagingClient) SendTemplateEmail(req SendTemplateEmailRequest) (*http.Response, error) {
	return client.Post(fmt.Sprintf("/messaging/emails/%s", req.TemplateID), req)
}
