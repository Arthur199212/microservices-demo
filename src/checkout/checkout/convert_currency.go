package checkout

import (
	"context"
	"fmt"

	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	currency "github.com/Arthur199212/microservices-demo/src/currency/pb"
	"github.com/rs/zerolog/log"
)

func (s *checkoutService) convertCurrency(
	ctx context.Context,
	fromCurrencyCode string,
	toCurrencyCode string,
	amount float32,
) (*pb.Money, error) {
	if fromCurrencyCode == toCurrencyCode {
		return &pb.Money{
			Amount:       amount,
			CurrencyCode: toCurrencyCode,
		}, nil
	}

	resp, err := s.currencyClient.Convert(ctx, &currency.ConvertRequest{
		From: &currency.Money{
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
	return &pb.Money{
		CurrencyCode: resp.GetCurrencyCode(),
		Amount:       resp.GetAmount(),
	}, nil
}
