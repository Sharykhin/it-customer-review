package handler

import (
	"net/http"

	"encoding/json"

	"log"

	"github.com/Sharykhin/it-customer-review/api/entity"
	"github.com/Sharykhin/it-customer-review/api/grpc"
	"github.com/Sharykhin/it-customer-review/api/queue"
	"github.com/Sharykhin/it-customer-review/api/util"
	"github.com/pkg/errors"
)

// create creates a new review in a database
func create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer util.Check(r.Body.Close)
	var rr entity.ReviewRequest
	if err := decoder.Decode(&rr); err != nil {
		util.JSONBadRequest(errors.New("please provide a valid json"), w)
		return
	}

	if err := rr.Validate(); err != nil {
		util.JSONBadRequest(err, w)
		return
	}

	review, err := grpc.ReviewService.Create(r.Context(), rr)

	if err != nil {
		log.Printf("could not create a new review: %v", err)
		util.JSON(util.Response{
			Success: false,
			Data:    nil,
			Error:   util.ErrorField{Err: errors.New(http.StatusText(http.StatusInternalServerError))},
			Meta:    nil,
		}, w, http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(entity.AnalyzeMessage{
		ID:      review.ID,
		Content: review.Content,
	})

	if err != nil {
		log.Printf("could not marshal analyze message struct: %v", err)
	}
	res, err = json.Marshal(entity.QueueMessage{
		Action:  "analyze tone",
		Payload: res,
	})

	if err != nil {
		log.Printf("could not marshal queue message struct: %v", err)
	}

	err = queue.Manager.Publish(res)
	if err != nil {
		log.Printf("could not publish a message %s: %v", res, err)
	}

	util.JSON(util.Response{
		Success: true,
		Data:    review,
		Error:   util.ErrorField{Err: nil},
		Meta:    nil,
	}, w, http.StatusCreated)
}
