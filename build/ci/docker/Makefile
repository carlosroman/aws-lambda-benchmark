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

pull-nodejs810:
	@($(DOCKER_SAM_PULL)nodejs8.10)

pull-nodejs10:
	@($(DOCKER_SAM_PULL)nodejs10.x)

info: 
	@echo Pulling latest lambda images...

all: info pull-go pull-python36 pull-python37 pull-nodejs810 pull-nodejs10
	@echo ... Lambda images updated!
