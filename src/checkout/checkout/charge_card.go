package checkout

import (
	"context"
	"fmt"

	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	payment "github.com/Arthur199212/microservices-demo/src/payment/pb"
	"github.com/rs/zerolog/log"
)

func (s *checkoutService) chargeCard(
	ctx context.Context,
	cardInfo CardInfo,
	money *pb.Money,
) (string, error) {
	resp, err := s.paymentClient.Charge(ctx, &payment.ChargeRequest{
		Money: &payment.Money{
			CurrencyCode: money.CurrencyCode,
			Amount:       money.Amount,
		},
		CardInfo: &payment.CardInfo{
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
