.PHONY: build fmt lint
.DEFAULT_GOAL := help

COMMIT := $(shell git rev-parse HEAD)
DATE := $(shell TZ=UTC date +'%Y-%m-%d %H:%M:%S UTC')
BUILD_OUTPUT := cmd/waffle/waffle

build: ## builds waffle
	@go build -v -trimpath \
		-ldflags="-X 'github.com/waffle-lua/waffle.Date=$(DATE)' \
		-X 'github.com/waffle-lua/waffle.Version=$(COMMIT)'" \
		-o $(BUILD_OUTPUT) \
		github.com/waffle-lua/waffle/cmd/waffle

lint:
	@golangci-lint run && echo "OK"

fmt:
	@gofumpt -extra -l -w .

help: ## prints out the help documentation (also will be printed by simply running `make` command with no arg)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

