package main

import (
	"context"
	"fmt"
	"os"
	"time"

	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	cartHandler "github.com/Arthur199212/microservices-demo/src/api_gateway/cart/handler"
	cartService "github.com/Arthur199212/microservices-demo/src/api_gateway/cart/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultPort = "4000"
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
	// checkoutServiceAddr := os.Getenv("CHECKOUT_SERVICE_ADDR")
	// currencyServiceAddr := os.Getenv("CURRENCY_SERVICE_ADDR")
	// productsServiceAddr := os.Getenv("PRODUCTS_SERVICE_ADDR")
	// shippingServiceAddr := os.Getenv("SHIPPING_SERVICE_ADDR")

	cartConn := mustDialGrpcClient(cartServiceAddr)
	defer cartConn.Close()
	cartClient := cartv1.NewCartServiceClient(cartConn)

	// checkoutClient := dialGrpcClient(checkoutServiceAddr)
	// currencyClient := dialGrpcClient(currencyServiceAddr)
	// productsClient := dialGrpcClient(productsServiceAddr)
	// shippingClient := dialGrpcClient(shippingServiceAddr)

	validate := validator.New()

	app := fiber.New()
	app.Use(cors.New())

	cartH := cartHandler.NewCartHandler(
		cartService.NewCartService(cartClient),
		validate,
	)
	cartH.AddRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Api Gateway")
	})

	app.Listen(fmt.Sprintf(":%s", port))
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
