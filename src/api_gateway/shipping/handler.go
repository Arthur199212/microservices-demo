package shipping

import (
	"fmt"
	"net/http"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShippingHandler interface {
	AddRoutes(app *fiber.App)
}

type shippingHandler struct {
	service  ShippingService
	validate *validator.Validate
}

func NewShippingHandler(
	service ShippingService,
	validate *validator.Validate,
) ShippingHandler {
	return &shippingHandler{
		service:  service,
		validate: validate,
	}
}

func (h *shippingHandler) AddRoutes(app *fiber.App) {
	shippingRoute := app.Group("/shipping")

	shippingRoute.Post("/quote", h.getQuote)
}

type Product struct {
	Id       int64 `json:"id" validate:"required,min=1"`
	Quantity int32 `json:"quantity" validate:"required,min=1"`
}

type GetQuoteInput struct {
	Address      models.Address `json:"address" validate:"required,dive"`
	Products     []Product      `json:"products" validate:"required,min=1,max=100,dive"`
	UserCurrency string         `json:"userCurrency" validate:"required,len=3"`
}

func (h *shippingHandler) getQuote(c *fiber.Ctx) error {
	input := GetQuoteInput{}
	if err := c.BodyParser(&input); err != nil {
		log.Error().Err(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}
	if input.UserCurrency == "" {
		input.UserCurrency = defaultCurrency
	}

	if err := h.validate.Struct(input); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("invalid argument: %+v", err),
		})
	}

	money, err := h.service.GetQuote(c.Context(), convertToGetQuoteArgs(input))
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
		msg := "cannot get shipping quote"
		log.Error().Err(err).Msg(msg)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": msg,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"quote":    money.GetAmount(),
		"currency": money.GetCurrencyCode(),
	})
}

func convertToGetQuoteArgs(input GetQuoteInput) GetQuoteArgs {
	state := ""
	if input.Address.State != nil {
		state = *input.Address.State
	}

	products := make([]*modelsv1.Product, len(input.Products))
	for i := range input.Products {
		products[i] = &modelsv1.Product{
			Id:       input.Products[i].Id,
			Quantity: input.Products[i].Quantity,
		}
	}

	return GetQuoteArgs{
		Address: &modelsv1.Address{
			City:          input.Address.City,
			Country:       input.Address.Country,
			State:         state,
			StreetAddress: input.Address.StreetAddress,
			ZipCode:       input.Address.ZipCode,
		},
		Products:     products,
		UserCurrency: input.UserCurrency,
	}
}
