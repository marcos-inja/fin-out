.PHONY: build test lint validate all

BINARY=fin
CMD=./cmd/fin

build:
	go build -o $(BINARY) $(CMD)

test:
	go test ./... -count=1 -cover

lint:
	go vet ./...

validate: build
	./$(BINARY) validate

all: lint build validate
