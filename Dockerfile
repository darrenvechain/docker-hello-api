# Build stage
FROM golang:1.22-alpine as builder

WORKDIR /app
COPY go.mod main.go /app/

RUN go build -o api

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/api /app/api

ENTRYPOINT ["/app/api"]