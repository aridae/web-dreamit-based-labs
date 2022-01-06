PROJECT_NAME := "web-dreamit-based-labs"
PKG := "github.com/aridae/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
 
.PHONY: all dep lint vet test test-coverage build clean
 
all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test-unit: ## Run unit tests
	@go test -v -p 1 --tags=unit ./...

test-integration: ## Run integration tests
	@go test -v -p 1 --tags=integration ./...

test-e2e: ## Run e2e tests
	@go test -v -p 1 --tags=e2e ./...

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt

# build: dep ## Build the binary file
# 	@go build -i -o build/main ./cmd/api_server/main.go
 
clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)/build
 
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run-local
run-local:
	sudo docker-compose up -d --force-recreate --remove-orphans

.PHONY: up-testing-containers-local
up-testing-containers-local:
	sudo docker-compose up -d --force-recreate --remove-orphans server_db redis

.PHONY: down-testing-containers-local
down-testing-containers-local:
	sudo docker-compose down

.PHONY: up-testing-containers
up-testing-containers:
	docker-compose up -d --force-recreate --remove-orphans server_db redis

.PHONY: down-testing-containers
down-testing-containers:
	docker-compose down

.PHONY: armageddon
armageddon:
	-make remove_containers
	-docker builder prune -f
	-docker network prune -f
	-docker volume rm $$(docker volume ls --filter dangling=true -q)
	-docker rmi $$(docker images -a -q) -f

# .PHONY: remove_containers
# remove_containers:
# 	-docker stop $$(docker ps -aq)
# 	-docker rm $$(docker ps -aq)
# .PHONY: integration_test
# integration_test:
# 	go test -tags=integration ./integration_tests -count=1 -run=$(INTEGRATION_TEST_SUITE_PATH) 
# .PHONY: test
# test:
# 	go test ./...
# .PHONY: cover
# cover:
# 	go test -coverprofile=coverage1.out -coverpkg=./... -cover ./...
# 	cat coverage1.out | grep -v mock | grep -v proto | grep -v cmd | grep -v models > cover.out
# 	go tool cover -func cover.out && go tool cover -html cover.out
# .DEFAULT_GOAL := run_local
