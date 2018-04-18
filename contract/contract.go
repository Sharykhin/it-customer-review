package contract

import (
	"context"

	"github.com/Sharykhin/it-customer-review/entity"
)

type (
	// StorageProvider is an interface that describe method for some sort of data storage
	StorageProvider interface {
		Create(ctx context.Context, r *entity.Review) (*entity.Review, error)
	}
)
