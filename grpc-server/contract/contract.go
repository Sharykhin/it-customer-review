package contract

import (
	"context"

	"github.com/Sharykhin/it-customer-review/grpc-server/entity"
)

type (
	// StorageProvider is an interface that describe method for some sort of data storage
	StorageProvider interface {
		Create(ctx context.Context, vr *entity.Review) (*entity.Review, error)
		Update(ctx context.Context, ru entity.ReviewUpdate, r *entity.ReviewM) (*entity.ReviewM, error)
		GetById(ctx context.Context, ID string) (*entity.ReviewM, error)
	}
)
