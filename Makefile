DOCKER=$(shell which docker)
DOCKER_COMPOSE=$(shell which docker-compose)
GO=$(shell which go)
GO_IMAGE=headphonista:go


.PHONY: docker/build docker/run build

docker/build:
	$(DOCKER) build -t $(GO_IMAGE) .

docker/run:
	$(DOCKER) run --rm $(GO_IMAGE)

docker/run/bash:
	$(DOCKER) run -it --rm $(GO_IMAGE) /bin/bash

up:
	$(DOCKER_COMPOSE) up --build  --remove-orphans -d

down:
	$(DOCKER_COMPOSE) down

