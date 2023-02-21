package products

import (
	"context"

	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
)

type ProductsService interface {
	GetProductById(ctx context.Context, id int64) (*productsv1.Product, error)
	ListProducts(ctx context.Context, args ListProductsArgs) ([]*productsv1.Product, error)
}

type productsService struct {
	productsClient productsv1.ProductsServiceClient
}

func NewProductsService(
	productsClient productsv1.ProductsServiceClient,
) ProductsService {
	return &productsService{
		productsClient: productsClient,
	}
}

func (s *productsService) GetProductById(
	ctx context.Context,
	id int64,
) (*productsv1.Product, error) {
	resp, err := s.productsClient.GetProduct(ctx, &productsv1.GetProductRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetProduct(), nil
}

type ListProductsArgs struct {
	Page     int32 `json:"page" validate:"required,min=1"`
	PageSize int32 `json:"pageSize" validate:"required,min=1,max=50"`
}

func (s *productsService) ListProducts(
	ctx context.Context,
	args ListProductsArgs,
) ([]*productsv1.Product, error) {
	resp, err := s.productsClient.ListProducts(ctx, &productsv1.ListProductsRequest{
		Page:     args.Page,
		PageSize: args.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetProducts(), nil
}
