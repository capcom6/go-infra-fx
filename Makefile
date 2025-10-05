.PHONY: all lint test deps coverage clean help

all: deps lint test coverage ## Runs all tasks

lint: ## Runs golangci-lint
	golangci-lint run --timeout=5m

test: ## Runs tests
	go test -race -coverprofile=coverage.out -covermode=atomic ./...

deps: ## Download dependencies
	go mod download

coverage: test ## Generate coverage
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

clean: ## Remove generated files
	rm -f coverage.out coverage.html

help:  ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
