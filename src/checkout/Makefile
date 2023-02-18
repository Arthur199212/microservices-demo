.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	gofmt -l -w .

.PHONY: proto
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: evans
evans:
	evans --host localhost --port 5002 -r repl

.PHONY: test
test:
	go test -v ./... -race -cover

.PHONY: mockgen
mockgen:
	@echo "deleting mocks..."
	@rm -f mock/*.go
	@rm -f checkout/mock/*.go
	@echo "mocking grpc clients..."
	@mockgen github.com/Arthur199212/microservices-demo/gen/services/cart/v1 CartServiceClient > mock/cart_service_mock.go
	@mockgen github.com/Arthur199212/microservices-demo/gen/services/currency/v1 CurrencyServiceClient > mock/currency_service_mock.go
	@mockgen github.com/Arthur199212/microservices-demo/gen/services/payment/v1 PaymentServiceClient > mock/payment_service_mock.go
	@mockgen github.com/Arthur199212/microservices-demo/gen/services/products/v1 ProductsServiceClient > mock/products_service_mock.go
	@mockgen github.com/Arthur199212/microservices-demo/gen/services/shipping/v1 ShippingServiceClient > mock/shipping_service_mock.go
	@echo "mocking checkout service..."
	@mockgen github.com/Arthur199212/microservices-demo/src/checkout/checkout CheckoutService > checkout/mock/checkout_service_mock.go
	@echo "done"
