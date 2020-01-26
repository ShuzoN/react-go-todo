GO=$(shell which go)
BIN=./bin

.PHONY: build fmt

build: fmt
	$(GO) $@ -o $(BIN)/main ./main.go

fmt:
	$(GO) fmt ./main.go

run: build
	$(BIN)/main
