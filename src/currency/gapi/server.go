package gapi

import (
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	"github.com/Arthur199212/microservices-demo/src/currency/data"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	currencyv1.UnimplementedCurrencyServiceServer
	grpc_health_v1.UnimplementedHealthServer
	cd data.CurrencyData
}

func NewServer(cd data.CurrencyData) *Server {
	return &Server{
		cd: cd,
	}
}
