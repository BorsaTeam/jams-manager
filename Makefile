# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOTOOLCOVER=$(GOCMD) tool cover
BIN=bin
BINARY_NAME=jams-manager
MAIN_PATH=./server/cmd/server/main.go

build:
	$(GOBUILD) -o ./$(BINARY_NAME) $(MAIN_PATH)

clean:
	rm -rf $(BINARY_NAME)

test:
	mkdir -p $(BIN)
	$(GOTEST) -short -coverprofile=$(BIN)/cov.out `go list ./... | grep -v vendor/`
	$(GOTOOLCOVER) -func=$(BIN)/cov.out