.DEFAULT_GOAL := help

.PHONY: build-web run-web help

build-web: ## Restores packages and builds web app
	cd web; yarn install
	cd web; npm run build

run-web: ## Runs web app locally
	cd web; npm start

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'