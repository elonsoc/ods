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
	SAML_CERT_PATH=../infra/saml/cert.pem \
	SAML_KEY_PATH=../infra/saml/key.pem \
	IDP_URL="http://localhost:8088/idp.xml" \
	SP_URL="http://localhost:3000" \
	WEB_URL="http://localhost:3001" \
	go run ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
