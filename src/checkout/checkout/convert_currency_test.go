package checkout

import (
	"context"
	"fmt"
	"testing"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	mock_v1 "github.com/Arthur199212/microservices-demo/src/checkout/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConvertCurrency(t *testing.T) {
	fromMoney := &modelsv1.Money{
		Amount:       100.29,
		CurrencyCode: "EUR",
	}
	toMoney := &modelsv1.Money{
		Amount:       110.45,
		CurrencyCode: "USD",
	}

	testCases := []struct {
		name           string
		money          *modelsv1.Money
		toCurrencyCode string
		setupMock      func(*mock_v1.MockCurrencyServiceClient)
		verify         func(*testing.T, *modelsv1.Money, error)
	}{
		{
			name:           "OK",
			money:          fromMoney,
			toCurrencyCode: toMoney.CurrencyCode,
			setupMock: func(s *mock_v1.MockCurrencyServiceClient) {
				s.EXPECT().Convert(
					gomock.Any(),
					&currencyv1.ConvertRequest{
						From:           fromMoney,
						ToCurrencyCode: toMoney.CurrencyCode,
					},
				).Times(1).Return(&currencyv1.ConvertResponse{
					Money: toMoney,
				}, nil)
			},
			verify: func(t *testing.T, res *modelsv1.Money, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, res)
				assert.Equal(t, toMoney, res)
			},
		},
		{
			name:           "same currency",
			money:          fromMoney,
			toCurrencyCode: fromMoney.CurrencyCode,
			setupMock:      func(s *mock_v1.MockCurrencyServiceClient) {},
			verify: func(t *testing.T, res *modelsv1.Money, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, res)
				assert.Equal(t, fromMoney, res)
			},
		},
		{
			name:           "failure",
			money:          fromMoney,
			toCurrencyCode: toMoney.CurrencyCode,
			setupMock: func(s *mock_v1.MockCurrencyServiceClient) {
				s.EXPECT().Convert(
					gomock.Any(),
					&currencyv1.ConvertRequest{
						From:           fromMoney,
						ToCurrencyCode: toMoney.CurrencyCode,
					},
				).Times(1).Return(nil, fmt.Errorf("mock error"))
			},
			verify: func(t *testing.T, res *modelsv1.Money, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "mock error")
				assert.Empty(t, res)
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

			test.setupMock(currencyClient)

			convertCurrency := (*checkoutService).convertCurrency
			res, err := convertCurrency(
				s.(*checkoutService),
				context.Background(),
				test.money,
				test.toCurrencyCode,
			)

			test.verify(t, res, err)
		})
	}
}
