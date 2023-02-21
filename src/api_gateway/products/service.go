package products

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/currency"
)

type ProductsService interface {
	GetProductById(ctx context.Context, args GetProductByIdArgs) (*productsv1.Product, error)
	ListProducts(ctx context.Context, args ListProductsArgs) ([]*productsv1.Product, error)
}

type productsService struct {
	productsClient  productsv1.ProductsServiceClient
	currencyService currency.CurrencyService
}

func NewProductsService(
	productsClient productsv1.ProductsServiceClient,
	currencyService currency.CurrencyService,
) ProductsService {
	return &productsService{
		productsClient:  productsClient,
		currencyService: currencyService,
	}
}

const (
	defaultCurrency = "EUR"
)

type GetProductByIdArgs struct {
	Id           int64  `json:"id" validate:"required,min=1"`
	UserCurrency string `json:"userCurrency" validate:"required,len=3"`
}

func (s *productsService) GetProductById(
	ctx context.Context,
	args GetProductByIdArgs,
) (*productsv1.Product, error) {
	resp, err := s.productsClient.GetProduct(ctx, &productsv1.GetProductRequest{
		Id: args.Id,
	})
	if err != nil {
		return nil, err
	}

	product := resp.GetProduct()
	if args.UserCurrency == defaultCurrency {
		return product, nil
	}

	money, err := s.currencyService.Convert(ctx, currency.ConvertArgs{
		From: &modelsv1.Money{
			Amount:       product.Price,
			CurrencyCode: product.Currency,
		},
		ToCurrencyCode: args.UserCurrency,
	})
	if err != nil {
		return nil, err
	}

	product.Price = money.GetAmount()
	product.Currency = money.GetCurrencyCode()

	return product, nil
}

type ListProductsArgs struct {
	Page         int32  `json:"page" validate:"required,min=1"`
	PageSize     int32  `json:"pageSize" validate:"required,min=1,max=50"`
	UserCurrency string `json:"userCurrency" validate:"required,len=3"`
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

	if args.UserCurrency == defaultCurrency {
		return resp.GetProducts(), nil
	}

	for _, product := range resp.GetProducts() {
		money, err := s.currencyService.Convert(ctx, currency.ConvertArgs{
			From: &modelsv1.Money{
				Amount:       product.Price,
				CurrencyCode: product.Currency,
			},
			ToCurrencyCode: args.UserCurrency,
		})
		if err != nil {
			return nil, err
		}

		product.Price = money.GetAmount()
		product.Currency = money.GetCurrencyCode()
	}

	return resp.Products, nil
}
