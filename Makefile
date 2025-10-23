.PHONY: all fmt lint test coverage benchmark deps release clean

all: fmt lint test benchmark ## Run all tests

fmt: ## Format the code
	golangci-lint fmt

lint: ## Lint the code
	golangci-lint run --timeout=5m

test: ## Run tests
	go test -race -shuffle=on -count=1 -covermode=atomic -coverpkg=./... -coverprofile=coverage.out ./...

coverage: test ## Generate coverage
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

benchmark: ## Run benchmarks
	go test -run=^$$ -bench=. -benchmem ./... | tee benchmark.txt

deps: ## Install dependencies
	go mod download

clean: ## Remove build artifacts
	rm -f coverage.* benchmark.txt
	rm -rf dist

help:  ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
