
.PHONY: github/build
github/build: ## Builds the app on GitHub
	mkdir bin/ && go build -o bin/tecli 

# GitHub actions doesn't allow internet connectivity
# Therefore integration tests that needs to download/pull data from the internet can't be executed
# Need to include helper.go as first argument as it implements helpful functions to test Cobra commands
.PHONY: github/test
github/test: go/generate ## Test the app on GitHub
	@cd tests && go test helper.go cmd_root_test.go
