package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	"github.com/Arthur199212/microservices-demo/src/products/db"
	"github.com/Arthur199212/microservices-demo/src/products/gapi"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = "5000"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	pdb := db.NewProductsDB()
	srv := gapi.NewServer(pdb)
	grpcServer := grpc.NewServer()
	productsv1.RegisterProductsServiceServer(grpcServer, srv)
	grpc_health_v1.RegisterHealthServer(grpcServer, srv)

	// to provide self-documentation
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	go func() {
		log.Info().Msgf("starting gRPC server at %s", listener.Addr().String())
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot start gRPC server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("server is shutting down ...")

	grpcServer.GracefulStop()
}
