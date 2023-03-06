package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/products/service"
	mock_service "github.com/Arthur199212/microservices-demo/src/api_gateway/products/service/mocks"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetProductByIdOK(t *testing.T) {
	product := &productsv1.Product{
		Id:          1,
		Name:        "tea",
		Description: "green",
		Picture:     "/tea",
		Price:       1.49,
		Currency:    "USD",
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	s := mock_service.NewMockProductsService(ctrl)
	validate := validator.New()
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.GetProductByIdArgs{
		Id:           product.Id,
		UserCurrency: "USD",
	}
	s.EXPECT().GetProductById(gomock.Any(), args).Times(1).
		Return(product, nil)

	u, err := url.Parse(fmt.Sprintf("/products/%d", product.Id))
	require.NoError(t, err)

	q := url.Values{}
	q.Add("currency", "USD")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotProduct *productsv1.Product
	err = json.Unmarshal(data, &gotProduct)
	require.NoError(t, err)
	require.Equal(t, gotProduct, product)
}

func TestGetProductByIdInvalidParamIdString(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	validate := validator.New()
	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	s := mock_service.NewMockProductsService(ctrl)
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	u := fmt.Sprintf("/products/%s", "abc")
	req, err := http.NewRequest(http.MethodGet, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetProductByIdInvalidParamId(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	validate := validator.New()
	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	s := mock_service.NewMockProductsService(ctrl)
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	u := fmt.Sprintf("/products/%d", -1)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetProductByIdNoCurrencyQueryParam(t *testing.T) {
	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	product := &productsv1.Product{
		Id:          2,
		Name:        "tea",
		Description: "black",
		Picture:     "/tea",
		Price:       2.56,
		Currency:    config.DefaultCurrency,
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	validate := validator.New()
	s := mock_service.NewMockProductsService(ctrl)
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.GetProductByIdArgs{
		Id:           product.Id,
		UserCurrency: config.DefaultCurrency,
	}
	s.EXPECT().GetProductById(gomock.Any(), args).Times(1).
		Return(product, nil)

	u := fmt.Sprintf("/products/%d", product.Id)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotProduct *productsv1.Product
	err = json.Unmarshal(data, &gotProduct)
	require.NoError(t, err)
	require.Equal(t, gotProduct, product)
}

func TestGetProductByIdNotFound(t *testing.T) {
	var productId int64 = 1

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	s := mock_service.NewMockProductsService(ctrl)
	validate := validator.New()
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.GetProductByIdArgs{
		Id:           productId,
		UserCurrency: config.DefaultCurrency,
	}
	s.EXPECT().GetProductById(gomock.Any(), args).Times(1).
		Return(nil, status.Errorf(codes.NotFound, "not found"))

	u := fmt.Sprintf("/products/%d", productId)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Error string
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)
	require.Contains(t, gotResp.Error, "not found")
}

func TestGetProductByIdServiceReturnsInvalidArgumentErr(t *testing.T) {
	var productId int64 = 1
	invalidCurrencyCode := "ERR"

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	s := mock_service.NewMockProductsService(ctrl)
	validate := validator.New()
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.GetProductByIdArgs{
		Id:           productId,
		UserCurrency: invalidCurrencyCode,
	}
	s.EXPECT().GetProductById(gomock.Any(), args).Times(1).
		Return(nil, status.Errorf(codes.InvalidArgument, "not found"))

	u, err := url.Parse(fmt.Sprintf("/products/%d", productId))
	require.NoError(t, err)

	q := url.Values{}
	q.Add("currency", invalidCurrencyCode)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Error string
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)
	require.Contains(t, gotResp.Error, "invalid argument")
}

func TestGetProductByIdInternalErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockProductsService(ctrl)
	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	validate := validator.New()
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	var productId int64 = 1
	args := service.GetProductByIdArgs{
		Id:           productId,
		UserCurrency: config.DefaultCurrency,
	}
	s.EXPECT().GetProductById(gomock.Any(), args).Times(1).
		Return(nil, status.Errorf(codes.Internal, "mock internal error"))

	u := fmt.Sprintf("/products/%d", productId)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Error string
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)
	require.Contains(t, gotResp.Error, "cannot get product")
}
