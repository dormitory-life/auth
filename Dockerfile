FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/auth-service .

COPY --from=builder /app/migrations ./migrations

COPY --from=builder /app/configs ./configs

EXPOSE 8080

CMD ["./auth-service", "/app/configs/config.yaml"]