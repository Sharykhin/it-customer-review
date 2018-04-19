package handler

import (
	"fmt"
	"net"

	"os"

	"log"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/Sharykhin/it-customer-review/grpc-server/database/mysql"
	"google.golang.org/grpc"
)

// ListenAndServe creates grps server and start listening income connections
func ListenAndServe() error {
	address := os.Getenv("GRPC_ADDRESS")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterReviewServer(s, &server{storage: mysql.Storage})
	s.Serve(lis)

	fmt.Printf("Start listening on %s\n", address)
	return s.Serve(lis)
}
