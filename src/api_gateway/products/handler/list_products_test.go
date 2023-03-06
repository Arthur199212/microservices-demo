package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

func TestListProductsOK(t *testing.T) {
	page := 1
	pageSize := 10
	userCurrency := "EUR"
	products := []*productsv1.Product{
		&productsv1.Product{
			Id:          1,
			Name:        "tea",
			Description: "green",
			Picture:     "/tea",
			Price:       1.20,
			Currency:    "EUR",
		},
		&productsv1.Product{
			Id:          2,
			Name:        "coffee",
			Description: "arabica",
			Picture:     "/coffee",
			Price:       4.73,
			Currency:    "EUR",
		},
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	validate := validator.New()
	s := mock_service.NewMockProductsService(ctrl)
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.ListProductsArgs{
		Page:         int32(page),
		PageSize:     int32(pageSize),
		UserCurrency: userCurrency,
	}
	s.EXPECT().ListProducts(gomock.Any(), args).Times(1).
		Return(products, nil)

	u, err := url.Parse("/products")
	require.NoError(t, err)

	q := url.Values{}
	q.Add("page", strconv.Itoa(page))
	q.Add("pageSize", strconv.Itoa(pageSize))
	q.Add("currency", userCurrency)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Products []*productsv1.Product
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)

	require.Len(t, gotResp.Products, len(products))
	require.Equal(t, gotResp.Products, products)
}

func TestListProductsNoQueryParams(t *testing.T) {
	products := []*productsv1.Product{
		&productsv1.Product{
			Id:          2,
			Name:        "tea",
			Description: "black",
			Picture:     "/tea",
			Price:       1.17,
			Currency:    "EUR",
		},
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	validate := validator.New()
	s := mock_service.NewMockProductsService(ctrl)
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.ListProductsArgs{
		Page:         int32(defaultPage),
		PageSize:     int32(defaultPageSize),
		UserCurrency: config.DefaultCurrency,
	}
	s.EXPECT().ListProducts(gomock.Any(), args).Times(1).
		Return(products, nil)

	req, err := http.NewRequest(http.MethodGet, "/products", nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Products []*productsv1.Product
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)

	require.Len(t, gotResp.Products, len(products))
	require.Equal(t, gotResp.Products, products)
}

func TestListProductsInvalidPageQueryParam(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	validate := validator.New()
	s := mock_service.NewMockProductsService(ctrl)
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	req, err := http.NewRequest(http.MethodGet, "/products", nil)
	require.NoError(t, err)

	q := req.URL.Query()
	q.Add("page", strconv.Itoa(-1))
	req.URL.RawQuery = q.Encode()

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestListProductsNoProducts(t *testing.T) {
	products := []*productsv1.Product{}
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	validate := validator.New()
	s := mock_service.NewMockProductsService(ctrl)
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.ListProductsArgs{
		Page:         int32(defaultPage),
		PageSize:     int32(defaultPageSize),
		UserCurrency: config.DefaultCurrency,
	}
	s.EXPECT().ListProducts(gomock.Any(), args).Times(1).
		Return(products, nil)

	req, err := http.NewRequest(http.MethodGet, "/products", nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Products []*productsv1.Product
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)

	require.Len(t, gotResp.Products, len(products))
	require.Equal(t, gotResp.Products, products)
}

func TestListProductsInvalidAgrument(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	validate := validator.New()
	s := mock_service.NewMockProductsService(ctrl)
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.ListProductsArgs{
		Page:         int32(defaultPage),
		PageSize:     int32(defaultPageSize),
		UserCurrency: config.DefaultCurrency,
	}
	s.EXPECT().ListProducts(gomock.Any(), args).Times(1).
		Return(nil, status.Errorf(codes.InvalidArgument, "mock error"))

	req, err := http.NewRequest(http.MethodGet, "/products", nil)
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

func TestListProductsInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	config := utils.Config{
		DefaultCurrency: "EUR",
	}
	validate := validator.New()
	s := mock_service.NewMockProductsService(ctrl)
	h := NewProductsHandler(config, s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.ListProductsArgs{
		Page:         int32(defaultPage),
		PageSize:     int32(defaultPageSize),
		UserCurrency: config.DefaultCurrency,
	}
	s.EXPECT().ListProducts(gomock.Any(), args).Times(1).
		Return(nil, status.Errorf(codes.Internal, "mock internal error"))

	req, err := http.NewRequest(http.MethodGet, "/products", nil)
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
	require.Contains(t, gotResp.Error, "cannot retrieve a list of products")
}
