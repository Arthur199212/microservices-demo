package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Arthur199212/microservices-demo/src/shipping/gapi"
	"github.com/Arthur199212/microservices-demo/src/shipping/pb"
	"github.com/Arthur199212/microservices-demo/src/shipping/shipping"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = "5004"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	shippingService := shipping.NewShippingService()
	srv := gapi.NewServer(shippingService)
	grpcServer := grpc.NewServer()
	pb.RegisterShippingServer(grpcServer, srv)

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
