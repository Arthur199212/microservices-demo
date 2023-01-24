package gapi

import (
	"github.com/Arthur199212/microservices-demo/src/products/db"
	"github.com/Arthur199212/microservices-demo/src/products/pb"
)

type Server struct {
	pb.UnimplementedProductsServer
	pdb db.ProductDB
}

func NewServer(pdb db.ProductDB) *Server {
	return &Server{
		pdb: pdb,
	}
}
