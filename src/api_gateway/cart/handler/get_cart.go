package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *cartHandler) getCart(c *fiber.Ctx) error {
	sessionId := c.Params("sessionId")
	if err := h.validate.Var(sessionId, "required,uuid"); err != nil {
		msg := "invalid argument \"sessionId\""
		log.Error().Err(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": msg,
		})
	}

	products, err := h.service.GetCart(c.Context(), sessionId)
	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			switch errStatus.Code() {
			case codes.NotFound:
				msg := fmt.Sprintf("cart with sessionId=%s not found", sessionId)
				log.Error().Err(err).Msg(msg)
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"error": msg,
				})
			default:
			}
		}
		msg := fmt.Sprintf("cannot retrieve cart with sessionId=%s", sessionId)
		log.Error().Err(err).Msgf(msg)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": msg,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"products": products,
	})
}
