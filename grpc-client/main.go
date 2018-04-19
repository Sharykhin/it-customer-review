package main

import (
	"log"

	"context"

	"fmt"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewReviewClient(conn)

	review, err := client.Create(context.Background(), &pb.ReviewRequest{
		Name:      "asdas",
		Email:     "asdas",
		Content:   "asdasd",
		Published: false,
	})

	if err != nil {
		log.Fatal("could not crete a new review")
	}

	fmt.Println(review)

	updatedReview, err := client.Update(context.Background(), &pb.ReviewRequest{
		Name:      "asdas",
		Email:     "asdas",
		Content:   "asdasd",
		Published: false,
		Score:     uint64(84),
		Category:  "positive",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedReview)

}
