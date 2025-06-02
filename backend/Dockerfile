# syntax=docker/dockerfile:1
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./s

RUN go mod download

COPY . .


RUN go build -o /main

WORKDIR /app

COPY --from=builder /main /main

CMD ["/main"]