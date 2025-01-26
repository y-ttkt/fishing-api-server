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

redis:
	docker exec -it benesse-mcm-api-server-redis redis-cli

ifeq ($(OS_NAME), "Linux")
shell:
	docker compose exec app su -s /bin/bash ${shell id -un}
else
shell:
	docker compose exec app bash
endif