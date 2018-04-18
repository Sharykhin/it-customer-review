package handler

import (
	"net/http"

	"github.com/Sharykhin/it-customer-review/util"
)

func pong(w http.ResponseWriter, r *http.Request) {
	util.JSON(util.Response{
		Success: true,
		Data:    "pong",
		Error:   util.ErrorField{Err: nil},
		Meta:    nil,
	}, w, http.StatusOK)
}
