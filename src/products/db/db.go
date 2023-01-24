package db

import (
	"database/sql"
)

type ProductDB interface {
	GetProduct(id int64) (Product, error)
	ListProducts(page, pageSize int32) ([]Product, error)
}

type FakeStore struct {
	// only read operations so no mutexes or kind of
	products []Product
}

func NewProductsDB() ProductDB {
	return &FakeStore{
		products: products,
	}
}

func (s *FakeStore) GetProduct(id int64) (Product, error) {
	// just linear search for demo
	for _, product := range s.products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, sql.ErrNoRows
}

func (s *FakeStore) ListProducts(page, pageSize int32) ([]Product, error) {
	page = page - 1 // pages go from 1 but indexes from 0
	if page < 0 {
		page = 0
	}
	if pageSize < 0 {
		pageSize = 0
	}

	l := page * pageSize
	if l > int32(len(s.products)) {
		return []Product{}, nil
	}

	r := l + pageSize
	if r > int32(len(s.products)) {
		r = int32(len(s.products))
	}

	pl := products[l:r]

	return pl, nil
}

var products = []Product{
	{
		ID:          1,
		Name:        "Product #1",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       0.99,
		Currency:    "EUR",
	},
	{
		ID:          2,
		Name:        "Product #2",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       3.47,
		Currency:    "EUR",
	},
	{
		ID:          3,
		Name:        "Product #3",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       1.45,
		Currency:    "EUR",
	},
	{
		ID:          4,
		Name:        "Product #4",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       3.47,
		Currency:    "EUR",
	},
	{
		ID:          5,
		Name:        "Product #5",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       5.38,
		Currency:    "EUR",
	},
	{
		ID:          6,
		Name:        "Product #6",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       7.20,
		Currency:    "EUR",
	},
	{
		ID:          7,
		Name:        "Product #7",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       0.99,
		Currency:    "EUR",
	},
	{
		ID:          8,
		Name:        "Product #8",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       3.47,
		Currency:    "EUR",
	},
	{
		ID:          9,
		Name:        "Product #9",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       1.45,
		Currency:    "EUR",
	},
	{
		ID:          10,
		Name:        "Product #10",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       3.47,
		Currency:    "EUR",
	},
	{
		ID:          11,
		Name:        "Product #11",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       5.38,
		Currency:    "EUR",
	},
	{
		ID:          12,
		Name:        "Product #12",
		Description: "Description of the product comes here",
		Picture:     "/url",
		Price:       7.20,
		Currency:    "EUR",
	},
}
