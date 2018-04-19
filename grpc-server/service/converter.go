package service

import (
	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/Sharykhin/it-customer-review/grpc-server/entity"
)

// ConvertReviewToResponse is helper fun that converts local entity to a grpc proto response
func ConvertReviewToResponse(r *entity.Review) (*pb.ReviewResponse, error) {
	res := pb.ReviewResponse{
		ID:        r.ID,
		Name:      r.Name,
		Email:     r.Email,
		Content:   r.Content,
		Score:     int64(r.Score),
		Category:  r.Category.String,
		CreatedAt: r.CreatedAt.String(),
	}

	return &res, nil
}
