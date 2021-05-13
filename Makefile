APP_NAME=twitter-clone
VERSION?=v0.1.0
BUILD=$(shell git rev-parse HEAD)

GO=go
GOOSS=darwin linux windows freebsd netbsd openbsd dragonfly
GOARCHS=386 arm arm64 amd64
LDFLAGS=-ldflags="-s -w -X"
BINARY_DIR=.bin

.DEFAULT_GOAL := help

.PHONY: help
# Source: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Displays all the available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: fmt
fmt: ## Format go files
	@$(GO) fmt ./...

.PHONY: vet
vet: ## go vet
	@$(GO) vet ./...

.PHONY: test
test: ## Runs unit tests [cmd: go test -v -bench . -benchmem ./...]
	@$(GO) test -v -bench . -benchmem ./...

.PHONY: clean
clean: ## Deletes all compiled / executable files
	@find .bin -type f -name '*' -print0 | xargs -0 rm --
	@echo ">> Deleted all build files!"

.PHONY: install
install: ## Installs the package
	@$(GO) install ${LDFLAGS} ./...

.PHONY: install-deps
install-deps: ## Install dependencies
	@$(GO) mod download

.PHONY: run
run: ## Runs the backend server
	@$(GO) run cmd/rest/main.go

.PHONY: dev
dev: ## Runs the backend server with hot-reload (Must have air installed https://github.com/cosmtrek/air)
	@air -c configs/.air.toml

.PHONY: build-rest
build-rest: ## Compiles the rest api server
	@$(GO) build $(LDFLAGS) -v -o $(BINARY_DIR)/$(APP_NAME)-$(VERSION)_server cmd/rest/main.go

.PHONY: build-rest-all
build-rest-all: ## Cross-compiles the rest api server
	@$(foreach GOOS, $(GOOSS),\
		$(foreach GOARCH, $(GOARCHS),\
			$(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH);\
			$(GO) build $(LDFLAGS) -v -o $(BINARY_DIR)/$(APP_NAME)-$(VERSION)_server-$(GOOS)-$(GOARCH) cmd/rest/main.go)))
