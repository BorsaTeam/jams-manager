RELEASE = $(RELEASE_VERSION)

# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test
GO_TOOL_COVER=$(GO_CMD) tool cover
BIN=bin
BINARY_NAME=jams-manager
MAIN_PATH=./server/cmd/server/main.go

build-linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GO_BUILD) -a -installsuffix cgo -ldflags "-X main.version=$(RELEASE)" -o ./$(BINARY_NAME) $(MAIN_PATH)

build-mac:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin $(GO_BUILD) -a -installsuffix cgo -ldflags "-X main.version=$(RELEASE)" -o ./$(BINARY_NAME) $(MAIN_PATH)

pack:
	docker build -t jams-manager/jams-manager_api:$(RELEASE) .

publish:
	heroku container:push jams-manager/jams-manager_api:$(RELEASE)

deploy:
	heroku container:release jams-manager/jams-manager_api:$(RELEASE)

release:
	curl --location --request POST 'https://api.github.com/repos/BorsaTeam/jams-manager/releases' --header 'Accept: application/vnd.github.inertia-preview+json' --header 'Authorization: token $(GITHUB_TOKEN)' --header 'Content-Type: application/json' --data-raw '{"tag_name": "$(RELEASE_VERSION)","target_commitish": "release-$(RELEASE_VERSION)","name": "Release $(RELEASE_VERSION)"}'

build-local:
	$(GO_BUILD) -o ./$(BINARY_NAME) $(MAIN_PATH)

clean:
	rm -rf $(BINARY_NAME)

test:
	mkdir -p $(BIN)
	$(GO_TEST) -short -coverprofile=$(BIN)/cov.out `go list ./... | grep -v vendor/`
	$(GO_TOOL_COVER) -func=$(BIN)/cov.out
