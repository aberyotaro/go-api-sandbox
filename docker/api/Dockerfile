FROM golang:1.24-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && \
    apk add --no-cache vim make git

RUN go install github.com/air-verse/air@latest && \
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# コンテナ内部で待ち受けるポート番号（ドキュメント的な意味合い）
EXPOSE 8080
