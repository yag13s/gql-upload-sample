BIN_NAME:=bin/gql-upload-sample
ENTRY_POINT:=cmd/main.go
BUILDFLAGS:=-gcflags 'all=-N -l'
VERSION:=$(shell git rev-parse --short HEAD)
LDFLAGS:=-s -w -extldflags="static"
GO_LDFLAGS_VERSION:=-X 'github.com/ShinsakuYagi/gql-upload-sample/model.Version=${VERSION}'

.PHONY: start
start: build
	./$(BIN_NAME)

.PHONY: build
build:
	go build -o $(BIN_NAME) -tags 'netgo' -installsuffix netgo \
		-ldflags '$(LDFLAGS) $(GO_LDFLAGS_VERSION)' \
		$(BUILDFLAGS) \
		$(ENTRY_POINT)

.PHONY: generate-graphql
generate-graphql:
	@go install github.com/99designs/gqlgen
	go generate ./graphql

.PHONY: bootstrap-golangci
bootstrap-golangci:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s

.PHONY: lint
lint:
	@if [ ! -e bin/golangci-lint ]; then $(MAKE) bootstrap-golangci ; fi
	bin/golangci-lint run ./...

.PHONY: test
test:
	CGO_ENABLED=0 go test -cover -coverprofile=./coverage.out ./... -count=1
	go tool cover -func=coverage.out | grep "total:"
