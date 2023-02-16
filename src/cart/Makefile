.PHONY: run
run: 
	go run main.go 

.PHONY: lint
lint:
	gofmt -l -w .

.PHONY: evans
evans:
	evans --host localhost --port 5001 -r repl

.PHONY: test
test:
	go test -v ./... -race -cover
