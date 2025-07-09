BINARY_NAME=codewise-cli
DOCKER_IMAGE=aryansharma04/codewise-cli

.PHONY: all build test docker-build docker-push

all: build

build:
	go build -o $(BINARY_NAME) main.go

test:
	go test ./tests/... -v

docker-build:
	docker build -t $(DOCKER_IMAGE):latest .

docker-push:
	docker push $(DOCKER_IMAGE):latest

clean:
	rm -f $(BINARY_NAME)
