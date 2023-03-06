package service

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/currency"
)

type ShippingService interface {
	GetQuote(ctx context.Context, args GetQuoteArgs) (*modelsv1.Money, error)
}

type shippingService struct {
	shippingClient  shippingv1.ShippingServiceClient
	currencyService currency.CurrencyService
}

func NewShippingService(
	shippingClient shippingv1.ShippingServiceClient,
	currencyService currency.CurrencyService,
) ShippingService {
	return &shippingService{
		shippingClient:  shippingClient,
		currencyService: currencyService,
	}
}

type GetQuoteArgs struct {
	Address      *modelsv1.Address
	Products     []*modelsv1.Product
	UserCurrency string
}

func (s *shippingService) GetQuote(
	ctx context.Context,
	args GetQuoteArgs,
) (*modelsv1.Money, error) {
	resp, err := s.shippingClient.GetQuote(ctx, &shippingv1.GetQuoteRequest{
		Address:  args.Address,
		Products: args.Products,
	})
	if err != nil {
		return nil, err
	}

	quote, quoteCurrency := resp.GetQuote(), resp.GetCurrencyCode()
	if args.UserCurrency == quoteCurrency {
		return nil, err
	}

	money, err := s.currencyService.Convert(ctx, currency.ConvertArgs{
		From: &modelsv1.Money{
			Amount:       quote,
			CurrencyCode: quoteCurrency,
		},
		ToCurrencyCode: args.UserCurrency,
	})
	if err != nil {
		return nil, err
	}
	return money, nil
}
