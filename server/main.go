package main

import (
	"context"
	"log"
	"net"
	"net/http"

	nifgenerators "github.com/dieguezz/nif-tools/pkg/generators"
	nifparser "github.com/dieguezz/nif-tools/pkg/parser"
	nifvalidator "github.com/dieguezz/nif-tools/pkg/validator"
	pb "github.com/dieguezz/nif-tools/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedNifApiServer
}

func (s *server) GetNIFControlDigit(ctx context.Context, in *pb.NIF) (*pb.ControlDigitResponse, error) {
	nif, err := nifparser.GetParsedNIF(in.GetNif())
	return &pb.ControlDigitResponse{ControlDigit: nif.Control}, err
}

func (s *server) GetType(ctx context.Context, in *pb.NIF) (*pb.TypeResponse, error) {
	nif, err := nifparser.GetParsedNIF(in.GetNif())
	return &pb.TypeResponse{Type: nif.Kind}, err
}

func (s *server) GetParsedNIF(ctx context.Context, in *pb.NIF) (*pb.ParsedNIFResponse, error) {
	nif, err := nifparser.GetParsedNIF(in.GetNif())
	return &pb.ParsedNIFResponse{Number: int32(nif.Number), Kind: nif.Kind, Control: nif.Control}, err
}

func (s *server) GenerateNIE(ctx context.Context, in *emptypb.Empty) (*pb.NIE, error) {
	nie := nifgenerators.GenerateNIE()
	return &pb.NIE{Nie: nie}, nil
}

func (s *server) GenerateNIF(ctx context.Context, in *emptypb.Empty) (*pb.NIF, error) {
	nif := nifgenerators.GenerateNIF(nifgenerators.NIFOptions{})
	return &pb.NIF{Nif: nif}, nil
}

func (s *server) ValidateNIF(ctx context.Context, in *pb.NIF) (*pb.IsValid, error) {
	isValid := nifvalidator.IsValidNIF(in.GetNif())
	return &pb.IsValid{IsValid: isValid}, nil
}

func (s *server) ValidateNIE(ctx context.Context, in *pb.NIE) (*pb.IsValid, error) {
	isValid := nifvalidator.IsValidNIE(in.GetNie())
	return &pb.IsValid{IsValid: isValid}, nil
}

func main() {
	// Grpc server
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	go func() {
		mux := runtime.NewServeMux()
		pb.RegisterNifApiHandlerServer(context.Background(), mux, &server{})
		// Rest server
		log.Fatalln(http.ListenAndServe("localhost:8080", mux))
	}()

	pb.RegisterNifApiServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
