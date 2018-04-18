package util

import (
	"encoding/json"
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
func JSON(r Response, w http.ResponseWriter, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(r)
}
