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

.PHONY: evans
evans:
	evans --host localhost --port 5002 -r repl

.PHONY: test
test:
	go test -v ./... -race -cover
