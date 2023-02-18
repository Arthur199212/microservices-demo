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
	evans --host localhost --port 5004 -r repl

.PHONY: test
test:
	go test -v ./... -race -cover

.PHONY: mockgen
mockgen:
	rm -f shipping/mocks/*.go
	mockgen github.com/Arthur199212/microservices-demo/src/shipping/shipping ShippingService > shipping/mocks/shipping_service_mock.go
