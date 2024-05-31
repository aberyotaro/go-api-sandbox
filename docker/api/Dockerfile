FROM golang:1.22.3-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY ../../ .

RUN go mod tidy

RUN go build -o app ./cmd/

EXPOSE 8080

CMD ["./app"]
