package checkout

import (
	"context"
	"fmt"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	"github.com/rs/zerolog/log"
)

const (
	defaultCurrency = "EUR"
)

func (s *checkoutService) prepareOrderItems(
	ctx context.Context,
	cartProducts []*modelsv1.Product,
	userCurrency string,
) ([]*checkoutv1.OrderItem, error) {
	orderItems := make([]*checkoutv1.OrderItem, len(cartProducts))
	for i, product := range cartProducts {
		resp, err := s.productsClient.GetProduct(ctx, &productsv1.GetProductRequest{
			Id: product.Id,
		})
		if err != nil {
			errMsg := fmt.Errorf("cannot get product with ID=%d: %+v", product.Id, err)
			log.Error().Err(err).
				Msgf(errMsg.Error())
			return nil, errMsg
		}

		money, err := s.convertCurrency(ctx, defaultCurrency, userCurrency, resp.Product.Price)
		if err != nil {
			errMsg := fmt.Errorf("failed to convert currency for product with ID=%d: %+v", product.Id, err)
			log.Error().Err(err).
				Msgf(errMsg.Error())
			return nil, errMsg
		}

		orderItems[i] = &checkoutv1.OrderItem{
			Product: &modelsv1.Product{
				Id:       product.Id,
				Quantity: product.Quantity,
			},
			Cost: money,
		}
	}

	return orderItems, nil
}
