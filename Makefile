SHELL:=/bin/bash

# project details
PRODUCT = nsm
APPNAME = microservice-golang-userdata
PACKAGE = github.com/artprocessors/nsm-microservice-golang-userdata

# build variables
BRANCH_NAME ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE  ?= $(shell date '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT  ?= $(shell git rev-list -1 HEAD)
VERSION     ?= 0.0.1
AUTHOR      ?= $(shell git log -1 --pretty=format:'%an')

BUILD_OVERRIDES = \
	-X "$(PACKAGE)/pkg/app.Name=$(APPNAME)" \
	-X "$(PACKAGE)/pkg/app.Product=$(PRODUCT)" \
	-X "$(PACKAGE)/pkg/app.Branch=$(BRANCH_NAME)" \
	-X "$(PACKAGE)/pkg/app.BuildDate=$(BUILD_DATE)" \
	-X "$(PACKAGE)/pkg/app.Commit=$(GIT_COMMIT)" \
	-X "$(PACKAGE)/pkg/app.Version=$(VERSION)" \
	-X "$(PACKAGE)/pkg/app.Author=$(AUTHOR)" \

# misc
CMDPATH = $(shell ls -1 cmd/*/main.go | head -1)
TMP = ._tmp
WORKSPACE = $(shell pwd)

# note regarding "GO111MODULE=on" :
# once we move our builds from $GOPATH, we can also remove this
export GO111MODULE=on

install:
	# get linter - make sure this version matches our CI tool
	command -v golangci-lint || go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.16.0

	# go-swagger
	command -v swagger || go get github.com/go-swagger/go-swagger/cmd/swagger@v0.18.0

	# download dependencies
	go mod download

build:
	go mod vendor
	@echo "Building with overrides '$(BUILD_OVERRIDES)'"
	CGO_ENABLED=0 GOARCH=amd64 \
		go build -a \
		-installsuffix cgo \
		-ldflags='-w -s $(BUILD_OVERRIDES)' \
		-o main $(CMDPATH)

clean-generated-server:
	rm -rf cmd
	rm -rf $(TMP)
	mkdir $(TMP)
	mv ./restapi/configure_*.go $(TMP)/ || true
	rm -rf restapi
	mkdir restapi
	mv $(TMP)/* ./restapi/ || true
	rm -rf $(TMP)

clean-generated-client:
	rm -rf ./client/sdk

clean: clean-generated-server clean-generated-client

build-dev: install
	@echo "Building with overrides '$(BUILD_OVERRIDES)'"
	go build \
		-ldflags='-w -s $(BUILD_OVERRIDES)' \
		-o main $(CMDPATH)

clean-swagger:
	rm -f swagger.{yaml,json}

build-swagger: clean-swagger
	swagger mixin --format=yaml --output ./api/swagger.yaml ./spec/userdata-swagger.yaml
	swagger mixin --format=json --output ./api/swagger.json ./spec/userdata-swagger.yaml

generate-server: clean-generated-server
	swagger generate server \
		--name nsm-microservice-golang-userdata \
		--model-package restapi/models \
		-f ./spec/userdata-swagger.yaml

generate: generate-server

run: build
	HOST=0.0.0.0 \
	PORT=8080 \
	LOG_LEVEL=debug \
	LOG_FORMAT=json \
		./main

run-dev:
	PORT=8080 \
	BASE_URL=http://localhost:8080 \
	PORT=8080 \
	LOG_LEVEL=debug \
	LOG_FORMAT=json \
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_USERNAME=nsm \
	DB_PASSWORD=nsm \
	DB_NAME=userdata \
	DB_SSL=disable \
		go run $(CMDPATH)

.PHONY: test
test:
	go test ./internal/... -cover -coverprofile=cover.out
	go tool cover -html=cover.out -o ./coverage.html


# Sonarqube friendly lint
# NOTE: If there are linting issues, this will return 0,
# but you will need to fix them before merging your PR (via sonarqube integration).
lint:
	golangci-lint run ./internal/... ./cmd/...

# Sonarqube friendly test
.PHONY: test-report
test-report:
	set -o pipefail
	go test -cover -v -coverprofile=coverage.out -json ./internal/... ./pkg/... | tee report.json

