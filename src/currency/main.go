package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/Arthur199212/microservices-demo/src/currencies/data"
	"github.com/Arthur199212/microservices-demo/src/currencies/gapi"
	"github.com/Arthur199212/microservices-demo/src/currencies/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = "5002"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cd, err := data.NewCurrencyData()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot setup currency data")
	}
	cd.MonitorRates(time.Hour) // checks the rates in ECB every hour

	srv := gapi.NewServer(cd)
	grpcServer := grpc.NewServer()
	pb.RegisterProductsServer(grpcServer, srv)

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
