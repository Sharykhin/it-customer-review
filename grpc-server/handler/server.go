package handler

import (
	"context"

	"fmt"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/Sharykhin/it-customer-review/grpc-server/contract"
	"github.com/Sharykhin/it-customer-review/grpc-server/entity"
	"github.com/Sharykhin/it-customer-review/grpc-server/service"
)

type server struct {
	storage contract.StorageProvider
}

func (s server) Create(ctx context.Context, request *pb.ReviewRequest) (*pb.ReviewResponse, error) {

	r := entity.NewReview()
	r.ReviewRequest = request
	r.Score = entity.Score(request.Score)

	if err := r.Validate(); err != nil {
		return nil, err
	}

	r, err := s.storage.Create(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("storage could not create a new review: %v", err)
	}
	res, err := service.ConvertReviewToResponse(r)
	if err != nil {
		return nil, fmt.Errorf("could not convert reiew entity to response: %v", err)
	}

	return res, nil
}

func (s server) Ping(ctx context.Context, in *pb.Empty) (*pb.Pong, error) {
	return &pb.Pong{Response: "pong"}, nil
}

func (s server) Update(ctx context.Context, request *pb.ReviewUpdateRequest) (*pb.ReviewResponse, error) {
	//review := new(pb.ReviewResponse)
	//review.ID = "asd"
	//review.Name = request.Name
	//review.Email = request.Email
	//review.Content = request.Content
	//review.Published = request.Published
	//review.Score = int64(request.Score)
	//review.Category = request.Category
	//review.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05")
	return nil, nil
}
