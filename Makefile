.PHONY: info setup setup-* install install-* build build-* clean clean-* sam-*

DOCKER_COMPOSE ?= docker-compose
DOCKER_COMPOSE_FILE := ./build/ci/codebuild/docker-compose.yml
DOCKER_COMPOSE_SAM_FILE := ./build/ci/codebuild/docker-compose.sam.yml
DOCKER_COMPOSE_LOAD_DATA_FILE := ./build/ci/codebuild/docker-compose.load-data.yml
DOCKER_COMPOSE_CMD := $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) -f $(DOCKER_COMPOSE_SAM_FILE)
DOCKER_COMPOSE_LOAD_CMD := $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) -f $(DOCKER_COMPOSE_LOAD_DATA_FILE)

ifeq ($(DETACH_ENABLED), true)
	DETACH := --detach
endif

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

docker-compose-env:
	@(echo "DOCKER_VOLUME_BASEDIR=$(CURDIR)" > .env)

sam-start: docker-compose-env
	@$(DOCKER_COMPOSE_CMD) up $(DETACH)

sam-stop: docker-compose-env
	@$(DOCKER_COMPOSE_CMD) stop

sam-create-table: docker-compose-env
	@$(DOCKER_COMPOSE_LOAD_CMD) run create-table

sam-data-load: docker-compose-env
	@$(DOCKER_COMPOSE_LOAD_CMD) run data-load

sam-pull:
	@$(MAKE) -C build/ci/docker
