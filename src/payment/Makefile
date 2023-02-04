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
	evans --host localhost --port 5003 -r repl
