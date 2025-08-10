# Variables
BINARY_NAME := bloggo
ZIP_FILE := $(BINARY_NAME)-$(shell date -I).tar.gz

all: build package

build:
	go build -o ./$(BINARY_NAME) cmd/blog/main.go


package: build
	mkdir build
	mv ./$(BINARY_NAME) ./build
	cp -r ./web ./build
	cp -r ./scripts ./build
	tar -czf "$(ZIP_FILE)" ./build

clean:
	rm -f *.tar.gz
	rm -rf ./build

