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

func (s server) Update(ctx context.Context, in *pb.ReviewUpdateRequest) (*pb.ReviewResponse, error) {

	ru := entity.ReviewUpdate{ReviewUpdateRequest: in}

	if err := ru.Validate(); err != nil {
		return nil, fmt.Errorf("validate error on update: %v", err)
	}

	review, err := s.storage.GetById(ctx, in.ID)
	if err != nil {
		return nil, fmt.Errorf("could not get a review by id %s: %v", in.ID, err)
	}
	if review == nil {
		return nil, fmt.Errorf("review with ID %s does not exist", in.ID)
	}

	review, err = s.storage.Update(ctx, ru, review)
	if err != nil {
		return nil, fmt.Errorf("storage could not update a review: %v", err)
	}

	res, err := service.ConvertReviewMToResponse(review)
	if err != nil {
		return nil, fmt.Errorf("could not convert reiew entity to response: %v", err)
	}

	return res, nil
}
