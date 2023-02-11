package main

import (
	"fmt"
	"net"

	"github.com/Arthur199212/microservices-demo/src/payment/gapi"
	"github.com/Arthur199212/microservices-demo/src/payment/pb"
	"github.com/Arthur199212/microservices-demo/src/payment/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := utils.LoadConfig()

	srv := gapi.NewServer(config)
	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServer(grpcServer, srv)

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