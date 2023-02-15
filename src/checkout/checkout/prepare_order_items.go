package checkout

import (
	"context"
	"fmt"

	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	products "github.com/Arthur199212/microservices-demo/src/products/pb"
	"github.com/rs/zerolog/log"
)

const (
	defaultCurrency = "EUR"
)

func (s *checkoutService) prepareOrderItems(
	ctx context.Context,
	cartProducts []*pb.Product,
	userCurrency string,
) ([]*pb.OrderItem, error) {
	orderItems := make([]*pb.OrderItem, len(cartProducts))
	for i, product := range cartProducts {
		resp, err := s.productsClient.GetProduct(ctx, &products.GetProductRequest{
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

		orderItems[i] = &pb.OrderItem{
			Product: &pb.Product{
				Id:       product.Id,
				Quantity: product.Quantity,
			},
			Cost: money,
		}
	}

	return orderItems, nil
}
