package gapi

import (
	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	"github.com/Arthur199212/microservices-demo/src/payment/utils"
)

type Server struct {
	paymentv1.UnimplementedPaymentServiceServer
	config utils.Config
}

func NewServer(config utils.Config) *Server {
	return &Server{
		config: config,
	}
}
