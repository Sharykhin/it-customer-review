package service

import (
	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/Sharykhin/it-customer-review/grpc-server/entity"
)

func ConvertReviewToResponse(r *entity.Review) (*pb.ReviewResponse, error) {
	res := pb.ReviewResponse{
		ID:        r.ID,
		Name:      r.Name,
		Email:     r.Email,
		Content:   r.Content,
		Score:     r.Score,
		Category:  r.Category,
		CreatedAt: r.CreatedAt.String(),
	}

	return &res, nil
}
