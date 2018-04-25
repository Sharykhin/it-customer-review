package rabbitmq

import (
	"log"
	"os"

	"fmt"

	"github.com/Sharykhin/it-customer-review/api/util"
	"github.com/streadway/amqp"
)

const (
	queueName = "analyze"
)

type rabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

// RabbitMQ is a reference to a private struct that implements Publish func
var RabbitMQ rabbitMQ

func init() {
	conn, err := amqp.Dial(os.Getenv("AMPQ_ADDRESS"))
	if err != nil {
		log.Fatalf("could not connect to rabbitmq: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not create a channel: %v", err)
	}

	notify := conn.NotifyClose(make(chan *amqp.Error))

	go listenClose(notify, ch, conn)

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		log.Fatalf("could not decale queue %s: %v", queueName, err)
	}
	RabbitMQ = rabbitMQ{conn: conn, ch: ch, q: q}
}

// Publish published messages into a queue
func (r rabbitMQ) Publish(body []byte) error {

	err := r.ch.Publish(
		"",       // exchange
		r.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("could not publish a message: %v", err)
	}
	return nil
}

func listenClose(notify chan *amqp.Error, ch *amqp.Channel, conn *amqp.Connection) {
	err := <-notify
	util.Check(ch.Close)
	util.Check(conn.Close)
	log.Fatalf("rabbitmq connection was broken: %v", err)
}
