package service

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/currency"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/utils"
)

type ProductsService interface {
	GetProductById(ctx context.Context, args GetProductByIdArgs) (*productsv1.Product, error)
	ListProducts(ctx context.Context, args ListProductsArgs) ([]*productsv1.Product, error)
}

type productsService struct {
	config          utils.Config
	productsClient  productsv1.ProductsServiceClient
	currencyService currency.CurrencyService
}

func NewProductsService(
	config utils.Config,
	productsClient productsv1.ProductsServiceClient,
	currencyService currency.CurrencyService,
) ProductsService {
	return &productsService{
		config:          config,
		productsClient:  productsClient,
		currencyService: currencyService,
	}
}

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
	if args.UserCurrency == s.config.DefaultCurrency {
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

	products := resp.GetProducts()
	if products == nil {
		products = make([]*productsv1.Product, 0)
		return products, nil
	}

	if args.UserCurrency == s.config.DefaultCurrency {
		return products, nil
	}

	for _, product := range products {
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

	return products, nil
}
