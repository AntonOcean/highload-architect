.PHONY: run
run:
	source .env && go run ./cmd/main

.PHONY: lint
lint:
	golangci-lint -v run --fix ./...

# make makemigration NAME=some_name
.PHONY: makemigration
makemigration:
	migrate create -ext sql -dir migrations $(NAME)

.PHONY: gen
gen:
	swag i -g swagger.go --parseInternal --parseDependency --dir ./internal

.PHONY: mock
mock:
	mockery --all --output internal/mocks

.PHONY: test
test:
	go test --race --vet= ./... -v

.PHONY: build
build:
	go build -tags=jsoniter -o ./tmp/main ./cmd/main
