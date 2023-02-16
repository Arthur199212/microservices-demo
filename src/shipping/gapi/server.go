package gapi

import (
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"github.com/Arthur199212/microservices-demo/src/shipping/shipping"
	"github.com/go-playground/validator/v10"
)

const (
	maxProductsToShip = 100
)

type Server struct {
	shippingv1.UnimplementedShippingServiceServer
	validate *validator.Validate
	shipping shipping.ShippingService
}

func NewServer(shipping shipping.ShippingService) *Server {
	return &Server{
		validate: validator.New(),
		shipping: shipping,
	}
}
