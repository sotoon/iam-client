GIT ?= git
DOCKER_IMAGE := registry.cafebazaar.ir:5000/infrastructure/integration/sib/bepa-client
COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= $(shell $(GIT) describe --tags ${COMMIT} 2> /dev/null || echo "$(COMMIT)")


.PHONY: resolve
resolve:
	GOPRIVATE=git.cafebazaar.ir CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go mod vendor
	go mod tidy


.PHONY: build
build:
	GOPRIVATE=git.cafebazaar.ir CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod=vendor -a -o bin/example.out ./examples/test.go

.PHONY: mockgen
mockgen:
	mockgen -destination=mocks/mock_client.go -package=mocks -source=pkg/client/interface.go Cache --build_flags=--mod=mod

.PHONY: test
test:
	go test -mod=vendor -v ./... -coverprofile cover.out

.PHONY: coverage
coverage: test
	go tool cover -func=cover.out

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: fmt-check
fmt-check: fmt
	git diff-index --quiet HEAD

.PHONY: vet
vet:
	go vet ./...

.PHONY: coverage-serve
coverage-serve:
	go tool cover -html=cover.out
