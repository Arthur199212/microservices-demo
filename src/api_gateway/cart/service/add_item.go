package service

import (
	"context"

	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	"github.com/google/uuid"
)

type AddItemsArgs struct {
	ProductId int64   `json:"productId" validate:"required,min=1"`
	Quantity  int32   `json:"quantity" validate:"required,min=1"`
	SessionId *string `json:"sessionId,omitempty" validate:"omitempty,uuid4"`
}

func (s *cartService) AddItem(
	ctx context.Context,
	args AddItemsArgs,
) (string, error) {
	sessionId := args.SessionId
	// create sessionId if a user doesn't have one
	if sessionId == nil {
		newSessionId := uuid.New().String()
		sessionId = &newSessionId
	}

	_, err := s.cartClient.AddItem(ctx, &cartv1.AddItemRequest{
		SessionId: *sessionId,
		ProducId:  args.ProductId,
		Quantity:  args.Quantity,
	})
	return *sessionId, err
}
