package handler

import (
	"log"
	"net/http"

	"github.com/Sharykhin/it-customer-review/util"
)

func pong(w http.ResponseWriter, r *http.Request) {
	err := util.JSON(util.Response{
		Success: true,
		Data:    "pong",
		Error:   util.ErrorField{Err: nil},
		Meta:    nil,
	}, w, http.StatusOK)
	if err != nil {
		log.Printf("could not return response, struct: %v, error: %v", r, err)
	}
}
