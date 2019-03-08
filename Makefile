.PHONY: client server

default: all

version=$(shell ver=$$(git log -n 1 --pretty=oneline --format=%D | awk -F, '{print $$1}' | awk '{print $$3}'); \
	if [ "$$ver" = "master" ] ; then \
	ver="master($$(git log -n 1 --pretty=oneline --format=%h))" ; \
	fi ; \
	echo $$ver)

client: 
	mkdir -p build
	go build -ldflags "-X main.version=${version}" ./cmd/mq-client 
	mv mq-client* ./build

server: 
	mkdir -p build
	go build -ldflags "-X main.version=${version}" ./cmd/mq-server
	mv mq-server* ./build

install:
	mv build/mq-* /usr/local/bin

all: client server

clean:
	rm -rf ./build/mq-*
