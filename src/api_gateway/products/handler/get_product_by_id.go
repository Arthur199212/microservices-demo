package handler

import (
	"fmt"
	"net/http"

	"github.com/Arthur199212/microservices-demo/src/api_gateway/products/service"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *productsHandler) getProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		msg := fmt.Sprintf("invalid parameter id=%s", c.Params("id"))
		log.Error().Msg(msg)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("%s: %+v", msg, err),
		})
	}

	userCurrency := c.Query("currency")
	if userCurrency == "" {
		userCurrency = h.config.DefaultCurrency
	}

	args := service.GetProductByIdArgs{
		Id:           int64(id),
		UserCurrency: userCurrency,
	}
	if err := h.validate.Struct(args); err != nil {
		msg := "invalid argument"
		log.Error().Err(err).Msg(msg)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("%s: %+v", msg, err),
		})
	}

	product, err := h.service.GetProductById(c.Context(), args)
	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			switch errStatus.Code() {
			case codes.NotFound:
				msg := fmt.Sprintf("product with id=%d not found", id)
				log.Error().Err(err).Msg(msg)
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"error": msg,
				})
			case codes.InvalidArgument:
				log.Error().Err(err).Msg("invalid argument")
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"error": fmt.Sprintf("invalid argument: %+v", err),
				})
			default:
			}
		}
		msg := fmt.Sprintf("cannot get product with id=%s", c.Params("id"))
		log.Error().Err(err).Msg(msg)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": msg,
		})
	}

	return c.Status(http.StatusOK).JSON(product)
}
