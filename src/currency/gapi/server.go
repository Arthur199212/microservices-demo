package gapi

import (
	"github.com/Arthur199212/microservices-demo/src/currency/data"
	"github.com/Arthur199212/microservices-demo/src/currency/pb"
)

type Server struct {
	pb.UnimplementedCurrencyServer
	cd data.CurrencyData
}

func NewServer(cd data.CurrencyData) *Server {
	return &Server{
		cd: cd,
	}
}
