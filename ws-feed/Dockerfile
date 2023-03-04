FROM golang:1.19-buster as builder

RUN apt-get update \
    && apt-get install -y --no-install-recommends git

ENV GO111MODULE "on"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o main ./cmd/main

FROM debian:buster-slim

WORKDIR /app

COPY --from=builder /app/main /app/

ENTRYPOINT ["/app/main"]
