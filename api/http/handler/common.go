package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fmt"

	"github.com/Sharykhin/it-customer-review/api/entity"
	"github.com/Sharykhin/it-customer-review/api/queue"
)

func publishAnalyzeJob(ID string, content string) error {
	res, err := json.Marshal(entity.AnalyzeMessage{
		ID:      ID,
		Content: content,
	})

	if err != nil {
		return fmt.Errorf("could not marshal analyze message struct: %v", err)
	}
	res, err = json.Marshal(entity.QueueMessage{
		Action:  "analyze tone",
		Payload: res,
	})

	if err != nil {
		return fmt.Errorf("could not marshal queue message struct: %v", err)
	}

	err = queue.Manager.Publish(res)
	if err != nil {
		return fmt.Errorf("could not publish a message %s: %v", res, err)
	}
	return nil
}

func queryParamInt(r *http.Request, key string, defaultValue int64) (int64, error) {
	v := r.FormValue(key)

	if v == "" {
		return defaultValue, nil
	}

	return strconv.ParseInt(v, 10, 64)
}

func queryCriteria(r *http.Request, key string) *entity.Criteria {
	v := r.FormValue(key)
	if v == "" {
		return nil
	}
	return &entity.Criteria{Key: key, Value: v}
}
