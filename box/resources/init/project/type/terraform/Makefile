.PHONY: aws/get-caller-identity
aws/get-caller-identity: ## Returns details  about the IAM user or role whose credentials are used to call the operation.
	aws sts get-caller-identity
	
.PHONY: terraform/init
terraform/init: ## Initializes a Terraform working directory
	terraform init

.PHONY: terraform/get
terraform/get: ## Downloads and installs modules for the configuration
	terraform get

.PHONY: terraform/fmt
terraform/fmt: ## Rewrites config files to canonical format
	terraform fmt

.PHONY: terraform/plan
terraform/plan: terraform/fmt terraform/validate terraform/workspace ## Generate and show an execution plan
ifdef environment
	terraform plan -out=plan.tfplan -var-file=environments/$(environment).tf
else
	@echo "argument environment NOT defined" && exit 1
endif

.PHONY: terraform/apply
terraform/apply:  ## Generates and shows the execution plan
	terraform apply plan.tfplan

.PHONY: terraform/validate
terraform/validate: ## Validates the Terraform files
	terraform validate

.PHONY: terraform/clean
terraform/clean: ## Removes local .terraform directory
	rm -rf .terraform/
	rm -rf terraform.tfstate.d
	rm -f *.tfplan
	rm -f terraform.tfstate*
	rm -f .terraform.tfstate*

.PHONY: terraform/workspace
terraform/workspace: ## Selects a workspace based on environment
ifdef environment
	@terraform workspace select $(environment) || (echo "terraform workspace not found, creating a new one..."; terraform workspace new $(environment))
else
	@echo "argument environment NOT defined" && exit 1
endif

.PHONY: terraform/destroy
terraform/destroy: ## Destroys Terraform-managed infrastructure
ifdef environment
	@terraform plan -destroy -out=destroy.tfplan -var-file=environments/$(environment).tf
	@terraform apply destroy.tfplan
else
	echo "argument environment NOT defined"
endif
