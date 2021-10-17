GO_VERSION_SHORT:=$(shell go version | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.17","$(shell printf "$(GO_VERSION_SHORT)\n1.17" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.17. Found: $(GO_VERSION_SHORT))
endif

###############################################################################

BOT_NAME=bot
BOT_MAIN=cmd/$(BOT_NAME)/main.go
BOT_EXE=./bin/$(BOT_NAME)$(shell go env GOEXE)

###############################################################################

.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: build
build: .build

.PHONY: lint
lint:
	@command -v golangci-lint 2>&1 > /dev/null || (echo "Install golangci-lint" && \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(shell go env GOPATH)/bin" v1.42.1)
	golangci-lint run ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: style
style:
	find . -iname *.go | xargs gofmt -w

###############################################################################

.build:
	go mod download && CGO_ENABLED=0 go build \
		-o $(BOT_EXE) $(BOT_MAIN)

###############################################################################
