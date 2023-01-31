package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Arthur199212/microservices-demo/src/cart/db"
	"github.com/Arthur199212/microservices-demo/src/cart/gapi"
	"github.com/Arthur199212/microservices-demo/src/cart/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = "5001"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := db.NewCartDB()
	srv := gapi.NewServer(db)
	grpcServer := grpc.NewServer()
	pb.RegisterCartServer(grpcServer, srv)

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
