package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	mock_service "github.com/Arthur199212/microservices-demo/src/api_gateway/cart/service/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestClearCartOK(t *testing.T) {
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
	s.EXPECT().ClearCart(gomock.Any(), sessionId).Times(1).
		Return(nil)

	u := fmt.Sprintf("/cart/%s", sessionId)
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestClearCartInvalidSessionId(t *testing.T) {
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
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestClearCartNotFound(t *testing.T) {
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
	s.EXPECT().ClearCart(gomock.Any(), sessionId).Times(1).
		Return(status.Errorf(codes.NotFound, "cart not found"))

	u := fmt.Sprintf("/cart/%s", sessionId)
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	require.NoError(t, err)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestClearCartInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	sessionId := "d1887de8-5894-4e18-9a33-ee45324b653b"
	s.EXPECT().ClearCart(gomock.Any(), sessionId).Times(1).
		Return(status.Errorf(codes.Internal, "mock error"))

	u := fmt.Sprintf("/cart/%s", sessionId)
	req, err := http.NewRequest(http.MethodDelete, u, nil)
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
	require.Contains(t, gotResp.Error, "cannot clear cart")
}
