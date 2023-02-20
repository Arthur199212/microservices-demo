package main

import (
	"fmt"
	"net"

	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	"github.com/Arthur199212/microservices-demo/src/payment/gapi"
	"github.com/Arthur199212/microservices-demo/src/payment/utils"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load .env file")
	}

	config := utils.LoadConfig()

	srv := gapi.NewServer(config)
	grpcServer := grpc.NewServer()
	paymentv1.RegisterPaymentServiceServer(grpcServer, srv)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("starting gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}
