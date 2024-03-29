package gapi

import (
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	"github.com/Arthur199212/microservices-demo/src/products/db"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	productsv1.UnimplementedProductsServiceServer
	grpc_health_v1.UnimplementedHealthServer
	pdb db.ProductDB
}

func NewServer(pdb db.ProductDB) *Server {
	return &Server{
		pdb: pdb,
	}
}
