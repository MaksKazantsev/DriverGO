COMPOSE_FILE_POSTGRES ?= $(CURDIR)/containers/docker-compose.postgres.yaml
COMPOSE_FILE_PROMETHEUS ?= $(CURDIR)/containers/docker-compose.prometheus.yaml
CONFIG_FILE ?= config/config.local.yaml
TESTS_SERVICE ?= internal/tests/service
TESTS_REPO ?= internal/tests/repository

# Target to start the PostgreSQL container using Docker Compose
postgres:
	@echo "Starting Docker Compose for Postgres with file located at: $(COMPOSE_FILE)"
	docker-compose -f $(COMPOSE_FILE_POSTGRES) up

prometheus:
	@echo "Starting Docker Compose for Prometheus with file located at: $(COMPOSE_FILE)"
	docker-compose -f $(COMPOSE_FILE_PROMETHEUS) up

# Target to run the local Go application
local:
	@echo "Starting the Go application with config: $(CONFIG_FILE)"
	go run main.go -c $(CONFIG_FILE)

build:
	@echo "Building the Go application and initializing Swagger documentation"
	go build -o bin/main main.go
	swag init

# Target to stop and remove all containers, networks, and volumes
clean:
	@echo "Stopping and removing all containers, networks, and volumes defined in: $(COMPOSE_FILE)"
	docker-compose -f $(COMPOSE_FILE) down -v

service-test:
	@echo "Running tests in the service directory: $(TESTS_SERVICE)"
	cd $(TESTS_SERVICE) && go test ./...

repo-test:
	@echo "Running tests in the repository directory: $(TESTS_REPO)"
	cd $(TESTS_REPO) && go test ./...

test:
	@echo "Starting linter test..."
	golangci-lint run


.PHONY: postgres local build clean service-test repo-test
