package checkout

import (
	"context"

	cart "github.com/Arthur199212/microservices-demo/src/cart/pb"
	"github.com/Arthur199212/microservices-demo/src/checkout/pb"
	currency "github.com/Arthur199212/microservices-demo/src/currency/pb"
	payment "github.com/Arthur199212/microservices-demo/src/payment/pb"
	products "github.com/Arthur199212/microservices-demo/src/products/pb"
	shipping "github.com/Arthur199212/microservices-demo/src/shipping/pb"
)

type CheckoutService interface {
	PlaceOrder(context.Context, PlaceOrderArgs) (*pb.Order, error)
}

type checkoutService struct {
	cartClient     cart.CartClient
	currencyClient currency.CurrencyClient
	paymentClient  payment.PaymentClient
	productsClient products.ProductsClient
	shippingClient shipping.ShippingClient
}

func NewCheckoutService(
	cartClient cart.CartClient,
	currencyClient currency.CurrencyClient,
	paymentClient payment.PaymentClient,
	productsClient products.ProductsClient,
	shippingClient shipping.ShippingClient,
) CheckoutService {
	return &checkoutService{
		cartClient:     cartClient,
		currencyClient: currencyClient,
		paymentClient:  paymentClient,
		productsClient: productsClient,
		shippingClient: shippingClient,
	}
}
