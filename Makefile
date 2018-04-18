#!make
.PHONY: test migrate-up migrate-down serve

include .docker/golang/.env
export $(shell sed 's/=.*//' .docker/golang/.env)

test:
	go run test -v ./...

migrate-up:
	docker-compose exec cr-golang-api migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ up

migrate-down:
	docker-compose exec cr-golang-api migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ down

lint:
	docker-compose exec cr-golang-api gometalinter.v2 ./...

serve:
	HTTP_ADDRESS=":8000" CORS_ORIGIN="*" go run *.go

