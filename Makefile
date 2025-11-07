BINARY=server
BIN_DIR=bin

.PHONY: build run tidy

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) ./cmd/server

run:
	go run ./cmd/server

tidy:
	go mod tidy
