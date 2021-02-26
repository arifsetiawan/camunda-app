# List of projects to build
PROJECTS := proxy external-task

all: build

clean: TARGET=clean
clean: default

build: TARGET=all
build: PROJECTS:=pkg $(PROJECTS)
build: default

release: TARGET=release
release: default 

docker-build: TARGET=docker-build
docker-build: default

docker-push: TARGET=docker-push
docker-push: default

LINTER := bin/golangci-lint
$(LINTER):
	wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.18.0

lint: TARGET=lint
# TODO(arif): Add test when we have some go files
lint: PROJECTS:=pkg $(PROJECTS)
lint: $(LINTER) default

default:
	@for PRJ in $(PROJECTS); do \
		echo "--- $$PRJ: $(TARGET) ---"; \
		$(MAKE) $(TARGET) -C $$PRJ; \
		if [ $$? -ne 0 ]; then \
			exit 1; \
		fi; \
	done

.PHONY: all $(PROJECTS) clean build docker-build docker-push release test lint default