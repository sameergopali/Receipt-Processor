# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test


# Main package name
MAIN=cmd/main.go

# Output binary name
BINARY_NAME=main

# Swagger parameters
SWAGGER_OUT=./docs
SWAGGER_DIR=./cmd,./internal

# Targets
all: test swagger build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN)

test:
	$(GOTEST) -v ./tests/...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GORUN)  $(MAIN)

swagger:
	swag init  -o $(SWAGGER_OUT) -d $(SWAGGER_DIR)

.PHONY: all build test clean run format swagger
