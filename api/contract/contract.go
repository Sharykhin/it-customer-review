package contract

import (
	"context"

	"github.com/Sharykhin/it-customer-review/api/entity"
)

type (
	// StorageProvider is an interface that describe method for some sort of data storage
	StorageProvider interface {
		Create(ctx context.Context, rr entity.ReviewRequest) (*entity.Review, error)
	}
)
