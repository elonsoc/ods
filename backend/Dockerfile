# syntax=docker/dockerfile:1.2
# dockerfile for ods-backend

# Build stage
FROM golang:1.20.1 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo  -o ods-backend .

# Run stage
FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/ods-backend ./ods-backend
ENTRYPOINT ["/app/ods-backend"]
