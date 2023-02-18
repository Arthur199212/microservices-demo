package gapi

import (
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	mock_checkout "github.com/Arthur199212/microservices-demo/src/checkout/checkout/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestPlaceOrderOK(t *testing.T) {
	address := &modelsv1.Address{
		City:          "city",
		Country:       "country",
		StreetAddress: "street address",
		ZipCode:       "00000",
	}
	items := []*checkoutv1.OrderItem{
		&checkoutv1.OrderItem{
			Product: &modelsv1.Product{
				Id:       1,
				Quantity: 2,
			},
			Cost: &modelsv1.Money{
				Amount:       1.20,
				CurrencyCode: "EUR",
			},
		},
		&checkoutv1.OrderItem{
			Product: &modelsv1.Product{
				Id:       2,
				Quantity: 3,
			},
			Cost: &modelsv1.Money{
				Amount:       1.99,
				CurrencyCode: "EUR",
			},
		},
	}
	transactionId := "transaction uuid"
	cardInfo := &modelsv1.CardInfo{
		Cvv:             "1111",
		ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
		ExpirationYear:  strconv.Itoa(time.Now().Year()),
		Number:          "123456789012",
	}
	shipping := &checkoutv1.Shipping{
		Cost: &modelsv1.Money{
			Amount:       0.98,
			CurrencyCode: "EUR",
		},
		Address: address,
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	// setup server
	checkoutService := mock_checkout.NewMockCheckoutService(ctrl)
	srv := NewServer(checkoutService)
	grpcServer := grpc.NewServer()
	t.Cleanup(func() {
		grpcServer.Stop()
	})

	checkoutv1.RegisterCheckoutServiceServer(grpcServer, srv)

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	go func() {
		err := grpcServer.Serve(lis)
		assert.NoError(t, err)
	}()

	// setup client
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	dialer := func(string, time.Duration) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		conn.Close()
	})

	client := checkoutv1.NewCheckoutServiceClient(conn)

	// setup mock
	checkoutService.EXPECT().PlaceOrder(
		gomock.Any(),
		gomock.Any(),
	).Times(1).Return(
		&checkoutv1.Order{
			TransactionId: transactionId,
			Shipping:      shipping,
			Items:         items,
		},
		nil,
	)

	// make a request
	res, err := client.PlaceOrder(
		context.Background(),
		&checkoutv1.PlaceOrderRequest{
			SessionId:    "01f8477c-9685-46f9-9a57-be6700b7a547",
			UserCurrency: "EUR",
			Address:      address,
			Email:        "person@company.com",
			CardInfo:     cardInfo,
		},
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Order)
	assert.Equal(t, res.Order.TransactionId, transactionId)
	assert.Len(t, res.Order.Items, len(items))
	for i := range res.Order.Items {
		assert.Equal(t, res.Order.Items[i].Product.Id, items[i].Product.Id)
		assert.Equal(t, res.Order.Items[i].Product.Quantity, items[i].Product.Quantity)
		assert.Equal(t, res.Order.Items[i].Cost.Amount, items[i].Cost.Amount)
		assert.Equal(t, res.Order.Items[i].Cost.CurrencyCode, items[i].Cost.CurrencyCode)
	}
	assert.Equal(t, res.Order.Shipping.Address, &modelsv1.Address{
		City:          shipping.Address.City,
		Country:       shipping.Address.Country,
		State:         "",
		StreetAddress: shipping.Address.StreetAddress,
		ZipCode:       shipping.Address.ZipCode,
	})
	assert.Equal(t, res.Order.Shipping.Cost.Amount, shipping.Cost.Amount)
	assert.Equal(t, res.Order.Shipping.Cost.CurrencyCode, shipping.Cost.CurrencyCode)
}

