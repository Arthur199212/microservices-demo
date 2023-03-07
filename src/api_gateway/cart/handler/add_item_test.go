package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/Arthur199212/microservices-demo/src/api_gateway/cart/service"
	mock_service "github.com/Arthur199212/microservices-demo/src/api_gateway/cart/service/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAddItemOK(t *testing.T) {
	var productId int64 = 1
	var quantity int32 = 3

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	sessionId := "b3cdd022-026d-41a8-a61c-ec4502ec15ef"
	args := service.AddItemsArgs{
		ProductId: productId,
		Quantity:  quantity,
	}
	s.EXPECT().AddItem(gomock.Any(), args).Times(1).
		Return(sessionId, nil)

	input := fiber.Map{
		"productId": productId,
		"quantity":  quantity,
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/cart", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})

	var gotResp struct {
		SessionId string
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)
	require.Equal(t, sessionId, gotResp.SessionId)
}

func TestAddItemInvalidContentType(t *testing.T) {
	var productId int64 = 1
	var quantity int32 = 3

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	input := fiber.Map{
		"productId": productId,
		"quantity":  quantity,
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/cart", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/xml")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAddItemInvalidProductId(t *testing.T) {
	var productId int64 = -1
	var quantity int32 = 3
	sessionId := "49344e5d-0a2d-45e5-9c6b-60628aed043b"

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	input := fiber.Map{
		"productId": productId,
		"quantity":  quantity,
		"sessionId": sessionId,
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/cart", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/xml")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAddItemWithSessionId(t *testing.T) {
	var productId int64 = 1
	var quantity int32 = 3
	sessionId := "49344e5d-0a2d-45e5-9c6b-60628aed043b"

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.AddItemsArgs{
		ProductId: productId,
		Quantity:  quantity,
		SessionId: &sessionId,
	}
	s.EXPECT().AddItem(gomock.Any(), args).Times(1).
		Return(sessionId, nil)

	input := fiber.Map{
		"productId": productId,
		"quantity":  quantity,
		"sessionId": sessionId,
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/cart", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})

	var gotResp struct {
		SessionId string
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)
	require.Equal(t, sessionId, gotResp.SessionId)
}

func TestAddItemInvalidSessionId(t *testing.T) {
	var productId int64 = 1
	var quantity int32 = 3
	sessionId := "000000-0000-invalid-uuid"

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	input := fiber.Map{
		"productId": productId,
		"quantity":  quantity,
		"sessionId": sessionId,
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/cart", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAddItemInternalError(t *testing.T) {
	var productId int64 = 1
	var quantity int32 = 3

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	s := mock_service.NewMockCartService(ctrl)
	validate := validator.New()
	h := NewCartHandler(s, validate)

	app := fiber.New()
	h.AddRoutes(app)

	args := service.AddItemsArgs{
		ProductId: productId,
		Quantity:  quantity,
	}
	s.EXPECT().AddItem(gomock.Any(), args).Times(1).
		Return("", status.Errorf(codes.Internal, "mock error"))

	input := fiber.Map{
		"productId": productId,
		"quantity":  quantity,
	}
	b, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/cart", bytes.NewReader(b))
	require.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})

	var gotResp struct {
		Error string
	}
	err = json.Unmarshal(data, &gotResp)
	require.NoError(t, err)
	require.Contains(t, gotResp.Error, "cannot add item to cart")
}
