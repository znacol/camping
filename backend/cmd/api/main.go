package main

import (
	"context"
	"github.com/znacol/camping/backend/api"
	"github.com/znacol/camping/backend/db"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	pb "github.com/znacol/camping/backend/proto"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip" // Allow gzip encoding of grpc responses
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// init sets up environment configuration
func init() {
	log.WithFields(log.Fields{"cpu": runtime.NumCPU(), "maxProcs": runtime.GOMAXPROCS(-1), "goroutines": runtime.NumGoroutine()}).Info("Initializing...")

	// Parse the environment variables
	if err := envconfig.Process("", &env); err != nil {
		log.WithError(err).Fatal("Unable to process environment config")
	}

	// Set verbose debugging
	if env.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug level logging enabled")
	}
}

// appEnv represents all environment/configuration needed for this application
type appEnv struct {
	Debug                  bool   `default:"false"`
	GrpcAddress    string `default:":50051" split_words:"true"`
	RestAddress    string `default:":8000" split_words:"true"`
	DbUser           string `default:"root" split_words:"true"`
	DbPassword       string `default:"password" split_words:"true"`
	DbHost           string `default:"camping-db" split_words:"true"`
	DbPort           string `default:"5432" split_words:"true"`
	DbName           string `default:"camping" split_words:"true"`
}

// env is the current environment configuration
var env appEnv

// main entrypoint
func main() {
	// Signal handlers and context
	sigCtx, sigCancelFunc := context.WithCancel(context.Background())
	setupSignalHandlers(sigCancelFunc)

	log := log.WithFields(log.Fields{
		"GrpcAddress": env.GrpcAddress,
		"RestAddress": env.RestAddress,
	})
	log.Info("Initializing Camping API...")

	// Connect to database
	log.Info("Connecting to the database")
	dbClient, err := db.New(env.DbUser, env.DbPassword, env.DbName, env.DbHost, env.DbPort)
	if err != nil {
		log.WithError(err).Fatal("Unable to connect to database")
	}

	// Create API
	apiHandler := api.New(dbClient)

	// Used to wait for all servers to gracefully shutdown
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Start servers
	go startGrpcServer(sigCtx, apiHandler, wg)
	go startRestGateway(sigCtx, apiHandler, wg)

	// Wait for termination signal and graceful shutdown
	wg.Wait()

	log.Info("Graceful shutdown")
}

// startGrpcServer configures then starts up the API GRPC Server in a new goroutine,
// then waits for the context to be done before gracefully terminating the server.
func startGrpcServer(ctx context.Context, apiHandler *api.Service, shutdownWG *sync.WaitGroup) {
		defer shutdownWG.Done()
	log := log.WithField("address", env.GrpcAddress)

	// Create and run GRPC Server
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time: 60 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			PermitWithoutStream: true,
			MinTime:             10 * time.Second,
		}),
	)

	pb.RegisterCampingServiceServer(grpcServer, apiHandler)
	reflection.Register(grpcServer)

	go func() {
		log.Info("Starting Camping API GRPC...")
		listener, listenErr := net.Listen("tcp", env.GrpcAddress)
		if listenErr != nil {
			log.WithError(listenErr).Fatal("Unable to listen for camping api grpc")
		}
		if serveErr := grpcServer.Serve(listener); serveErr != nil {
			log.WithError(serveErr).Fatal("Unable to start camping api grpc")
		}
		log.Debug("Camping API GRPC server closed")
	}()

	// Wait for shutdown signal
	<-ctx.Done()

	// Gracefully handle all current grpc connections (unfortunately no way to put a timeout on this)
	grpcServer.GracefulStop()
	log.Debug("Camping API GRPC server shutdown")
}

// startRestGateway configures then starts up the API HTTP REST->GRPC Gateway Server in a new goroutine,
// then waits for the context to be done before gracefully terminating the server.
func startRestGateway(ctx context.Context, apiHandler *api.Service, shutdownWG *sync.WaitGroup) {
	defer shutdownWG.Done()

	// Create GRPC Gateway Router/Mux
	gwMux := gwruntime.NewServeMux(gwruntime.WithMarshalerOption(gwruntime.MIMEWildcard, &gwruntime.JSONPb{OrigName: true, EmitDefaults: true}))

	// Register GRPC Gateway with GRPC Server
	gatewayCtx, gatewayCancelFunc := context.WithCancel(context.Background())
	if err := pb.RegisterCampingServiceHandlerFromEndpoint(gatewayCtx, gwMux, env.GrpcAddress, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.WithError(err).Fatal("Unable to register camping gateway with grpc server")
	}

	// Create router
	router, err := api.CreateRouter(apiHandler, gwMux.ServeHTTP)
	if err != nil {
		log.WithError(err).Fatal("Unable to create camping router")
	}

	// Create and run HTTP REST Server
	httpRestServer := &http.Server{
		Addr:         env.RestAddress,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		// Wrap to support http/2 cleartext
		// This is required if we want http/2, because we are using https termination at our load balancer.
		Handler: h2c.NewHandler(router, &http2.Server{}),
	}
	go func() {
		log.Info("Starting Camping API REST...")
		if serveErr := httpRestServer.ListenAndServe(); serveErr != nil && serveErr != http.ErrServerClosed {
			log.WithError(serveErr).Fatal("Unable to start camping api rest")
		}
		log.Debug("Camping API REST server closed")
	}()

	// Wait for shutdown signal
	<-ctx.Done()

	// Gracefully handle all current http rest connections
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer shutdownCancel()
	if shutdownErr := httpRestServer.Shutdown(shutdownCtx); shutdownErr != nil {
		log.WithError(shutdownErr).Warn("Unable to gracefully handle all http rest connections")
	}
	gatewayCancelFunc()
	log.Debug("Camping API REST server shutdown")
}


// setupSignalHandlers sets up common application signal handlers
// to cancel the context upon interrupt or sigterm.
func setupSignalHandlers(cancel context.CancelFunc) {
	// Setup a signal handler.
	sigChan := make(chan os.Signal, 8)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.WithField("signal", sig.String()).Warn("Caught cancellation signal")
		cancel()
	}()

	log.Info("Signal handlers setup")
}

