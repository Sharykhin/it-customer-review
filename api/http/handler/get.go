package handler

import (
	"net/http"

	"github.com/Sharykhin/it-customer-review/api/grpc"
	"github.com/Sharykhin/it-customer-review/api/util"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Get returns a review by its ID
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	review, err := grpc.ReviewService.Get(r.Context(), id)

	if err != nil {
		err := status.Convert(err)
		if err.Code() == codes.NotFound {
			util.JSON(util.Response{
				Success: false,
				Data:    nil,
				Error:   util.ErrorField{Err: errors.New(err.Message())},
				Meta:    nil,
			}, w, http.StatusNotFound)
			return
		}

		util.JSON(util.Response{
			Success: false,
			Data:    nil,
			Error:   util.ErrorField{Err: errors.New(http.StatusText(http.StatusInternalServerError))},
			Meta:    nil,
		}, w, http.StatusInternalServerError)
		return
	}

	util.JSON(util.Response{
		Success: true,
		Data:    review,
		Error:   util.ErrorField{Err: nil},
		Meta:    nil,
	}, w, http.StatusOK)
}
