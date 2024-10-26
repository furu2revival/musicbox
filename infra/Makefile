.PHONY: format
format:
	terraform fmt -recursive

.PHONY: lint
lint:
	tflint
	terraform validate
	terraform fmt -check -recursive
