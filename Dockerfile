
FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o log-aggregator-svc ./cmd/main.go

CMD ["./log-aggregator-svc"]