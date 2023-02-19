package service

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
)

type CartService interface {
	AddItem(ctx context.Context, args AddItemsArgs) (string, error)
	ClearCart(ctx context.Context, sessionId string) error
	GetCart(ctx context.Context, sessionId string) ([]*modelsv1.Product, error)
}

type cartService struct {
	cartClient cartv1.CartServiceClient
}

func NewCartService(cartClient cartv1.CartServiceClient) CartService {
	return &cartService{
		cartClient: cartClient,
	}
}
