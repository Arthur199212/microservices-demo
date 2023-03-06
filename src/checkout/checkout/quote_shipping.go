package checkout

import (
	"context"
	"fmt"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"github.com/rs/zerolog/log"
)

func (s *checkoutService) quoteShipping(
	ctx context.Context,
	address Address,
	cartItems []*modelsv1.Product,
	userCurrency string,
) (*modelsv1.Money, error) {
	products := make([]*modelsv1.Product, len(cartItems))
	for i, p := range cartItems {
		products[i] = &modelsv1.Product{
			Id:       p.Id,
			Quantity: p.Quantity,
		}
	}
	state := ""
	if address.State != nil {
		state = *address.State
	}
	shippingAddress := &modelsv1.Address{
		StreetAddress: address.StreetAddress,
		City:          address.City,
		State:         state,
		Country:       address.Country,
		ZipCode:       address.ZipCode,
	}

	resp, err := s.shippingClient.GetQuote(ctx, &shippingv1.GetQuoteRequest{
		Address:  shippingAddress,
		Products: products,
	})
	if err != nil {
		err = fmt.Errorf("cannot shipping quote: %+v", err)
		log.Error().Err(err)
		return nil, err
	}

	quote := &modelsv1.Money{
		Amount:       resp.GetQuote(),
		CurrencyCode: s.config.DefaultCurrency,
	}
	money, err := s.convertCurrency(ctx, quote, userCurrency)
	if err != nil {
		err = fmt.Errorf("cannot conver currency for quote shipping: %+v", err)
		log.Error().Err(err)
		return nil, err
	}
	return money, nil
}
