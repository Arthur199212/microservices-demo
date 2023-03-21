package gapi

import (
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	"github.com/Arthur199212/microservices-demo/src/checkout/checkout"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	checkoutv1.UnimplementedCheckoutServiceServer
	grpc_health_v1.UnimplementedHealthServer
	checkoutService checkout.CheckoutService
	validate        *validator.Validate
}

func NewServer(checkoutService checkout.CheckoutService) *Server {
	return &Server{
		checkoutService: checkoutService,
		validate:        validator.New(),
	}
}
