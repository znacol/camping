package api

import (
	"compress/flate"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/znacol/camping/backend/db"
	pb "github.com/znacol/camping/backend/proto"
	"log"
	"net/http"
	"strings"
)

var _ pb.CampingServiceServer = &Service{}

type Service struct {
	dbClient *db.DB
}

func New(db *db.DB) *Service {
	return &Service{
		dbClient: db,
	}
}

// CreateRouter creates a new Router / Mux / http.Handler to route all traffic to the proper handlers
func CreateRouter(gatewayHandler http.HandlerFunc) (http.Handler, error) {

	// Router
	r := chi.NewRouter()

	// Middlewares for all endpoints (include GRPC Gateway)
	r.Use(
		// Redirect requests with trailing path slash
		middleware.RedirectSlashes,

		// CORS
		cors.AllowAll().Handler,

		// Compress responses if not already compressed, and request Accept-Encoding allows it
		middleware.Compress(flate.BestSpeed),
	)

	// GRPC->REST Gateway handles any routes/methods not defined below
	r.MethodNotAllowed(gatewayHandler)
	r.NotFound(gatewayHandler)

	return allowCORS(r), nil
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Printf("preflight request for %s", r.URL.Path)
	return
}

// allowCORS allows Cross Origin Resource Sharing from any origin
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
