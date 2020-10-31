.PHONY: help

help: ## Prints help for targets with comments
	@grep -E '^[a-zA-Z0-9.\ _-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lifecycle: ## Possible lifecycle of the app
	@echo "init-project -> check-tools -> update -> test -> build -> docker-build"

init: ## Initialize project, add some more needed dependencies
	@go get -u golang.org/x/lint/golint
	@go get github.com/gavv/httpexpect/v2
	@go get github.com/swaggo/swag/cmd/swag
	@go get gopkg.in/go-playground/validator.v10
	@go mod download
	@go mod vendor


check-tools: ## Print all available versions
	@go version
	@golint --version
	@yamllint --version
	@docker --version

update: ## Update go dependencies
	@go get -v all
	@go mod download
	@go mod tidy
	@go mod vendor

test: ## Execute linter and go tests
	@yamllint -c .yamllint .  || echo "🔥 yamllint syntax-check failed"
	@golint || echo "🔥 golint syntax-check failed"

build: ## Build This service
#	@go generate **/*.go
	@go build cmd/main.go || echo "🔥 go build failed"

docker-build: ## Build application and put it in docker container
	@docker build . || echo "🔥 docker build failed"