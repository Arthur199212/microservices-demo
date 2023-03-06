package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/shipping/service"
	mock_shipping "github.com/Arthur199212/microservices-demo/src/api_gateway/shipping/service/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetQuote(t *testing.T) {
	userCurrency := "USD"
	products := []*modelsv1.Product{
		&modelsv1.Product{Id: 1, Quantity: 1},
		&modelsv1.Product{Id: 2, Quantity: 6},
	}
	address := &modelsv1.Address{
		City:          "city",
		Country:       "country",
		State:         "ny",
		StreetAddress: "street address",
		ZipCode:       "10001",
	}
	var quote float32 = 3.56

	testCases := []struct {
		name       string
		input      fiber.Map
		setupMock  func(s *mock_shipping.MockShippingService)
		newRequest func() *http.Request
		verify     func(t *testing.T, resp *http.Response, err error)
	}{
		{
			name: "OK",
			setupMock: func(s *mock_shipping.MockShippingService) {
				s.EXPECT().GetQuote(gomock.Any(), service.GetQuoteArgs{
					Address:      address,
					Products:     products,
					UserCurrency: userCurrency,
				}).Times(1).Return(
					&modelsv1.Money{
						Amount:       quote,
						CurrencyCode: userCurrency,
					},
					nil,
				)
			},
			newRequest: func() *http.Request {
				input := fiber.Map{
					"userCurrency": userCurrency,
					"address": fiber.Map{
						"city":          address.City,
						"country":       address.Country,
						"state":         address.State,
						"streetAddress": address.StreetAddress,
						"zipCode":       address.ZipCode,
					},
					"products": products,
				}
				body, err := json.Marshal(input)
				require.NoError(t, err)

				req, err := http.NewRequest("POST", "/shipping/quote", bytes.NewReader(body))
				require.NoError(t, err)
				req.Header.Add("content-type", "application/json")
				return req
			},
			verify: func(t *testing.T, resp *http.Response, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, resp.StatusCode)
				require.NotEmpty(t, resp.Body)

				data, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				t.Cleanup(func() {
					resp.Body.Close()
				})

				var gotResp struct {
					Currency string  `json:"currency"`
					Quote    float32 `json:"quote"`
				}
				err = json.Unmarshal(data, &gotResp)
				require.NoError(t, err)

				require.Equal(t, userCurrency, gotResp.Currency)
				require.Equal(t, quote, gotResp.Quote)
			},
		},
		{
			name:      "invalid payload",
			setupMock: func(s *mock_shipping.MockShippingService) {},
			newRequest: func() *http.Request {
				body, err := json.Marshal(fiber.Map{})
				require.NoError(t, err)

				req, err := http.NewRequest("POST", "/shipping/quote", bytes.NewReader(body))
				require.NoError(t, err)
				req.Header.Add("content-type", "application/xml")
				return req
			},
			verify: func(t *testing.T, resp *http.Response, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
		{
			name:      "products are empty",
			setupMock: func(s *mock_shipping.MockShippingService) {},
			newRequest: func() *http.Request {
				input := fiber.Map{
					"userCurrency": userCurrency,
					"address": fiber.Map{
						"city":          address.City,
						"country":       address.Country,
						"state":         address.State,
						"streetAddress": address.StreetAddress,
						"zipCode":       address.ZipCode,
					},
					"products": []*modelsv1.Product{},
				}
				body, err := json.Marshal(input)
				require.NoError(t, err)

				req, err := http.NewRequest("POST", "/shipping/quote", bytes.NewReader(body))
				require.NoError(t, err)
				req.Header.Add("content-type", "application/json")
				return req
			},
			verify: func(t *testing.T, resp *http.Response, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
		{
			name:      "missing userCurrency",
			setupMock: func(s *mock_shipping.MockShippingService) {},
			newRequest: func() *http.Request {
				input := fiber.Map{
					"address": fiber.Map{
						"city":          address.City,
						"country":       address.Country,
						"state":         address.State,
						"streetAddress": address.StreetAddress,
						"zipCode":       address.ZipCode,
					},
					"products": products,
				}
				body, err := json.Marshal(input)
				require.NoError(t, err)

				req, err := http.NewRequest("POST", "/shipping/quote", bytes.NewReader(body))
				require.NoError(t, err)
				req.Header.Add("content-type", "application/json")
				return req
			},
			verify: func(t *testing.T, resp *http.Response, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
		{
			name:      "invalid zipCode",
			setupMock: func(s *mock_shipping.MockShippingService) {},
			newRequest: func() *http.Request {
				input := fiber.Map{
					"address": fiber.Map{
						"city":          address.City,
						"country":       address.Country,
						"state":         address.State,
						"streetAddress": address.StreetAddress,
						"zipCode":       "abcdef",
					},
					"products": products,
				}
				body, err := json.Marshal(input)
				require.NoError(t, err)

				req, err := http.NewRequest("POST", "/shipping/quote", bytes.NewReader(body))
				require.NoError(t, err)
				req.Header.Add("content-type", "application/json")
				return req
			},
			verify: func(t *testing.T, resp *http.Response, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
		{
			name: "empty address.state",
			setupMock: func(s *mock_shipping.MockShippingService) {
				s.EXPECT().GetQuote(gomock.Any(), service.GetQuoteArgs{
					Address: &modelsv1.Address{
						City:          address.City,
						Country:       address.Country,
						State:         "",
						StreetAddress: address.StreetAddress,
						ZipCode:       address.ZipCode,
					},
					Products:     products,
					UserCurrency: userCurrency,
				}).Times(1).Return(
					&modelsv1.Money{
						Amount:       quote,
						CurrencyCode: userCurrency,
					},
					nil,
				)
			},
			newRequest: func() *http.Request {
				input := fiber.Map{
					"userCurrency": userCurrency,
					"address": fiber.Map{
						"city":          address.City,
						"country":       address.Country,
						"streetAddress": address.StreetAddress,
						"zipCode":       address.ZipCode,
					},
					"products": products,
				}
				body, err := json.Marshal(input)
				require.NoError(t, err)

				req, err := http.NewRequest("POST", "/shipping/quote", bytes.NewReader(body))
				require.NoError(t, err)
				req.Header.Add("content-type", "application/json")
				return req
			},
			verify: func(t *testing.T, resp *http.Response, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, resp.StatusCode)
				require.NotEmpty(t, resp.Body)

				data, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				t.Cleanup(func() {
					resp.Body.Close()
				})

				var gotResp struct {
					Currency string  `json:"currency"`
					Quote    float32 `json:"quote"`
				}
				err = json.Unmarshal(data, &gotResp)
				require.NoError(t, err)

				require.Equal(t, userCurrency, gotResp.Currency)
				require.Equal(t, quote, gotResp.Quote)
			},
		},
		{
			name: "shipping service returns invalid-argument error",
			setupMock: func(s *mock_shipping.MockShippingService) {
				s.EXPECT().GetQuote(gomock.Any(), service.GetQuoteArgs{
					Address: &modelsv1.Address{
						City:          address.City,
						Country:       address.Country,
						State:         address.State,
						StreetAddress: address.StreetAddress,
						ZipCode:       address.ZipCode,
					},
					Products:     products,
					UserCurrency: userCurrency,
				}).Times(1).Return(nil, status.Errorf(codes.InvalidArgument, "mock error"))
			},
			newRequest: func() *http.Request {
				input := fiber.Map{
					"userCurrency": userCurrency,
					"address": fiber.Map{
						"city":          address.City,
						"country":       address.Country,
						"state":         address.State,
						"streetAddress": address.StreetAddress,
						"zipCode":       address.ZipCode,
					},
					"products": products,
				}
				body, err := json.Marshal(input)
				require.NoError(t, err)

				req, err := http.NewRequest("POST", "/shipping/quote", bytes.NewReader(body))
				require.NoError(t, err)
				req.Header.Add("content-type", "application/json")
				return req
			},
			verify: func(t *testing.T, resp *http.Response, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)

				data, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				t.Cleanup(func() {
					resp.Body.Close()
				})

				var gotResp struct {
					Error string `json:"error"`
				}
				err = json.Unmarshal(data, &gotResp)
				require.NoError(t, err)
				require.Contains(t, gotResp.Error, "mock error")
			},
		},
		{
			name: "shipping service returns internal error",
			setupMock: func(s *mock_shipping.MockShippingService) {
				s.EXPECT().GetQuote(gomock.Any(), service.GetQuoteArgs{
					Address: &modelsv1.Address{
						City:          address.City,
						Country:       address.Country,
						State:         address.State,
						StreetAddress: address.StreetAddress,
						ZipCode:       address.ZipCode,
					},
					Products:     products,
					UserCurrency: userCurrency,
				}).Times(1).Return(nil, status.Errorf(codes.Internal, "mock internal error"))
			},
			newRequest: func() *http.Request {
				input := fiber.Map{
					"userCurrency": userCurrency,
					"address": fiber.Map{
						"city":          address.City,
						"country":       address.Country,
						"state":         address.State,
						"streetAddress": address.StreetAddress,
						"zipCode":       address.ZipCode,
					},
					"products": products,
				}
				body, err := json.Marshal(input)
				require.NoError(t, err)

				req, err := http.NewRequest("POST", "/shipping/quote", bytes.NewReader(body))
				require.NoError(t, err)
				req.Header.Add("content-type", "application/json")
				return req
			},
			verify: func(t *testing.T, resp *http.Response, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusInternalServerError, resp.StatusCode)

				data, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				t.Cleanup(func() {
					resp.Body.Close()
				})

				var gotResp struct {
					Error string `json:"error"`
				}
				err = json.Unmarshal(data, &gotResp)
				require.NoError(t, err)
				require.Contains(t, gotResp.Error, "cannot get shipping quote")
			},
		},
	}

	for _, test := range testCases {
		ctrl := gomock.NewController(t)
		t.Cleanup(func() {
			ctrl.Finish()
		})

		s := mock_shipping.NewMockShippingService(ctrl)
		validate := validator.New()
		handler := NewShippingHandler(s, validate)

		app := fiber.New()
		handler.AddRoutes(app)

		test.setupMock(s)

		resp, err := app.Test(test.newRequest())
		test.verify(t, resp, err)
	}
}
