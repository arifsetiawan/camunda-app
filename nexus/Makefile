
NAME := nexus
SERVER := $(NAME)-service

BUILD_DIR := build/bin
BINARIES := $(SERVER)
LOCATION := github.com/arifsetiawan/camunda-app/$(NAME)

HUB ?= ghcr.io/arifsetiawan
TAG ?= $(shell git rev-parse HEAD)

all: $(BINARIES)

test:
	npm run test:unit
	npm run test:e2e

lint:
	npm run lint

docker-build:
	docker build -t $(HUB)/$(SERVER):$(TAG) -f Dockerfile .

docker-push:
	docker push $(HUB)/$(SERVER):$(TAG)

docker-run:
	docker run $(HUB)/$(SERVER):$(TAG)

.PNOHY: all test lint docker-build docker-push docker-run
