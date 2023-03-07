package service

import (
	"context"

	modelsv1 "github.com/Arthur199212/microservices-demo/gen/models/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	"github.com/Arthur199212/microservices-demo/src/api_gateway/models"
)

type CheckoutService interface {
	PlaceOrder(ctx context.Context, args CheckoutServiceArgs) (*checkoutv1.Order, error)
}

type checkoutService struct {
	checkoutClient checkoutv1.CheckoutServiceClient
}

func NewCheckoutService(
	checkoutClient checkoutv1.CheckoutServiceClient,
) CheckoutService {
	return &checkoutService{
		checkoutClient: checkoutClient,
	}
}

type CardInfo struct {
	Cvv             string `json:"cvv" validate:"required,numeric,min=3,max=4"`
	ExpirationMonth string `json:"expirationMonth" validate:"required,numeric,min=1,max=2"`
	ExpirationYear  string `json:"expirationYear" validate:"required,numeric,len=4"`
	Number          string `json:"number" validate:"required,numeric,min=8,max=19"`
}

type CheckoutServiceArgs struct {
	Email        string         `json:"email" validate:"required,email"`
	SessionId    string         `json:"sessionId" validate:"required,uuid4"`
	UserCurrency string         `json:"userCurrency" validate:"required,len=3"`
	Address      models.Address `json:"address"`
	CardInfo     CardInfo       `json:"cardInfo"`
}

func (s *checkoutService) PlaceOrder(
	ctx context.Context,
	args CheckoutServiceArgs,
) (*checkoutv1.Order, error) {
	state := ""
	if args.Address.State != nil {
		state = *args.Address.State
	}
	resp, err := s.checkoutClient.PlaceOrder(ctx, &checkoutv1.PlaceOrderRequest{
		SessionId:    args.SessionId,
		UserCurrency: args.UserCurrency,
		Address: &modelsv1.Address{
			City:          args.Address.City,
			Country:       args.Address.Country,
			State:         state,
			StreetAddress: args.Address.StreetAddress,
			ZipCode:       args.Address.ZipCode,
		},
		Email: args.Email,
		CardInfo: &modelsv1.CardInfo{
			Cvv:             args.CardInfo.Cvv,
			ExpirationMonth: args.CardInfo.ExpirationMonth,
			ExpirationYear:  args.CardInfo.ExpirationYear,
			Number:          args.CardInfo.Number,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.GetOrder(), nil
}
