package gapi

import (
	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	"github.com/Arthur199212/microservices-demo/src/cart/db"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	cartv1.UnimplementedCartServiceServer
	grpc_health_v1.UnimplementedHealthServer
	cartDB db.CartDB
}

func NewServer(cartDB db.CartDB) *Server {
	return &Server{
		cartDB: cartDB,
	}
}
