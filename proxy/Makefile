
NAME := proxy
SERVER := $(NAME)-service
SERVER_STATIC := $(NAME)-service-static

BUILD_DIR := build/bin
BINARIES := $(SERVER)
DEBUG_BINARIES := debug-$(SERVER)
STATIC_BINARIES := $(SERVER_STATIC)
LOCATION := github.com/arifsetiawan/camunda-app/$(NAME)

HUB ?= docker.io/arifsetiawan
TAG ?= $(shell git rev-parse HEAD)

all: $(BINARIES)

$(BINARIES): $(NAME)-%: $(BUILD_DIR)/$(NAME)-%
$(addprefix $(BUILD_DIR)/,$(BINARIES)): $(BUILD_DIR)/$(NAME)-%:
	@echo @ "$@"
	mkdir -p $(BUILD_DIR)
	go build -v -o $@ $(LOCATION)/server
	chmod +x $@
	@echo "Done building $(NAME) $*"

debug: $(DEBUG_BINARIES)
$(DEBUG_BINARIES): debug-$(NAME)-%: $(BUILD_DIR)/debug-$(NAME)-%
$(addprefix $(BUILD_DIR)/,$(DEBUG_BINARIES)): $(BUILD_DIR)/debug-$(NAME)-%:
	@echo @ "$@"
	mkdir -p $(BUILD_DIR)
	go build -v -gcflags='all=-N -l' -o $@ $(LOCATION)/server
	chmod +x $@
	@echo "Done building debug-$(NAME) $*"

$(STATIC_BINARIES): $(NAME)-%-static: $(BUILD_DIR)/$(NAME)-%-static
$(addprefix $(BUILD_DIR)/,$(STATIC_BINARIES)): $(BUILD_DIR)/$(NAME)-%-static:
	@echo "Building static binary"
	CGO_ENABLED=0 GOOS=linux go build \
		-a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo \
		-o $(BUILD_DIR)/$(NAME)-$*-static $(LOCATION)/server
	chmod +x $(BUILD_DIR)/$(NAME)-$*-static
	@echo "Done building static $(NAME) $*"

release: $(STATIC_BINARIES)

test:
	go test `go list -f '{{if .TestGoFiles}}{{.ImportPath}}{{end}}' ./...`

lint:
	../bin/golangci-lint run --config ../golangci.yml

docker-build: $(STATIC_BINARIES)
	docker build -t $(HUB)/$(SERVER):$(TAG) -f Dockerfile .

docker-push:
	docker push $(HUB)/$(SERVER):$(TAG)

docker-run:
	docker run $(HUB)/$(SERVER):$(TAG)

clean:
	rm -f $(BUILD_DIR)/$(SERVER) $(BUILD_DIR)/$(SERVER_STATIC) $(BUILD_DIR)/debug-$(SERVER)
	@echo "Cleaned $(SERVER) $(SERVER_STATIC) debug-$(SERVER)"

.PNOHY: all test lint clean release docker-build docker-push docker-run
