package grpc

import (
	"context"
	"log"
	"os"

	"fmt"
	"io"

	"github.com/Sharykhin/it-customer-review/api/contract"
	"github.com/Sharykhin/it-customer-review/api/entity"
	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"google.golang.org/grpc"
)

type (
	reviewService struct {
		client pb.ReviewClient
	}

	listResult struct {
		err  error
		list []entity.Review
	}

	countResult struct {
		err   error
		total int64
	}
)

// ReviewService is a grpc service that would be responsible for managing reviews.
var ReviewService contract.ReviewProvider

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

func (ctrl reviewService) Index(ctx context.Context, criteria []entity.Criteria, limit, offset int64) ([]entity.Review, int64, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	chList := ctrl.runList(ctx, criteria, limit, offset)
	chCount := ctrl.runCount(ctx, criteria)

	var rl []entity.Review
	var total int64

	for {
		if chList == nil && chCount == nil {
			return rl, total, nil
		}
		select {
		case listRes, ok := <-chList:
			if !ok {
				chList = nil
				continue
			}
			if listRes.err != nil {
				cancel()
				return nil, 0, listRes.err
			}
			rl = listRes.list
		case countRes, ok := <-chCount:
			if !ok {
				chCount = nil
				continue
			}
			if countRes.err != nil {
				cancel()
				return nil, 0, countRes.err
			}
			total = countRes.total
		}
	}
}

func (ctrl reviewService) runList(ctx context.Context, criteria []entity.Criteria, limit, offset int64) <-chan listResult {
	chListResult := make(chan listResult)
	go func() {
		defer close(chListResult)
		var lr listResult
		var rl []entity.Review
		var c []*pb.Criteria
		for _, filter := range criteria {
			c = append(c, &pb.Criteria{Key: filter.Key, Value: filter.Value})
		}
		stream, err := ctrl.client.GetReviewList(ctx, &pb.ReviewListFilter{Criteria: c, Limit: limit, Offset: offset})

		if err != nil {
			lr.err = err
		} else {
			for {
				res, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					lr.err = fmt.Errorf("failed to get a review from a stream: %v", err)
					break
				}
				r := convert(res)
				rl = append(rl, *r)
			}
		}

		lr.list = rl
		select {
		case <-ctx.Done():
			return
		case chListResult <- lr:
		}
	}()
	return chListResult
}

func (ctrl reviewService) runCount(ctx context.Context, criteria []entity.Criteria) <-chan countResult {
	chCountResult := make(chan countResult)
	go func() {
		defer close(chCountResult)
		var cr countResult
		var c []*pb.Criteria
		for _, filter := range criteria {
			c = append(c, &pb.Criteria{Key: filter.Key, Value: filter.Value})
		}

		t, err := ctrl.client.CountReviews(ctx, &pb.ReviewCountFilter{Criteria: c})
		if err != nil {
			cr.err = err
		}

		cr.total = t.Total
		select {
		case <-ctx.Done():
			return
		case chCountResult <- cr:
		}
	}()
	return chCountResult
}

func (ctrl reviewService) Count(ctx context.Context, criteria *pb.ReviewCountFilter) (int64, error) {
	res, err := ctrl.client.CountReviews(ctx, criteria)
	if err != nil {
		return 0, err
	}
	return res.Total, nil
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
