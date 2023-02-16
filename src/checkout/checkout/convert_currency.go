package checkout

import (
	"context"
	"fmt"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	"github.com/rs/zerolog/log"
)

func (s *checkoutService) convertCurrency(
	ctx context.Context,
	fromCurrencyCode string,
	toCurrencyCode string,
	amount float32,
) (*modelsv1.Money, error) {
	if fromCurrencyCode == toCurrencyCode {
		return &modelsv1.Money{
			Amount:       amount,
			CurrencyCode: toCurrencyCode,
		}, nil
	}

	resp, err := s.currencyClient.Convert(ctx, &currencyv1.ConvertRequest{
		From: &modelsv1.Money{
			CurrencyCode: fromCurrencyCode,
			Amount:       amount,
		},
		ToCurrencyCode: toCurrencyCode,
	})
	if err != nil {
		errMsg := fmt.Errorf("cannot convert currency: %+v", err)
		log.Error().Err(err).
			Msgf(errMsg.Error())
		return nil, errMsg
	}
	return &modelsv1.Money{
		CurrencyCode: resp.GetCurrencyCode(),
		Amount:       resp.GetAmount(),
	}, nil
}
