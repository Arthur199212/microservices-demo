package handler

import (
	"github.com/Arthur199212/microservices-demo/src/api_gateway/cart/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CartHandler interface {
	AddRoutes(app *fiber.App)
}

type cartHandler struct {
	service  service.CartService
	validate *validator.Validate
}

func NewCartHandler(
	service service.CartService,
	validate *validator.Validate,
) CartHandler {
	return &cartHandler{
		service:  service,
		validate: validate,
	}
}

func (h *cartHandler) AddRoutes(app *fiber.App) {
	cart := app.Group("/cart")

	cart.Get("/:sessionId", h.getCart)
	cart.Post("/", h.addItem)
	cart.Delete("/:sessionId", h.clearCart)
}
