package service

import (
	"context"

	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
)

func (s *cartService) ClearCart(ctx context.Context, sessionId string) error {
	_, err := s.cartClient.ClearCart(ctx, &cartv1.ClearCartRequest{
		SessionId: sessionId,
	})
	return err
}
