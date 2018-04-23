package queue

import (
	"github.com/Sharykhin/it-customer-review/api/contract"
	"github.com/Sharykhin/it-customer-review/api/queue/rabbitmq"
)

// manager is a private struct that include interface
// so concrete implementation must satisfy this interface
type manager struct {
	contract.QueueMessageProvider
}

// Manager is a concrete implementation of a queue manager
var Manager = manager{QueueMessageProvider: rabbitmq.RabbitMQ}
