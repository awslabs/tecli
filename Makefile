# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

include lib/make/*/Makefile

.PHONY: tecli/test
tecli/test:
	@cd tests && go test -v

.PHONY: tecli/build
tecli/build: go/mod/tidy go/version go/get go/fmt go/generate go/build ## Builds the app

.PHONY: tecli/build/docker
tecli/build/docker: tecli/build
	GOOS=linux GOARCH=amd64 go build -o dist/tecli-linux-amd64 main.go
	docker build -t tecli:latest .

.PHONY: tecli/install
tecli/install: go/get go/fmt go/generate go/install ## Builds the app and install all dependencies

.PHONY: tecli/run
tecli/run: go/fmt ## Run a command
ifdef command
	make go/run command='$(command)'
else
	make go/run
endif

.PHONY: tecli/release
tecli/release: ## Compile to multiple architectures
	@mkdir -p dist
	@echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o dist/tecli-darwin-amd64 main.go
	GOOS=solaris GOARCH=amd64 go build -o dist/tecli-solaris-amd64 main.go

	GOOS=freebsd GOARCH=386 go build -o dist/tecli-freebsd-386 main.go
	GOOS=freebsd GOARCH=amd64 go build -o dist/tecli-freebsd-amd64 main.go
	GOOS=freebsd GOARCH=arm go build -o dist/tecli-freebsd-arm main.go

	GOOS=openbsd GOARCH=386 go build -o dist/tecli-openbsd-386 main.go
	GOOS=openbsd GOARCH=amd64 go build -o dist/tecli-openbsd-amd64 main.go
	GOOS=openbsd GOARCH=arm go build -o dist/tecli-openbsd-arm main.go

	GOOS=linux GOARCH=386 go build -o dist/tecli-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/tecli-linux-amd64 main.go
	GOOS=linux GOARCH=arm go build -o dist/tecli-linux-arm main.go

	GOOS=windows GOARCH=386 go build -o dist/tecli-windows-386 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/tecli-windows-amd64 main.go
	chmod +x dist/tecli-*

.PHONY: tecli/clean
tecli/clean: ## Removes unnecessary files and directories
	rm -rf downloads/
	rm -rf generated-*/
	rm -rf dist/
	rm -rf build/


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

.DEFAULT_GOAL := tecli/help

.PHONY: tecli/help
tecli/help: ## This HELP message
	@fgrep -h ": ##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
