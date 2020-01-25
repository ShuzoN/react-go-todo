HTTP_PORT=80
HTTPS_PORT=443
MYSQL_PORT=3306
DOCKER=$(shell which docker)
DOCKER_COMPOSE=HTTP_PORT=$(HTTP_PORT) HTTPS_PORT=$(HTTPS_PORT) MYSQL_PORT=$(MYSQL_PORT) MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) $(shell which docker-compose)
DOCKER_COMPOSE_SERVICES=$(shell cat docker-compose.yml|awk '/^services/,/^network/' | grep -E '^\s{2}\S+' | sed 's/://g' | xargs)
GO=$(shell which go)
GO_IMAGE=headphonista:go
MYSQL_ROOT_PASSWORD=mysqlrootpassword
MYSQL_CONFIG=my.cnf
MAKE=$(shell which make)
MIGRATE=./mysql/bin/migrate
GO_CONTAINER_ID=$(shell docker ps | grep 'web' | awk '{print $$1}')

.PHONY: docker/build docker/run build

docker/build:
	$(DOCKER) build -t $(GO_IMAGE) .

docker/run:
	$(DOCKER) run --rm $(GO_IMAGE)

go/build:
	$(DOCKER) exec $(GO_CONTAINER_ID) make -C ./app build

docker/run/bash:
	$(DOCKER) run -it --rm $(GO_IMAGE) /bin/bash

server:
	$(MAKE) server/up mysql/setup

server/up:
	$(DOCKER_COMPOSE) up --build  --remove-orphans -d $(DOCKER_COMPOSE_SERVICES)

server/down:
	$(DOCKER_COMPOSE) down


$(MYSQL_CONFIG):
	echo "[client]" >> $@
	echo "user=root" >> $@
	echo "password=$(MYSQL_ROOT_PASSWORD)" >> $@
	echo "host=127.0.0.1" >> $@
	echo "port=$(MYSQL_PORT)" >> $@

mysql/wait:
	which mysqladmin
	@echo "waiting boot mysql..." &&  sleep 30


mysql/init: $(MYSQL_CONFIG) mysql/wait
	mysql --defaults-extra-file=$(MYSQL_CONFIG) -e "CREATE DATABASE IF NOT EXISTS headphonista"

$(MIGRATE):
	$(MAKE) -C ./mysql install

mysql/setup: $(MIGRATE) mysql/init
	$(MIGRATE) -path ./mysql/sql -database 'mysql://root:$(MYSQL_ROOT_PASSWORD)@tcp(127.0.0.1:$(MYSQL_PORT))/headphonista' up

