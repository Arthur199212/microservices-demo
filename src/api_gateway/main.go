package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	cartHandler "github.com/Arthur199212/microservices-demo/src/api_gateway/cart/handler"
	cartService "github.com/Arthur199212/microservices-demo/src/api_gateway/cart/service"
	checkoutHandler "github.com/Arthur199212/microservices-demo/src/api_gateway/checkout/handler"
	checkoutService "github.com/Arthur199212/microservices-demo/src/api_gateway/checkout/service"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/currency"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/health"
	productsHandler "github.com/Arthur199212/microservices-demo/src/api_gateway/products/handler"
	productsService "github.com/Arthur199212/microservices-demo/src/api_gateway/products/service"
	shippingHandler "github.com/Arthur199212/microservices-demo/src/api_gateway/shipping/handler"
	shippingService "github.com/Arthur199212/microservices-demo/src/api_gateway/shipping/service"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config, err := utils.LoadConfig("configs")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	cartConn := mustDialGrpcClient(config.CartServiceAddr)
	defer cartConn.Close()
	cartClient := cartv1.NewCartServiceClient(cartConn)

	checkoutConn := mustDialGrpcClient(config.CheckoutServiceAddr)
	defer checkoutConn.Close()
	checkoutClient := checkoutv1.NewCheckoutServiceClient(checkoutConn)

	currencyConn := mustDialGrpcClient(config.CurrencyServiceAddr)
	defer currencyConn.Close()
	currencyClient := currencyv1.NewCurrencyServiceClient(currencyConn)

	productsConn := mustDialGrpcClient(config.ProductsServiceAddr)
	defer productsConn.Close()
	productsClient := productsv1.NewProductsServiceClient(productsConn)

	shippingConn := mustDialGrpcClient(config.ShippingServiceAddr)
	defer shippingConn.Close()
	shippingClient := shippingv1.NewShippingServiceClient(shippingConn)

	validate := validator.New()

	app := fiber.New()
	app.Use(cors.New())

	currencyService := currency.NewCurrencyService(currencyClient)

	cartH := cartHandler.NewCartHandler(
		cartService.NewCartService(cartClient),
		validate,
	)
	cartH.AddRoutes(app)

	checkoutH := checkoutHandler.NewCheckoutHandler(
		checkoutService.NewCheckoutService(checkoutClient),
		validate,
	)
	checkoutH.AddRoutes(app)

	currencyH := currency.NewCurrencyHandler(
		currencyService,
		validate,
	)
	currencyH.AddRoutes(app)

	productsH := productsHandler.NewProductsHandler(
		config,
		productsService.NewProductsService(config, productsClient, currencyService),
		validate,
	)
	productsH.AddRoutes(app)

	shippingH := shippingHandler.NewShippingHandler(
		shippingService.NewShippingService(
			shippingClient,
			currencyService,
		),
		validate,
	)
	shippingH.AddRoutes(app)

	healthH := health.NewHealthHandler()
	healthH.AddRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Api Gateway")
	})

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", config.Port)); err != nil {
			log.Fatal().Err(err).
				Msg("error while trying to listen to port for the incoming requests")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("server is shutting down ...")

	if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
		log.Fatal().Err(err).Msg("error while shutting down")
	}
}

func mustDialGrpcClient(addr string) *grpc.ClientConn {
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
