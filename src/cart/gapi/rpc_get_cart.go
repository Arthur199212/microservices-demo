package gapi

import (
	"context"
	"database/sql"

	"github.com/Arthur199212/microservices-demo/src/cart/db"
	"github.com/Arthur199212/microservices-demo/src/cart/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCart(
	ctx context.Context,
	req *pb.GetCartRequest,
) (*pb.GetCartResponse, error) {
	sessionId := req.GetSessionId()
	if sessionId == "" {
		return &pb.GetCartResponse{}, status.Errorf(codes.InvalidArgument, "sessionId cannot be empty")
	}

	cart, err := s.cartDB.GetCart(sessionId)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.GetCartResponse{},
				status.Errorf(codes.NotFound, "cart with sessionId=%s not found", sessionId)
		}
		return &pb.GetCartResponse{}, status.Errorf(codes.Internal, "cannot clear cart: %v", err)
	}

	return &pb.GetCartResponse{
		SessionId: cart.SessionID,
		Products:  convertProducts(cart.Products),
	}, nil
}

func convertProducts(products []*db.Product) []*pb.Product {
	res := make([]*pb.Product, len(products))
	for i, p := range products {
		res[i] = &pb.Product{
			Id:       p.ID,
			Quantity: p.Quantity,
		}
	}
	return res
}
