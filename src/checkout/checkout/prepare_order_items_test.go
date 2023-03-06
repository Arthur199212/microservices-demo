package checkout

import (
	"context"
	"fmt"
	"testing"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	mock_v1 "github.com/Arthur199212/microservices-demo/src/checkout/mocks"
	"github.com/Arthur199212/microservices-demo/src/checkout/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPrepareOrderItems(t *testing.T) {
	currency := "EUR"
	products := []*productsv1.Product{
		&productsv1.Product{
			Id:          1,
			Name:        "product",
			Description: "desc",
			Picture:     "url",
			Price:       1.23,
			Currency:    currency,
		},
		&productsv1.Product{
			Id:          2,
			Name:        "product 2",
			Description: "desc",
			Picture:     "url",
			Price:       2.27,
			Currency:    currency,
		},
	}

	testCases := []struct {
		name         string
		cartItems    []*modelsv1.Product
		userCurrency string
		setupMock    func(*mock_v1.MockProductsServiceClient)
		verify       func(*testing.T, []*checkoutv1.OrderItem, error)
	}{
		{
			name: "OK",
			cartItems: []*modelsv1.Product{
				&modelsv1.Product{
					Id:       products[0].Id,
					Quantity: 2,
				},
				&modelsv1.Product{
					Id:       products[1].Id,
					Quantity: 3,
				},
			},
			userCurrency: currency,
			setupMock: func(pc *mock_v1.MockProductsServiceClient) {
				c1 := pc.EXPECT().GetProduct(
					gomock.Any(),
					&productsv1.GetProductRequest{
						Id: products[0].Id,
					},
				).Times(1).Return(&productsv1.GetProductResponse{
					Product: products[0],
				}, nil)

				c2 := pc.EXPECT().GetProduct(
					gomock.Any(),
					&productsv1.GetProductRequest{
						Id: products[1].Id,
					},
				).Times(1).Return(&productsv1.GetProductResponse{
					Product: products[1],
				}, nil)

				gomock.InOrder(c1, c2)
			},
			verify: func(t *testing.T, res []*checkoutv1.OrderItem, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			},
		},
		{
			name: "cannot get product",
			cartItems: []*modelsv1.Product{
				&modelsv1.Product{
					Id:       products[0].Id,
					Quantity: 2,
				},
			},
			userCurrency: currency,
			setupMock: func(pc *mock_v1.MockProductsServiceClient) {
				pc.EXPECT().GetProduct(
					gomock.Any(),
					&productsv1.GetProductRequest{
						Id: products[0].Id,
					},
				).Times(1).Return(nil, fmt.Errorf("mock error"))
			},
			verify: func(t *testing.T, res []*checkoutv1.OrderItem, err error) {
				assert.Error(t, err)
				assert.Empty(t, res)
				assert.ErrorContains(t, err, "mock error")
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(func() {
				ctrl.Finish()
			})

			cartClient := mock_v1.NewMockCartServiceClient(ctrl)
			currencyClient := mock_v1.NewMockCurrencyServiceClient(ctrl)
			paymentClient := mock_v1.NewMockPaymentServiceClient(ctrl)
			productsClient := mock_v1.NewMockProductsServiceClient(ctrl)
			shippingClient := mock_v1.NewMockShippingServiceClient(ctrl)

			config := utils.Config{
				DefaultCurrency: currency,
			}
			s := NewCheckoutService(
				config,
				cartClient,
				currencyClient,
				paymentClient,
				productsClient,
				shippingClient,
			)

			test.setupMock(productsClient)

			prepareOrderItems := (*checkoutService).prepareOrderItems
			res, err := prepareOrderItems(
				s.(*checkoutService),
				context.Background(),
				test.cartItems,
				test.userCurrency,
			)

			test.verify(t, res, err)
		})
	}

}
