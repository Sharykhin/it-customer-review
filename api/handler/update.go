package handler

import (
	"net/http"

	"encoding/json"

	"log"

	"fmt"

	"github.com/Sharykhin/it-customer-review/api/entity"
	"github.com/Sharykhin/it-customer-review/api/grpc"
	"github.com/Sharykhin/it-customer-review/api/util"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	decoder := json.NewDecoder(r.Body)
	defer util.Check(r.Body.Close)
	var rr entity.ReviewUpdateRequest
	if err := decoder.Decode(&rr); err != nil {
		util.JSONBadRequest(errors.New("please provide a valid json"), w)
		return
	}

	review, err := grpc.ReviewService.Update(r.Context(), id, rr)
	fmt.Println("got review", review)
	if err != nil {
		log.Printf("could not update a review: %v", err)
		util.JSON(util.Response{
			Success: false,
			Data:    nil,
			Error:   util.ErrorField{},
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
