.PHONY: all test clean build

all: test build

run:
	./takehome

build:
	go build -o ./takehome ./cmd

clean:
	rm -f ./takehome store.db

test:
	go test ./...
