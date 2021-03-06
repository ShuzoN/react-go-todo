HTTP_PORT=80
HTTPS_PORT=443
MYSQL_PORT=3306
REALIZE_PORT=3000
UI_PORT=8080
DOCKER=$(shell which docker)
DOCKER_COMPOSE=HTTP_PORT=$(HTTP_PORT) HTTPS_PORT=$(HTTPS_PORT) MYSQL_PORT=$(MYSQL_PORT) REALIZE_PORT=$(REALIZE_PORT) UI_PORT=$(UI_PORT) MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) $(shell which docker-compose)
DOCKER_COMPOSE_SERVICES=$(shell cat docker-compose.yml|awk '/^services/,/^network/' | grep -E '^\s{2}\S+' | sed 's/://g' | xargs)
GO=$(shell which go)
MYSQL_ROOT_PASSWORD=mysqlrootpassword
MYSQL_CONFIG=my.cnf
MAKE=$(shell which make)
MIGRATE=./mysql/bin/migrate
GO_CONTAINER_ID=$(shell docker ps | grep 'web' | awk '{print $$1}')
GO_VERSION=go1.13.7
GO_PKG=https://dl.google.com/go/$(GO_VERSION).darwin-amd64.tar.gz

install/local/go:
	[ -f ./$(GO_VERSION).tar.gz ] || wget $(GO_PKG) -O ./$(GO_VERSION).tar.gz
	sudo tar -C /usr/local -zxvf ./$(GO_VERSION).tar.gz
	[ -f ./$(GO_VERSION).tar.gz ] && rm ./$(GO_VERSION).tar.gz

install/local: install/local/go
	@echo "export PATH=$PATH:/usr/local/go/bin >> ~/.bashrc"
	@echo "export GOPATH=~/go >> ~/.bashrc"

logs:
	$(DOCKER_COMPOSE) logs -f

server:
	$(MAKE) server/up mysql/setup

server/up:
	$(DOCKER_COMPOSE) up --build  --remove-orphans -d $(DOCKER_COMPOSE_SERVICES)

server/down:
	$(DOCKER_COMPOSE) down

web/rebuild:
	$(DOCKER_COMPOSE) rm --force --stop web
	$(DOCKER_COMPOSE) up --build -d web

ui/rebuild:
	$(DOCKER_COMPOSE) rm --force --stop ui
	$(DOCKER_COMPOSE) up --build -d ui


$(MYSQL_CONFIG):
	echo "[client]" >> $@
	echo "user=root" >> $@
	echo "password=$(MYSQL_ROOT_PASSWORD)" >> $@
	echo "host=127.0.0.1" >> $@
	echo "port=$(MYSQL_PORT)" >> $@


mysql/start:
	$(DOCKER_COMPOSE) start mysqld

mysql/stop:
	$(DOCKER_COMPOSE) stop mysqld

mysql/rebuild:
	$(DOCKER_COMPOSE) rm --force --stop mysqld
	$(DOCKER_COMPOSE) up --build -d mysqld

mysql/lifecheck: 
	until (mysqladmin ping &>/dev/null) do echo '.' && sleep 1; done

mysql/wait:
	which mysqladmin
	@echo "waiting boot mysql..."
	$(MAKE) mysql/lifecheck


mysql/init: $(MYSQL_CONFIG) mysql/wait
	mysql --defaults-extra-file=$(MYSQL_CONFIG) -e "CREATE DATABASE IF NOT EXISTS headphonista"

$(MIGRATE):
	$(MAKE) -C ./mysql install

mysql/setup: $(MIGRATE) mysql/init
	$(MAKE) mysql/migrate/up


mysql/migrate/up:
	$(MIGRATE) -path ./mysql/sql -database 'mysql://root:$(MYSQL_ROOT_PASSWORD)@tcp(127.0.0.1:$(MYSQL_PORT))/headphonista' up


