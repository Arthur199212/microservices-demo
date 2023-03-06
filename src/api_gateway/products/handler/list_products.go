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

const (
	defaultPage     = 1
	defaultPageSize = 10
)

func (h *productsHandler) listProducts(c *fiber.Ctx) error {
	page, pageSize := c.QueryInt("page", defaultPage), c.QueryInt("pageSize", defaultPageSize)
	userCurrency := c.Query("currency", h.config.DefaultCurrency)

	args := service.ListProductsArgs{
		Page:         int32(page),
		PageSize:     int32(pageSize),
		UserCurrency: userCurrency,
	}
	if err := h.validate.Struct(args); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("invalid argument: %+v", err),
		})
	}

	products, err := h.service.ListProducts(c.Context(), args)
	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			switch errStatus.Code() {
			case codes.InvalidArgument:
				log.Error().Err(err).Msg("invalid argument")
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"error": fmt.Sprintf("invalid argument: %+v", err),
				})
			default:
			}
		}
		msg := "cannot retrieve a list of products"
		log.Error().Err(err).Msgf(msg)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": msg,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"products": products,
	})
}
