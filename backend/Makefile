# Makefiles are to make our lives easier!
.PHONY: init
init:
	go get .

.PHONY: build
build:
	go build -o bin/backend .

.PHONY: run
run:
	PORT=3000 \
	DATABASE_URL=postgres://postgres:postgrespw@localhost:5432/postgres \
	LOGGING_URL=http://localhost:3100/loki/api/v1/push \
	PROMETHEUS_URL=http://localhost:9090 \
	REDIS_URL=redis://localhost:6379 \
	STATSD_URL=localhost:9125 \
	go run ./...
