package checkout

import (
	"context"
	"fmt"

	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	ss "github.com/Arthur199212/microservices-demo/src/shipping/shipping"
)

type CardInfo struct {
	Cvv             int32  `validate:"required,numeric,min=3,max=4"`
	ExpirationMonth int32  `validate:"required,numeric,len=2"`
	ExpirationYear  int32  `validate:"required,numeric,len=2"`
	Number          string `validate:"required,numeric,min=8,max=19"`
}

type PlaceOrderArgs struct {
	Address      ss.Address
	CardInfo     CardInfo
	Email        string `validate:"required,email"`
	SessionId    string `validate:"required,uuid4"`
	UserCurrency string `validate:"required,len=3"`
}

func (s *checkoutService) PlaceOrder(
	ctx context.Context,
	args PlaceOrderArgs,
) (*pb.Order, error) {
	products, err := s.getCartProducts(ctx, args.SessionId)
	if err != nil {
		return nil, fmt.Errorf("cart failure: %+v", err)
	}

	orderItems, err := s.prepareOrderItems(ctx, products, args.UserCurrency)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare order items: %+v", err)
	}

	shippingCost, err := s.quoteShipping(ctx, args.Address, products, args.UserCurrency)
	if err != nil {
		return nil, err
	}

	totalSum := &pb.Money{
		CurrencyCode: defaultCurrency,
		Amount:       shippingCost.Amount,
	}
	for _, item := range orderItems {
		totalSum.Amount += item.Cost.Amount
	}

	transactionId, err := s.chargeCard(ctx, args.CardInfo, totalSum)
	if err != nil {
		return nil, err
	}

	order := &pb.Order{
		TransactionId: transactionId,
		Shipping: &pb.Shipping{
			Cost: shippingCost,
			Address: &pb.Address{
				StreetAddress: args.Address.StreetAddress,
				City:          args.Address.City,
				State:         *args.Address.State,
				Country:       args.Address.Country,
				ZipCode:       args.Address.ZipCode,
			},
		},
		Items: orderItems,
	}
	return order, nil
}
