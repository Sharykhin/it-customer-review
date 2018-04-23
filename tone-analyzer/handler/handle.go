package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"context"

	"github.com/Sharykhin/it-customer-review/tone-analyzer/analyzer"
	"github.com/Sharykhin/it-customer-review/tone-analyzer/entity"
	"github.com/Sharykhin/it-customer-review/tone-analyzer/grpc"
)

const (
	analyzeToneAction = "analyze tone"
)

func handle(msg []byte) error {
	qm := entity.QueueMessage{}
	err := json.Unmarshal(msg, &qm)
	if err != nil {
		return fmt.Errorf("could not parse income message: %v", err)
	}

	switch qm.Action {
	case analyzeToneAction:
		go analyze(qm.Payload)
	default:
		fmt.Println("There is no action")
	}
	return nil
}

func analyze(p []byte) {
	var am entity.AnalyzeMessage
	err := json.Unmarshal(p, &am)
	if err != nil {
		log.Printf("could not unmarshal analyze message struct: %v", err)
	}
	s, err := analyzer.Analyzer.Analyze(am.Content)
	if err != nil {
		log.Fatal(err)
	}
	r := entity.ReviewRequestUpdate{
		Score: s,
	}

	if s >= 50 {
		r.Category = "positive"
	} else {
		r.Category = "negative"
	}

	err = grpc.ReviewService.Update(context.Background(), am.ID, r)
	if err != nil {
		log.Printf("could not update a review: %v", err)
	}
}
