REGISTRY=techotron
VERSION=$(shell git rev-parse --short HEAD)
APP=online-quote-book:be-$(VERSION)
IMAGE=golang:1.16 go
GOLANGCI_TAG=v1.28.3
ENVIRONMENT=dev
LOCAL_GO=go

build: 
	docker build -t $(REGISTRY)/$(APP) .
run-db:
# Command to run db container for debugging of unit tests that rely on one
	docker-compose -f docker-compose.yaml -f docker-compose-debug.yaml up --build -d db
run-backend:
	docker-compose -f docker-compose.yaml -f docker-compose-debug.yaml up --build -d backend db
down:
	docker-compose -f docker-compose.yaml -f docker-compose-debug.yaml down -v

# https://docs.docker.com/compose/reference/up/
# Specify --exit-code-from test here to pass down the exit code to the calling process (docker-compose up). 
#   When this happens, the docker-compose up returns an exit code 1 which will attempt to stop all containers because it
#   "implies --abort-on-container-exit". Then the make down runs which results in a second SIGTERM getting sent, meaning
#   the resulting exit code is 2 - because: 
#     "If SIGINT or SIGTERM is sent again during this shutdown phase, the running containers are killed, and the exit code is 2.
# So if make test command exits with code 2, that means tests failed and make down command triggered an additional SIGINT signal 
#   which returned the code 2 instead of code 1 exit code
all-tests:
	docker-compose up --exit-code-from test test db
test: all-tests down
lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:$(GOLANGCI_TAG) golangci-lint run
publish: docker-login docker-push
docker-push:
	docker push $(REGISTRY)/$(APP)
docker-login:
	docker login -u techotron
coverage: test coverage-report
coverage-report: 
	$(LOCAL_GO) tool cover -html=c.out
coverage-ci:
	docker run --rm -v $(shell pwd):/app -w /app $(IMAGE) tool cover -func=c.out -o report.out
