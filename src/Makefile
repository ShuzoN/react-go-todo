GO=$(shell which go)

.PHONY: build fmt

build: fmt
	$(GO) $@ main.go

fmt:
	$(GO) fmt ./main.go
