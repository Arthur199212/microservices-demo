package currency

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
)

type CurrencyService interface {
	Convert(ctx context.Context, args ConvertArgs) (*modelsv1.Money, error)
	GetSupportedCurrencies(ctx context.Context) ([]string, error)
}

type currencyService struct {
	client currencyv1.CurrencyServiceClient
}

func NewCurrencyService(
	client currencyv1.CurrencyServiceClient,
) CurrencyService {
	return &currencyService{
		client: client,
	}
}

type ConvertArgs struct {
	From           *modelsv1.Money
	ToCurrencyCode string
}

func (s *currencyService) Convert(
	ctx context.Context,
	args ConvertArgs,
) (*modelsv1.Money, error) {
	resp, err := s.client.Convert(ctx, &currencyv1.ConvertRequest{
		From:           args.From,
		ToCurrencyCode: args.ToCurrencyCode,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetMoney(), nil
}

func (s *currencyService) GetSupportedCurrencies(ctx context.Context) ([]string, error) {
	resp, err := s.client.GetSupportedCurrencies(ctx, &currencyv1.GetSupportedCurrenciesRequest{})
	if err != nil {
		return nil, err
	}
	return resp.GetCurrencyCodes(), nil
}
