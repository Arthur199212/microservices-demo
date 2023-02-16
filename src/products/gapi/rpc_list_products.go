package gapi

import (
	"context"

	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	"github.com/Arthur199212/microservices-demo/src/products/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	defaultPageSize = 10
)

func (s *Server) ListProducts(
	ctx context.Context,
	req *productsv1.ListProductsRequest,
) (*productsv1.ListProductsResponse, error) {
	page := req.GetPage()
	pageSize := req.GetPageSize()

	if page < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "page should be greater than zero")
	}
	if pageSize < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "limit should be greater or equal than zero")
	}
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	products, err := s.pdb.ListProducts(page, pageSize)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to get a list of products: %s",
			err,
		)
	}

	return &productsv1.ListProductsResponse{
		Products: convertProductsList(products),
	}, nil
}

func convertProductsList(products []db.Product) []*productsv1.Product {
	pl := make([]*productsv1.Product, len(products))

	for i, product := range products {
		pl[i] = convertProduct(product)
	}

	return pl
}
