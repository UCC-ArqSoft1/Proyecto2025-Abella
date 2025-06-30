FROM golang:latest

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o backend

EXPOSE 8523

CMD ["./backend"]