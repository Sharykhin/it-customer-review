package grpc

import (
	"context"
	"log"
	"os"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/Sharykhin/it-customer-review/tone-analyzer/entity"
	"google.golang.org/grpc"
)

type (
	reviewService struct {
		client pb.ReviewClient
	}
)

// ReviewService is a grpc service that would be responsible for managing reviews.
var ReviewService reviewService

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

func (ctrl reviewService) Update(ctx context.Context, ID string, r entity.ReviewRequestUpdate) error {
	var fu pb.FieldToUpdate

	fu.Name = &pb.FieldToUpdate_NameNull{NameNull: true}
	fu.Content = &pb.FieldToUpdate_ContentNull{ContentNull: true}
	fu.Score = &pb.FieldToUpdate_ScoreValue{ScoreValue: r.Score}
	fu.Category = &pb.FieldToUpdate_CategoryValue{CategoryValue: r.Category}
	fu.Published = &pb.FieldToUpdate_PublishedNull{PublishedNull: true}

	in := pb.ReviewUpdateRequest{
		ID:             ID,
		FieldsToUpdate: &fu,
	}

	_, err := ctrl.client.Update(ctx, &in)
	if err != nil {
		return err
	}
	return nil
}
