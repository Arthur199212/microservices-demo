package gapi

import (
	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	"github.com/Arthur199212/microservices-demo/src/cart/db"
)

type Server struct {
	cartv1.UnimplementedCartServiceServer
	cartDB db.CartDB
}

func NewServer(cartDB db.CartDB) *Server {
	return &Server{
		cartDB: cartDB,
	}
}
