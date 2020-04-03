package clients

import (
	"github.com/covidlocaldev/common/keys"
	"github.com/hellodoctordev/gotwilio"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

func NewTwilioClient() *gotwilio.Twilio {
	return gotwilio.NewTwilioClient(keys.TwilioKeys.AccountSID, keys.TwilioKeys.AuthToken)
}

func NewApplicationAccessToken(twilio *gotwilio.Twilio, identity string) *gotwilio.AccessToken {
	token := twilio.NewAccessToken()
	token.APIKeySid = keys.TwilioKeys.APIKeySID
	token.APIKeySecret = keys.TwilioKeys.APIKeySecret
	token.ExpiresAt = time.Now().Add(time.Hour * 4)
	token.NotBefore = time.Now()
	token.Identity = identity

	return token
}

func UnmarshalTwilioRequestBody(r *http.Request, o interface{}) (err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}

	params, _ := url.ParseQuery(string(body))

	v := reflect.ValueOf(o).Elem()
	typeOfO := v.Type()

	for i := 0; i < v.NumField(); i++ {
		typeField := typeOfO.Field(i)

		value := params.Get(typeField.Name)

		v.FieldByName(typeField.Name).SetString(value)
	}

	return
}
