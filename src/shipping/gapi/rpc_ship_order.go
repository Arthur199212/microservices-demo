package gapi

import (
	"context"

	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.opentelemetry.io/otel/attribute"
)

func (s *Server) ShipOrder(ctx context.Context, req *shippingv1.ShipOrderRequest) (*shippingv1.ShipOrderResponse, error) {
	_, span := s.tracer.Start(ctx, "ShipOrder")
	defer span.End()

	products := convertToProducts(req.GetProducts())
	if len(products) < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "product list is empty")
	}
	if len(products) > maxProductsToShip {
		return nil, status.Errorf(codes.InvalidArgument, "product list has more then %d items", maxProductsToShip)
	}
	if err := s.validate.Var(products, "required,dive"); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	address := convertToAddress(req.GetAddress())
	if err := s.validate.Struct(address); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	trackingId, err := s.shipping.ShipOrder(address, products)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	span.SetAttributes(
		attribute.String("shipping.tracking_id", trackingId),
	)

	return &shippingv1.ShipOrderResponse{
		TrackingId: trackingId,
	}, nil
}
