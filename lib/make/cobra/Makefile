# Use camelCase (not snake_case/snake-case) for command names. 
# Otherwise, you will encounter errors. 
# For example, make cobra/add add-user is incorrect, but make cobra/add addUser is valid.
.PHONY: cobra/add
cobra/add: ## Add a Cobra command
ifdef command
	cd cobra && cobra add $(command)
else
	@echo "command not defined"
endif

.PHONY: cobra/del
cobra/del: ## Delete a Cobra command
ifdef command
	rm -f cmd/$(command).go
else
	@echo "command not defined"
endif


.PHONY: cobra/install
cobra/install:
	go get github.com/spf13/cobra/cobra

