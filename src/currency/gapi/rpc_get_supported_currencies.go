package gapi

import (
	"context"

	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetSupportedCurrencies(
	ctx context.Context,
	req *emptypb.Empty,
) (*currencyv1.GetSupportedCurrenciesResponse, error) {
	currencies := s.cd.GetSupportedCurrencies()

	return &currencyv1.GetSupportedCurrenciesResponse{
		CurrencyCodes: currencies,
	}, nil
}
