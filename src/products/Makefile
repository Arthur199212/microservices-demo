# -include .env
# export

.PHONY: run
run:
	go run main.go

# up:
# 	docker compose up -d

# down:
# 	docker compose down

.PHONY: proto
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 5000 -r repl

