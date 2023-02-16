package gapi

import (
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	"github.com/Arthur199212/microservices-demo/src/products/db"
)

type Server struct {
	productsv1.UnimplementedProductsServiceServer
	pdb db.ProductDB
}

func NewServer(pdb db.ProductDB) *Server {
	return &Server{
		pdb: pdb,
	}
}
