include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## mod/versions: checks for new modules versions
.PHONY: mod/versions
mod/versions:
	@go list -m -versions go.mongodb.org/mongo-driver
	@go list -m -versions github.com/v8tix/kit

## run/temp-srv: runs the API server (from code)
.PHONY: run/temp-srv
run/temp-srv:
	@go run ./cmd/api -port=${WEB_PORT} -con-file=${DB_HOST_CONN_DIR} -db=${DB_NAME} -env=${ENV} -cors-trusted-origins=${CORS_ORIGINS_DEV} >>${LOG_DIR} 2>>${LOG_DIR}

## run/temp-srv-help: runs the API server help (from code)
.PHONY: run/temp-srv-help
run/temp-srv-help::
	@go run ./cmd/api -h

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy and vendor dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	@go fmt ./...
	@echo 'Vetting code...'
	@go vet ./...
	@echo 'Running tests...'
	@go clean -testcache
	@go test -race -failfast -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	@go mod tidy
	@go mod verify
	@echo 'Vendoring dependencies...'
	@go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

CURRENT_TIME = $(shell date --iso-8601=seconds)
GIT_DESCRIPTION = $(shell git describe --always --dirty --tags --long)
LINKER_FLAGS = '-s -X main.buildTime=${CURRENT_TIME} -X main.version=${GIT_DESCRIPTION}'

## build/srv: builds the API server
.PHONY: build/srv
build/srv:
	@echo 'Building ...'
	@GOOS=linux GOARCH=amd64 go build -ldflags=${LINKER_FLAGS} -o=./docker/api/${BIN_NAME} ./cmd/api
	@echo 'Artifact ready !'

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #

## run/help: runs the API server help
.PHONY: run/help
run/help:
	@./docker/api/${BIN_NAME} -h

## run/srv-log: runs the API server. App logs to external file.
.PHONY: run/srv-log
run/srv-log:
	@echo 'Running ...'
	@./docker/api/${BIN_NAME} -port=${WEB_PORT} -con-file=${DB_HOST_CONN_DIR} -db=${DB_NAME} -env=${ENV} -cors-trusted-origins=${CORS_ORIGINS_DEV} >>${LOG_DIR} 2>>${LOG_DIR}

## run/srv: runs the API server.
.PHONY: run/srv
run/srv:
	@echo 'Running ...'
	@./docker/api/${BIN_NAME} -port=${WEB_PORT} -con-file=${DB_HOST_CONN_DIR} -db=${DB_NAME} -env=${ENV} -cors-trusted-origins=${CORS_ORIGINS_DEV}

# ==================================================================================== #
# DOCKER
# ==================================================================================== #

## cntr/run: runs the container in the background
.PHONY: cntr/run
cntr/run:
	@./scripts/cntr_run.sh "${CPUS}" "${MEMORY}" "${IMAGE}" "${DB_HOST_CONN_DIR}" "${DB_CNTR_CONN_DIR}" "${HOST_GRPC_PORT}" "${CNTR_GRPC_PORT}" "${IMAGE_TAG}" "${GRPC_PORT}" "${DB_CNTR_CONN_DIR}" "${DB_NAME}" "${GRPC_PORT_TYPE}" "${CNTR_INFO_LOG_DIR}" "${CNTR_ERROR_LOG_DIR}"

## cntr/attach: attaches the running container
.PHONY: cntr/attach
cntr/attach:
	@docker exec -it "${IMAGE}" bash

## cntr/stop: stop the running container
.PHONY: cntr/stop
cntr/stop:
	@docker stop "${IMAGE}"

## cntr/delete: stop and delete the image
.PHONY: cntr/delete
cntr/delete: cntr/stop
	@docker rmi "${IMAGE_TAG}"

## cntr/build: builds the container image
.PHONY: cntr/build
cntr/build:
	@docker build --no-cache  -f ./docker/Dockerfile -t ${IMAGE_TAG} .

## cntr/push: push the container image to Docker Hub
.PHONY: cntr/push
cntr/push:
	@docker tag ${IMAGE_TAG} ${DOCKER_TAG}
	@docker push ${DOCKER_TAG}
