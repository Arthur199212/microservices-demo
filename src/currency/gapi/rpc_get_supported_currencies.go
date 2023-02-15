package gapi

import (
	"context"

	"github.com/Arthur199212/microservices-demo/src/currency/pb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetSupportedCurrencies(
	ctx context.Context,
	req *emptypb.Empty,
) (*pb.GetSupportedCurrenciesResponse, error) {
	currencies := s.cd.GetSupportedCurrencies()

	return &pb.GetSupportedCurrenciesResponse{
		CurrencyCodes: currencies,
	}, nil
}
