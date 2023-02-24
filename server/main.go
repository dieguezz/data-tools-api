package main

import (
	"context"
	"log"
	"net"
	"net/http"

	cifgenerators "github.com/dieguezz/nif-tools/pkg/cif/generators"
	documentparser "github.com/dieguezz/nif-tools/pkg/document-parser"
	niegenerators "github.com/dieguezz/nif-tools/pkg/nie/generators"
	nievalidator "github.com/dieguezz/nif-tools/pkg/nie/validators"
	nifgenerators "github.com/dieguezz/nif-tools/pkg/nif/generators"
	nifvalidator "github.com/dieguezz/nif-tools/pkg/nif/validators"
	pb "github.com/dieguezz/nif-tools/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedNifApiServer
}

func (s *server) GetNIFControlDigit(ctx context.Context, in *pb.NIF) (*pb.ControlDigitResponse, error) {
	nif, err := documentparser.GetParsedDocument(in.GetNif())
	return &pb.ControlDigitResponse{ControlDigit: nif.Control}, err
}

func (s *server) GetType(ctx context.Context, in *pb.NIF) (*pb.TypeResponse, error) {
	nif, err := documentparser.GetParsedDocument(in.GetNif())
	return &pb.TypeResponse{Type: nif.Kind}, err
}

func (s *server) GetParsedDocument(ctx context.Context, in *pb.NIF) (*pb.ParsedDocumentResponse, error) {
	nif, err := documentparser.GetParsedDocument(in.GetNif())
	return &pb.ParsedDocumentResponse{Number: int32(nif.Number), Kind: nif.Kind, Control: nif.Control}, err
}

func (s *server) GenerateNIF(ctx context.Context, in *emptypb.Empty) (*pb.NIF, error) {
	nif := nifgenerators.GenerateNIF(nifgenerators.NIFOptions{})
	return &pb.NIF{Nif: nif}, nil
}

func (s *server) ValidateNIF(ctx context.Context, in *pb.NIF) (*pb.IsValid, error) {
	isValid := nifvalidator.IsValidNIF(in.GetNif())
	return &pb.IsValid{IsValid: isValid}, nil
}

func (s *server) GenerateNIE(ctx context.Context, in *emptypb.Empty) (*pb.NIE, error) {
	nie := niegenerators.GenerateNIE()
	return &pb.NIE{Nie: nie}, nil
}

func (s *server) ValidateNIE(ctx context.Context, in *pb.NIE) (*pb.IsValid, error) {
	isValid := nievalidator.IsValidNIE(in.GetNie())
	return &pb.IsValid{IsValid: isValid}, nil
}

func (s *server) GetCIFControlDigit(ctx context.Context, in *pb.CIF) (*pb.ControlDigitResponse, error) {
	cif, err := documentparser.GetParsedDocument(in.GetCif())
	return &pb.ControlDigitResponse{ControlDigit: cif.Control}, err
}

func (s *server) GenerateCIF(ctx context.Context, in *emptypb.Empty) (*pb.CIF, error) {
	cif := cifgenerators.GenerateCIF()
	return &pb.CIF{Cif: cif}, nil
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
