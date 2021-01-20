.PHONY: go/build
go/build: ## Compile packages and dependencies
	mkdir -p build
	go build -o build/tfe-cli

.PHONY: go/install
go/install: ## Compile and install packages and dependencies
	go install

.PHONY: go/run
go/run: ## Compile and run Go program
ifdef command
	go run main.go $(command)
else
	go run main.go
endif

.PHONY: go/fmt
go/fmt: ## gofmt (reformat) package sources
	go fmt github.com/awslabs/tfe-cli/...
	go fmt ./...
	gofmt -s -w .

.PHONY: go/generate
go/generate: ## Generate Go files by processing source
	go generate ./...

.PHONY: go/mod/tidy
go/mod/tidy: ## Add missing and remove unused modules
	go mod tidy

.PHONY: go/get
go/get: ## Add dependencies to current module and install them
	go get ./...

.PHONY: go/clean
go/clean: ## Remove object files and cached files
	go clean -cache -modcache -i -r

.PHONY: go/list
go/list: ## List packages or modules
	go list -m -versions $(package)

.PHONY: go/version
go/version: ## Print Go version
	go version

.PHONY: go/env
go/env: ## Print Go environment information
	go env

# GitHub actions doesn't allow egress internet connectivity
# Therefore integration tests that needs to download/pull data from the internet can't be executed
.PHONY: go/test
go/test:
	@cd tests && go test pkg_test.go cmd_root_test.go cmd_init_test.go cmd_render_test.go cmd_version_test.go cmd_gitignore_test.go
