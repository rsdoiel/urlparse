

all: install

build:
	go build -o bin/urlparse cmds/urlparse/urlparse.go

install: 
	go install cmds/urlparse/urlparse.go

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

release:
	./mk-release.sh
