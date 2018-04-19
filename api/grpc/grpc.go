package grpc

import (
	"context"
	"log"
	"os"

	"github.com/Sharykhin/it-customer-review/api/entity"
	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"google.golang.org/grpc"
)

type (
	reviewService struct {
		client pb.ReviewClient
	}
)

// ReviewService is a grpc service that would be responsible for managing reviews.
var ReviewService reviewService

func (ctrl reviewService) Create(ctx context.Context, rr entity.ReviewRequest) (*pb.ReviewResponse, error) {

	response, err := ctrl.client.Create(ctx, &pb.ReviewRequest{
		Name:    rr.Name,
		Email:   rr.Email,
		Content: rr.Content,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func init() {
	conn, err := grpc.Dial(os.Getenv("GRPC_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connet to a grpc server: %v", err)
	}
	client := pb.NewReviewClient(conn)

	if _, err := client.Ping(context.Background(), &pb.Empty{}); err != nil {
		log.Fatalf("Could not ping a grpc server: %v", err)
	}
	ReviewService = reviewService{client: client}
}
