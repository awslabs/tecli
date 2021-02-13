
.PHONY: github/build
github/build: ## Builds the app on Github
	mkdir bin/ && go build -o bin/tecli 

# GitHub actions doesn't allow egress internet connectivity
# Therefore integration tests that needs to download/pull data from the internet can't be executed
.PHONY: github/test
github/test:
	@cd tests && go test pkg_test.go cmd_root_test.go cmd_init_test.go cmd_render_test.go cmd_version_test.go

.PHONY: go/test
go/test:
	@cd tests && go test -v