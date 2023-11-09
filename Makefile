# Go parameters
GOBUILD=go build
GOCLEAN=go clean -cache
GOLINT=$(GOCMD)lint
GOTEST=go test
GOGET=go get
BINARY_NAME=graph-ql-api

.PHONY: pre-build docker-test

pre-build:
	echo "building graph-ql-api application"

build: pre-build
	$(GOBUILD) -tags "redis" -o $(BINARY_NAME) ./cmd
test:
	REDIS_URL=localhost:80 $(GOTEST) ./... --cover

run: export LOGGER_MODE=dev
run: export ENABLE_CORS=true
run: export GIN_MODE=
run: export MYSQL_USER=graph-ql-api-user
run: export MYSQL_PASSWORD=graph-ql-api-password
run: export MYSQL_HOST=127.0.0.1
run: export MYSQL_PORT=8141
run: export MYSQL_DATABASE=graph-ql-api
run:
	go run ./cmd/main.go
