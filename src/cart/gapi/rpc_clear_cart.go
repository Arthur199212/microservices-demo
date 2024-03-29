package gapi

import (
	"context"
	"database/sql"

	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ClearCart(
	ctx context.Context,
	req *cartv1.ClearCartRequest,
) (*cartv1.ClearCartResponse, error) {
	emptyResp := &cartv1.ClearCartResponse{}

	sessionId := req.GetSessionId()
	if sessionId == "" {
		return emptyResp, status.Errorf(codes.InvalidArgument, "sessionId cannot be empty")
	}

	err := s.cartDB.ClearCart(sessionId)
	if err != nil {
		if err == sql.ErrNoRows {
			return emptyResp,
				status.Errorf(codes.NotFound, "cart with sessionId=%s not found", sessionId)
		}
		return emptyResp, status.Errorf(codes.Internal, "cannot clear cart: %v", err)
	}

	return emptyResp, nil
}
