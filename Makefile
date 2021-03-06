.PHONY: info setup setup-* install install-* build build-* clean clean-* sam-* setup/*

DOCKER_COMPOSE ?= docker-compose
DOCKER_COMPOSE_FILE := ./build/ci/codebuild/docker-compose.yml
DOCKER_COMPOSE_SAM_FILE := ./build/ci/codebuild/docker-compose.sam.yml
DOCKER_COMPOSE_LOAD_DATA_FILE := ./build/ci/codebuild/docker-compose.load-data.yml
DOCKER_COMPOSE_TEST_FILE := ./build/ci/codebuild/docker-compose.tests.yml
DOCKER_COMPOSE_CMD := $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) -f $(DOCKER_COMPOSE_SAM_FILE)
DOCKER_COMPOSE_LOAD_CMD := $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) -f $(DOCKER_COMPOSE_LOAD_DATA_FILE)
DOCKER_COMPOSE_TEST_CMD := $(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE_TEST_FILE)

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

clean-java8:
	@$(MAKE) -C lambdas/java8 clean

build-java8:
	@$(MAKE) -C lambdas/java8 fix-gradlew
	@$(MAKE) -C lambdas/java8 build

sam-lint:
	@(cfn-lint api/aws-sam/template.yaml)

docker-compose-env:
	@(echo "DOCKER_VOLUME_BASEDIR=$(CURDIR)" > .env)
	@(echo "LOCAL_M2=$(LOCAL_M2)" >> .env)

sam-start: docker-compose-env
	@$(DOCKER_COMPOSE_CMD) up $(DETACH)

sam-stop: docker-compose-env
	@$(DOCKER_COMPOSE_CMD) stop

sam-create-table: docker-compose-env
	@$(DOCKER_COMPOSE_LOAD_CMD) run create-table

sam-data-load: docker-compose-env
	@$(DOCKER_COMPOSE_LOAD_CMD) run data-load

sam-test-clean:
	@rm -rf test/bdd/target
	@echo "All clean"
	@mkdir -p test/bdd/target

sam-test-aggregate: docker-compose-env sam-test-clean
	@$(DOCKER_COMPOSE_TEST_CMD) run sam-test-aggregate

sam-test: docker-compose-env sam-test-clean
	@$(DOCKER_COMPOSE_TEST_CMD) run sam-test

sam-pull:
	@$(MAKE) -C build/ci/docker

setup/python/venv:
	@(test -d venv || python3 -m venv venv)
	@(./venv/bin/pip install --upgrade --requirement ./build/ci/codebuild/requirements.txt)

setup/python: setup/python/venv
	@(./venv/bin/python --version)

