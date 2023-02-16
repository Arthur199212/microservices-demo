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
	cartProducts []*modelsv1.Product,
	userCurrency string,
) (*modelsv1.Money, error) {
	products := make([]*modelsv1.Product, len(cartProducts))
	for i, p := range cartProducts {
		products[i] = &modelsv1.Product{
			Id:       p.Id,
			Quantity: p.Quantity,
		}
	}
	shippingAddress := &modelsv1.Address{
		StreetAddress: address.StreetAddress,
		City:          address.City,
		State:         *address.State,
		Country:       address.Country,
		ZipCode:       address.ZipCode,
	}

	resp, err := s.shippingClient.GetQuote(ctx, &shippingv1.GetQuoteRequest{
		Address:  shippingAddress,
		Products: products,
	})
	if err != nil {
		errMsg := fmt.Errorf("cannot quote shipping: %+v", err)
		log.Error().Err(errMsg)
		return nil, errMsg
	}

	money, err := s.convertCurrency(ctx, defaultCurrency, userCurrency, resp.GetQuote())
	if err != nil {
		errMsg := fmt.Errorf("cannot conver currency for quote shipping: %+v", err)
		log.Error().Err(errMsg)
		return nil, errMsg
	}
	return money, nil
}
