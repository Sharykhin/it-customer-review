#!make
.PHONY: test lint

test:
	go test -v ./...

lint:
	gometalinter.v2 ./... --deadline=60s
