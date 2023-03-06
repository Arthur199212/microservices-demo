package main

import (
	"context"
	"fmt"
	"net"
	"time"

	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"github.com/Arthur199212/microservices-demo/src/checkout/checkout"
	"github.com/Arthur199212/microservices-demo/src/checkout/gapi"
	"github.com/Arthur199212/microservices-demo/src/checkout/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := utils.LoadConfig("configs")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	cartConn := dialGrpcClient(config.CartServiceAddr)
	defer cartConn.Close()
	cartClient := cartv1.NewCartServiceClient(cartConn)

	currencyConn := dialGrpcClient(config.CurrencyServiceAddr)
	defer currencyConn.Close()
	currencyClient := currencyv1.NewCurrencyServiceClient(currencyConn)

	paymentConn := dialGrpcClient(config.PaymentServiceAddr)
	defer paymentConn.Close()
	paymentClient := paymentv1.NewPaymentServiceClient(paymentConn)

	productsConn := dialGrpcClient(config.ProductsServiceAddr)
	defer productsConn.Close()
	productsClient := productsv1.NewProductsServiceClient(productsConn)

	shippingConn := dialGrpcClient(config.ShippingServiceAddr)
	defer shippingConn.Close()
	shippingClient := shippingv1.NewShippingServiceClient(shippingConn)

	checkoutService := checkout.NewCheckoutService(
		config,
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

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))
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
