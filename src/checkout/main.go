package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	cart "github.com/Arthur199212/microservices-demo/src/cart/pb"
	"github.com/Arthur199212/microservices-demo/src/checkout/checkout"
	"github.com/Arthur199212/microservices-demo/src/checkout/gapi"
	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	currency "github.com/Arthur199212/microservices-demo/src/currency/pb"
	payment "github.com/Arthur199212/microservices-demo/src/payment/pb"
	products "github.com/Arthur199212/microservices-demo/src/products/pb"
	shipping "github.com/Arthur199212/microservices-demo/src/shipping/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = "5005"
)

func main() {
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
	cartClient := cart.NewCartClient(cartConn)

	currencyConn := dialGrpcClient(currencyServiceAddr)
	defer currencyConn.Close()
	currencyClient := currency.NewCurrencyClient(currencyConn)

	paymentConn := dialGrpcClient(paymentServiceAddr)
	defer paymentConn.Close()
	paymentClient := payment.NewPaymentClient(paymentConn)

	productsConn := dialGrpcClient(productsServiceAddr)
	defer productsConn.Close()
	productsClient := products.NewProductsClient(productsConn)

	shippingConn := dialGrpcClient(shippingServiceAddr)
	defer shippingConn.Close()
	shippingClient := shipping.NewShippingClient(shippingConn)

	checkoutService := checkout.NewCheckoutService(
		cartClient,
		currencyClient,
		paymentClient,
		productsClient,
		shippingClient,
	)
	srv := gapi.NewServer(checkoutService)
	grpcServer := grpc.NewServer()
	pb.RegisterCheckoutServer(grpcServer, srv)

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
