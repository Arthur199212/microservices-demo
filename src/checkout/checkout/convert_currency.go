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
	money *modelsv1.Money,
	toCurrencyCode string,
) (*modelsv1.Money, error) {
	if money.CurrencyCode == toCurrencyCode {
		return money, nil
	}

	resp, err := s.currencyClient.Convert(ctx, &currencyv1.ConvertRequest{
		From:           money,
		ToCurrencyCode: toCurrencyCode,
	})
	if err != nil {
		err = fmt.Errorf("cannot convert currency: %+v", err)
		log.Error().Err(err).
			Msgf(err.Error())
		return nil, err
	}

	return resp.GetMoney(), nil
}
