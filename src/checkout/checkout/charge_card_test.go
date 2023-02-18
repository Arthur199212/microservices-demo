package checkout

import (
	"context"
	"fmt"
	"testing"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	mock_v1 "github.com/Arthur199212/microservices-demo/src/checkout/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCargeCard(t *testing.T) {
	cardInfo := CardInfo{
		Cvv:             "0123",
		ExpirationMonth: "2",
		ExpirationYear:  "2023",
		Number:          "0123456789012",
	}
	money := &modelsv1.Money{
		Amount:       100,
		CurrencyCode: "EUR",
	}
	mockTransactionId := "uuid mock"

	testCases := []struct {
		name          string
		cardInfo      CardInfo
		money         *modelsv1.Money
		transactionId string
		setupMock     func(*mock_v1.MockPaymentServiceClient)
		verify        func(*testing.T, string, error)
	}{
		{
			name:          "OK",
			cardInfo:      cardInfo,
			money:         money,
			transactionId: mockTransactionId,
			setupMock: func(ps *mock_v1.MockPaymentServiceClient) {
				ps.EXPECT().Charge(gomock.Any(), &paymentv1.ChargeRequest{
					Money: money,
					CardInfo: &modelsv1.CardInfo{
						Cvv:             cardInfo.Cvv,
						ExpirationMonth: cardInfo.ExpirationMonth,
						ExpirationYear:  cardInfo.ExpirationYear,
						Number:          cardInfo.Number,
					},
				}).Times(1).Return(&paymentv1.ChargeResponse{
					TransactionId: mockTransactionId,
				}, nil)
			},
			verify: func(t *testing.T, transactionId string, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, transactionId)
				assert.Equal(t, mockTransactionId, transactionId)
			},
		},
		{
			name:          "failure",
			cardInfo:      cardInfo,
			money:         money,
			transactionId: mockTransactionId,
			setupMock: func(ps *mock_v1.MockPaymentServiceClient) {
				ps.EXPECT().Charge(gomock.Any(), &paymentv1.ChargeRequest{
					Money: money,
					CardInfo: &modelsv1.CardInfo{
						Cvv:             cardInfo.Cvv,
						ExpirationMonth: cardInfo.ExpirationMonth,
						ExpirationYear:  cardInfo.ExpirationYear,
						Number:          cardInfo.Number,
					},
				}).Times(1).Return(&paymentv1.ChargeResponse{
					TransactionId: "",
				}, fmt.Errorf("mock error"))
			},
			verify: func(t *testing.T, transactionId string, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "mock error")
				assert.Empty(t, transactionId)
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

			test.setupMock(paymentClient)

			chargeCard := (*checkoutService).chargeCard
			transactionId, err := chargeCard(s.(*checkoutService), context.Background(), cardInfo, money)

			test.verify(t, transactionId, err)
		})
	}
}
