DOCKER=$(shell which docker)
GO=$(shell which go)
GO_IMAGE=headphonista:go


.PHONY: docker/build docker/run build

docker/build:
	$(DOCKER) build -t $(GO_IMAGE) .

docker/run:
	$(DOCKER) run -it --rm $(GO_IMAGE) /bin/bash

build:
	$(GO) build *.go

