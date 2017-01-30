.PHONY: build run

NAME =middlewares

test:
	go test -v $(shell go list ./... | grep -v /vendor/)
