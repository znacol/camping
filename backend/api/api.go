package api

import (
	"compress/flate"
	"github.com/znacol/camping/backend/db"
	pb "github.com/znacol/camping/backend/proto"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
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
func CreateRouter(api *Service, gatewayHandler http.HandlerFunc) (http.Handler, error) {

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

	return r, nil
}
