.PHONY: info setup setup-* install install-* build build-* clean clean-* sam-*

load-data:
	@(go run tools/dataloader/dataloader.go)

info:
	@(env | sort -f)
	@(go version)

clean-golang:
	@(make --directory=lambdas/golang clean)

build-golang:
	@(make --directory=lambdas/golang build)

sam-lint:
	@(cfn-lint api/aws-sam/template.yaml)
