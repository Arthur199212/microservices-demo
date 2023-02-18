package gapi

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"github.com/Arthur199212/microservices-demo/src/shipping/shipping"
	mock_shipping "github.com/Arthur199212/microservices-demo/src/shipping/shipping/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestGetQuote(t *testing.T) {
	stateMock := "state"
	mockAddress := shipping.Address{
		StreetAddress: "some address",
		City:          "some city",
		Country:       "country",
		ZipCode:       "00000",
		State:         &stateMock,
	}
	productOne := shipping.Product{
		ID:       1,
		Quantity: 5,
	}
	productTwo := shipping.Product{
		ID:       2,
		Quantity: 10,
	}
	quoteMock := shipping.Quote{
		Quote:        5.49,
		CurrencyCode: "EUR",
	}

	testCases := []struct {
		name      string
		address   *modelsv1.Address
		products  []*modelsv1.Product
		setupMock func(*mock_shipping.MockShippingService)
		verify    func(*shippingv1.GetQuoteResponse, error)
	}{
		{
			name: "OK",
			address: &modelsv1.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products: []*modelsv1.Product{
				&modelsv1.Product{
					Id:       productOne.ID,
					Quantity: productOne.Quantity,
				},
				&modelsv1.Product{
					Id:       productTwo.ID,
					Quantity: productTwo.Quantity,
				},
			},
			setupMock: func(s *mock_shipping.MockShippingService) {
				s.
					EXPECT().
					GetQuote(mockAddress, []shipping.Product{productOne, productTwo}).
					Times(1).
					Return(quoteMock, nil)
			},
			verify: func(resp *shippingv1.GetQuoteResponse, err error) {
				assert.NoError(t, err)
				assert.Equal(t, quoteMock.CurrencyCode, resp.CurrencyCode)
				assert.Equal(t, quoteMock.Quote, resp.Quote)
			},
		},
		{
			name: "state in the address is empty",
			address: &modelsv1.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         "",
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products: []*modelsv1.Product{
				&modelsv1.Product{
					Id:       productOne.ID,
					Quantity: productOne.Quantity,
				},
				&modelsv1.Product{
					Id:       productTwo.ID,
					Quantity: productTwo.Quantity,
				},
			},
			setupMock: func(s *mock_shipping.MockShippingService) {
				s.
					EXPECT().
					GetQuote(
						shipping.Address{
							StreetAddress: mockAddress.StreetAddress,
							City:          mockAddress.City,
							State:         nil,
							Country:       mockAddress.Country,
							ZipCode:       mockAddress.ZipCode,
						},
						[]shipping.Product{productOne, productTwo},
					).
					Times(1).
					Return(quoteMock, nil)
			},
			verify: func(resp *shippingv1.GetQuoteResponse, err error) {
				assert.NoError(t, err)
				assert.Equal(t, quoteMock.CurrencyCode, resp.CurrencyCode)
				assert.Equal(t, quoteMock.Quote, resp.Quote)
			},
		},
		{
			name: "no products to ship",
			address: &modelsv1.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products:  []*modelsv1.Product{},
			setupMock: func(mss *mock_shipping.MockShippingService) {},
			verify: func(resp *shippingv1.GetQuoteResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "too many products to ship",
			address: &modelsv1.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products:  make([]*modelsv1.Product, maxProductsToShip+1),
			setupMock: func(mss *mock_shipping.MockShippingService) {},
			verify: func(resp *shippingv1.GetQuoteResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "invalid quantity",
			address: &modelsv1.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products: []*modelsv1.Product{
				&modelsv1.Product{
					Id:       productOne.ID,
					Quantity: -1,
				},
			},
			setupMock: func(s *mock_shipping.MockShippingService) {},
			verify: func(resp *shippingv1.GetQuoteResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "invalid streetAddress in the address",
			address: &modelsv1.Address{
				StreetAddress: "aa",
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products: []*modelsv1.Product{
				&modelsv1.Product{
					Id:       productOne.ID,
					Quantity: productOne.Quantity,
				},
			},
			setupMock: func(s *mock_shipping.MockShippingService) {},
			verify: func(resp *shippingv1.GetQuoteResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "internal service error",
			address: &modelsv1.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products: []*modelsv1.Product{
				&modelsv1.Product{
					Id:       productOne.ID,
					Quantity: productOne.Quantity,
				},
				&modelsv1.Product{
					Id:       productTwo.ID,
					Quantity: productTwo.Quantity,
				},
			},
			setupMock: func(s *mock_shipping.MockShippingService) {
				s.
					EXPECT().
					GetQuote(mockAddress, []shipping.Product{productOne, productTwo}).
					Times(1).
					Return(
						shipping.Quote{},
						fmt.Errorf("mock error"),
					)
			},
			verify: func(resp *shippingv1.GetQuoteResponse, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, codes.Internal.String())
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// setup mock server
			ctrl := gomock.NewController(t)
			t.Cleanup(func() {
				ctrl.Finish()
			})

			shippingService := mock_shipping.NewMockShippingService(ctrl)
			srv := NewServer(shippingService)
			grpcSrv := grpc.NewServer()
			t.Cleanup(func() {
				grpcSrv.Stop()
			})
			shippingv1.RegisterShippingServiceServer(grpcSrv, srv)

			lis := bufconn.Listen(1024 * 1024)
			t.Cleanup(func() {
				lis.Close()
			})

			go func() {
				err := grpcSrv.Serve(lis)
				assert.NoError(t, err)
			}()

			// setup client
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			t.Cleanup(func() {
				cancel()
			})

			dialer := func(context.Context, string) (net.Conn, error) {
				return lis.Dial()
			}

			conn, err := grpc.DialContext(
				ctx,
				"",
				grpc.WithContextDialer(dialer),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			assert.NoError(t, err)
			t.Cleanup(func() {
				conn.Close()
			})

			client := shippingv1.NewShippingServiceClient(conn)

			// setup mock
			test.setupMock(shippingService)

			// make a request
			resp, err := client.GetQuote(ctx, &shippingv1.GetQuoteRequest{
				Address:  test.address,
				Products: test.products,
			})

			// verify
			test.verify(resp, err)
		})
	}
}
