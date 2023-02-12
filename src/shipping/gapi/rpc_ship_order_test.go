package gapi

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/Arthur199212/microservices-demo/src/shipping/pb"
	"github.com/Arthur199212/microservices-demo/src/shipping/shipping"
	mock_shipping "github.com/Arthur199212/microservices-demo/src/shipping/shipping/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestShipOrder(t *testing.T) {
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
	transactionId := "mock-transaction-id"

	testCases := []struct {
		name      string
		address   *pb.Address
		products  []*pb.Product
		setupMock func(*mock_shipping.MockShippingService)
		verify    func(*pb.ShipOrderResponse, error)
	}{
		{
			name: "OK",
			address: &pb.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products: []*pb.Product{
				&pb.Product{
					Id:       productOne.ID,
					Quantity: productOne.Quantity,
				},
				&pb.Product{
					Id:       productTwo.ID,
					Quantity: productTwo.Quantity,
				},
			},
			setupMock: func(s *mock_shipping.MockShippingService) {
				s.
					EXPECT().
					ShipOrder(mockAddress, []shipping.Product{productOne, productTwo}).
					Times(1).
					Return(
						transactionId,
						nil,
					)
			},
			verify: func(resp *pb.ShipOrderResponse, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp.TrackingId)
				assert.Equal(t, transactionId, resp.TrackingId)
			},
		},
		{
			name: "no products to ship",
			address: &pb.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products:  []*pb.Product{},
			setupMock: func(mss *mock_shipping.MockShippingService) {},
			verify: func(resp *pb.ShipOrderResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "too many products to ship",
			address: &pb.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products:  make([]*pb.Product, maxProductsToShip+1),
			setupMock: func(mss *mock_shipping.MockShippingService) {},
			verify: func(resp *pb.ShipOrderResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "invalid product ID",
			address: &pb.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products: []*pb.Product{
				&pb.Product{
					Id:       0,
					Quantity: productOne.Quantity,
				},
			},
			setupMock: func(s *mock_shipping.MockShippingService) {},
			verify: func(resp *pb.ShipOrderResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "invalid zipCode in the address",
			address: &pb.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       "10",
			},
			products: []*pb.Product{
				&pb.Product{
					Id:       productOne.ID,
					Quantity: productOne.Quantity,
				},
			},
			setupMock: func(s *mock_shipping.MockShippingService) {},
			verify: func(resp *pb.ShipOrderResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "internal service error",
			address: &pb.Address{
				StreetAddress: mockAddress.StreetAddress,
				City:          mockAddress.City,
				State:         *mockAddress.State,
				Country:       mockAddress.Country,
				ZipCode:       mockAddress.ZipCode,
			},
			products: []*pb.Product{
				&pb.Product{
					Id:       productOne.ID,
					Quantity: productOne.Quantity,
				},
				&pb.Product{
					Id:       productTwo.ID,
					Quantity: productTwo.Quantity,
				},
			},
			setupMock: func(s *mock_shipping.MockShippingService) {
				s.
					EXPECT().
					ShipOrder(mockAddress, []shipping.Product{productOne, productTwo}).
					Times(1).
					Return(
						"",
						fmt.Errorf("mock error"),
					)
			},
			verify: func(resp *pb.ShipOrderResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
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
			pb.RegisterShippingServer(grpcSrv, srv)

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
			t.Cleanup(func() {
				conn.Close()
			})
			assert.NoError(t, err)

			client := pb.NewShippingClient(conn)

			// setup mock
			test.setupMock(shippingService)

			// make a request
			resp, err := client.ShipOrder(ctx, &pb.ShipOrderRequest{
				Address:  test.address,
				Products: test.products,
			})

			// verify
			test.verify(resp, err)
		})
	}
}
