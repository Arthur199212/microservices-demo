.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	gofmt -l -w .

.PHONY: test
test:
	go test -v ./... -race -cover

.PHONY: mockgen
mockgen:
	@rm -f shipping/service/mocks/*.go
	@mockgen github.com/Arthur199212/microservices-demo/src/api_gateway/shipping/service ShippingService > shipping/service/mocks/shipping_service_mock.go
