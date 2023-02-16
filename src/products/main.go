package main

import (
	"fmt"
	"net"
	"os"

	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	"github.com/Arthur199212/microservices-demo/src/products/db"
	"github.com/Arthur199212/microservices-demo/src/products/gapi"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
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

	// to provide self-documentation
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("starting gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}
