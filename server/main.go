package main

import (
	// "context"
	"github.com/dieguezz/nif-tools/control-digit"
	// pb "github.com/dieguezz/nif-tools/proto"
	// "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	// "google.golang.org/grpc"
	// "log"
	// "net"
	// "net/http"
)

// type server struct {
// 	pb.UnimplementedNifApiServer
// }

// func (s *server) GetControlDigit(ctx context.Context, in *pb.ControlDigitRequest) (*pb.ControlDigitResponse, error) {
// 	controlDigit := controldigit.GetControlDigit(in.GetNif())
// 	return &pb.ControlDigitResponse{ControlDigit: controlDigit}, nil
// }

func main() {

	// lis, err := net.Listen("tcp", ":9000")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// grpcServer := grpc.NewServer()

	// go func() {
	// 	mux := runtime.NewServeMux()
	// 	pb.RegisterNifApiHandlerServer(context.Background(), mux, &server{})
	// 	log.Fatalln(http.ListenAndServe("localhost:8080", mux))
	// }()

	// pb.RegisterNifApiServer(grpcServer, &server{})

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	// }
	controldigit.GetControlDigit("L51069845r")
	controldigit.GetControlDigit("x5555555")
	controldigit.GetControlDigit("y5555555")
	controldigit.GetControlDigit("z5555555")
}
