package checkout

import (
	"context"
	"fmt"
	"testing"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	mock_v1 "github.com/Arthur199212/microservices-demo/src/checkout/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestQuoteShipping(t *testing.T) {
	state := "state"
	address := Address{
		City:          "some city",
		Country:       "country",
		State:         &state,
		StreetAddress: "street address 1",
		ZipCode:       "00000",
	}
	products := []*modelsv1.Product{
		&modelsv1.Product{
			Id:       3,
			Quantity: 7,
		},
		&modelsv1.Product{
			Id:       4,
			Quantity: 3,
		},
	}
	money := &modelsv1.Money{
		Amount:       100.10,
		CurrencyCode: defaultCurrency,
	}

	testCases := []struct {
		name         string
		address      Address
		cartItems    []*modelsv1.Product
		userCurrency string
		setupMock    func(*mock_v1.MockShippingServiceClient)
		verify       func(*testing.T, *modelsv1.Money, error)
	}{
		{
			name:         "OK",
			address:      address,
			cartItems:    products,
			userCurrency: defaultCurrency,
			setupMock: func(s *mock_v1.MockShippingServiceClient) {
				s.EXPECT().GetQuote(
					gomock.Any(),
					&shippingv1.GetQuoteRequest{
						Address: &modelsv1.Address{
							City:          address.City,
							Country:       address.Country,
							State:         *address.State,
							StreetAddress: address.StreetAddress,
							ZipCode:       address.ZipCode,
						},
						Products: products,
					},
				).Times(1).Return(&shippingv1.GetQuoteResponse{
					Quote:        money.Amount,
					CurrencyCode: money.CurrencyCode,
				}, nil)
			},
			verify: func(t *testing.T, res *modelsv1.Money, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, res)
				assert.Equal(t, money, res)
			},
		},
		{
			name:         "cannot get shipping quote",
			address:      address,
			cartItems:    products,
			userCurrency: defaultCurrency,
			setupMock: func(s *mock_v1.MockShippingServiceClient) {
				s.EXPECT().GetQuote(
					gomock.Any(),
					&shippingv1.GetQuoteRequest{
						Address: &modelsv1.Address{
							City:          address.City,
							Country:       address.Country,
							State:         *address.State,
							StreetAddress: address.StreetAddress,
							ZipCode:       address.ZipCode,
						},
						Products: products,
					},
				).Times(1).Return(nil, fmt.Errorf("mock err"))
			},
			verify: func(t *testing.T, res *modelsv1.Money, err error) {
				assert.Error(t, err)
				assert.Empty(t, res)
				assert.ErrorContains(t, err, "mock err")
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
				cartClient,
				currencyClient,
				paymentClient,
				productsClient,
				shippingClient,
			)

			test.setupMock(shippingClient)

			quoteShipping := (*checkoutService).quoteShipping
			res, err := quoteShipping(
				s.(*checkoutService),
				context.Background(),
				test.address,
				test.cartItems,
				test.userCurrency,
			)

			test.verify(t, res, err)
		})
	}
}
