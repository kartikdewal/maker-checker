DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= maker_checker
DB_PORT ?= 5432
DATABASE = "postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable"
.PHONY: build
build: go-build

.PHONY: format
format:
	@echo "Formatting with gofmt"
	gofmt -w .

.PHONY: clean
clean: go-clean ## Clean build cache and dependencies

.PHONY: services-up
services-up:
	@echo "Starting all docker services..."
	docker-compose -f docker-compose.yaml up --force-recreate --detach

.PHONY: services-down
services-down:
	@echo "Stopping all services..."
	docker-compose -f docker-compose.yaml down

db-setup:
	@echo "Creating databases"
	docker compose exec db psql -U postgres -c 'CREATE DATABASE maker_checker'
	docker compose exec db psql -U postgres -c 'CREATE DATABASE maker_checker_test'

migrate-tool-download:
	@echo "downloading migrate tool"
	@curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.darwin-arm64.tar.gz | tar xvz -C ./tools/mac

migrate-up:
	@echo "Running migrations up..."
	./tools/mac/migrate -database ${DATABASE} -path ./store/psql/migrations -verbose up

migrate-down:
	@echo "Running migrations down..."
	./tools/mac/migrate -database ${DATABASE} -path ./store/psql/migrations -verbose down $(COUNT)

migration-create:
	@echo "Creating migration files"
	./tools/mac/migrate create --ext sql --dir=./store/psql/migrations $(NAME)

go-build:
	@echo "Building Go services..."
	@rm -rf build
	@mkdir build
	go build -o build -v ./...
	@echo "Go services available at ./build"

go-clean: go-clean-cache go-clean-deps

go-clean-cache:
	@echo "Cleaning build cache..."
	go clean -cache

go-clean-test-cache:
	@echo "Cleaning test cache..."
	go clean -testcache

go-clean-deps:
	@echo "Cleaning dependencies..."
	go mod tidy

go-deps:
	@echo "Installing dependencies..."
	go mod download

go-gen-mocks:
	@echo "Generating mocks..."
	# requires mockery tool https://github.com/vektra/mockery
	mockery --all --keeptree --dir=pkg --output=pkg/mocks

go-unit-test:
	@echo "Running unit tests..."
	gotestsum -- -v -tags="unit" ./...

http-start-dev:
	@echo "Starting HTTP server"
	PROFILE=dev go run cmd/maker-checker/main.go --profile=dev

http-start:
	@echo "Starting HTTP server"
	./build/maker-checker