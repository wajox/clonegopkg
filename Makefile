BUILD_DIR = build

test-all:
	go test -v ./...

.PHONY: build
build:
	go mod verify
	go mod tidy
	go build ${GOARGS}  -o ${BUILD_DIR}/main

gen:
	go generate ./...

deps:
	wire ./...


install-tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0
	go get github.com/google/wire/cmd/wire
	go get -u github.com/onsi/ginkgo/ginkgo
