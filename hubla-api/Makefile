.PHONY: build test clean up down migrate create-db drop-db run

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=myapp

DOCKER=docker
DOCKER_COMPOSE=docker-compose
COMPOSE_UP=$(DOCKER_COMPOSE) up
COMPOSE_DOWN=$(DOCKER_COMPOSE) down
DOCKER_BUILD=$(DOCKER) buildx build --platform linux/amd64,linux/arm64 -t $(BINARY_NAME) .

DB_HOST=localhost
DB_PORT=5432
DB_NAME=hubla
MIGRATIONS_PATH=database/migrations

build:
	$(GOBUILD) -o $(BINARY_NAME) -v
	$(DOCKER_BUILD)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

up:
	$(COMPOSE_UP)

down:
	$(COMPOSE_DOWN)

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database postgres://$(DB_HOST):$(DB_PORT)/$(DB_NAME) up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database postgres://$(DB_HOST):$(DB_PORT)/$(DB_NAME) down

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) $(filename)
