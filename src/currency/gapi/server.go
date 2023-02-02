package gapi

import (
	"github.com/Arthur199212/microservices-demo/src/currencies/data"
	"github.com/Arthur199212/microservices-demo/src/currencies/pb"
)

type Server struct {
	pb.UnimplementedProductsServer
	cd data.CurrencyData
}

func NewServer(cd data.CurrencyData) *Server {
	return &Server{
		cd: cd,
	}
}
