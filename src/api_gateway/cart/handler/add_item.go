package handler

import (
	"fmt"
	"net/http"

	"github.com/Arthur199212/microservices-demo/src/api_gateway/cart/service"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (h *cartHandler) addItem(c *fiber.Ctx) error {
	args := service.AddItemsArgs{}
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

	sessionId, err := h.service.AddItem(c.Context(), args)
	if err != nil {
		msg := fmt.Sprintf("cannot add item to cart with sessionId=%s", sessionId)
		log.Error().Err(err).Msg(msg)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": msg,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"sessionId": sessionId,
	})
}
