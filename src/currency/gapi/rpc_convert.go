package gapi

import (
	"context"

	"github.com/Arthur199212/microservices-demo/src/currencies/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Convert(
	ctx context.Context,
	req *pb.ConvertRequest,
) (*pb.Money, error) {
	money := req.GetFrom()
	if money.GetAmount() < 0 {
		return &pb.Money{}, status.Errorf(codes.InvalidArgument, "amount should be greater than zero")
	}

	if err := s.cd.VerifySupportedCurrency(money.CurrencyCode); err != nil {
		return &pb.Money{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err := s.cd.VerifySupportedCurrency(req.GetToCurrencyCode()); err != nil {
		return &pb.Money{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	convertedAmount, err := s.cd.Convert(money.CurrencyCode, req.GetToCurrencyCode(), money.GetAmount())
	if err != nil {
		return &pb.Money{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.Money{
		CurrencyCode: req.GetToCurrencyCode(),
		Amount:       convertedAmount,
	}, nil
}
