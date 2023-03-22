package health

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler interface {
	AddRoutes(app *fiber.App)
}

type healthHandler struct{}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

func (h *healthHandler) AddRoutes(app *fiber.App) {
	app.Get("/health", h.healthCheck)
}

func (h *healthHandler) healthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "OK",
	})
}
