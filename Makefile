HTTP_PORT=80
HTTPS_PORT=443
MYSQL_PORT=3306
DOCKER=$(shell which docker)
DOCKER_COMPOSE=HTTP_PORT=$(HTTP_PORT) HTTPS_PORT=$(HTTPS_PORT) MYSQL_PORT=$(MYSQL_PORT) $(shell which docker-compose)
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

