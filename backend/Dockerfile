# syntax=docker/dockerfile:1.2
# dockerfile for launchpad-backend

# Build stage
FROM golang:1.20.1 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo  -o launchpad-backend .

# Run stage
FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/launchpad-backend ./launchpad-backend
ENTRYPOINT ["/app/launchpad-backend"]
