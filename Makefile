SRC_PATHS := ./cmd ./internal ./pkg

.PHONY: build
build: ## Build application binaries.
	go build -race ./cmd/...

.PHONY: tidy
tidy: ## Clean up deps.
	go mod tidy

.PHONY: image
image: ## Build Docker image.
	docker build --platform linux/amd64 -t navigationsvc .

.PHONY: install
install: ## Install application.
	go install -race ./cmd/navigationsvc

.PHONY: lint
lint: ## Run lint tool.
	golangci-lint run

.PHONY: test
test: ## Run unit tests.
	go test -v -race $(foreach p, $(SRC_PATHS), $(p)/...)

.PHONY: run
run: build
	go run ./cmd/service/main.go

.PHONY: ci
ci: tidy lint test
