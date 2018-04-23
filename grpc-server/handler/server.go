package handler

import (
	"context"

	"fmt"

	"database/sql"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/Sharykhin/it-customer-review/grpc-server/contract"
	"github.com/Sharykhin/it-customer-review/grpc-server/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	storage contract.ReviewManager
}

func (s server) Ping(ctx context.Context, in *pb.Empty) (*pb.Pong, error) {
	return &pb.Pong{Response: "pong"}, nil
}

func (s server) Create(ctx context.Context, in *pb.ReviewCreateRequest) (*pb.ReviewResponse, error) {

	r := entity.NewReview()
	r.Name = in.Name
	r.Email = in.Email
	r.Content = in.Content
	r.Published = in.Published

	if in.GetScoreNull() {
		r.Score = sql.NullInt64{}
	} else {
		r.Score = sql.NullInt64{Valid: true, Int64: in.GetScoreValue()}
	}
	if in.GetCategoryNull() {
		r.Category = sql.NullString{Valid: false, String: ""}
	} else {
		r.Category = sql.NullString{Valid: true, String: in.GetCategoryValue()}
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	r, err := s.storage.Create(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("storage could not create a new review: %v", err)
	}
	res, err := convert(r)
	if err != nil {
		return nil, fmt.Errorf("could not convert reiew entity to response: %v", err)
	}

	return res, nil
}

func (s server) Update(ctx context.Context, in *pb.ReviewUpdateRequest) (*pb.ReviewResponse, error) {

	ru := entity.ReviewUpdate{ReviewUpdateRequest: in}

	if err := ru.Validate(); err != nil {
		return nil, err
	}

	review, err := s.storage.GetByID(ctx, in.ID)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("review with ID %s counlt not be found", in.ID))
	}
	if review == nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("review with ID %s counlt not be found", in.ID))
	}

	review, err = s.storage.Update(ctx, ru, review)

	if err != nil {
		return nil, fmt.Errorf("storage could not update a review: %v", err)
	}

	res, err := convert(review)
	if err != nil {
		return nil, fmt.Errorf("could not convert reiew entity to response: %v", err)
	}

	return res, nil
}

func convert(r *entity.Review) (*pb.ReviewResponse, error) {
	var res pb.ReviewResponse
	res.ID = r.ID
	res.Name = r.Name
	res.Email = r.Email
	res.Content = r.Content
	res.Published = r.Published
	if r.Score.Valid {
		res.Score = &pb.ReviewResponse_ScoreValue{ScoreValue: r.Score.Int64}
	} else {
		res.Score = &pb.ReviewResponse_ScoreNull{ScoreNull: true}
	}

	if r.Category.Valid {
		res.Category = &pb.ReviewResponse_CategoryValue{CategoryValue: r.Category.String}
	} else {
		res.Category = &pb.ReviewResponse_CategoryNull{CategoryNull: true}
	}

	res.CreatedAt = r.CreatedAt.String()
	res.UpdatedAt = r.UpdatedAt.String()

	return &res, nil
}
