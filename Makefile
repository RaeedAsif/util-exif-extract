# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

# Main package parameters
PKG=./cmd/
MAIN=main.go
BINARY_NAME=bin/app

all: test build
run:
		$(GOBUILD) -o $(BINARY_NAME) $(PKG)
		./$(BINARY_NAME)

.PHONY: all build test clean run
