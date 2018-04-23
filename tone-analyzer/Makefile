#!make
.PHONY: test lint

test:
	go test -v ./...

lint:
	gometalinter.v2 ./...

local-serve:
	GRPC_ADDRESS=localhost:50051 AMPQ_ADDRESS="amqp://localhost:5672" go run *.go