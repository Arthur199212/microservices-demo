package checkout

import (
	"context"
	"strconv"
	"testing"
	"time"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	mock_v1 "github.com/Arthur199212/microservices-demo/src/checkout/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPlaceOrder(t *testing.T) {
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

	args := PlaceOrderArgs{
		Address: Address{
			StreetAddress: "some street address",
			City:          "city",
			Country:       "country",
			ZipCode:       "00000",
		},
		CardInfo: CardInfo{
			Cvv:             "1111",
			ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
			ExpirationYear:  strconv.Itoa(time.Now().Year()),
			Number:          "0123456789",
		},
		Email:        "person@company.com",
		SessionId:    "session uuid",
		UserCurrency: defaultCurrency,
	}
	products := []*productsv1.Product{
		&productsv1.Product{
			Id:          1,
			Name:        "name",
			Description: "desc",
			Picture:     "url",
			Price:       1.49,
			Currency:    defaultCurrency,
		},
	}
	transactionId := "transaction uuid"

	// setup mocks
	cartClient.EXPECT().GetCart(
		gomock.Any(),
		&cartv1.GetCartRequest{
			SessionId: args.SessionId,
		},
	).Times(1).Return(&cartv1.GetCartResponse{
		SessionId: args.SessionId,
		Products: []*modelsv1.Product{
			&modelsv1.Product{
				Id:       products[0].Id,
				Quantity: 3,
			},
		},
	}, nil)

	productsClient.EXPECT().GetProduct(
		gomock.Any(),
		&productsv1.GetProductRequest{
			Id: products[0].Id,
		},
	).Times(1).Return(&productsv1.GetProductResponse{
		Product: products[0],
	}, nil)

	shippingClient.EXPECT().GetQuote(
		gomock.Any(),
		&shippingv1.GetQuoteRequest{
			Address: &modelsv1.Address{
				City:          args.Address.City,
				Country:       args.Address.Country,
				StreetAddress: args.Address.StreetAddress,
				ZipCode:       args.Address.ZipCode,
			},
			Products: []*modelsv1.Product{
				&modelsv1.Product{
					Id:       products[0].Id,
					Quantity: 3,
				},
			},
		},
	).Times(1).Return(&shippingv1.GetQuoteResponse{
		Quote:        0.68,
		CurrencyCode: defaultCurrency,
	}, nil)

	paymentClient.EXPECT().Charge(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&paymentv1.ChargeResponse{
			TransactionId: transactionId,
		}, nil)

	// verify
	order, err := s.PlaceOrder(context.Background(), args)

	assert.NoError(t, err)
	assert.NotEmpty(t, order)
	assert.Equal(t, order.TransactionId, transactionId)
}
