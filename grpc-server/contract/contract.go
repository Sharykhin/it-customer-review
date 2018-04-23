package contract

import (
	"context"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/Sharykhin/it-customer-review/grpc-server/entity"
)

type (
	//ReviewManager contains all method for managing reviews
	ReviewManager interface {
		StorageProvider
		ReviewRepositoryProvider
	}
	// StorageProvider is an interface that describe method for some sort of data storage
	StorageProvider interface {
		Create(ctx context.Context, vr *entity.Review) (*entity.Review, error)
		Update(ctx context.Context, ru entity.ReviewUpdate, r *entity.Review) (*entity.Review, error)
	}
	// ReviewRepositoryProvider describes methods for getting reviews by different criteria
	ReviewRepositoryProvider interface {
		GetByID(ctx context.Context, ID string) (*entity.Review, error)
		Count(ctx context.Context, criteria []*pb.Criteria) (int64, error)
		GetList(ctx context.Context, criteria []*pb.Criteria, limit, offset int64) ([]entity.Review, error)
	}
)
