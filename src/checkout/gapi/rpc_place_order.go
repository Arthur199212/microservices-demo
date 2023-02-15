package gapi

import (
	"context"

	"github.com/Arthur199212/microservices-demo/src/checkout/checkout"
	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	"github.com/Arthur199212/microservices-demo/src/shipping/shipping"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) PlaceOrder(ctx context.Context, req *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	args := checkout.PlaceOrderArgs{
		Address:      convertToAddress(req.GetAddress()),
		CardInfo:     convertToCardInfo(req.GetCardInfo()),
		Email:        req.GetEmail(),
		SessionId:    req.GetSessionId(),
		UserCurrency: req.GetUserCurrency(),
	}
	if err := s.validate.Struct(args); err != nil {
		status.Errorf(codes.InvalidArgument, err.Error())
	}

	order, err := s.checkoutService.PlaceOrder(ctx, args)
	if err != nil {
		status.Errorf(codes.Internal, err.Error())
	}

	return &pb.PlaceOrderResponse{
		Order: order,
	}, nil
}

func convertToAddress(address *pb.Address) shipping.Address {
	return shipping.Address{
		StreetAddress: address.StreetAddress,
		City:          address.City,
		Country:       address.Country,
		ZipCode:       address.ZipCode,
		State:         &address.State,
	}
}

func convertToCardInfo(cardInfo *pb.CardInfo) checkout.CardInfo {
	return checkout.CardInfo{
		Cvv:             cardInfo.Cvv,
		ExpirationMonth: cardInfo.ExpirationMonth,
		ExpirationYear:  cardInfo.ExpirationYear,
		Number:          cardInfo.Number,
	}
}
