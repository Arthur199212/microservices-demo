package gapi

import (
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	"github.com/Arthur199212/microservices-demo/src/currency/data"
)

type Server struct {
	currencyv1.UnimplementedCurrencyServiceServer
	cd data.CurrencyData
}

func NewServer(cd data.CurrencyData) *Server {
	return &Server{
		cd: cd,
	}
}
