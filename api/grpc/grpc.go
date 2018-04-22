package grpc

import (
	"context"
	"log"
	"os"

	"fmt"

	"github.com/Sharykhin/it-customer-review/api/contract"
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
var ReviewService contract.ServiceProvider

func (ctrl reviewService) Create(ctx context.Context, rr entity.ReviewRequest) (*entity.Review, error) {

	r, err := ctrl.client.Create(ctx, &pb.ReviewRequest{
		Name:    rr.Name,
		Email:   rr.Email,
		Content: rr.Content,
		Score:   -1,
	})

	if err != nil {
		return nil, err
	}
	fmt.Println(r)
	return &entity.Review{
		ID:        r.ID,
		Name:      r.Name,
		Email:     r.Email,
		Content:   r.Content,
		Published: r.Published,
		Score:     entity.Score(r.Score),
		Category:  entity.NullString(r.Category),
		CreatedAt: r.CreatedAt,
	}, nil
}

func (ctrl reviewService) Update(ctx context.Context, ID string, rr entity.ReviewUpdateRequest) (*entity.Review, error) {
	in := pb.ReviewUpdateRequest{
		ID:    ID,
		Name:  rr.Name,
		Email: rr.Email,
	}

	if rr.Published.Valid {
		in.Published = &pb.ReviewUpdateRequest_PublishedValue{PublishedValue: rr.Published.Value}
	} else {
		in.Published = &pb.ReviewUpdateRequest_PublishedNull{PublishedNull: true}
	}
	r, err := ctrl.client.Update(ctx, &in)
	if err != nil {
		return nil, err
	}

	fmt.Println("BABA", r)
	return &entity.Review{
		ID:        r.ID,
		Name:      r.Name,
		Email:     r.Email,
		Content:   r.Content,
		Published: r.Published,
		Score:     entity.Score(r.Score),
		Category:  entity.NullString(r.Category),
		CreatedAt: r.CreatedAt,
	}, nil
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
