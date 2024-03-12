# syntax=docker/dockerfile:1.2
# dockerfile for ods-auth

# Build stage
FROM golang:1.20.1 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo  -o ods-auth .

# Run stage
FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/ods-auth ./ods-auth
ENTRYPOINT ["/app/ods-auth"]
