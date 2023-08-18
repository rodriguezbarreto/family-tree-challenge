
FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o seed ./cmd/seed/seed.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

COPY --from=builder /app/seed /app/seed

COPY --from=builder /app/.env /app/.env

EXPOSE 8080

CMD ["/app/main"]
