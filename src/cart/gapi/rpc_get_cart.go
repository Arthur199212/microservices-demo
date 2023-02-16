package gapi

import (
	"context"
	"database/sql"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	"github.com/Arthur199212/microservices-demo/src/cart/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCart(
	ctx context.Context,
	req *cartv1.GetCartRequest,
) (*cartv1.GetCartResponse, error) {
	sessionId := req.GetSessionId()
	if sessionId == "" {
		return &cartv1.GetCartResponse{}, status.Errorf(codes.InvalidArgument, "sessionId cannot be empty")
	}

	cart, err := s.cartDB.GetCart(sessionId)
	if err != nil {
		if err == sql.ErrNoRows {
			return &cartv1.GetCartResponse{},
				status.Errorf(codes.NotFound, "cart with sessionId=%s not found", sessionId)
		}
		return &cartv1.GetCartResponse{}, status.Errorf(codes.Internal, "cannot clear cart: %v", err)
	}

	return &cartv1.GetCartResponse{
		SessionId: cart.SessionID,
		Products:  convertProducts(cart.Products),
	}, nil
}

func convertProducts(products []*db.Product) []*modelsv1.Product {
	res := make([]*modelsv1.Product, len(products))
	for i, p := range products {
		res[i] = &modelsv1.Product{
			Id:       p.ID,
			Quantity: p.Quantity,
		}
	}
	return res
}
