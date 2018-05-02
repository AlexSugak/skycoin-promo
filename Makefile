.DEFAULT_GOAL := help

.PHONY: build-web test-web run-web help stop-docker docker-up run-docker promo lint check install-linters

build-web: ## Restores packages and builds web app
	cd web; yarn install
	cd web; npm run build

run-web: ## Runs web app locally
	cd web; npm start

test-web: ## Run UI tests
	cd web; yarn install
	cd web; CI=true yarn test

stop-docker: ## Stops all docker containers
	docker-compose kill

docker-up: ## Starts docker containers
	docker-compose up -d

run-docker: build-web docker-up ## Run all containers

promo: ## Run skycoin promo back-end. To add arguments, do 'make ARGS="--foo" promo'.
	go run cmd/promo/promo.go ${ARGS}&

lint: ## Run linters. Use make install-linters first.
	vendorcheck ./...
	gometalinter --deadline=3m -j 2 --disable-all --tests --vendor \
		-E deadcode \
		-E errcheck \
		-E gas \
		-E goconst \
		-E gofmt \
		-E goimports \
		-E golint \
		-E ineffassign \
		-E interfacer \
		-E maligned \
		-E megacheck \
		-E misspell \
		-E nakedret \
		-E structcheck \
		-E unconvert \
		-E unparam \
		-E varcheck \
		-E vet \
		./...

check: lint test-web ## Run tests and linters

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	go get -u github.com/alecthomas/gometalinter
	gometalinter --vendored-linters --install

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'