
.PHONY: build up down all
# all - Build and start the development environment
all: build up
# build - Build the development environment
build: 
	docker compose -f ./infra/docker-compose.yml build
# up - Start the development environment
# will detach and run in the background
up:
	docker compose -f ./infra/docker-compose.yml up -d

# down - Stop the development environment
down: 
	docker compose -f ./infra/docker-compose.yml down