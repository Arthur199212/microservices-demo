package gapi

import (
	"github.com/Arthur199212/microservices-demo/src/checkout/checkout"
	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	pb.UnimplementedCheckoutServer
	checkoutService checkout.CheckoutService
	validate        *validator.Validate
}

func NewServer(checkoutService checkout.CheckoutService) *Server {
	return &Server{
		checkoutService: checkoutService,
		validate:        validator.New(),
	}
}
