
FROM golang:1.20 as builder

WORKDIR /app

COPY . .

COPY .env .env

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/.env /app/.env

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]
