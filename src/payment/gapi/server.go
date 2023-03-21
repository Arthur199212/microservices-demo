package gapi

import (
	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	"github.com/Arthur199212/microservices-demo/src/payment/utils"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	paymentv1.UnimplementedPaymentServiceServer
	grpc_health_v1.UnimplementedHealthServer
	config utils.Config
}

func NewServer(config utils.Config) *Server {
	return &Server{
		config: config,
	}
}
