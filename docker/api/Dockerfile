FROM golang:1.22.3-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update

RUN apk add --no-cache -t .build-dependencies \
    mysql-client \
    mysql-dev

RUN set -eux && \
  apk add --no-cache git vim && \
  go install github.com/cosmtrek/air@latest && \
  go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ../../ .

RUN go mod download

EXPOSE 8080
