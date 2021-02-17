
.PHONY: github/build
github/build: ## Builds the app on GitHub
	mkdir bin/ && go build -o bin/tecli 

# GitHub actions doesn't allow internet connectivity
# Therefore integration tests that needs to download/pull data from the internet can't be executed
.PHONY: github/test
github/test: go/generate ## Test the app on GitHub
	@cd tests && go test cmd_root_test.go
