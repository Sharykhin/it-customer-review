package contract

import (
	"context"

	"github.com/Sharykhin/it-customer-review/api/entity"
)

type (
	// ServiceProvider is an interface that describe method for some sort of data storage
	ServiceProvider interface {
		Create(ctx context.Context, rr entity.ReviewRequest) (*entity.Review, error)
		Update(ctx context.Context, ID string, rr entity.ReviewUpdateRequest) (*entity.Review, error)
		Get(ctx context.Context, ID string) (*entity.Review, error)
	}
	// QueueMessageProvider describes funcs that should be implemented to make queue work as expected.
	QueueMessageProvider interface {
		Publish(body []byte) error
	}
)
