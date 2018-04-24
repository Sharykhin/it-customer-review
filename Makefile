#!make
.PHONY: install

install:
	cp api/.docker/golang/.env.example api/.docker/golang/.env
	cp grpc-server/.docker/golang/.env.example grpc-server/.docker/golang/.env
	cp grpc-server/.docker/mysql/.env.example grpc-server/.docker/mysql/.env
	cp tone-analyzer/.docker/golang/.env.example tone-analyzer/.docker/golang/.env
	cp tone-analyzer/.docker/rabbitmq/.env.example tone-analyzer/.docker/rabbitmq/.env
	docker-compose build
