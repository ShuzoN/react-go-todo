GO=$(shell which go)
PWD=$(shell pwd)
DIR=bin

.PHONY: build fmt

$(DIR): clear
	mkdir -p $(PWD)/$@

build: $(DIR) fmt
	$(GO) $@ -o $(PWD)/$(DIR)/main $(PWD)/main.go

fmt:
	$(GO) fmt $(PWD)/main.go

clear:
	rm -rf $(PWD)/$(DIR)
