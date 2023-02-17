package gapi

import (
	"context"

	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Charge(
	ctx context.Context,
	req *paymentv1.ChargeRequest,
) (*paymentv1.ChargeResponse, error) {
	money := req.GetMoney()
	if money.GetAmount() < 0 {
		return &paymentv1.ChargeResponse{}, status.Errorf(codes.InvalidArgument, "amount should be greater than 0")
	}

	cardInfo := req.GetCardInfo()
	card := creditcard.Card{
		Number: cardInfo.GetNumber(),
		Cvv:    cardInfo.GetCvv(),
		Month:  cardInfo.GetExpirationMonth(),
		Year:   cardInfo.GetExpirationYear(),
	}
	if err := card.Validate(s.config.AllowTestCardNumbers); err != nil {
		return &paymentv1.ChargeResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err := card.Method(); err != nil {
		return &paymentv1.ChargeResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if !(card.Company.Short == "visa" || card.Company.Short == "mastercard") {
		return &paymentv1.ChargeResponse{}, status.Errorf(codes.InvalidArgument, "only visa of mastercard is accepted")
	}

	transactionId := uuid.New().String()
	return &paymentv1.ChargeResponse{
		TransactionId: transactionId,
	}, nil
}
