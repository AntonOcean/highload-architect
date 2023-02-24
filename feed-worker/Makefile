.PHONY: run
run:
	source .env && go run ./cmd/main

.PHONY: lint
lint:
	golangci-lint -v run --fix ./...

.PHONY: gen
gen:
	swag i -g swagger.go --parseInternal --parseDependency --dir ./internal

.PHONY: test
test:
	go test --race --vet= ./... -v

.PHONY: build
build:
	go build -tags=jsoniter -o ./tmp/main ./cmd/main
