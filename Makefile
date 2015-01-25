
build: urlparse

all: install

urlparse: urlparse.go
	go build urlparse.go

install: urlparse.go
	go install urlparse.go

clean:
	rm urlparse

