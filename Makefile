HELL=/bin/bash

ifeq ($(OS), Windows_NT)
OS_NAME="Windows"
else
UNAME=$(shell uname)
ifeq ($(UNAME),Linux)
OS_NAME="Linux"
else
ifeq ($(UNAME),Darwin)
OS_NAME="MacOS"
else
OS_NAME="Other"
endif
endif
endif

DB_DSN = "mysql://fishing_api_server:password@tcp(mysql:3306)/fishing_api_server"

MIGRATIONS_PATH = /migrations

build:
	docker compose build

install:
	cp .env.example .env
	make build
	make up

up:
	USER_NAME=$(shell id -nu) USER_ID=$(shell id -u) GROUP_NAME=$(shell id -ng) GROUP_ID=$(shell id -g) OS_NAME=$(OS_NAME) docker compose up

stop:
	docker compose stop

down:
	docker compose down

ps:
	docker compose ps

.PHONY: migrate-up migrate-down migrate-create
migrate-up:
	docker-compose run --rm migrate -path=$(MIGRATIONS_PATH) -database $(DB_DSN) up

migrate-down:
	docker-compose run --rm migrate -path=$(MIGRATIONS_PATH) -database $(DB_DSN) down

migrate-create:
	@echo "Usage: make migrate-create NAME=<migration_name> TYPE=<go|sql>"
	@echo "Example: make migrate-create NAME=create_users_table TYPE=sql"

.PHONY: redis
redis:
	docker exec -it fishing-api-server-redis /bin/bash -c "redis-cli"

ifeq ($(OS_NAME), "Linux")
shell:
	docker compose exec app su -s /bin/bash ${shell id -un}
else
shell:
	docker compose exec app bash
endif