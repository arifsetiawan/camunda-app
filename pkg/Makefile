PKG := ./...

all: common

common:
	go build -v $(PKG)

test:
	go test -v `go list -f '{{if .TestGoFiles}}{{.ImportPath}}{{end}}' $(PKG)`

lint:
	../bin/golangci-lint run --config ../golangci.yml

.PHONY: all common test lint
