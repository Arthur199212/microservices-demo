package gapi

import (
	"github.com/Arthur199212/microservices-demo/src/cart/db"
	"github.com/Arthur199212/microservices-demo/src/cart/pb"
)

type Server struct {
	pb.UnimplementedCartServer
	cartDB db.CartDB
}

func NewServer(cartDB db.CartDB) *Server {
	return &Server{
		cartDB: cartDB,
	}
}
