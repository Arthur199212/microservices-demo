package gapi

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/Arthur199212/microservices-demo/src/payment/pb"
	"github.com/Arthur199212/microservices-demo/src/payment/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestChargeHandler(t *testing.T) {
	testCases := []struct {
		name     string
		money    *pb.Money
		cardInfo *pb.CardInfo
		verifyFn func(t *testing.T, resp *pb.ChargeResponse, err error)
	}{
		{
			name: "OK",
			money: &pb.Money{
				CurrencyCode: "EUR",
				Amount:       100,
			},
			cardInfo: &pb.CardInfo{
				Number:          "4242424242424242",
				Cvv:             100,
				ExpirationYear:  int32(time.Now().Year() + 1),
				ExpirationMonth: int32(time.Now().Month()),
			},
			verifyFn: func(t *testing.T, resp *pb.ChargeResponse, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp.TransactionId)
			},
		},
		{
			name: "invalid amount",
			money: &pb.Money{
				CurrencyCode: "EUR",
				Amount:       -100,
			},
			cardInfo: &pb.CardInfo{
				Number:          "4242424242424242",
				Cvv:             100,
				ExpirationYear:  int32(time.Now().Year() + 1),
				ExpirationMonth: int32(time.Now().Month()),
			},
			verifyFn: func(t *testing.T, resp *pb.ChargeResponse, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "amount")
			},
		},
		{
			name: "invalid card number",
			money: &pb.Money{
				CurrencyCode: "EUR",
				Amount:       100,
			},
			cardInfo: &pb.CardInfo{
				Number:          "000000000000000",
				Cvv:             100,
				ExpirationYear:  int32(time.Now().Year() + 1),
				ExpirationMonth: int32(time.Now().Month()),
			},
			verifyFn: func(t *testing.T, resp *pb.ChargeResponse, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "invalid method",
			money: &pb.Money{
				CurrencyCode: "EUR",
				Amount:       100,
			},
			cardInfo: &pb.CardInfo{
				Number:          "378282246310005", // American Express (https://stripe.com/docs/testing)
				Cvv:             100,
				ExpirationYear:  int32(time.Now().Year() + 1),
				ExpirationMonth: int32(time.Now().Month()),
			},
			verifyFn: func(t *testing.T, resp *pb.ChargeResponse, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "lack of cvv",
			money: &pb.Money{
				CurrencyCode: "EUR",
				Amount:       100,
			},
			cardInfo: &pb.CardInfo{
				Number:          "4242424242424242",
				ExpirationYear:  int32(time.Now().Year() + 1),
				ExpirationMonth: int32(time.Now().Month()),
			},
			verifyFn: func(t *testing.T, resp *pb.ChargeResponse, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// setup test server
			grpcServer := grpc.NewServer()
			t.Cleanup(func() {
				grpcServer.Stop()
			})

			config := &utils.Config{
				AllowTestCardNumbers: true,
				Port:                 "5000",
			}
			svr := NewServer(config)
			pb.RegisterPaymentServer(grpcServer, svr)

			lis := bufconn.Listen(1024 * 1024)
			t.Cleanup(func() {
				lis.Close()
			})

			go func() {
				err := grpcServer.Serve(lis)
				assert.NoError(t, err)
			}()

			// setup test client
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

			client := pb.NewPaymentClient(conn)

			// make a request
			resp, err := client.Charge(ctx, &pb.ChargeRequest{
				Money:    test.money,
				CardInfo: test.cardInfo,
			})

			// verify result
			test.verifyFn(t, resp, err)
		})
	}
}