func TestPlaceOrderInvalidState(t *testing.T) {
	address := &modelsv1.Address{
		City:          "city",
		Country:       "country",
		StreetAddress: "street address",
		State:         "s",
		ZipCode:       "00000",
	}
	cardInfo := &modelsv1.CardInfo{
		Cvv:             "1111",
		ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
		ExpirationYear:  strconv.Itoa(time.Now().Year()),
		Number:          "123456789012",
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	// setup server
	checkoutService := mock_checkout.NewMockCheckoutService(ctrl)
	srv := NewServer(checkoutService)
	grpcServer := grpc.NewServer()
	t.Cleanup(func() {
		grpcServer.Stop()
	})

	checkoutv1.RegisterCheckoutServiceServer(grpcServer, srv)

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	go func() {
		err := grpcServer.Serve(lis)
		assert.NoError(t, err)
	}()

	// setup client
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	dialer := func(string, time.Duration) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		conn.Close()
	})

	client := checkoutv1.NewCheckoutServiceClient(conn)

	// make a request
	res, err := client.PlaceOrder(
		context.Background(),
		&checkoutv1.PlaceOrderRequest{
			SessionId:    "01f8477c-9685-46f9-9a57-be6700b7a547",
			UserCurrency: "EUR",
			Address:      address,
			Email:        "person@company.com",
			CardInfo:     cardInfo,
		},
	)

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestPlaceOrderInvalidZipCode(t *testing.T) {
	address := &modelsv1.Address{
		City:          "city",
		Country:       "country",
		StreetAddress: "street address",
		State:         "state",
		ZipCode:       "invalid",
	}
	cardInfo := &modelsv1.CardInfo{
		Cvv:             "1111",
		ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
		ExpirationYear:  strconv.Itoa(time.Now().Year()),
		Number:          "123456789012",
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	// setup server
	checkoutService := mock_checkout.NewMockCheckoutService(ctrl)
	srv := NewServer(checkoutService)
	grpcServer := grpc.NewServer()
	t.Cleanup(func() {
		grpcServer.Stop()
	})

	checkoutv1.RegisterCheckoutServiceServer(grpcServer, srv)

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	go func() {
		err := grpcServer.Serve(lis)
		assert.NoError(t, err)
	}()

	// setup client
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	dialer := func(string, time.Duration) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		conn.Close()
	})

	client := checkoutv1.NewCheckoutServiceClient(conn)

	// make a request
	res, err := client.PlaceOrder(
		context.Background(),
		&checkoutv1.PlaceOrderRequest{
			SessionId:    "01f8477c-9685-46f9-9a57-be6700b7a547",
			UserCurrency: "EUR",
			Address:      address,
			Email:        "person@company.com",
			CardInfo:     cardInfo,
		},
	)

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestPlaceOrderInvalidCurrency(t *testing.T) {
	address := &modelsv1.Address{
		City:          "city",
		Country:       "country",
		StreetAddress: "street address",
		State:         "state",
		ZipCode:       "00000",
	}
	cardInfo := &modelsv1.CardInfo{
		Cvv:             "1111",
		ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
		ExpirationYear:  strconv.Itoa(time.Now().Year()),
		Number:          "123456789012",
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	// setup server
	checkoutService := mock_checkout.NewMockCheckoutService(ctrl)
	srv := NewServer(checkoutService)
	grpcServer := grpc.NewServer()
	t.Cleanup(func() {
		grpcServer.Stop()
	})

	checkoutv1.RegisterCheckoutServiceServer(grpcServer, srv)

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	go func() {
		err := grpcServer.Serve(lis)
		assert.NoError(t, err)
	}()

	// setup client
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	dialer := func(string, time.Duration) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		conn.Close()
	})

	client := checkoutv1.NewCheckoutServiceClient(conn)

	// make a request
	res, err := client.PlaceOrder(
		context.Background(),
		&checkoutv1.PlaceOrderRequest{
			SessionId:    "01f8477c-9685-46f9-9a57-be6700b7a547",
			UserCurrency: "invalid",
			Address:      address,
			Email:        "person@company.com",
			CardInfo:     cardInfo,
		},
	)

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestPlaceOrderInvalidEmail(t *testing.T) {
	address := &modelsv1.Address{
		City:          "city",
		Country:       "country",
		StreetAddress: "street address",
		State:         "state",
		ZipCode:       "00000",
	}
	cardInfo := &modelsv1.CardInfo{
		Cvv:             "1111",
		ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
		ExpirationYear:  strconv.Itoa(time.Now().Year()),
		Number:          "123456789012",
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	// setup server
	checkoutService := mock_checkout.NewMockCheckoutService(ctrl)
	srv := NewServer(checkoutService)
	grpcServer := grpc.NewServer()
	t.Cleanup(func() {
		grpcServer.Stop()
	})

	checkoutv1.RegisterCheckoutServiceServer(grpcServer, srv)

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	go func() {
		err := grpcServer.Serve(lis)
		assert.NoError(t, err)
	}()

	// setup client
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	dialer := func(string, time.Duration) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		conn.Close()
	})

	client := checkoutv1.NewCheckoutServiceClient(conn)

	// make a request
	res, err := client.PlaceOrder(
		context.Background(),
		&checkoutv1.PlaceOrderRequest{
			SessionId:    "01f8477c-9685-46f9-9a57-be6700b7a547",
			UserCurrency: "EUR",
			Address:      address,
			Email:        "personcompany.com",
			CardInfo:     cardInfo,
		},
	)

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestPlaceOrderError(t *testing.T) {
	address := &modelsv1.Address{
		City:          "city",
		Country:       "country",
		StreetAddress: "street address",
		ZipCode:       "00000",
	}
	cardInfo := &modelsv1.CardInfo{
		Cvv:             "1111",
		ExpirationMonth: strconv.Itoa(int(time.Now().Month())),
		ExpirationYear:  strconv.Itoa(time.Now().Year()),
		Number:          "123456789012",
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	// setup server
	checkoutService := mock_checkout.NewMockCheckoutService(ctrl)
	srv := NewServer(checkoutService)
	grpcServer := grpc.NewServer()
	t.Cleanup(func() {
		grpcServer.Stop()
	})

	checkoutv1.RegisterCheckoutServiceServer(grpcServer, srv)

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	go func() {
		err := grpcServer.Serve(lis)
		assert.NoError(t, err)
	}()

	// setup client
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	dialer := func(string, time.Duration) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		conn.Close()
	})

	client := checkoutv1.NewCheckoutServiceClient(conn)

	// setup mock
	checkoutService.EXPECT().PlaceOrder(
		gomock.Any(),
		gomock.Any(),
	).Times(1).Return(nil, fmt.Errorf("mock error"))

	// make a request
	res, err := client.PlaceOrder(
		context.Background(),
		&checkoutv1.PlaceOrderRequest{
			SessionId:    "01f8477c-9685-46f9-9a57-be6700b7a547",
			UserCurrency: "EUR",
			Address:      address,
			Email:        "person@company.com",
			CardInfo:     cardInfo,
		},
	)

	assert.Error(t, err)
	assert.Empty(t, res)
	assert.ErrorContains(t, err, "mock error")
}
