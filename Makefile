HOST=erdos.localhost
PORT=3001
CONTAINER_PORT=3001
export APIURL=http://$(HOST):$(PORT)/v1/api

#export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
GOOS=linux
GOARCH=amd64
APP=fiber-rw
APP_STATIC=$(APP)-static
LDFLAGS="-w -s -extldflags=-static"

USERNAME=u$(shell date +%s)
EMAIL=$(USERNAME)@mail.com
PASSWORD=password
NEWMAN_URL=https://github.com/gothinkster/realworld/raw/main/api/Conduit.postman_collection.json


help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

download: ## Download go dependency
	go mod download

docs:
	go install github.com/swaggo/swag/cmd/swag@latest
	go generate .

generate: ## Generate swagger docs. Required https://github.com/gofiber/swagger
	go generate .

build: ## Build project with dynamic library(see shared lib with "ldd <your_file>")
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -race -o $(APP) .

build-static: ## Build project as single static linked executable file
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0  go build -ldflags=$(LDFLAGS)  -o $(APP_STATIC) .

build-image: ## Build docker image. Required https://www.docker.com
	docker build -t fiber-rw .

run: docs ## Run project
	go run -race .


test: ## Run unit test without race detection
	go test -v ./...

test-race: ## Run unit test with race detection
	go test -v -race  ./...

rtest: ## Run unit test with rich format.
	richgo test -v ./...

rtest-race: ## Run unit test with race detection in rich format.
	richgo test -v -race  ./...


newman: ## Run integration test. Required https://github.com/postmanlabs/newman
	newman run  --delay-request $(DELAY_REQUEST) --global-var "APIURL=$(APIURL)"  --global-var "USERNAME=$(USERNAME)" --global-var "EMAIL=$(EMAIL)" --global-var "PASSWORD=$(PASSWORD)" $(NEWMAN_URL)

newmanx: ## Run integration test when Node.js installed. Required https://nodejs.org
	npx newman run --bail --verbose  --delay-request $(DELAY_REQUEST) --global-var "APIURL=$(APIURL)"  --global-var "USERNAME=$(USERNAME)" --global-var "EMAIL=$(EMAIL)" --global-var "PASSWORD=$(PASSWORD)" $(NEWMAN_URL)


