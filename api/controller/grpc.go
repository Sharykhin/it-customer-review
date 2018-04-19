package controller

import (
	"context"
	"log"
	"os"

	"time"

	"github.com/Sharykhin/it-customer-review/api/contract"
	"github.com/Sharykhin/it-customer-review/api/entity"
	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"google.golang.org/grpc"
)

type (
	reviewCtrl struct {
		client pb.ReviewClient
	}
)

// ReviewCtrl is a review controller that handles all necessary requests.
// It implements contract.StorageProvider interface so the packages that may use it
// would be aware of that it provides
var ReviewCtrl contract.StorageProvider

func (ctrl reviewCtrl) Create(ctx context.Context, rr entity.ReviewRequest) (*entity.Review, error) {
	review := entity.NewReview()

	response, err := ctrl.client.Create(ctx, &pb.ReviewRequest{
		ID:        review.ID,
		Name:      rr.Name,
		Email:     rr.Email,
		Content:   rr.Content,
		Published: review.Published,
	})

	if err != nil {
		return nil, err
	}

	review.Content = response.Content
	review.Name = response.Name
	review.Email = response.Email
	review.Category = entity.NullString(response.Category)
	t, err := time.Parse(entity.JSONTimeFormat, response.CreatedAt)
	if err != nil {
		log.Printf("could not parse time: %v", err)
	}
	review.CreatedAt = entity.JSONTime(t)
	return review, nil
}

func init() {
	conn, err := grpc.Dial(os.Getenv("GRPC_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connet to a grpc server: %v", err)
	}
	//TODO: we need to close the connection
	client := pb.NewReviewClient(conn)
	ReviewCtrl = reviewCtrl{client: client}
}
