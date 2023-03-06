package checkout

import (
	"context"

	cartv1 "github.com/Arthur199212/microservices-demo/gen/services/cart/v1"
	checkoutv1 "github.com/Arthur199212/microservices-demo/gen/services/checkout/v1"
	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	paymentv1 "github.com/Arthur199212/microservices-demo/gen/services/payment/v1"
	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"github.com/Arthur199212/microservices-demo/src/checkout/utils"
)

type CheckoutService interface {
	PlaceOrder(context.Context, PlaceOrderArgs) (*checkoutv1.Order, error)
}

type checkoutService struct {
	config         utils.Config
	cartClient     cartv1.CartServiceClient
	currencyClient currencyv1.CurrencyServiceClient
	paymentClient  paymentv1.PaymentServiceClient
	productsClient productsv1.ProductsServiceClient
	shippingClient shippingv1.ShippingServiceClient
}

func NewCheckoutService(
	config utils.Config,
	cartClient cartv1.CartServiceClient,
	currencyClient currencyv1.CurrencyServiceClient,
	paymentClient paymentv1.PaymentServiceClient,
	productsClient productsv1.ProductsServiceClient,
	shippingClient shippingv1.ShippingServiceClient,
) CheckoutService {
	return &checkoutService{
		config:         config,
		cartClient:     cartClient,
		currencyClient: currencyClient,
		paymentClient:  paymentClient,
		productsClient: productsClient,
		shippingClient: shippingClient,
	}
}
