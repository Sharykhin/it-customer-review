#!make
.PHONY: test lint

test:
	go test -v ./...

lint:
	gometalinter.v2 ./...

local-serve:
	GRPC_ADDRESS="localhost:50051" HTTP_ADDRESS=":8081" CORS_ORIGIN="*" SESSION_KEY="test" go run *.go

