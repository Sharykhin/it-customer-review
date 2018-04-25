package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Sharykhin/it-customer-review/api/entity"
	"github.com/Sharykhin/it-customer-review/api/grpc"
	"github.com/Sharykhin/it-customer-review/api/logger"
	"github.com/Sharykhin/it-customer-review/api/session"
	"github.com/Sharykhin/it-customer-review/api/util"
	"github.com/pkg/errors"
)

// Create creates a new review
func Create(w http.ResponseWriter, r *http.Request) {

	sess, err := session.Store.Get(r, "it-customer-review")
	if err != nil {
		logger.Logger.Errorf("could not get cookie from a request: %v", err)
		util.JSONError(err, w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer util.Check(r.Body.Close)
	var rr entity.ReviewRequest
	if err = decoder.Decode(&rr); err != nil {
		util.JSONBadRequest(errors.New("please provide a valid json"), w)
		return
	}

	if err = rr.Validate(); err != nil {
		util.JSONBadRequest(err, w)
		return
	}

	review, err := grpc.ReviewService.Create(r.Context(), rr)

	if err != nil {
		logger.Logger.Errorf("could not create a new review, request: %v, error: %v", rr, err)
		util.JSONError(err, w)
		return
	}

	err = publishAnalyzeJob(review.ID, review.Content)
	if err != nil {
		logger.Logger.Errorf("could not dispatch analyzer job: %v", err)
	}

	sess.Values[review.ID] = true
	err = sess.Save(r, w)
	if err != nil {
		logger.Logger.Errorf("could not write a session value: %v", err)
	}

	util.JSON(util.Response{
		Success: true,
		Data:    review,
		Error:   util.ErrorField{Err: nil},
		Meta:    nil,
	}, w, http.StatusCreated)
}
