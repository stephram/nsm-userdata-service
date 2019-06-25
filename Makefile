SHELL:=/bin/bash

# project details
PRODUCT = nsm
APPNAME = microservice-golang-userdata
PACKAGE = github.com/artprocessors/nsm-microservice-golang-userdata

# build variables
BRANCH_NAME ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE  ?= $(shell date '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT  ?= $(shell git rev-list -1 HEAD)
VERSION     ?= 0.0.0
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
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
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

build-docker: build
	docker build -t nsm-microservice-golang-userdata -f Dockerfile .

clean-swagger:
	rm -f swagger-{private,public}.{yaml,json}

build-swagger: clean-swagger
	swagger mixin --format=yaml --output swagger.yaml ./spec/userdata-swagger.yaml
	swagger mixin --format=json --output swagger.json ./spec/userdata-swagger.yaml

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
	LOG_FORMAT=text \
		./main

run-dev:
	PORT=8080 \
	API_KEY=TESTKEY123 \
	BASE_URL=http://localhost:8081 \
	PORT=8080 \
	ENVIRONMENT_NAME=local \
	ENVIRONMENT_NUMBER=0 \
	LOG_LEVEL=debug \
	LOG_FORMAT=text \
	AWS_ENDPOINT=MOCK \
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_USERNAME=vti \
	DB_PASSWORD=vti \
	DB_NAME=pdp \
	DB_SSL=disable \
		go run $(CMDPATH)

run-docker:
	docker run -it \
		--publish 8080:8080 \
		--env AWS_DEFAULT_REGION=us-west-2 \
		--env PORT=8080 \
		--env HOST=0.0.0.0 \
		--env LOG_LEVEL=debug \
		--env LOG_FORMAT=text \
		trf-microservice-golang-pdp

build-codebuild-image:
	# nbos image
	cd /tmp; [ -d "/tmp/nbos-cloud-docker-images" ] || git clone git@github.com:transurbantech/nbos-cloud-docker-images.git
	cd /tmp/nbos-cloud-docker-images/golang; docker build -t nbos-codebuild-golang .

	# aws image
	cd /tmp; [ -d "aws-codebuild-docker-images" ] || git clone git@github.com:aws/aws-codebuild-docker-images.git
	cd /tmp/aws-codebuild-docker-images/ubuntu/golang/1.10/; docker build -t aws/codebuild/golang:1.10 .

run-codebuild-local: build-codebuild-image

	[ -d "./artifacts" ] || mkdir artifacts
	cd /tmp; [ -d "aws-codebuild-docker-images" ] || git clone git@github.com:aws/aws-codebuild-docker-images.git

	/tmp/aws-codebuild-docker-images/local_builds/codebuild_build.sh \
		-i aws/codebuild/golang:1.10 \
		-s $(WORKSPACE) \
		-a $(WORKSPACE)/artifacts \
		-e $(WORKSPACE)/buildenv

build-mocks:
	mockery -name SQSAPI -dir ./vendor/github.com/aws/aws-sdk-go -recursive
	mockery -name TripStatusRepo -dir ./internal -recursive
	mockery -name Artifact -dir ./internal -recursive
	mockery -name IService -dir ./internal/hclient -recursive
	mockery -name IConfig -dir ./config -recursive

.PHONY: test
test:
	go test ./internal/... -cover -coverprofile=cover.out
	go tool cover -html=cover.out -o ./coverage.html

test-integration: local-db
	go test -cover -v  ./test/...

xray:
	cd /tmp \
	&& [ -d "./xray" ] \
	|| ( \
		curl -O https://s3.dualstack.us-east-2.amazonaws.com/aws-xray-assets.us-east-2/xray-daemon/aws-xray-daemon-macos-3.x.zip \
		&& unzip -d xray aws-xray-daemon-macos-3.x.zip \
	) \
	&& cd xray \
	&& ./xray_mac -o -n us-west-2

# Sonarqube friendly lint
# NOTE: If there are linting issues, this will return 0,
# but you will need to fix them before merging your PR (via sonarqube integration).
lint-report:
	golangci-lint -v run ./internal/... ./cmd/... ./config/...



# executes pact tests
.PHONY: test-pact
test-pact:
	rm -rf pacts || true
	go clean -testcache
	go test ./test/pact/consumer/... -v

# Sonarqube friendly test
.PHONY: test-report
test-report: test-pact
	set -o pipefail
	go test -cover -v -coverprofile=coverage.out -json ./internal/... ./pkg/... | tee report.json

# Jenkins CI required code for deployment
.ONESHELL:
deployci:
	if [ "$(ARTIFACT_NAME)" == "" ]
	then
	 	echo "Error: ARTIFACT_NAME is required"
		exit 1
	fi
	OPS_ACCOUNT_ID=216871930768
	HEALTH_CHECK_PATH='/health'
	DEPLOY_DATE=$$(date -u +"%Y-%m-%dT%H:%M:%SZ")

	ansible-playbook -vvv aws/deploy-ecs-container.yml -e "
	opsaccountid=$${OPS_ACCOUNT_ID}
	awsaccountid=$(AWS_ACCOUNT_ID)
	credstashNo=$(CREDSTASH_NO)
	healthcheckpath='$${HEALTH_CHECK_PATH}'
	containername=$(CONTAINERNAME)
	region=$(AWS_REGION)
	tagproduct=$(PRODUCT)
	tagenvironment=$(ENV)
	tagenvironmentnumber=$(ENV_NO)
	ecrname=$(ECRNAME)
	containertag=$(ARTIFACT_NAME)
	deploydate=$${DEPLOY_DATE}
	kmskeyarn=arn:aws:kms:$(AWS_REGION):$(AWS_ACCOUNT_ID):alias/$(PRODUCT)-$(ENV)-credstash
	loglevel=info
	jenkins_workspace="$(WORKSPACE)"
	cluster_env_no=$(CLUSTER_ENV_NO)
	version="$(ARTIFACT_NAME)"
	accountname=$(ACCOUNT_NAME)
	"

.SILENT:
cleanci:
	rm -f build.json
	rm -f buildoutput.json
	rm -f deployparameters.txt
	rm -f containersinventory.yml

.ONESHELL:
compileci: cleanci
	if [ "$(ARTIFACT_NAME)" == "" ]
	then
	 	echo "Error: ARTIFACT_NAME is required"
		exit 1
	fi

	DEPLOYSPEC=$$(cat deployspec.json | tr -d '\n')
	DOCKER_IMAGES=$$(aws ecr list-images --repository-name $(ECRNAME) --query "imageIds[].imageTag"  | grep $(ARTIFACT_NAME) | awk -F '"' '{print $$2}')
	echo "images:" > containersinventory.yml
	for img in $$DOCKER_IMAGES; do
		SERVICENAME=$$(echo $$img | cut -d"_" -f3)
		CONTAINERFOLDER=$$(echo $${DEPLOYSPEC} | jq '.[]|.[]| select(.servicename|contains("'$${SERVICENAME}'"))| .containerfolder')
		RULEPATH=$$(echo $${DEPLOYSPEC} | jq '.[]|.[]| select(.servicename|contains("'$${SERVICENAME}'"))| .rulepath')
		DESIREDCOUNT=$$(echo $${DEPLOYSPEC} | jq '.[]|.[]| select(.servicename|contains("'$${SERVICENAME}'"))| .desiredcount')
		CONTAINERPORT=$$(echo $${DEPLOYSPEC} | jq '.[]|.[]| select(.servicename|contains("'$${SERVICENAME}'"))| .containerport')
		MEMORYMIN=$$(echo $${DEPLOYSPEC} | jq '.[]|.[]| select(.servicename|contains("'$${SERVICENAME}'"))| .memorymin')
		MEMORYMAX=$$(echo $${DEPLOYSPEC} | jq '.[]|.[]| select(.servicename|contains("'$${SERVICENAME}'"))| .memorymax')

		if [ "$${CONTAINERFOLDER}" == "" ]
		then
		 echo "Error: deployspec.json missing configuration for $${SERVICENAME}"
		exit 1
		fi

		# Generate the inventory for the ansible playbook

		echo "- containertag: $${img}" >> containersinventory.yml
		if [ "$${RULEPATH}" != "" ]
		then
		echo "  rulepath: $${RULEPATH}" >> containersinventory.yml
		fi
		echo "  desiredcount: $${DESIREDCOUNT}" >> containersinventory.yml
		echo "  containerfolder: $${CONTAINERFOLDER}" >> containersinventory.yml
		echo "  containerport: $${CONTAINERPORT}" >> containersinventory.yml
		echo "  memorymin: $${MEMORYMIN}" >> containersinventory.yml
		echo "  memorymax: $${MEMORYMAX}" >> containersinventory.yml
	done
