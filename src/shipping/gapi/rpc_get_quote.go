package gapi

import (
	"context"

	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetQuote(ctx context.Context, req *shippingv1.GetQuoteRequest) (*shippingv1.GetQuoteResponse, error) {
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

	quote, err := s.shipping.GetQuote(address, products)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &shippingv1.GetQuoteResponse{
		Quote:        quote.Quote,
		CurrencyCode: quote.CurrencyCode,
	}, nil
}
