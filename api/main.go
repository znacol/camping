package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/znacol/camping/api/db"
	pb "github.com/znacol/camping/api/proto"
)

type server struct {
	dbClient DB
}

func main() {
	db := db.New("root", "password", "camping", "localhost", "5302")

	addr := ":30251"
	clientAddr := fmt.Sprintf("localhost%s", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	go runGRPC(lis)
	runHTTP(clientAddr)
}

func runGRPC(lis net.Listener, dbClient DB) {
	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	campingServer := &server{
		dbClient: dbClient,
	}

	pb.RegisterCampingServiceServer(grpcServer, &server{})

	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	grpcServer.Serve(lis)
}

// TODO cleanup
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Printf("preflight request for %s", r.URL.Path)
	return
}

// TODO cleanup
// allowCORS allows Cross Origin Resoruce Sharing from any origin.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func runHTTP(clientAddr string) {
	addr := ":8081"

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := pb.RegisterCampingServiceHandlerFromEndpoint(context.Background(), mux, clientAddr, opts); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	log.Printf("HTTP Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, allowCORS(mux)))
}

/*
* Endpoints
 */

// GetSavedSites returns information on campsites
func (s *server) GetSavedSites(c context.Context, request *pb.GetSavedSitesRequest) (response *pb.GetSavedSitesResponse, err error) {
	locations, err := svc.dbClient.GetSavedSites(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting s")
	}

	response = &pb.GetSavedSitesResponse{
		Message: "Hello world",
	}

	return response, nil
}
