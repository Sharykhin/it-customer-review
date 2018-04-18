package controller

import (
	"context"

	"github.com/Sharykhin/it-customer-review/contract"
	"github.com/Sharykhin/it-customer-review/database/mysql"
	"github.com/Sharykhin/it-customer-review/entity"
)

type (
	reviewCtrl struct {
		storage contract.StorageProvider
	}
)

// ReviewCtrl is a review controller that would handle reviews requests
var ReviewCtrl = reviewCtrl{storage: mysql.Storage}

func (rc reviewCtrl) Create(ctx context.Context, rr entity.ReviewRequest) (*entity.Review, error) {
	review := entity.NewReview()
	review.Name = rr.Name
	review.Email = rr.Email
	review.Content = rr.Content

	return rc.storage.Create(ctx, review)

}
