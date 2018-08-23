.PHONY: info setup setup-* install install-* build build-* clean clean-* sam-*

info:
	@(env | sort --ignore-case)

clean-golang:
	@(make --directory=lambdas/golang clean)

setup-golang:
	@(make --directory=lambdas/golang setup)

install-golang:
	@(make --directory=lambdas/golang install)

build-golang:
	@(make --directory=lambdas/golang build)

sam-lint:
	@(cfn-lint api/aws-sam/template.yaml)
