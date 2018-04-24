package util

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// Response struct represents base response format
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   ErrorField  `json:"error"`
	Meta    interface{} `json:"meta"`
}

// ErrorField provides nullable error type
type ErrorField struct {
	Err error
}

var appEnv = os.Getenv("APP_ENV")

// MarshalJSON implements Marshaler interface to return nil in case there was no error
func (ef ErrorField) MarshalJSON() ([]byte, error) {
	if ef.Err != nil {
		return json.Marshal(ef.Err.Error())
	}
	return json.Marshal(nil)
}

// JSON sends a json response to a client
func JSON(r Response, w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Fatalf("could not sent a response to a client, error: %v. Struct: %v", err, r)
	}
}

// JSONError is a helper that wraps 500 error response
func JSONError(err error, w http.ResponseWriter) {
	if appEnv == "prod" {
		err = errors.New(http.StatusText(http.StatusInternalServerError))
	}
	JSON(Response{
		Success: false,
		Data:    nil,
		Error:   ErrorField{Err: err},
		Meta:    nil,
	}, w, http.StatusInternalServerError)
}

// JSONBadRequest is a helper func that returns 400 (bad request) response
func JSONBadRequest(err error, w http.ResponseWriter) {
	JSON(Response{
		Success: false,
		Data:    nil,
		Error:   ErrorField{Err: err},
		Meta:    nil,
	}, w, http.StatusBadRequest)
}
