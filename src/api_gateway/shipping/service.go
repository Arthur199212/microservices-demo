package shipping

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
)

type ShippingService interface {
	GetQuote(ctx context.Context, args GetQuoteArgs) (float32, error)
}

type shippingService struct {
	shippingClient shippingv1.ShippingServiceClient
}

func NewShippingService(
	shippingClient shippingv1.ShippingServiceClient,
) ShippingService {
	return &shippingService{
		shippingClient: shippingClient,
	}
}

type GetQuoteArgs struct {
	Address  *modelsv1.Address
	Products []*modelsv1.Product
}

func (s *shippingService) GetQuote(
	ctx context.Context,
	args GetQuoteArgs,
) (float32, error) {
	resp, err := s.shippingClient.GetQuote(ctx, &shippingv1.GetQuoteRequest{
		Address:  args.Address,
		Products: args.Products,
	})
	if err != nil {
		return 0, err
	}
	// todo: convert currency if needed
	return resp.GetQuote(), nil
}
