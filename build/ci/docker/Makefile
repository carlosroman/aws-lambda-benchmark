.PHONY : all pull-*
.DEFAULT_GOAL := all

DOCKER ?= docker
DOCKER_SAM_PULL ?= $(DOCKER) pull lambci/lambda:

pull-go:
	@($(DOCKER_SAM_PULL)go1.x)

pull-python36:
	@($(DOCKER_SAM_PULL)python3.6)

pull-python37:
	@($(DOCKER_SAM_PULL)python3.7)

pull-node810:
	@($(DOCKER_SAM_PULL)nodejs8.10)

info: 
	@echo Pulling latest lambda images...

all: info pull-go pull-python36 pull-python37 pull-node810
	@echo ... Lambda images updated!
