package gapi

import (
	"context"
	"database/sql"

	"github.com/Arthur199212/microservices-demo/src/cart/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ClearCart(
	ctx context.Context,
	req *pb.ClearCartRequest,
) (*emptypb.Empty, error) {
	emptyResp := &emptypb.Empty{}

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
		return emptyResp, status.Errorf(codes.Internal, "cannot clear cart:", err)
	}

	return emptyResp, nil
}
