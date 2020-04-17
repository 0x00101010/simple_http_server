GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=simple_http_server

all: test build
build:
	echo "Starting to build the whole project..."
	if [ ! -d "bin" ]; then mkdir bin; fi
	$(GOBUILD) -v -o bin ./...
test:
	echo "Start testing the whole project..."
	$(GOTEST) -v ./...
run: build
	bin/$(BINARY_NAME) scu_websites 8888