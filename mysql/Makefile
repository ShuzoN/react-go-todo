MAKE=$(shell which make)
DIRS=tmp bin
VERSION=v4.8.0
OS_TYPE=$(shell echo $(shell uname) | tr A-Z a-z)
OS_ARCH=amd64
out=./bin/migrate

.PHONY: install

install: $(DIRS)
	curl -L https://github.com/golang-migrate/migrate/releases/download/$(VERSION)/migrate.$(OS_TYPE)-$(OS_ARCH).tar.gz > $</migrate.$(OS_TYPE)-$(OS_ARCH).tar.gz
	cd $< && tar -zxvf migrate.$(OS_TYPE)-$(OS_ARCH).tar.gz
	mv $</migrate.$(OS_TYPE)-$(OS_ARCH) $(out)


clean:
	rm -rf $(DIRS)
	
$(DIRS):
		mkdir -p $@
