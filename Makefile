#!make
.PHONY: test migrate-up migrate-down serve docker-migrate-up docker-migrate-down docker-lint

include .docker/golang/.env
export $(shell sed 's/=.*//' .docker/golang/.env)

test:
	go run test -v ./...

docker-migrate-up:
	docker-compose exec cr-golang-api migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ up

docker-migrate-down:
	docker-compose exec cr-golang-api migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ down

docker-lint:
	docker-compose exec cr-golang-api gometalinter.v2 ./...

lint:
	gometalinter.v2 ./...

migrate-up:
	migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ up

migrate-down:
	migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}" -path migration/ down

serve:
	HTTP_ADDRESS=":8000" CORS_ORIGIN="*" go run *.go

