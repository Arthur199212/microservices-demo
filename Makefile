.PHONY: proto
proto:
	rm -f gen/*.go
	buf generate proto

.PHONY: proto_lint
proto_lint:
	buf lint proto
