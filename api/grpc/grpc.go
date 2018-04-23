package grpc

import (
	"context"
	"log"
	"os"

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

func (ctrl reviewService) Create(ctx context.Context, rr entity.ReviewRequest) (*entity.Review, error) {

	res, err := ctrl.client.Create(ctx, &pb.ReviewCreateRequest{
		Name:      rr.Name,
		Email:     rr.Email,
		Content:   rr.Content,
		Published: false,
		Score:     &pb.ReviewCreateRequest_ScoreNull{ScoreNull: true},
		Category:  &pb.ReviewCreateRequest_CategoryNull{CategoryNull: true},
	})

	if err != nil {
		return nil, err
	}

	r := convert(res)
	return r, nil
}

func (ctrl reviewService) Update(ctx context.Context, ID string, rr entity.ReviewUpdateRequest) (*entity.Review, error) {
	var fu pb.FieldToUpdate

	if rr.Name.Valid {
		fu.Name = &pb.FieldToUpdate_NameValue{NameValue: rr.Name.Value}
	} else {
		fu.Name = &pb.FieldToUpdate_NameNull{NameNull: true}
	}

	if rr.Content.Valid {
		fu.Content = &pb.FieldToUpdate_ContentValue{ContentValue: rr.Content.Value}
	} else {
		fu.Content = &pb.FieldToUpdate_ContentNull{ContentNull: true}
	}

	if rr.Published.Valid {
		fu.Published = &pb.FieldToUpdate_PublishedValue{PublishedValue: rr.Published.Value}
	} else {
		fu.Published = &pb.FieldToUpdate_PublishedNull{PublishedNull: true}
	}

	fu.Score = &pb.FieldToUpdate_ScoreNull{ScoreNull: true}
	fu.Category = &pb.FieldToUpdate_CategoryNull{CategoryNull: true}

	in := pb.ReviewUpdateRequest{
		ID:             ID,
		FieldsToUpdate: &fu,
	}

	res, err := ctrl.client.Update(ctx, &in)
	if err != nil {
		return nil, err
	}

	r := convert(res)
	return r, nil
}

func (ctrl reviewService) Get(ctx context.Context, ID string) (*entity.Review, error) {
	res, err := ctrl.client.Get(ctx, &pb.ReviewID{ID: ID})
	if err != nil {
		return nil, err
	}
	r := convert(res)
	return r, nil
}

func convert(res *pb.ReviewResponse) *entity.Review {

	r := entity.Review{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		Content:   res.Content,
		Published: res.Published,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	if res.GetCategoryNull() {
		r.Category = entity.NullString{Valid: false}
	} else {
		r.Category = entity.NullString{Valid: true, Value: res.GetCategoryValue()}
	}

	if res.GetScoreNull() {
		r.Score = entity.Score(-1)
	} else {
		r.Score = entity.Score(res.GetScoreValue())
	}
	return &r
}
