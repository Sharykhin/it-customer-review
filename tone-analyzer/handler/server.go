package handler

import (
	"fmt"

	"github.com/Sharykhin/it-customer-review/tone-analyzer/logger"
	"github.com/Sharykhin/it-customer-review/tone-analyzer/queue"
)

// ListenAndServe listens income messages for a specific queue
func ListenAndServe() error {
	msgs, err := queue.Manager.Listen()
	if err != nil {
		return err
	}
	var done chan struct{}

	go func() {
		for msg := range msgs {
			fmt.Printf("Received a message: %s\n", msg)
			err := handle(msg)
			if err != nil {
				logger.Logger.Errorf("could not handle income message: %s, error: %v", msg, err)
			}
		}
	}()
	<-done
	return nil
}
