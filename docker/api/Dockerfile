FROM golang:1.22.3-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update

COPY ../../ .

RUN go mod download

RUN set -eux && \
  apk add --no-cache git vim && \
  go install github.com/cosmtrek/air@latest

EXPOSE 8080
