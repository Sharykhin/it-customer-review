package rabbitmq

import (
	"fmt"
	"log"
	"os"

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

func (r rabbitMQ) Listen() (<-chan []byte, error) {
	msgs, err := r.ch.Consume(
		r.q.Name, // queue
		"",       // consumer
		true,     // auto-ack
		true,     // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)

	if err != nil {
		return nil, fmt.Errorf("could not consume a queue %s: %v", queueName, err)
	}

	ch := createMessageChannel(msgs)

	return ch, nil
}

func createMessageChannel(msgs <-chan amqp.Delivery) <-chan []byte {
	ch := make(chan []byte)
	go func() {
		defer close(ch)
		for d := range msgs {
			ch <- d.Body
		}
	}()
	return ch
}

func listenClose(notify chan *amqp.Error, ch *amqp.Channel, conn *amqp.Connection) {
	err := <-notify
	util.Check(ch.Close)
	util.Check(conn.Close)
	log.Fatalf("rabbitmq connection was broken: %v", err)
}
