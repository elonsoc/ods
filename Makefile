
.PHONY: build up down all
# up - Start the development environment
# will detach and run in the background
build: 
	docker compose -f ./infra/docker-compose.yml build
up:
	docker compose -d -f ./infra/docker-compose.yml up

# down - Stop the development environment
down: 
	docker compose -f ./infra/docker-compose.yml down

all: up