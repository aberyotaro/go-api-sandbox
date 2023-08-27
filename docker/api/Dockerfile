FROM golang:1.20.0-alpine

WORKDIR /app

COPY ../../ .

RUN go mod download

RUN go build -o app ./cmd/

EXPOSE 8080

CMD ["./app"]
