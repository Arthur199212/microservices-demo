package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/checkout/service"
	mock_service "github.com/Arthur199212/microservices-demo/src/api_gateway/checkout/service/mocks"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestPlaceOrderOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCheckoutService(ctrl)
	validate := validator.New()

	h := NewCheckoutHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	state := "ny"
	address := models.Address{
		City:          "city",
		Country:       "country",
		State:         &state,
		StreetAddress: "street address",
		ZipCode:       "10001",
	}
	order := &checkoutv1.Order{
		TransactionId: "uuid",
		Shipping: &checkoutv1.Shipping{
			Cost: &modelsv1.Money{
				Amount:       1.34,
				CurrencyCode: "EUR",
			},
			Address: &modelsv1.Address{
				City:          address.City,
				Country:       address.Country,
				State:         *address.State,
				StreetAddress: address.StreetAddress,
				ZipCode:       address.ZipCode,
			},
		},
		Items: []*checkoutv1.OrderItem{
			&checkoutv1.OrderItem{
				Product: &modelsv1.Product{
					Id:       1,
					Quantity: 2,
				},
			},
			&checkoutv1.OrderItem{
				Product: &modelsv1.Product{
					Id:       2,
					Quantity: 6,
				},
			},
		},
	}
	args := service.CheckoutServiceArgs{
		Email:        "person@company.com",
		SessionId:    "ceaad9c4-4ebc-437e-a71c-538a5738ce11",
		UserCurrency: "EUR",
		Address:      address,
		CardInfo: service.CardInfo{
			Cvv:             "111",
			ExpirationMonth: "2",
			ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
			Number:          "4242424242424242",
		},
	}
	s.EXPECT().PlaceOrder(gomock.Any(), args).Times(1).
		Return(order, nil)

	input := fiber.Map{
		"email":        args.Email,
		"sessionId":    args.SessionId,
		"userCurrency": args.UserCurrency,
		"address": fiber.Map{
			"city":          args.Address.City,
			"country":       args.Address.Country,
			"state":         args.Address.State,
			"streetAddress": args.Address.StreetAddress,
			"zipCode":       args.Address.ZipCode,
		},
		"cardInfo": fiber.Map{
			"cvv":             args.CardInfo.Cvv,
			"expirationMonth": args.CardInfo.ExpirationMonth,
			"expirationYear":  args.CardInfo.ExpirationYear,
			"number":          args.CardInfo.Number,
		},
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/checkout/place-order", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Order Order `json:"order"`
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)

	expectedOrder := Order{
		TransactionId: order.TransactionId,
		Items:         order.Items,
		Shipping: Shipping{
			Cost: order.Shipping.Cost,
			Address: models.Address{
				City:          order.Shipping.Address.City,
				Country:       order.Shipping.Address.Country,
				State:         &order.Shipping.Address.State,
				StreetAddress: order.Shipping.Address.StreetAddress,
				ZipCode:       order.Shipping.Address.ZipCode,
			},
		},
	}
	require.Equal(t, expectedOrder, gotResp.Order)
}

func TestPlaceOrderInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCheckoutService(ctrl)
	validate := validator.New()

	h := NewCheckoutHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	state := "ny"
	address := models.Address{
		City:          "city",
		Country:       "country",
		State:         &state,
		StreetAddress: "street address",
		ZipCode:       "10001",
	}
	args := service.CheckoutServiceArgs{
		Email:        "person@company.com",
		SessionId:    "ceaad9c4-4ebc-437e-a71c-538a5738ce11",
		UserCurrency: "EUR",
		Address:      address,
		CardInfo: service.CardInfo{
			Cvv:             "111",
			ExpirationMonth: "2",
			ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
			Number:          "4242424242424242",
		},
	}
	s.EXPECT().PlaceOrder(gomock.Any(), args).Times(1).
		Return(nil, status.Errorf(codes.Internal, "mock internal error"))

	input := fiber.Map{
		"email":        args.Email,
		"sessionId":    args.SessionId,
		"userCurrency": args.UserCurrency,
		"address": fiber.Map{
			"city":          args.Address.City,
			"country":       args.Address.Country,
			"state":         args.Address.State,
			"streetAddress": args.Address.StreetAddress,
			"zipCode":       args.Address.ZipCode,
		},
		"cardInfo": fiber.Map{
			"cvv":             args.CardInfo.Cvv,
			"expirationMonth": args.CardInfo.ExpirationMonth,
			"expirationYear":  args.CardInfo.ExpirationYear,
			"number":          args.CardInfo.Number,
		},
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/checkout/place-order", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Error string
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)
	require.Contains(t, gotResp.Error, "cannot place order")
}

