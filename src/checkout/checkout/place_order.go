package checkout

import (
	"context"
	"fmt"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
)

type Address struct {
	StreetAddress string  `validate:"required,min=5,max=64"`
	City          string  `validate:"required,min=2,max=64"`
	Country       string  `validate:"required,min=2,max=64"`
	ZipCode       string  `validate:"required,numeric,min=4,max=10"`
	State         *string `validate:"omitempty,min=2,max=64"`
}

type CardInfo struct {
	Cvv             string `validate:"required,numeric,min=3,max=4"`
	ExpirationMonth string `validate:"required,numeric,min=1,max=2"`
	ExpirationYear  string `validate:"required,numeric,len=4"`
	Number          string `validate:"required,numeric,min=8,max=19"`
}

type PlaceOrderArgs struct {
	Address      Address
	CardInfo     CardInfo
	Email        string `validate:"required,email"`
	SessionId    string `validate:"required,uuid4"`
	UserCurrency string `validate:"required,len=3"`
}

func (s *checkoutService) PlaceOrder(
	ctx context.Context,
	args PlaceOrderArgs,
) (*checkoutv1.Order, error) {
	items, err := s.getCartItems(ctx, args.SessionId)
	if err != nil {
		return nil, fmt.Errorf("cart failure: %+v", err)
	}

	orderItems, err := s.prepareOrderItems(ctx, items, args.UserCurrency)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare order items: %+v", err)
	}

	shippingCost, err := s.quoteShipping(ctx, args.Address, items, args.UserCurrency)
	if err != nil {
		return nil, err
	}

	totalSum := &modelsv1.Money{
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

	state := ""
	if args.Address.State != nil {
		state = *args.Address.State
	}
	order := &checkoutv1.Order{
		TransactionId: transactionId,
		Shipping: &checkoutv1.Shipping{
			Cost: shippingCost,
			Address: &modelsv1.Address{
				StreetAddress: args.Address.StreetAddress,
				City:          args.Address.City,
				State:         state,
				Country:       args.Address.Country,
				ZipCode:       args.Address.ZipCode,
			},
		},
		Items: orderItems,
	}
	return order, nil
}
