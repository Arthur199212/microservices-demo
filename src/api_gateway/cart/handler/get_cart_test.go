package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	mock_service "github.com/Arthur199212/microservices-demo/src/api_gateway/cart/service/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetCartOK(t *testing.T) {
	products := []*modelsv1.Product{
		&modelsv1.Product{
			Id:       1,
			Quantity: 2,
		},
		&modelsv1.Product{
			Id:       2,
			Quantity: 4,
		},
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	sessionId := "49344e5d-0a2d-45e5-9c6b-60628aed043b"
	s.EXPECT().GetCart(gomock.Any(), sessionId).Times(1).
		Return(products, nil)

	u := fmt.Sprintf("/cart/%s", sessionId)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Products []*modelsv1.Product
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)
	require.Len(t, gotResp.Products, len(products))
	require.Equal(t, gotResp.Products, products)
}

func TestGetCartInvalidSessionId(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	sessionId := "invalid-uuid"
	u := fmt.Sprintf("/cart/%s", sessionId)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetCartNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	sessionId := "49344e5d-0a2d-45e5-9c6b-60628aed043b"
	s.EXPECT().GetCart(gomock.Any(), sessionId).Times(1).
		Return(nil, status.Errorf(codes.NotFound, "cart not found"))

	u := fmt.Sprintf("/cart/%s", sessionId)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetCartInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	sessionId := "49344e5d-0a2d-45e5-9c6b-60628aed043b"
	s.EXPECT().GetCart(gomock.Any(), sessionId).Times(1).
		Return(nil, status.Errorf(codes.Internal, "mock error"))

	u := fmt.Sprintf("/cart/%s", sessionId)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gotResp struct {
		Error string `json:"error"`
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)
	require.Contains(t, gotResp.Error, "cannot retrieve cart")
}
