.DEFAULT_GOAL := help

.PHONY: build-web run-web help stop-docker docker-up run-docker

build-web: ## Restores packages and builds web app
	cd web; yarn install
	cd web; npm run build

run-web: ## Runs web app locally
	cd web; npm start

stop-docker: ## Stops all docker containers
	docker-compose kill

docker-up: ## Starts docker containers
	docker-compose up -d

run-docker: build-web docker-up ## Run all containers

promo: ## Run skycoin promo back-end. To add arguments, do 'make ARGS="--foo" promo'.
	go run cmd/promo/promo.go ${ARGS}&

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'