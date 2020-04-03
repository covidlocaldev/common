package keys

type KeyType interface {
	GetKeyFilePrefix()	string
}

type TwilioKey struct {
	KeyType
	AccountSID        	string
	AuthToken         	string
	APIKeySID         	string
	APIKeySecret      	string
	PushCredentialSID 	string
	TwimlApplicationSID string
}

func (k *TwilioKey) GetKeyFilePrefix() string {
	return "twilio"
}

type InternalServiceKey struct {
	KeyType
	ServiceToken	string
	ServiceSecret	string
}

func (k *InternalServiceKey) GetKeyFilePrefix() string {
	return "internal"
}

type AdminKey struct {
	KeyType
	AdminTokenIssuer string
	AdminTokenSecret string
}

func (k *AdminKey) GetKeyFilePrefix() string {
	return "admin"
}

type GoogleOAuthKey struct {
	KeyType
	DoctorClientID		string
	DoctorClientSecret	string
	PatientClientID		string
	PatientClientSecret	string
}

func (k *GoogleOAuthKey) GetKeyFilePrefix() string {
	return "google-oauth"
}

type GoogleApiKey struct {
	GoogleApiServerKey string
}

func (k *GoogleApiKey) GetKeyFilePrefix() string {
	return "google"
}
