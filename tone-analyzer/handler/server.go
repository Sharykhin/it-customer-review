package handler

import (
	"fmt"
	"log"

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
				log.Println(err)
			}
		}
	}()
	<-done
	return nil
}
