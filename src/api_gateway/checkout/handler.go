package checkout

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type CheckoutHandler interface {
	AddRoutes(app *fiber.App)
}

type checkoutHandler struct {
	service  CheckoutService
	validate *validator.Validate
}

func NewCheckoutHandler(
	service CheckoutService,
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

func (h *checkoutHandler) placeOrder(c *fiber.Ctx) error {
	args := CheckoutServiceArgs{}
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
		"order": order,
	})
}
