# syntax=docker/dockerfile:1.2
# dockerfile for ods-api

# Build stage
FROM golang:1.20.1 AS builder
WORKDIR /app
# this gets the full workspace since we need it.
COPY .. . 
CMD ["mv", "api"]
RUN CGO_ENABLED=0 go build -a -installsuffix cgo  -o ods-api .

# Run stage
FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/ods-api ./ods-api
ENTRYPOINT ["/app/ods-api"]
