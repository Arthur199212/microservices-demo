package gapi

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Convert(
	ctx context.Context,
	req *currencyv1.ConvertRequest,
) (*modelsv1.Money, error) {
	money := req.GetFrom()
	if money.GetAmount() < 0 {
		return &modelsv1.Money{}, status.Errorf(codes.InvalidArgument, "amount should be greater than zero")
	}

	if err := s.cd.VerifySupportedCurrency(money.CurrencyCode); err != nil {
		return &modelsv1.Money{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err := s.cd.VerifySupportedCurrency(req.GetToCurrencyCode()); err != nil {
		return &modelsv1.Money{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	convertedAmount, err := s.cd.Convert(money.CurrencyCode, req.GetToCurrencyCode(), money.GetAmount())
	if err != nil {
		return &modelsv1.Money{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &modelsv1.Money{
		CurrencyCode: req.GetToCurrencyCode(),
		Amount:       convertedAmount,
	}, nil
}
