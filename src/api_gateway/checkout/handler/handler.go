package handler

import (
	"fmt"
	"net/http"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/checkout/service"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type CheckoutHandler interface {
	AddRoutes(app *fiber.App)
}

type checkoutHandler struct {
	service  service.CheckoutService
	validate *validator.Validate
}

func NewCheckoutHandler(
	service service.CheckoutService,
	validate *validator.Validate,
) CheckoutHandler {
	return &checkoutHandler{
		service:  service,
		validate: validate,
	}
}

func (h *checkoutHandler) AddRoutes(app *fiber.App) {
	checkout := app.Group("/checkout")

	checkout.Post("/place-order", h.placeOrder)
}

type Shipping struct {
	Cost    *modelsv1.Money `json:"cost"`
	Address models.Address  `json:"address"`
}

type Order struct {
	TransactionId string                  `json:"transactionId"`
	Shipping      Shipping                `json:"shipping"`
	Items         []*checkoutv1.OrderItem `json:"items"`
}

func (h *checkoutHandler) placeOrder(c *fiber.Ctx) error {
	args := service.CheckoutServiceArgs{}
	if err := c.BodyParser(&args); err != nil {
		log.Error().Err(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}

	if err := h.validate.Struct(args); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("invalid argument: %+v", err),
		})
	}

	order, err := h.service.PlaceOrder(c.Context(), args)
	if err != nil {
		msg := "cannot place order"
		log.Error().Err(err).Msg(msg)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": msg,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"order": convertToResponseOrder(order),
	})
}

func convertToResponseOrder(o *checkoutv1.Order) Order {
	order := Order{
		TransactionId: o.TransactionId,
		Items:         o.Items,
		Shipping: Shipping{
			Cost: o.Shipping.Cost,
			Address: models.Address{
				City:          o.Shipping.Address.City,
				Country:       o.Shipping.Address.Country,
				State:         nil,
				StreetAddress: o.Shipping.Address.StreetAddress,
				ZipCode:       o.Shipping.Address.ZipCode,
			},
		},
	}
	if o.Shipping.Address.State != "" {
		order.Shipping.Address.State = &o.Shipping.Address.State
	}
	return order
}
