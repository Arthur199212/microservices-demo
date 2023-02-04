package gapi

import (
	"context"
	"strconv"

	"github.com/Arthur199212/microservices-demo/src/payment/pb"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Charge(
	ctx context.Context,
	req *pb.ChargeRequest,
) (*pb.ChargeResponse, error) {
	money := req.GetMoney()
	if money.GetAmount() < 0 {
		return &pb.ChargeResponse{}, status.Errorf(codes.InvalidArgument, "amount should be greater than 0")
	}

	cardInfo := req.GetCardInfo()
	card := creditcard.Card{
		Number: cardInfo.GetNumber(),
		Cvv:    strconv.Itoa(int(cardInfo.GetCvv())),
		Month:  strconv.Itoa(int(cardInfo.GetExpirationMonth())),
		Year:   strconv.Itoa(int(cardInfo.GetExpirationYear())),
	}
	if err := card.Validate(s.config.AllowTestCardNumbers); err != nil {
		return &pb.ChargeResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err := card.Method(); err != nil {
		return &pb.ChargeResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if !(card.Company.Short == "visa" || card.Company.Short == "mastercard") {
		return &pb.ChargeResponse{}, status.Errorf(codes.InvalidArgument, "only visa of mastercard is accepted")
	}

	transactionId := uuid.New().String()
	return &pb.ChargeResponse{
		TransactionId: transactionId,
	}, nil
}
