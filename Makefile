# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Modernized Makefile (Go 1.25). The previous version included files from a
# `habits/lib/make/` tree that is no longer present in the repository, which
# made every target fail. This rewrite keeps the same target names where they
# are referenced by docs / CI but implements them inline.

SHELL := /bin/bash

WORKSPACE := $(shell pwd)
DIST_DIR  := $(WORKSPACE)/dist
BINARY    := tecli
MAIN      := main.go

GO        ?= go
GOFLAGS   ?=

# Cross-compile target list: GOOS/GOARCH pairs used by `tecli/compile`.
PLATFORMS := \
	darwin/amd64 \
	solaris/amd64 \
	freebsd/386 freebsd/amd64 freebsd/arm \
	openbsd/386 openbsd/amd64 openbsd/arm \
	linux/386 linux/amd64 linux/arm \
	windows/386 windows/amd64

.DEFAULT_GOAL := help

# ---------------------------------------------------------------------------
# Help
# ---------------------------------------------------------------------------

.PHONY: help
help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "; printf "Usage: make <target>\n\nTargets:\n"} \
	/^[a-zA-Z0-9_\/\-]+:.*?## / {printf "  \033[36m%-28s\033[0m %s\n", $$1, $$2}' \
	$(MAKEFILE_LIST)

# ---------------------------------------------------------------------------
# Build / install
# ---------------------------------------------------------------------------

.PHONY: tecli/build
tecli/build: tecli/clean go/tidy go/fmt go/vet go/generate go/install ## Build the CLI for the host platform and install it.

.PHONY: tecli/install
tecli/install: go/fmt go/generate go/install ## Build and install the CLI on the host platform.

.PHONY: tecli/run
tecli/run: go/fmt ## Run the CLI. Pass `command='args...'` to forward args.
	@$(GO) run $(MAIN) $(command)

.PHONY: tecli/compile
tecli/compile: ## Cross-compile to every supported OS/arch into ./dist.
	@mkdir -p $(DIST_DIR)
	@echo "Compiling for every OS and Platform"
	@for p in $(PLATFORMS); do \
		os=$${p%/*}; arch=$${p#*/}; \
		out=$(DIST_DIR)/$(BINARY)-$$os-$$arch; \
		if [ "$$os" = "windows" ]; then out=$$out.exe; fi; \
		echo "  -> GOOS=$$os GOARCH=$$arch -> $$out"; \
		GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 $(GO) build $(GOFLAGS) -o $$out $(MAIN) || exit 1; \
	done

# ---------------------------------------------------------------------------
# Test
# ---------------------------------------------------------------------------

.PHONY: tecli/test
tecli/test: go/generate ## Run all Go tests.
	@$(GO) test ./...

.PHONY: tecli/test/configure
tecli/test/configure: ## Run the `configure` command integration tests.
	@cd tests/commands && $(GO) test -run ConfigureCmdFlags
	@cd tests/commands && $(GO) test -run ConfigureCreate
	@cd tests/commands && $(GO) test -run ConfigureList
	@cd tests/commands && $(GO) test -run ConfigureRead
	@cd tests/commands && $(GO) test -run ConfigureUpdate
	@cd tests/commands && $(GO) test -run ConfigureDelete

# ---------------------------------------------------------------------------
# Clean
# ---------------------------------------------------------------------------

.PHONY: tecli/clean
tecli/clean: ## Remove build artifacts.
	@rm -rf downloads/ generated-*/ dist/ build/
	@rm -f box/blob.go clencli/log.json

.PHONY: tecli/clean/all
tecli/clean/all: tecli/clean ## Remove build artifacts AND user config dirs.
	@rm -rf .tecli ~/.tecli

# ---------------------------------------------------------------------------
# Go helpers
# ---------------------------------------------------------------------------

.PHONY: go/tidy
go/tidy: ## Tidy go.mod / go.sum.
	@$(GO) mod tidy

.PHONY: go/fmt
go/fmt: ## Run gofmt.
	@$(GO) fmt ./...

.PHONY: go/vet
go/vet: ## Run go vet.
	@$(GO) vet ./...

.PHONY: go/generate
go/generate: ## Run go generate.
	@$(GO) generate ./...

.PHONY: go/install
go/install: ## go install the module to $GOBIN.
	@$(GO) install ./...

.PHONY: go/build
go/build: ## go build all packages (sanity check).
	@$(GO) build ./...
