package checkout

import (
	"context"
	"fmt"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	"github.com/rs/zerolog/log"
)

func (s *checkoutService) getCartProducts(
	ctx context.Context,
	sessionId string,
) ([]*modelsv1.Product, error) {
	cartResp, err := s.cartClient.GetCart(ctx, &cartv1.GetCartRequest{
		SessionId: sessionId,
	})
	if err != nil {
		errMsg := fmt.Errorf("cannot retrieve cart with sessionId=%s: %+v", sessionId, err)
		log.Error().Err(err).
			Msgf(errMsg.Error())
		return nil, errMsg
	}

	products := make([]*modelsv1.Product, len(cartResp.GetProducts()))
	for i, p := range cartResp.Products {
		products[i] = &modelsv1.Product{
			Id:       p.GetId(),
			Quantity: p.GetQuantity(),
		}
	}
	return products, nil
}
