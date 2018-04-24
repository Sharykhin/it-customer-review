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

// Manager is a general queue manager that returns a concrete implementation
var Manager = manager{MessageProvider: rabbitmq.RabbitMQ}
