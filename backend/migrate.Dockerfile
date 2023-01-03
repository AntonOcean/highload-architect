FROM golang:1.19-buster as builder

RUN apt-get update \
    && apt-get install -y --no-install-recommends git && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY scripts/migrate.sh migrate.sh
COPY migrations /app/migrations

ENTRYPOINT ["sh", "migrate.sh"]