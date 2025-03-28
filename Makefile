# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=minisapi

# Docker parameters
DOCKER_COMPOSE=docker-compose

.PHONY: all build test clean run dev setup

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

dev:
	$(DOCKER_COMPOSE) up --build

setup:
	cp .env.example .env
	go mod download
	go mod tidy

lint:
	golangci-lint run

swagger:
	swag init

# Docker commands
docker-build:
	docker build -t minisapi .

docker-run:
	docker run -p 8080:8080 minisapi 