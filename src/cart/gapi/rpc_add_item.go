package gapi

import (
	"context"

	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddItem(
	ctx context.Context,
	req *cartv1.AddItemRequest,
) (*cartv1.AddItemResponse, error) {
	emptyResp := &cartv1.AddItemResponse{}

	sessionId := req.GetSessionId()
	if sessionId == "" {
		return emptyResp, status.Errorf(codes.InvalidArgument, "sessionId cannot be empty")
	}

	quantity := req.GetQuantity()
	if quantity <= 0 {
		return emptyResp, status.Errorf(codes.InvalidArgument, "quantity shold be more then zero")
	}

	productId := req.GetProducId()
	if productId <= 0 {
		return emptyResp, status.Errorf(codes.InvalidArgument, "productId shold be more then zero")
	}

	err := s.cartDB.AddItem(sessionId, productId, quantity)
	if err != nil {
		status.Errorf(codes.Internal, "cannot add item to cart: %v", err)
	}

	return emptyResp, nil
}
