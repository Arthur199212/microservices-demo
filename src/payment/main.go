package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	"github.com/Arthur199212/microservices-demo/src/payment/gapi"
	"github.com/Arthur199212/microservices-demo/src/payment/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := utils.LoadConfig("configs")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	srv := gapi.NewServer(config)
	grpcServer := grpc.NewServer()
	paymentv1.RegisterPaymentServiceServer(grpcServer, srv)
	grpc_health_v1.RegisterHealthServer(grpcServer, srv)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))
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
