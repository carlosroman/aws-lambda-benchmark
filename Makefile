.PHONY: info setup setup-* install install-* build build-* clean clean-* sam-*

load-data:
	@$(MAKE) --directory=tools/dataloader load-data

info:
	@(printenv | sort -if)
	@(go version)

clean-golang:
	@$(MAKE) -C lambdas/golang clean

build-golang:
	@$(MAKE) -C lambdas/golang build

sam-lint:
	@(cfn-lint api/aws-sam/template.yaml)
