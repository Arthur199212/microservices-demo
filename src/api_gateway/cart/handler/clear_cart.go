package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *cartHandler) clearCart(c *fiber.Ctx) error {
	sessionId := c.Params("sessionId")
	if err := h.validate.Var(sessionId, "required,uuid4"); err != nil {
		msg := "invalid argument \"sessionId\""
		log.Error().Err(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": msg,
		})
	}

	err := h.service.ClearCart(c.Context(), sessionId)
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
		msg := fmt.Sprintf("cannot clear cart with sessionId=%s", sessionId)
		log.Error().Err(err).Msg(msg)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": msg,
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
