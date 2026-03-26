IMAGE ?= auth-svc
TAG ?= local

.PHONY: all local-build local-run db-build db-up db-down start gen-proto docs

all: db-build db-up local-build local-start

start: build run

build: gen-proto
	@echo "Building auth svc..."
	@cd $(CURDIR) && go build -o .bin/main cmd/main.go

run:
	@echo "Starting auth svc..."
	@cd $(CURDIR) && go run cmd/main.go configs/config.yaml

local: gen-proto local-build local-run

db-build:
	@echo "DB build"
	@cd $(CURDIR) && docker compose build --no-cache

db-up:
	@echo "DB up"
	@cd $(CURDIR) && docker compose up db -d

local-build: gen-proto
	@echo "Building auth svc..."
	@cd $(CURDIR) && go build -o .bin/main cmd/main.go

local-run:
	@echo "Starting auth svc..."
	@cd $(CURDIR) && go run cmd/main.go configs/local.yaml

	
local-stop:
	@echo "Stopping auth svc..."
	@-lsof -ti:8081 | xargs kill -9 2>/dev/null
	@echo "Port 8081 is free"

db-down:
	@echo "DB down"
	@cd $(CURDIR) && docker compose down -v

gen-proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	  proto/auth.proto

docs:
	@echo "Generating swagger docs..."
	@cd $(CURDIR) && swag init -g cmd/main.go -o docs --parseInternal --parseDependency


.PHONY: docker-build
docker-build:
	@echo "Building docker image..."
	@docker build -t $(IMAGE):$(TAG) .
