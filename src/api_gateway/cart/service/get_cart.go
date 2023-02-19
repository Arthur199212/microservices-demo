package service

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
)

func (s *cartService) GetCart(
	ctx context.Context,
	sessionId string,
) ([]*modelsv1.Product, error) {
	resp, err := s.cartClient.GetCart(ctx, &cartv1.GetCartRequest{
		SessionId: sessionId,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetProducts(), nil
}
