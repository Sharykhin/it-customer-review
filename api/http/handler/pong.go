package handler

import (
	"net/http"

	"github.com/Sharykhin/it-customer-review/api/util"
)

// Pong is a simple live health method
func Pong(w http.ResponseWriter, r *http.Request) {
	util.JSON(util.Response{
		Success: true,
		Data:    "pong",
		Error:   util.ErrorField{Err: nil},
		Meta:    nil,
	}, w, http.StatusOK)
}
