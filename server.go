package main

import (
	"fmt"
	pb "github.com/dieguezz/nif-tools/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); ett != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
