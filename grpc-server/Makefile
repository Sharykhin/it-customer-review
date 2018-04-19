#!make
.PHONY: test lint migrate-up migrate-down

include .docker/golang/.env
export $(shell sed 's/=.*//' .docker/golang/.env)

test:
	go test -v ./...

lint:
	gometalinter.v2 ./...

migrate-up:
	migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ up

migrate-down:
	migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ down


