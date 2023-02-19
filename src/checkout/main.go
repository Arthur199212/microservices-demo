package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"github.com/Arthur199212/microservices-demo/src/checkout/checkout"
	"github.com/Arthur199212/microservices-demo/src/checkout/gapi"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"github.com/joho/godotenv"
)

const (
	defaultPort = "5005"
)

func main() {
	err := godotenv.Load()
  if err != nil {
		log.Fatal().Err(err).Msg("cannot load .env file")
  }

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cartServiceAddr := os.Getenv("CART_SERVICE_ADDR")
	currencyServiceAddr := os.Getenv("CURRENCY_SERVICE_ADDR")
	paymentServiceAddr := os.Getenv("PAYMENT_SERVICE_ADDR")
	productsServiceAddr := os.Getenv("PRODUCTS_SERVICE_ADDR")
	shippingServiceAddr := os.Getenv("SHIPPING_SERVICE_ADDR")

	cartConn := dialGrpcClient(cartServiceAddr)
	defer cartConn.Close()
	cartClient := cartv1.NewCartServiceClient(cartConn)

	currencyConn := dialGrpcClient(currencyServiceAddr)
	defer currencyConn.Close()
	currencyClient := currencyv1.NewCurrencyServiceClient(currencyConn)

	paymentConn := dialGrpcClient(paymentServiceAddr)
	defer paymentConn.Close()
	paymentClient := paymentv1.NewPaymentServiceClient(paymentConn)

	productsConn := dialGrpcClient(productsServiceAddr)
	defer productsConn.Close()
	productsClient := productsv1.NewProductsServiceClient(productsConn)

	shippingConn := dialGrpcClient(shippingServiceAddr)
	defer shippingConn.Close()
	shippingClient := shippingv1.NewShippingServiceClient(shippingConn)

	checkoutService := checkout.NewCheckoutService(
		cartClient,
		currencyClient,
		paymentClient,
		productsClient,
		shippingClient,
	)
	srv := gapi.NewServer(checkoutService)
	grpcServer := grpc.NewServer()
	checkoutv1.RegisterCheckoutServiceServer(grpcServer, srv)

	// to provide self-documentation
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("starting gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func dialGrpcClient(addr string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to gRPC server")
	}
	return conn
}
