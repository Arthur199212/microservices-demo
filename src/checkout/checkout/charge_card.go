package checkout

import (
	"context"
	"fmt"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	"github.com/rs/zerolog/log"
)

func (s *checkoutService) chargeCard(
	ctx context.Context,
	cardInfo CardInfo,
	money *modelsv1.Money,
) (string, error) {
	resp, err := s.paymentClient.Charge(ctx, &paymentv1.ChargeRequest{
		Money: &modelsv1.Money{
			CurrencyCode: money.CurrencyCode,
			Amount:       money.Amount,
		},
		CardInfo: &modelsv1.CardInfo{
			Number:          cardInfo.Number,
			Cvv:             cardInfo.Cvv,
			ExpirationYear:  cardInfo.ExpirationYear,
			ExpirationMonth: cardInfo.ExpirationMonth,
		},
	})
	if err != nil {
		errMsg := fmt.Errorf("cannot charge card: %+v", err)
		log.Error().Err(errMsg)
		return "", errMsg
	}
	return resp.GetTransactionId(), nil
}
