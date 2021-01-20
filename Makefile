# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

include lib/make/*/Makefile

.PHONY: tfe-cli/test
tfe-cli/test:
	@cd tests && go test -v

.PHONY: tfe-cli/build
tfe-cli/build: go/mod/tidy go/version go/get go/fmt go/generate go/build ## Builds the app

.PHONY: tfe-cli/install
tfe-cli/install: go/get go/fmt go/generate go/install ## Builds the app and install all dependencies

.PHONY: tfe-cli/run
tfe-cli/run: go/fmt ## Run a command
ifdef command
	make go/run command='$(command)'
else
	make go/run
endif

.PHONY: tfe-cli/compile
tfe-cli/compile: ## Compile to multiple architectures
	@mkdir -p dist
	@echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o dist/tfe-cli-darwin-amd64 main.go
	GOOS=solaris GOARCH=amd64 go build -o dist/tfe-cli-solaris-amd64 main.go

	GOOS=freebsd GOARCH=386 go build -o dist/tfe-cli-freebsd-386 main.go
	GOOS=freebsd GOARCH=amd64 go build -o dist/tfe-cli-freebsd-amd64 main.go
	GOOS=freebsd GOARCH=arm go build -o dist/tfe-cli-freebsd-arm main.go

	GOOS=openbsd GOARCH=386 go build -o dist/tfe-cli-openbsd-386 main.go
	GOOS=openbsd GOARCH=amd64 go build -o dist/tfe-cli-openbsd-amd64 main.go
	GOOS=openbsd GOARCH=arm go build -o dist/tfe-cli-openbsd-arm main.go

	GOOS=linux GOARCH=386 go build -o dist/tfe-cli-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/tfe-cli-linux-amd64 main.go
	GOOS=linux GOARCH=arm go build -o dist/tfe-cli-linux-arm main.go

	GOOS=windows GOARCH=386 go build -o dist/tfe-cli-windows-386 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/tfe-cli-windows-amd64 main.go

.PHONY: tfe-cli/tag
tfe-cli/tag: ## Tag a version
ifdef version
	git tag -a v$(version) -m 'Release version v$(version)'
else
	@echo "version not specified"
endif

.PHONY: tfe-cli/clean
tfe-cli/clean: ## Removes unnecessary files and directories
	rm -rf downloads/
	rm -rf generated-*/
	rm -rf dist/
	rm -rf build/

.PHONY: tfe-cli/update-readme
tfe-cli/update-readme: ## Renders template readme.tmpl with additional documents
	@echo "Updating README.tmpl to the latest version"
	@cp box/resources/init/tfe-cli/readme.tmpl tfe-cli/readme.tmpl
	@echo "Generate COMMANDS.md"
	@echo "## Commands" > COMMANDS.md
	@echo '```' >> COMMANDS.md
	@tfe-cli --help >> COMMANDS.md
	@echo '```' >> COMMANDS.md
	@echo "COMMANDS.md generated successfully"
	@tfe-cli render template --name readme

.PHONY: tfe-cli/test
tfe-cli/test: go/test

.DEFAULT_GOAL := tfe-cli/help

.PHONY: tfe-cli/help
tfe-cli/help: ## This HELP message
	@fgrep -h ": ##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

split = $(word $2,$(subst $3, ,$1))
word-slash = $(word $2,$(subst /, ,$1))
word-dot = $(word $2,$(subst ., ,$1))

CURRENT_BRANCH := $(shell git branch --show-current)
CURRENT_COMMIT := $(shell git rev-parse --short HEAD)
LATEST_TAG := $(shell git describe --tags --abbrev=0)
# LATEST_TAG := $(shell git describe --tags `git rev-list --tags --max-count=1`)  # gets tags across all branches, not just the current branch
LATEST_CANDIDATE_TAG := $(shell git describe --tags --abbrev=0 --match "*-rc.*")

RELEASE_VERSION=v$(call word-slash,$(CURRENT_BRANCH),2)
CANDIDATE_VERSION=$(LATEST_TAG)-rc


.PHONY: tfe-cli/release
tfe-cli/release: go/mod/tidy
	@echo CURRENT BRANCH IS: $(CURRENT_BRANCH)
	@echo CURRENT COMMIT IS: $(CURRENT_COMMIT)
	@echo LATEST TAG IS: $(LATEST_TAG)
	@echo LATEST_CANDIDATE_TAG IS : $(LATEST_CANDIDATE_TAG)
ifneq (,$(findstring release,$(CURRENT_BRANCH)))
	@echo RELEASE FINAL VERSION
	git tag $(RELEASE_VERSION)
else ifneq (,$(findstring develop,$(CURRENT_BRANCH)))
	@echo RELEASE CANDIDATE VERSION
ifeq ($(strip $(LATEST_CANDIDATE_TAG)),) # not found
	git tag $(CANDIDATE_VERSION).1
else
	$(eval major=$(call word-dot,$(LATEST_CANDIDATE_TAG),1))
	$(eval minor=$(call word-dot,$(LATEST_CANDIDATE_TAG),2))
	$(eval patch=$(call word-dot,$(LATEST_CANDIDATE_TAG),3))

	$(eval n_release_candidates=$(call word-dot,$(LATEST_CANDIDATE_TAG),4))
	$(eval n_release_candidates=$(shell echo $$(($(n_release_candidates)+1))))
	git tag $(major).$(minor).$(patch).$(n_release_candidates)
endif
endif
