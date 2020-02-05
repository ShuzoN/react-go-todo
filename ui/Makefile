YARN=$(shell which yarn)
NODE_MODULES=./node_modules


.PHONY: install build run


install: $(NODE_MODULES)
	[ -d $(NODE_MODULES) ] || $(YARN) install

build:
	$(YARN) build

run:
	$(YARN) start


