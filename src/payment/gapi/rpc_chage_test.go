package gapi

import (
	"context"
	"net"
	"strconv"
	"testing"
	"time"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
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
		money    *modelsv1.Money
		cardInfo *modelsv1.CardInfo
		verifyFn func(t *testing.T, resp *paymentv1.ChargeResponse, err error)
	}{
		{
			name: "OK",
			money: &modelsv1.Money{
				CurrencyCode: "EUR",
				Amount:       100,
			},
			cardInfo: &modelsv1.CardInfo{
				Number:          "4242424242424242",
				Cvv:             "100",
				ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
				ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
			},
			verifyFn: func(t *testing.T, resp *paymentv1.ChargeResponse, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp.TransactionId)
			},
		},
		{
			name: "invalid amount",
			money: &modelsv1.Money{
				CurrencyCode: "EUR",
				Amount:       -100,
			},
			cardInfo: &modelsv1.CardInfo{
				Number:          "4242424242424242",
				Cvv:             "100",
				ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
				ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
			},
			verifyFn: func(t *testing.T, resp *paymentv1.ChargeResponse, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "amount")
			},
		},
		{
			name: "invalid card number",
			money: &modelsv1.Money{
				CurrencyCode: "EUR",
				Amount:       100,
			},
			cardInfo: &modelsv1.CardInfo{
				Number:          "000000000000000",
				Cvv:             "100",
				ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
				ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
			},
			verifyFn: func(t *testing.T, resp *paymentv1.ChargeResponse, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "invalid method",
			money: &modelsv1.Money{
				CurrencyCode: "EUR",
				Amount:       100,
			},
			cardInfo: &modelsv1.CardInfo{
				Number:          "378282246310005", // American Express (https://stripe.com/docs/testing)
				Cvv:             "100",
				ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
				ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
			},
			verifyFn: func(t *testing.T, resp *paymentv1.ChargeResponse, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, codes.InvalidArgument.String())
			},
		},
		{
			name: "lack of cvv",
			money: &modelsv1.Money{
				CurrencyCode: "EUR",
				Amount:       100,
			},
			cardInfo: &modelsv1.CardInfo{
				Number:          "4242424242424242",
				ExpirationYear:  strconv.Itoa(time.Now().Year() + 1),
				ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
			},
			verifyFn: func(t *testing.T, resp *paymentv1.ChargeResponse, err error) {
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

			config := utils.Config{
				AllowTestCardNumbers: true,
				Port:                 "5000",
			}
			svr := NewServer(config)
			paymentv1.RegisterPaymentServiceServer(grpcServer, svr)

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

			client := paymentv1.NewPaymentServiceClient(conn)

			// make a request
			resp, err := client.Charge(ctx, &paymentv1.ChargeRequest{
				Money:    test.money,
				CardInfo: test.cardInfo,
			})

			// verify result
			test.verifyFn(t, resp, err)
		})
	}
}
