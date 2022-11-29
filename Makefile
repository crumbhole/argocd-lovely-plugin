.PHONY:	all clean code-vet code-fmt test get docker

DEPS := $(shell find . -type f -name "*.go" -printf "%p ")

all: code-vet code-fmt test build/argocd-lovely-plugin

docker:
	docker build .

clean:
	$(RM) -rf build

get: $(DEPS)
	go get ./...

test: get
	go test ./...

test_verbose: get
	go test -v ./...

build/argocd-lovely-plugin: $(DEPS) get
	mkdir -p build
	go build -buildvcs=false -o build ./...

code-vet: $(DEPS) get
## Run go vet for this project. More info: https://golang.org/cmd/vet/
	@echo go vet
	go vet $$(go list ./... )

code-fmt: $(DEPS) get
## Run go fmt for this project
	@echo go fmt
	go fmt $$(go list ./... )

lint: $(DEPS) get
## Run golint for this project
	@echo golint
	golint $$(go list ./... )

coverage: $(DEPS) get
#	go test -v ./... -coverpkg=./pkg/processor/... -coverprofile=coverage.out
	go test -v ./... -parallel=1 -coverpkg=./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
