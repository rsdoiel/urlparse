#
# Simple Makefile
#
all: clean build install

build:
	go build -o bin/urlparse cmds/urlparse/urlparse.go

install: 
	env GOBIN=$(HOME)/bin go install cmds/urlparse/urlparse.go

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

save:
	./mk-website.bash
	git commit -am "Quick save"
	git push origin master

website:
	./mk-website.bash

release:
	./mk-release.bash

publish:
	./mk-website.bash
	./publish.bash
