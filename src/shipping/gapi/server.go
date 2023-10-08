package gapi

import (
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"github.com/Arthur199212/microservices-demo/src/shipping/shipping"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/health/grpc_health_v1"

	"go.opentelemetry.io/otel/trace"
)

const (
	maxProductsToShip = 100
)

type Server struct {
	shippingv1.UnimplementedShippingServiceServer
	grpc_health_v1.UnimplementedHealthServer
	validate *validator.Validate
	shipping shipping.ShippingService
	tracer   trace.Tracer
}

func NewServer(tracer trace.Tracer, shipping shipping.ShippingService) *Server {
	return &Server{
		validate: validator.New(),
		shipping: shipping,
		tracer:   tracer,
	}
}