func TestPlaceOrderUnsupportedContentType(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCheckoutService(ctrl)
	validate := validator.New()

	h := NewCheckoutHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	state := "ny"
	address := models.Address{
		City:          "city",
		Country:       "country",
		State:         &state,
		StreetAddress: "street address",
		ZipCode:       "10001",
	}
	cardInfo := service.CardInfo{
		Cvv:             "111",
		ExpirationMonth: "2",
		ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
		Number:          "4242424242424242",
	}
	input := fiber.Map{
		"email":        "person@company.com",
		"sessionId":    "ceaad9c4-4ebc-437e-a71c-538a5738ce11",
		"userCurrency": "EUR",
		"address": fiber.Map{
			"city":          address.City,
			"country":       address.Country,
			"state":         address.State,
			"streetAddress": address.StreetAddress,
			"zipCode":       address.ZipCode,
		},
		"cardInfo": fiber.Map{
			"cvv":             cardInfo.Cvv,
			"expirationMonth": cardInfo.ExpirationMonth,
			"expirationYear":  cardInfo.ExpirationYear,
			"number":          cardInfo.Number,
		},
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/checkout/place-order", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/xml")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestPlaceOrderInvalidEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCheckoutService(ctrl)
	validate := validator.New()

	h := NewCheckoutHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	state := "ny"
	address := models.Address{
		City:          "city",
		Country:       "country",
		State:         &state,
		StreetAddress: "street address",
		ZipCode:       "10001",
	}
	cardInfo := service.CardInfo{
		Cvv:             "111",
		ExpirationMonth: "2",
		ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
		Number:          "4242424242424242",
	}
	input := fiber.Map{
		"email":        "personcompany.com",
		"sessionId":    "ceaad9c4-4ebc-437e-a71c-538a5738ce11",
		"userCurrency": "EUR",
		"address": fiber.Map{
			"city":          address.City,
			"country":       address.Country,
			"state":         address.State,
			"streetAddress": address.StreetAddress,
			"zipCode":       address.ZipCode,
		},
		"cardInfo": fiber.Map{
			"cvv":             cardInfo.Cvv,
			"expirationMonth": cardInfo.ExpirationMonth,
			"expirationYear":  cardInfo.ExpirationYear,
			"number":          cardInfo.Number,
		},
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/checkout/place-order", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestPlaceOrderInvalidSessionId(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCheckoutService(ctrl)
	validate := validator.New()

	h := NewCheckoutHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	state := "ny"
	address := models.Address{
		City:          "city",
		Country:       "country",
		State:         &state,
		StreetAddress: "street address",
		ZipCode:       "10001",
	}
	cardInfo := service.CardInfo{
		Cvv:             "111",
		ExpirationMonth: "2",
		ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
		Number:          "4242424242424242",
	}
	input := fiber.Map{
		"email":        "person@company.com",
		"sessionId":    "00000-invalid-uuid",
		"userCurrency": "EUR",
		"address": fiber.Map{
			"city":          address.City,
			"country":       address.Country,
			"state":         address.State,
			"streetAddress": address.StreetAddress,
			"zipCode":       address.ZipCode,
		},
		"cardInfo": fiber.Map{
			"cvv":             cardInfo.Cvv,
			"expirationMonth": cardInfo.ExpirationMonth,
			"expirationYear":  cardInfo.ExpirationYear,
			"number":          cardInfo.Number,
		},
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/checkout/place-order", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestPlaceOrderWithNoStateInAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCheckoutService(ctrl)
	validate := validator.New()

	h := NewCheckoutHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	address := &modelsv1.Address{
		City:          "city",
		Country:       "country",
		StreetAddress: "street address",
		ZipCode:       "10001",
	}
	order := &checkoutv1.Order{
		TransactionId: "uuid",
		Shipping: &checkoutv1.Shipping{
			Cost: &modelsv1.Money{
				Amount:       1.34,
				CurrencyCode: "EUR",
			},
			Address: address,
		},
		Items: []*checkoutv1.OrderItem{
			&checkoutv1.OrderItem{
				Product: &modelsv1.Product{
					Id:       1,
					Quantity: 2,
				},
			},
		},
	}
	args := service.CheckoutServiceArgs{
		Email:        "person@company.com",
		SessionId:    "ceaad9c4-4ebc-437e-a71c-538a5738ce11",
		UserCurrency: "EUR",
		Address: models.Address{
			City:          address.City,
			Country:       address.Country,
			StreetAddress: address.StreetAddress,
			ZipCode:       address.ZipCode,
		},
		CardInfo: service.CardInfo{
			Cvv:             "111",
			ExpirationMonth: "2",
			ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
			Number:          "4242424242424242",
		},
	}
	s.EXPECT().PlaceOrder(gomock.Any(), args).Times(1).
		Return(order, nil)

	input := fiber.Map{
		"email":        args.Email,
		"sessionId":    args.SessionId,
		"userCurrency": args.UserCurrency,
		"address": fiber.Map{
			"city":          args.Address.City,
			"country":       args.Address.Country,
			"streetAddress": args.Address.StreetAddress,
			"zipCode":       args.Address.ZipCode,
		},
		"cardInfo": fiber.Map{
			"cvv":             args.CardInfo.Cvv,
			"expirationMonth": args.CardInfo.ExpirationMonth,
			"expirationYear":  args.CardInfo.ExpirationYear,
			"number":          args.CardInfo.Number,
		},
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/checkout/place-order", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Order Order `json:"order"`
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)

	expectedOrder := Order{
		TransactionId: order.TransactionId,
		Items:         order.Items,
		Shipping: Shipping{
			Cost: order.Shipping.Cost,
			Address: models.Address{
				City:          order.Shipping.Address.City,
				Country:       order.Shipping.Address.Country,
				State:         nil,
				StreetAddress: order.Shipping.Address.StreetAddress,
				ZipCode:       order.Shipping.Address.ZipCode,
			},
		},
	}
	require.Equal(t, expectedOrder, gotResp.Order)
}
