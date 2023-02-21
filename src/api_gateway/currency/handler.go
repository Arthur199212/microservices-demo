package currency

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type CurrencyHandler interface {
	AddRoutes(app *fiber.App)
}

type currencyHandler struct {
	service  CurrencyService
	validate *validator.Validate
}

func NewCurrencyHandler(
	service CurrencyService,
	validate *validator.Validate,
) CurrencyHandler {
	return &currencyHandler{
		service:  service,
		validate: validate,
	}
}

func (h *currencyHandler) AddRoutes(app *fiber.App) {
	currencies := app.Group("/currencies")

	currencies.Get("/", h.getSupportedCurrencies)
}

func (h *currencyHandler) getSupportedCurrencies(c *fiber.Ctx) error {
	currencies, err := h.service.GetSupportedCurrencies(c.Context())
	if err != nil {
		msg := "cannot get list of supported currencies"
		log.Error().Err(err).Msg(msg)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": msg,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"currencies": currencies,
	})
}
