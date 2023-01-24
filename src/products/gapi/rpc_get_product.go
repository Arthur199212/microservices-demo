package gapi

import (
	"context"
	"database/sql"

	"github.com/Arthur199212/microservices-demo/src/products/db"
	"github.com/Arthur199212/microservices-demo/src/products/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetProduct(
	ctx context.Context,
	req *pb.GetProductRequest,
) (*pb.GetProductResponse, error) {
	id := req.GetId()
	if id < 1 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"product ID should be greater than 0",
		)
	}

	product, err := s.pdb.GetProduct(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(
				codes.NotFound,
				"product with ID=%d not found",
				id,
			)
		}

		return nil, status.Errorf(
			codes.Internal,
			"failed to get a list of products: %s",
			err,
		)
	}

	return &pb.GetProductResponse{
		Product: convertProduct(product),
	}, nil
}

func convertProduct(product db.Product) *pb.Product {
	return &pb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Picture:     product.Picture,
		Price:       product.Price,
		Currency:    product.Currency,
	}
}
