package handler

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/Sharykhin/it-customer-review/api/entity"
	"github.com/Sharykhin/it-customer-review/api/grpc"
	"github.com/Sharykhin/it-customer-review/api/logger"
	"github.com/Sharykhin/it-customer-review/api/session"
	"github.com/Sharykhin/it-customer-review/api/util"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Update updates an existing review
func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sess, err := session.Store.Get(r, "it-customer-review")
	if err != nil {
		logger.Logger.Errorf("could not get cookie from a request: %v", err)
		util.JSONError(err, w)
		return
	}

	if access, ok := sess.Values[id].(bool); !ok || !access {
		util.JSON(util.Response{
			Success: false,
			Data:    nil,
			Error:   util.ErrorField{Err: errors.New(http.StatusText(http.StatusForbidden))},
			Meta:    nil,
		}, w, http.StatusForbidden)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer util.Check(r.Body.Close)
	var rr entity.ReviewUpdateRequest
	if err = decoder.Decode(&rr); err != nil {
		util.JSONBadRequest(errors.New("please provide a valid json"), w)
		return
	}

	if err = rr.Validate(); err != nil {
		util.JSONBadRequest(err, w)
		return
	}

	review, err := grpc.ReviewService.Update(r.Context(), id, rr)

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
		logger.Logger.Errorf("could not update a review, request: %v, error: %v", rr, err.Err())
		util.JSONError(err.Err(), w)
		return
	}

	if rr.Content.Valid {
		err := publishAnalyzeJob(review.ID, review.Content)
		if err != nil {
			log.Printf("could not dispatch analyzer job: %v", err)
		}
	}

	util.JSON(util.Response{
		Success: true,
		Data:    review,
		Error:   util.ErrorField{Err: nil},
		Meta:    nil,
	}, w, http.StatusOK)
}
