FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o blockchain-client .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/blockchain-client .

EXPOSE 8080

CMD ["./blockchain-client"]