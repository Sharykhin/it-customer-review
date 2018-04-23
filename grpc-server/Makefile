#!make
.PHONY: test lint migrate-up migrate-down

include .docker/golang/.env
export $(shell sed 's/=.*//' .docker/golang/.env)

test:
	go test -v ./...

lint:
	gometalinter.v2 ./... --deadline=60s

migrate-up:
	migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ up

migrate-down:
	migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ down

local-serve:
	DB_USER=root DB_PASS=root DB_HOST=localhost DB_NAME=test GRPC_ADDRESS="localhost:50051" go run *.go


