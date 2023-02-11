package gapi

import (
	"context"

	"github.com/Arthur199212/microservices-demo/src/cart/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AddItem(
	ctx context.Context,
	req *pb.AddItemRequest,
) (*emptypb.Empty, error) {
	emptyResp := &emptypb.Empty{}

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
