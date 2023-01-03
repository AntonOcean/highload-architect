#!/bin/bash

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/vektra/mockery/v2@latest