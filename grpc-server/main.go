package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"google.golang.org/grpc"
)

type server struct {
}

func (s server) Create(ctx context.Context, request *pb.ReviewRequest) (*pb.ReviewResponse, error) {
	review := new(pb.ReviewResponse)
	review.ID = request.ID
	review.Name = request.Name
	review.Email = request.Email
	review.Content = request.Content
	review.Published = request.Published
	review.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05")

	return review, nil
}

func (s server) Update(ctx context.Context, request *pb.ReviewRequest) (*pb.ReviewResponse, error) {
	review := new(pb.ReviewResponse)
	review.ID = request.ID
	review.Name = request.Name
	review.Email = request.Email
	review.Content = request.Content
	review.Published = request.Published
	review.Score = request.Score
	review.Category = request.Category
	review.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05")
	return review, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterReviewServer(s, &server{})
	s.Serve(lis)
}
