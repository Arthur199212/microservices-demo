package gapi

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	"github.com/Arthur199212/microservices-demo/src/checkout/checkout"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) PlaceOrder(ctx context.Context, req *checkoutv1.PlaceOrderRequest) (*checkoutv1.PlaceOrderResponse, error) {
	args := checkout.PlaceOrderArgs{
		Address:      convertToAddress(req.GetAddress()),
		CardInfo:     convertToCardInfo(req.GetCardInfo()),
		Email:        req.GetEmail(),
		SessionId:    req.GetSessionId(),
		UserCurrency: req.GetUserCurrency(),
	}
	if err := s.validate.Struct(args); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	order, err := s.checkoutService.PlaceOrder(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &checkoutv1.PlaceOrderResponse{
		Order: order,
	}, nil
}

func convertToAddress(address *modelsv1.Address) checkout.Address {
	return checkout.Address{
		StreetAddress: address.StreetAddress,
		City:          address.City,
		Country:       address.Country,
		ZipCode:       address.ZipCode,
		State:         &address.State,
	}
}

func convertToCardInfo(cardInfo *modelsv1.CardInfo) checkout.CardInfo {
	return checkout.CardInfo{
		Cvv:             cardInfo.Cvv,
		ExpirationMonth: cardInfo.ExpirationMonth,
		ExpirationYear:  cardInfo.ExpirationYear,
		Number:          cardInfo.Number,
	}
}
