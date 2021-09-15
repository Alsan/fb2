SHELL := /bin/bash
BASE_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
VERSION ?= $(shell git describe --tags --always --match=v* 2> /dev/null || \
					 			cat $(CURDIR)/.version 2> /dev/null || echo v0)
VERSION_HASH = $(shell git rev-parse HEAD)

BIN = $(BASE_PATH)/bin
PATH := $(BIN):$(PATH)
export PATH

# printing
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

GO = GOGC=off go
# go module
MODULE = $(shell env GO111MODULE=on $(GO) list -m)

DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
VERSION_HASH = 	$(shell git rev-parse --short HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

LDFLAGS += -X "$(MODULE)/version.Version=$(VERSION)" -X "$(MODULE)/version.CommitSHA=$(VERSION_HASH)"

# tools
$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) installing $(PACKAGE)…)
	$Q env GOBIN=$(BIN) $(GO) install $(PACKAGE)

GOLANGCI_LINT = $(BIN)/golangci-lint
$(BIN)/golangci-lint: PACKAGE=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1

GOIMPORTS = $(BIN)/goimports
$(BIN)/goimports: PACKAGE=golang.org/x/tools/cmd/goimports@v0.1.5

all: build run

update:
	go mod tidy

clean:
	rm $(BIN)/*

build:
	go build -ldflags '$(LDFLAGS)' -o $(BIN)/fb-cli

run:
	$(BIN)/fb-cli downloadfile test.txt
