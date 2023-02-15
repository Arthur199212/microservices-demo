package checkout

import (
	"context"
	"fmt"

	cart "github.com/Arthur199212/microservices-demo/src/cart/pb"
	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	"github.com/rs/zerolog/log"
)

func (s *checkoutService) getCartProducts(ctx context.Context, sessionId string) ([]*pb.Product, error) {
	cartResp, err := s.cartClient.GetCart(ctx, &cart.GetCartRequest{
		SessionId: sessionId,
	})
	if err != nil {
		errMsg := fmt.Errorf("cannot retrieve cart with sessionId=%s: %+v", sessionId, err)
		log.Error().Err(err).
			Msgf(errMsg.Error())
		return nil, errMsg
	}

	products := make([]*pb.Product, len(cartResp.GetProducts()))
	for i, p := range cartResp.Products {
		products[i] = &pb.Product{
			Id:       p.GetId(),
			Quantity: p.GetQuantity(),
		}
	}
	return products, nil
}
