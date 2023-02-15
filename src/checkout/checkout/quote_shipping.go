package checkout

import (
	"context"
	"fmt"

	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	shipping "github.com/Arthur199212/microservices-demo/src/shipping/pb"
	ss "github.com/Arthur199212/microservices-demo/src/shipping/shipping"
	"github.com/rs/zerolog/log"
)

func (s *checkoutService) quoteShipping(
	ctx context.Context,
	address ss.Address,
	cartProducts []*pb.Product,
	userCurrency string,
) (*pb.Money, error) {
	products := make([]*shipping.Product, len(cartProducts))
	for i, p := range cartProducts {
		products[i] = &shipping.Product{
			Id:       p.Id,
			Quantity: p.Quantity,
		}
	}
	shippingAddress := &shipping.Address{
		StreetAddress: address.StreetAddress,
		City:          address.City,
		State:         *address.State,
		Country:       address.Country,
		ZipCode:       address.ZipCode,
	}

	resp, err := s.shippingClient.GetQuote(ctx, &shipping.GetQuoteRequest{
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
