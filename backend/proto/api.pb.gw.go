// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: backend/proto/api.proto

/*
Package camping is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package camping

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_CampingService_GetAllSites_0(ctx context.Context, marshaler runtime.Marshaler, client CampingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAllSitesRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetAllSites(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterCampingServiceHandlerFromEndpoint is same as RegisterCampingServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCampingServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterCampingServiceHandler(ctx, mux, conn)
}

// RegisterCampingServiceHandler registers the http handlers for service CampingService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCampingServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCampingServiceHandlerClient(ctx, mux, NewCampingServiceClient(conn))
}

// RegisterCampingServiceHandlerClient registers the http handlers for service CampingService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CampingServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CampingServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CampingServiceClient" to call the correct interceptors.
func RegisterCampingServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CampingServiceClient) error {

	mux.Handle("GET", pattern_CampingService_GetAllSites_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CampingService_GetAllSites_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CampingService_GetAllSites_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CampingService_GetAllSites_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "camping", "sites"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_CampingService_GetAllSites_0 = runtime.ForwardResponseMessage
)