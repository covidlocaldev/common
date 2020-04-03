package utils

import (
	"encoding/json"
	"github.com/covidlocaldev/common/logging"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadBody(body io.Reader, o interface{}) (err error) {
	bodyData, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("error reading body: %s", err)
		return
	}

	err = json.Unmarshal(bodyData, o)
	if err != nil {
		log.Printf("error unmarshaling payload: %s", err)
	}

	return
}

func WriteJSONResponse(w http.ResponseWriter, response interface{}) {
	js, err := json.Marshal(response)
	if err != nil {
		logging.Error("could not marshal CreateChatConsultationResponse: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}