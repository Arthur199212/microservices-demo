package checkout

import (
	"context"
	"fmt"
	"testing"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	mock_v1 "github.com/Arthur199212/microservices-demo/src/checkout/mocks"
	"github.com/Arthur199212/microservices-demo/src/checkout/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetCartItems(t *testing.T) {
	sessionId := "mock session id"
	products := []*modelsv1.Product{
		&modelsv1.Product{
			Id:       1,
			Quantity: 3,
		},
		&modelsv1.Product{
			Id:       2,
			Quantity: 7,
		},
	}

	testCases := []struct {
		name           string
		money          *modelsv1.Money
		toCurrencyCode string
		setupMock      func(*mock_v1.MockCartServiceClient)
		verify         func(*testing.T, []*modelsv1.Product, error)
	}{
		{
			name: "OK",
			setupMock: func(s *mock_v1.MockCartServiceClient) {
				s.EXPECT().GetCart(
					gomock.Any(),
					&cartv1.GetCartRequest{
						SessionId: sessionId,
					},
				).Times(1).Return(&cartv1.GetCartResponse{
					SessionId: sessionId,
					Products:  products,
				}, nil)
			},
			verify: func(t *testing.T, res []*modelsv1.Product, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, res)
				assert.Equal(t, products, res)
			},
		},
		{
			name: "failure",
			setupMock: func(s *mock_v1.MockCartServiceClient) {
				s.EXPECT().GetCart(
					gomock.Any(),
					&cartv1.GetCartRequest{
						SessionId: sessionId,
					},
				).Times(1).Return(nil, fmt.Errorf("mock error"))
			},
			verify: func(t *testing.T, res []*modelsv1.Product, err error) {
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

			s := NewCheckoutService(
				utils.Config{},
				cartClient,
				currencyClient,
				paymentClient,
				productsClient,
				shippingClient,
			)

			test.setupMock(cartClient)

			getCartItems := (*checkoutService).getCartItems
			res, err := getCartItems(
				s.(*checkoutService),
				context.Background(),
				sessionId,
			)

			test.verify(t, res, err)
		})
	}
}
