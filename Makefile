.PHONY: all local-build local-run db-build db-up db-down start

all: db-build db-up local-build local-start

start: build run

build:
	@echo "Building auth svc..."
	@cd $(CURDIR) && go build -o .bin/main cmd/main.go

run:
	@echo "Starting auth svc..."
	@cd $(CURDIR) && go run cmd/main.go configs/config.yaml


local: local-build local-run

db-build:
	@echo "DB build"
	@cd $(CURDIR) && docker compose build --no-cache

db-up:
	@echo "DB up"
	@cd $(CURDIR) && docker compose up db -d

local-build:
	@echo "Building auth svc..."
	@cd $(CURDIR) && go build -o .bin/main cmd/main.go

local-run:
	@echo "Starting auth svc..."
	@cd $(CURDIR) && go run cmd/main.go configs/local.yaml

db-down:
	@echo "DB down"
	@cd $(CURDIR) && docker compose down -v