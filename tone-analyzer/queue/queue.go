package queue

import (
	"github.com/Sharykhin/it-customer-review/tone-analyzer/contract"
	"github.com/Sharykhin/it-customer-review/tone-analyzer/queue/rabbitmq"
)

type (
	manager struct {
		contract.MessageProvider
	}
)

var Manager = manager{MessageProvider: rabbitmq.RabbitMQ}
