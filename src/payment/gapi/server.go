package gapi

import (
	"github.com/Arthur199212/microservices-demo/src/payment/pb"
	"github.com/Arthur199212/microservices-demo/src/payment/utils"
)

type Server struct {
	pb.UnimplementedPaymentServer
	config *utils.Config
}

func NewServer(config *utils.Config) *Server {
	return &Server{
		config: config,
	}
}
