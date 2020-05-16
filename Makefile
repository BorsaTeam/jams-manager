# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test
GO_TOOL_COVER=$(GO_CMD) tool cover
BIN=bin
BINARY_NAME=jams-manager
MAIN_PATH=./server/cmd/server/main.go

TAG=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export $(TAG)

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GO_BUILD) -a -installsuffix cgo -ldflags "-X main.version=$(TAG)" -o ./$(BINARY_NAME) $(MAIN_PATH)

pack:
	docker build -t gcr.io/jams-manager/jams-manager_api:$(TAG) .

upload:
	docker push gcr.io/jams-manager/jams-manager_api:$(TAG)

deploy:
	envsubst < k8s/deployment.yml | kubectl apply -f -

ship: test pack upload deploy

build-local:
	$(GO_BUILD) -o ./$(BINARY_NAME) $(MAIN_PATH)

clean:
	rm -rf $(BINARY_NAME)

test:
	mkdir -p $(BIN)
	$(GO_TEST) -short -coverprofile=$(BIN)/cov.out `go list ./... | grep -v vendor/`
	$(GO_TOOL_COVER) -func=$(BIN)/cov.out