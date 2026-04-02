.PHONY: build test lint clean

build:
	go build -o bin/cli-template .

test:
	go test -v -cover ./...

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/
