PROJECT_NAME := go_boilerplate
PKG_LIST := $(shell go list ${PROJECT_NAME}/tests/testing/... | grep -v /vendor/)


.PHONY: all dep build core test

all: build

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


########################
### DEVELOP and TEST ###
########################
development:
	# booting up dependency containers
	@docker-compose up -d --build consul db

	# wait for consul container be ready
	@while ! curl --request GET -sL --url 'http://localhost:8500/' > /dev/null 2>&1; do printf .; sleep 1; done

	# setting KV, dependency of app
	@curl --request PUT --data-binary @config.local.json http://localhost:8500/v1/kv/${PROJECT_NAME}

	# building user_api
	@docker-compose up -d --build ${PROJECT_NAME}

test: ## Run unittests
	@go test -cover -short ${PKG_LIST} -v

coverage: ## Generate global code coverage report
	@go tool cover -func=cov.out

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)
	@docker-compose down

run: ## Run application
	@go run . serve
