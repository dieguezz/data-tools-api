// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/nif.proto

/*
Package protopb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package protopb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_NifApi_GetControlDigit_0(ctx context.Context, marshaler runtime.Marshaler, client NifApiClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ControlDigitRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["nif"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "nif")
	}

	protoReq.Nif, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "nif", err)
	}

	msg, err := client.GetControlDigit(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NifApi_GetControlDigit_0(ctx context.Context, marshaler runtime.Marshaler, server NifApiServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ControlDigitRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["nif"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "nif")
	}

	protoReq.Nif, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "nif", err)
	}

	msg, err := server.GetControlDigit(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterNifApiHandlerServer registers the http handlers for service NifApi to "mux".
// UnaryRPC     :call NifApiServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterNifApiHandlerFromEndpoint instead.
func RegisterNifApiHandlerServer(ctx context.Context, mux *runtime.ServeMux, server NifApiServer) error {

	mux.Handle("GET", pattern_NifApi_GetControlDigit_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/proto.NifApi/GetControlDigit", runtime.WithHTTPPathPattern("/controldigit/{nif}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NifApi_GetControlDigit_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NifApi_GetControlDigit_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterNifApiHandlerFromEndpoint is same as RegisterNifApiHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterNifApiHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterNifApiHandler(ctx, mux, conn)
}

// RegisterNifApiHandler registers the http handlers for service NifApi to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterNifApiHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterNifApiHandlerClient(ctx, mux, NewNifApiClient(conn))
}

// RegisterNifApiHandlerClient registers the http handlers for service NifApi
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "NifApiClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "NifApiClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "NifApiClient" to call the correct interceptors.
func RegisterNifApiHandlerClient(ctx context.Context, mux *runtime.ServeMux, client NifApiClient) error {

	mux.Handle("GET", pattern_NifApi_GetControlDigit_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/proto.NifApi/GetControlDigit", runtime.WithHTTPPathPattern("/controldigit/{nif}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NifApi_GetControlDigit_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NifApi_GetControlDigit_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_NifApi_GetControlDigit_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"controldigit", "nif"}, ""))
)

var (
	forward_NifApi_GetControlDigit_0 = runtime.ForwardResponseMessage
)
