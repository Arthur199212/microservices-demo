package handler

import (
	"github.com/Arthur199212/microservices-demo/src/api_gateway/products/service"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductsHandler interface {
	AddRoutes(app *fiber.App)
}

type productsHandler struct {
	config   utils.Config
	service  service.ProductsService
	validate *validator.Validate
}

func NewProductsHandler(
	config utils.Config,
	service service.ProductsService,
	validate *validator.Validate,
) ProductsHandler {
	return &productsHandler{
		config:   config,
		service:  service,
		validate: validate,
	}
}

func (h *productsHandler) AddRoutes(app *fiber.App) {
	productsRoute := app.Group("/products")

	productsRoute.Get("/", h.listProducts)
	productsRoute.Get("/:id", h.getProductById)
}
