# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

include lib/make/*/Makefile

.PHONY: tecli/test
tecli/test:
	@cd tests && go test -v

.PHONY: tecli/build
tecli/build: go/mod/tidy go/version go/get go/fmt go/generate go/build tecli/update-readme ## Builds the app

.PHONY: tecli/install
tecli/install: go/get go/fmt go/generate go/install ## Builds the app and install all dependencies

.PHONY: tecli/run
tecli/run: go/fmt ## Run a command
ifdef command
	make go/run command='$(command)'
else
	make go/run
endif

.PHONY: tecli/clean
tecli/clean: ## Removes unnecessary files and directories
	rm -rf downloads/
	rm -rf generated-*/
	rm -rf dist/
	rm -rf build/
	rm -f box/blob.go
	rm -f clencli/log.json

.PHONY: tecli/terminalizer
tecli/terminalizer:
ifdef command
	terminalizer record terminalizer-$(command) --config clencli/terminalizer.yml --skip-sharing
	terminalizer render terminalizer-$(command) --output clencli/terminalizer/$(command).gif
else
	@echo 'Need to pass "command" parameter'
endif	

.PHONY: tecli/update-readme
tecli/update-readme: ## Renders template readme.tmpl with additional documents
	@echo "Generate COMMANDS.md"
	@echo "## Commands" > COMMANDS.md
	@echo '```' >> COMMANDS.md
	@build/tecli --help >> COMMANDS.md
	@echo '```' >> COMMANDS.md
	@echo "COMMANDS.md generated successfully"
	@clencli render template --name readme

.PHONY: tecli/test
tecli/test: go/test

.DEFAULT_GOAL := help

.PHONY: help
help: ## This HELP message
	@fgrep -h ": ##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
