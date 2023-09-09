.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	@if [ -n "$$(gofmt -l .)" ]; \
		then echo 'Forgot to run "make lint_fix"?' && exit 1; \
	fi

.PHONY: lint_fix
lint_fix:
	@gofmt -l -w .

.PHONY: test
test:
	go test -v ./... -race -cover

.PHONY: mockgen
mockgen:
	@rm -f shipping/service/mocks/*.go
	@mockgen github.com/Arthur199212/microservices-demo/src/api_gateway/cart/service CartService > cart/service/mocks/cart_service_mock.go
	@mockgen github.com/Arthur199212/microservices-demo/src/api_gateway/checkout/service CheckoutService > checkout/service/mocks/checkout_service_mock.go
	@mockgen github.com/Arthur199212/microservices-demo/src/api_gateway/products/service ProductsService > products/service/mocks/products_service_mock.go
	@mockgen github.com/Arthur199212/microservices-demo/src/api_gateway/shipping/service ShippingService > shipping/service/mocks/shipping_service_mock.go
