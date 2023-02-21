package products

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductsHandler interface {
	AddRoutes(app *fiber.App)
}

type productsHandler struct {
	service  ProductsService
	validate *validator.Validate
}

func NewProductsHandler(
	service ProductsService,
	validate *validator.Validate,
) ProductsHandler {
	return &productsHandler{
		service:  service,
		validate: validate,
	}
}

func (h *productsHandler) AddRoutes(app *fiber.App) {
	productsRoute := app.Group("/products")

	productsRoute.Get("/", h.listProducts)
	productsRoute.Get("/:id", h.getProductById)
}

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
		userCurrency = defaultCurrency
	}

	args := GetProductByIdArgs{
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

func (h *productsHandler) listProducts(c *fiber.Ctx) error {
	page, pageSize := c.QueryInt("page", 1), c.QueryInt("pageSize", 10)
	userCurrency := c.Query("currency", defaultCurrency)

	args := ListProductsArgs{
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
