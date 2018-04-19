package util

import (
	"encoding/json"
	"log"
	"net/http"
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

// JSONBadRequest is a helper func that returns 400 (bad request) response
func JSONBadRequest(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	r := Response{
		Success: false,
		Data:    nil,
		Error:   ErrorField{Err: err},
		Meta:    nil,
	}
	w.WriteHeader(http.StatusBadRequest)
	err = json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Fatalf("could not sent a response to a client, error: %v. Struct: %v", err, r)
	}
}
