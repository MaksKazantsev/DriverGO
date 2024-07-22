COMPOSE_FILE ?= $(CURDIR)/containers/docker-compose.yml
CONFIG_FILE ?= config/config.local.yaml

# Target to start the PostgreSQL container using Docker Compose
postgres:
	@echo "Starting Docker Compose with file located at: $(COMPOSE_FILE)"
	docker-compose -f $(COMPOSE_FILE) up -d

# Target to run the local Go application
start-local:
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

.PHONY: postgres start-local clean
