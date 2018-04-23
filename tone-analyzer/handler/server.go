package handler

import (
	"fmt"
	"log"

	"github.com/Sharykhin/it-customer-review/tone-analyzer/queue/rabbitmq"
)

// ListenAndServe listens income messages for a specific queue
func ListenAndServe() error {
	msgs, err := rabbitmq.RabbitMQ.Listen()
	if err != nil {
		return err
	}
	var done chan struct{}

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)
			err := handle(d.Body)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	<-done
	return nil
}
